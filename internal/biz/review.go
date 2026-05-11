package biz

import (
	"context"

	v1 "github.com/open-portfolios/review/api/review/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/open-portfolios/review/internal/data/model"
)

type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	GetReview(context.Context, int64) (*model.ReviewInfo, error)
	SaveReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
	GetReviewReply(context.Context, int64) (*model.ReviewReplyInfo, error)
	AuditReview(context.Context, *AuditParam) error
	AppealReview(context.Context, *AppealParam) error
	AuditAppeal(context.Context, *AuditAppealParam) error
	ListReviewByUserID(ctx context.Context, userID int64, offset, limit int) ([]*model.ReviewInfo, error)
}

type ReviewUsecase struct {
	repo  ReviewRepo
	flake SnowflakeRepo
	log   *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, flake SnowflakeRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo:  repo,
		flake: flake,
		log:   log.NewHelper(logger),
	}
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReview %v", review.OrderID)

	// Check existence
	exist, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, v1.ErrorDatabaseFailure("database failure: %v", err)
	}
	if len(exist) > 0 {
		return nil, v1.ErrorAlreadyReviewed("order %v already reviewed", review.OrderID)
	}

	// Generate Snowflake ID
	review.ReviewID, err = uc.flake.Generate(ctx)
	if err != nil {
		return nil, err
	}

	return uc.repo.SaveReview(ctx, review)
}

func (uc *ReviewUsecase) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] GetReview reviewID:%v", reviewID)
	return uc.repo.GetReview(ctx, reviewID)
}

func (uc *ReviewUsecase) CreateReply(ctx context.Context, param *ReplyParam) (*model.ReviewReplyInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReply param:%v", param)
	id, err := uc.flake.Generate(ctx)
	if err != nil {
		return nil, err
	}
	reply := &model.ReviewReplyInfo{
		ReplyID:   id,
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	}
	return uc.repo.SaveReply(ctx, reply)
}

func (uc *ReviewUsecase) AuditReview(ctx context.Context, param *AuditParam) error {
	uc.log.WithContext(ctx).Debugf("[biz] AuditReview param:%v", param)
	return uc.repo.AuditReview(ctx, param)
}

func (uc ReviewUsecase) AppealReview(ctx context.Context, param *AppealParam) error {
	uc.log.WithContext(ctx).Debugf("[biz] AppealReview param:%v", param)
	return uc.repo.AppealReview(ctx, param)
}
