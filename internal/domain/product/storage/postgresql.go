package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Striker87/notes/internal/domain/product/model"
	"github.com/Striker87/notes/pkg/client/postgresql"
	db "github.com/Striker87/notes/pkg/client/postgresql/model"
	"github.com/Striker87/notes/pkg/logging"
)

const (
	scheme = "public"
	table  = "product"
)

type ProductStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}

func NewProductStorage(client PostgreSQLClient, logger *logging.Logger) ProductStorage {
	return ProductStorage{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
		logger:       logger,
	}
}

func (s *ProductStorage) queryLogger(sql string, table string, args []any) *logging.Logger {
	return s.logger.ExtraFields(map[string]any{
		"sql":   sql,
		"table": table,
		"args":  args,
	})
}

func (s *ProductStorage) All(ctx context.Context) ([]model.Product, error) {
	sql, args, err := s.queryBuilder.Select("id").
		Column("name").
		Column("description").
		Column("image_id").
		Column("price").
		Column("currency_id").
		Column("rating").
		Column("created_at").
		Column("updated_at").
		From(scheme + "." + table).ToSql()

	// todo filtering and sort

	logger := s.queryLogger(sql, table, args)
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err)
		return nil, err
	}

	logger.Trace("do query")
	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	list := make([]model.Product, 0)

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.ImageID, &p.Price, &p.CurrencyID, &p.Rating, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			err = db.ErrScan(postgresql.ParsePgError(err))
			logger.Error(err)
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}
