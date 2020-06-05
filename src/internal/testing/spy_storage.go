package testing

import (
	"context"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/storage"
)

type SpyStorage struct {
	querier storage.Querier
}

func (s *SpyStorage) Querier(ctx context.Context, mint int64, maxt int64) (storage.Querier, error) {
	return s.querier, nil
}

func (s *SpyStorage) StartTime() (int64, error) {
	panic("not implemented")
}

func (s *SpyStorage) Appender() (storage.Appender, error) {
	//	panic("not implemented")
	return &SpyAppender{}, nil
}

func (s *SpyStorage) Close() error {
	//	panic("not implemented")
	return nil
}

func NewSpyStorage(querier storage.Querier) *SpyStorage {
	return &SpyStorage{
		querier: querier,
	}
}

type SpyAppender struct {
}

func (s SpyAppender) Add(l labels.Labels, t int64, v float64) (uint64, error) {
	//	panic("implement me")
	return 0, nil
}

func (s SpyAppender) AddFast(l labels.Labels, ref uint64, t int64, v float64) error {
	panic("implement me")
}

func (s SpyAppender) Commit() error {
	// panic("implement me")
	return nil
}

func (s SpyAppender) Rollback() error {
	panic("implement me")
}
