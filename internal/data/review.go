package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/open-portfolios/review/internal/biz"
	"github.com/open-portfolios/review/internal/data/model"
	"github.com/open-portfolios/review/internal/data/query"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	if err := r.data.q.ReviewInfo.WithContext(ctx).Save(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	t := r.data.q.ReviewInfo
	return t.WithContext(ctx).Where(t.OrderID.Eq(orderID)).Find()
}

func (r *reviewRepo) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	t := r.data.q.ReviewInfo
	return t.WithContext(ctx).Where(t.ReviewID.Eq(reviewID)).First()
}

func (r *reviewRepo) SaveReply(ctx context.Context, reply *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	t := r.data.q.ReviewInfo
	review, err := t.WithContext(ctx).Where(t.ReviewID.Eq(reply.ReviewID)).First()
	if err != nil {
		return nil, err
	}

	if review.HasReply == 1 {
		return nil, ErrAlreadyReplied
	}
	if review.StoreID != reply.StoreID {
		return nil, ErrHorizontalOverreach
	}
	err = r.data.q.Transaction(func(tx *query.Query) error {
		if err := tx.ReviewReplyInfo.
			WithContext(ctx).
			Save(reply); err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply create reply fail, err:%v", err)
			return err
		}
		if _, err := tx.ReviewInfo.
			WithContext(ctx).
			Where(tx.ReviewInfo.ReviewID.Eq(reply.ReviewID)).
			Update(tx.ReviewInfo.HasReply, 1); err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply update review fail, err:%v", err)
			return err
		}
		return nil
	})
	return reply, err
}

func (r *reviewRepo) GetReviewReply(ctx context.Context, reviewID int64) (*model.ReviewReplyInfo, error) {
	t := r.data.q.ReviewReplyInfo
	return t.WithContext(ctx).Where(t.ReviewID.Eq(reviewID)).First()
}

func (r *reviewRepo) AuditReview(context.Context, *biz.AuditParam) error {
	panic("unimplemented")
}

func (r *reviewRepo) AppealReview(context.Context, *biz.AppealParam) error {
	panic("unimplemented")
}

func (r *reviewRepo) AuditAppeal(context.Context, *biz.AuditAppealParam) error {
	panic("unimplemented")
}

func (r *reviewRepo) ListReviewByUserID(ctx context.Context, userID int64, offset, limit int) ([]*model.ReviewInfo, error) {
	t := r.data.q.ReviewInfo
	return t.WithContext(ctx).
		Where(t.UserID.Eq(userID)).
		Order(t.ID.Desc()).
		Offset(offset).
		Limit(limit).
		Find()
}
