package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysConfigModel = (*customSysConfigModel)(nil)

type (
	// SysConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysConfigModel.
	SysConfigModel interface {
		sysConfigModel
		FindAll(ctx context.Context) ([]*SysConfig, error)
	}

	customSysConfigModel struct {
		*defaultSysConfigModel
	}
)

// NewSysConfigModel returns a model for the database table.
func NewSysConfigModel(conn sqlx.SqlConn, c cache.CacheConf) SysConfigModel {
	return &customSysConfigModel{
		defaultSysConfigModel: newSysConfigModel(conn, c),
	}
}

func (m *customSysConfigModel) FindAll(ctx context.Context) ([]*SysConfig, error) {
	var resp []*SysConfig
	query := fmt.Sprintf("select %s from %s", sysConfigRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
