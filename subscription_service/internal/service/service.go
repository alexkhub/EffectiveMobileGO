package service

import (
	"context"
	subscriptionservice "effective_mobile"
	"effective_mobile/internal/repository"

	"github.com/google/uuid"
)




type Subsription interface{
	CreateSubscriptionService(ctx context.Context, request subscriptionservice.CreateSubscription)(uuid.UUID, error)
	GetSubscriptionService(ctx context.Context, subId uuid.UUID) (subscriptionservice.Subscription, error)
	ListSubscriptionService(ctx context.Context) ([]subscriptionservice.Subscription, error)
	DeleteSubscriptionService(ctx context.Context, subId uuid.UUID) error
	UpdateSubscriptionService(ctx context.Context, request subscriptionservice.UpdateSubscription) error
	TotalPriceService(ctx context.Context, filter subscriptionservice.FilterSubscription) (int64, error)
}

type Sevice struct{
	Subsription
}


type ServiceDeps struct{
	Repos *repository.Repository
}

func NewService(deps *ServiceDeps) *Sevice{
	return &Sevice{
		Subsription: NewSubsriptionService(deps.Repos.Subsription),
	}
}