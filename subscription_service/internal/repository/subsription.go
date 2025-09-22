package repository

import (
	"context"
	subscriptionservice "effective_mobile"
	"fmt"
	"strings"
	"time"

	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) CreateSubscriptionRepository(ctx context.Context, request subscriptionservice.CreateSubscription) (uuid.UUID, error) {
	var id uuid.UUID

	query := "INSERT INTO subscription (service_name, price, user_id, start_date) VALUES ($1, $2, $3, $4) RETURNING id;"

	startDate, err := time.Parse("01-2006", request.StartDate)
	if err != nil {
		return uuid.Nil, err
	}
	row := r.db.QueryRowContext(ctx, query, request.ServiceName, request.Price, request.UserId, startDate)

	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *SubscriptionRepository) GetSubscriptionRepository(ctx context.Context, subId uuid.UUID) (subscriptionservice.PreparationSubscription, error) {
	var response subscriptionservice.PreparationSubscription
	query := "SELECT id, service_name, price, user_id, start_date FROM subscription WHERE id = $1;"

	if err := r.db.GetContext(ctx, &response, query, subId); err != nil {
		if err == sql.ErrNoRows {
			return subscriptionservice.PreparationSubscription{}, nil
		}
		return subscriptionservice.PreparationSubscription{}, err
	}
	return response, nil

}

func (r *SubscriptionRepository) ListSubscriptionRepository(ctx context.Context) ([]subscriptionservice.PreparationSubscription, error) {
	var response []subscriptionservice.PreparationSubscription
	query := "SELECT id, service_name, price, user_id, start_date FROM subscription;"

	if err := r.db.SelectContext(ctx, &response, query); err != nil {
		return nil, err
	}
	return response, nil

}

func (r *SubscriptionRepository) DeleteSubscriptionRepository(ctx context.Context, subId uuid.UUID) error {
	query := "DELETE FROM subscription WHERE id = $1;"
	_, err := r.db.ExecContext(ctx, query, subId)
	return err

}

func (r *SubscriptionRepository) UpdateSubscriptionRepository(ctx context.Context, request subscriptionservice.UpdateSubscription) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if request.ServiceName.Valid {
		setValue = append(setValue, fmt.Sprintf("service_name=$%d", argId))
		args = append(args, request.ServiceName.String)
		argId++
	}
	if request.Price.Valid {
		setValue = append(setValue, fmt.Sprintf("price=$%d", argId))
		args = append(args, request.Price.Int64)
		argId++
	}
	if request.UserId.Valid {
		setValue = append(setValue, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, request.UserId.String)
		argId++
	}
	if request.StartDate.Valid {
		setValue = append(setValue, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, request.StartDate.String)
		argId++
	}
	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE subscription SET %s WHERE id=$%d;", setQuery, argId)
	args = append(args, request.Id)

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *SubscriptionRepository) TotalPriceRepository(ctx context.Context, filter subscriptionservice.FilterSubscription) (int64, error) {
	var total null.Int
	reqId := ctx.Value("req_id")
	if reqId == "" {
		reqId = "none"
	}

	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	startDateFrom, err := time.Parse("01-2006", filter.StartDateFrom.String)
	if err != nil {
		return 0, err
	}
	startDateTo, err := time.Parse("01-2006", filter.StartDateTo.String)
	if err != nil {
		return 0, err
	}

	setValue = append(setValue, fmt.Sprintf("start_date BETWEEN $%d AND $%d", argId, argId+1))
	args = append(args, startDateFrom, startDateTo)
	argId += 2

	if filter.UserId.Valid {
		setValue = append(setValue, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, filter.UserId.String)
		argId++
	}

	if filter.ServiceName.Valid {
		setValue = append(setValue, fmt.Sprintf("service_name=$%d", argId))
		args = append(args, filter.ServiceName.String)
		argId++
	}

	setQuery := strings.Join(setValue, " AND ")
	query := fmt.Sprintf("SELECT SUM(price) AS total FROM subscription WHERE %s;", setQuery)

	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "TotalPriceRepository",
		"sql":    query,
	}).Debug()

	if err := r.db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, err
	}

	return total.Int64, nil

}
