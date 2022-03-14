package aggregate

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/airdb/sls-bbhj/pkg/util"
	"github.com/google/uuid"
)

// LostAggr defines functions used to handle lost request.
type LostAggr interface {
	List(ctx context.Context, opts schema.LostListRequest) ([]*schema.LostItem, error)
	GetByID(ctx context.Context, id uint) (*schema.LostDetail, error)
	GetWxMpCode(ctx context.Context, id uint) []byte
	Create(ctx context.Context, opts schema.LostCreateRequest) error
}

type lostAggr struct {
	repo repository.Factory
}

const (
	LOST_WXMP_CODE_FILENAME = `wxmp_code.jpg`
)

var _ LostAggr = (*lostAggr)(nil)

func newLosts(aggr *aggregate) *lostAggr {
	return &lostAggr{repo: aggr.repo}
}

// List returns lost list in the storage. This function has a good performance.
func (u *lostAggr) List(ctx context.Context, opts schema.LostListRequest) ([]*schema.LostItem, error) {
	if categoryId, err := strconv.Atoi(opts.Category); err == nil {
		category, err := u.repo.Categories().GetById(ctx, categoryId)
		if err == nil {
			opts.Category = category.Name
		}
	}

	items, err := u.repo.Losts().List(ctx, opts)
	if err != nil {
		log.Printf("list losts from storage failed: %s", err.Error())

		return nil, err
	}

	data := make([]*schema.LostItem, len(items))
	for k, v := range items {
		data[k] = &schema.LostItem{
			ID:            v.ID,
			Title:         v.Subject,
			Category:      v.Category,
			Name:          v.Nickname,
			Babyid:        v.Babyid,
			Introduction:  v.Subject,
			MissedAt:      v.MissedAt.Format(defaultTimeFormat),
			MissedAddress: v.MissedAddress,
		}
	}

	return data, nil
}

// GetByID returns a lost detail in the storage. This function has a good performance.
func (u *lostAggr) GetByID(ctx context.Context, id uint) (*schema.LostDetail, error) {
	item, err := u.repo.Losts().GetByID(ctx, id)
	if err != nil {
		log.Printf("get losts from storage failed: %s", err.Error())

		return nil, err
	}

	stat, err := u.repo.Losts().GetStatByID(ctx, id)
	if err != nil {
		log.Printf("get losts from storage failed: %s", err.Error())

		return nil, err
	}

	return &schema.LostDetail{
		ID:    item.ID,
		Title: item.Subject,

		Name:         item.Nickname,
		Babyid:       item.Babyid,
		Introduction: item.Subject,
		ShareCount:   stat.ShareCount,
		ShowCount:    stat.ShowCount,

		// 基础信息
		NameMore: item.Nickname,
		Gender: func() string {
			switch item.Gender {
			case 1:
				return "男"
			case 2:
				return "女"
			default:
				return "未知"
			}
		}(),
		BirthedAt: item.BirthedAt.Format(defaultTimeFormat),
		Carousel:  []schema.CarouselItem{},

		// 失踪信息
		MissAt:     item.MissedAt.Format(defaultTimeFormat),
		MissAddr:   item.MissedAddress,
		MissHeight: item.Height,
		Details:    item.Details,
		Character:  item.Characters + "<br>" + item.Details,

		// 寻亲信息
		Category: item.Category,
		DataFrom: item.DataFrom,
		Follower: item.Follower,

		WxMore: &schema.WxMore{
			ShareAppMessage: &schema.WxShareAppMessage{
				ShareKey: uuid.New().String(),
				Title:    item.Title,
				ImageURL: item.AvatarURL,
			},
			ShareTimeline: &schema.WxShareTimeline{
				ShareKey: uuid.New().String(),
				Title:    item.Title,
				Query: func() string {
					query := url.Values{}
					query.Add("lost_id", strconv.Itoa(int(item.ID)))
					return query.Encode()
				}(),
				ImageURL: item.AvatarURL,
			},
			CodeUnlimit: &schema.WxCodeUnlimit{
				URL: fmt.Sprintf(`/v1/lost/%d/%s`, item.ID, LOST_WXMP_CODE_FILENAME),
			},
		},
	}, nil
}

func (u *lostAggr) GetWxMpCode(ctx context.Context, id uint) []byte {
	wx := util.NewWechatMiniProgram(util.NewWechat())
	code, err := wx.CodeUnlimit(
		`pages/article/detail/index`,
		fmt.Sprintf("id=%d", id),
	)
	if err != nil {
		return []byte("error")
	}

	return code
}

func (u *lostAggr) Create(ctx context.Context, in schema.LostCreateRequest) error {
	return u.repo.Losts().Create(ctx, in)
}
