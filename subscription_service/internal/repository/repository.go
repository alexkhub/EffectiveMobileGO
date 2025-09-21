package repository

import (
	"context"
	subscriptionservice "effective_mobile"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)



type Subsription interface{
	CreateSubscriptionRepository(ctx context.Context, request subscriptionservice.CreateSubscription)(uuid.UUID, error)
	GetSubscriptionRepository(ctx context.Context, subId uuid.UUID) (subscriptionservice.PreparationSubscription, error)
	ListSubscriptionRepository(ctx context.Context) ([]subscriptionservice.PreparationSubscription, error)
	DeleteSubscriptionRepository(ctx context.Context, subId uuid.UUID) (error)
	UpdateSubscriptionRepository(ctx context.Context, request subscriptionservice.UpdateSubscription) error
	TotalPriceRepository(ctx context.Context, filter subscriptionservice.FilterSubscription)(int64, error)
}

type Repository struct{
	Subsription
}

type RepositoryDeps struct{
	DB *sqlx.DB
}

func NewRepository(deps *RepositoryDeps) *Repository{
	return &Repository{
		Subsription: NewSubsriptionRepository(deps.DB),
	}
}


