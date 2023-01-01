package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDictDataModel = (*customSysDictDataModel)(nil)

type (
	// SysDictDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictDataModel.
	SysDictDataModel interface {
		sysDictDataModel
		FindByDictType(context.Context, string) ([]*SysDictData, error)
	}

	customSysDictDataModel struct {
		*defaultSysDictDataModel
	}
)

// NewSysDictDataModel returns a model for the database table.
func NewSysDictDataModel(conn sqlx.SqlConn, c cache.CacheConf) SysDictDataModel {
	return &customSysDictDataModel{
		defaultSysDictDataModel: newSysDictDataModel(conn, c),
	}
}

func (m *customSysDictDataModel) FindByDictType(ctx context.Context, dictType string) ([]*SysDictData, error) {
	var resp []*SysDictData
	query := fmt.Sprintf("select %s from %s where dict_type = ?", sysDictDataRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, dictType)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
