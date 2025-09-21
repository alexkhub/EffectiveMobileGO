package subscriptionservice

import (
	"time"


	"github.com/google/uuid"
	"gopkg.in/guregu/null.v3"
)


type CreateSubscription struct{
	ServiceName string `json:"service_name" binding:"required"`
	Price int `json:"price" binding:"required"`
	UserId uuid.UUID `json:"user_id" binding:"required,uuid"`
	StartDate string `json:"start_date" binding:"required" time_format:"01-2006"`
}

type Subscription struct{
	Id uuid.UUID `json:"id" binding:"omitempty"`
	ServiceName string `json:"service_name" binding:"omitempty"`
	Price int `json:"price" binding:"omitempty"`
	UserId uuid.UUID `json:"user_id" binding:"omitempty,uuid"`
	StartDate string `json:"start_date" binding:"omitempty" time_format:"01-2006"`
}

type UpdateSubscription struct{
	Id uuid.UUID
	ServiceName null.String `json:"service_name" binding:"omitempty"`
	Price null.Int `json:"price" binding:"omitempty"`
	UserId null.String`json:"user_id" binding:"omitempty,uuid"`
	StartDate null.String `json:"start_date" binding:"omitempty" time_format:"01-2006"`
}

type PreparationSubscription struct{
	Id uuid.UUID `db:"id"`
	ServiceName string `db:"service_name"`
	Price int `db:"price"`
	UserId uuid.UUID `db:"user_id"`
	StartDate time.Time `db:"start_date"`
}

type FilterSubscription struct{
	StartDateFrom null.String `binding:"omitempty" time_format:"01-2006"`
	StartDateTo null.String `binding:"omitempty" time_format:"01-2006"`
	UserId null.String `binding:"omitempty,uuid"`
	ServiceName null.String `binding:"omitempty"`
}
