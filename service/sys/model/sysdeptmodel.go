package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDeptModel = (*customSysDeptModel)(nil)

type (
	// SysDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDeptModel.
	SysDeptModel interface {
		sysDeptModel
		FindAll(context.Context, int64, int64) ([]*SysDept, error)
		Count(context.Context) (int64, error)
	}

	customSysDeptModel struct {
		*defaultSysDeptModel
	}
)

// NewSysDeptModel returns a model for the database table.
func NewSysDeptModel(conn sqlx.SqlConn, c cache.CacheConf) SysDeptModel {
	return &customSysDeptModel{
		defaultSysDeptModel: newSysDeptModel(conn, c),
	}
}

func (m *customSysDeptModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysDept, error) {
	var resp []*SysDept
	query := fmt.Sprintf("select %s from %s order by dept_id limit ?,?", sysDeptRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysDeptModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(dept_id) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}
