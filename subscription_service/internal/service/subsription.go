package service

import (
	"context"
	subscriptionservice "effective_mobile"
	"effective_mobile/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)


type SubsriptionService struct{
	repos repository.Subsription
}

func NewSubsriptionService(repos repository.Subsription) *SubsriptionService{
	return &SubsriptionService{
		repos: repos,
	}
}


func (s *SubsriptionService) CreateSubscriptionService(ctx context.Context, request subscriptionservice.CreateSubscription)(uuid.UUID, error){
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "CreateSubscriptionService",
	}).Debug()

	return s.repos.CreateSubscriptionRepository(ctx, request)

}

func (s *SubsriptionService) GetSubscriptionService(ctx context.Context, subId uuid.UUID) (subscriptionservice.Subscription, error){
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "GetSubscriptionService",
	}).Debug()

	data, err := s.repos.GetSubscriptionRepository(ctx, subId)
	if err != nil{
		return subscriptionservice.Subscription{}, err
	}
	startDate := data.StartDate.Format("01-2006")

	return subscriptionservice.Subscription{
		Id: data.Id,
		ServiceName: data.ServiceName,
		Price: data.Price,
		UserId: data.UserId,
		StartDate: startDate,
	}, nil

}

func (s *SubsriptionService) ListSubscriptionService(ctx context.Context) ([]subscriptionservice.Subscription, error){
	
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "ListSubscriptionService",
	}).Debug()

	data, err := s.repos.ListSubscriptionRepository(ctx)
	if err != nil{
		return nil, err
	}
	response := make([]subscriptionservice.Subscription, 0, len(data))
	for _, point := range data{
		startDate := point.StartDate.Format("01-2006")

		response = append(response, subscriptionservice.Subscription{
			Id: point.Id,
			ServiceName: point.ServiceName,
			Price: point.Price,
			UserId: point.UserId,
			StartDate: startDate,
	})
	}

	return response, nil
}


func (s *SubsriptionService) DeleteSubscriptionService(ctx context.Context, subId uuid.UUID) error{
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "GetSubscriptionService",
	}).Debug()
	return s.repos.DeleteSubscriptionRepository(ctx, subId)

}

func (s *SubsriptionService) UpdateSubscriptionService(ctx context.Context, request subscriptionservice.UpdateSubscription) error{
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "UpdateSubscriptionService",
	}).Debug()

	return s.repos.UpdateSubscriptionRepository(ctx, request)
}

func (s *SubsriptionService) TotalPriceService(ctx context.Context, filter subscriptionservice.FilterSubscription) (int64, error){
	reqId := ctx.Value("req_id")
	if reqId == ""{
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId, 
		"method": "TotalPriceService",
	}).Debug()

	return s.repos.TotalPriceRepository(ctx, filter)
}
