package repository

import (
	"context"

	"github.com/faridlan/go-mongo-mock/model/domain"
)

type InsertRepository interface {
	Insert(ctx context.Context, rekapData domain.Rekap) (*domain.Rekap, error)
}
