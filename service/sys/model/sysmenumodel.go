package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		FindMenuList(context.Context, []int64) ([]*SysMenu, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn, c),
	}
}

func (m *customSysMenuModel) FindMenuList(ctx context.Context, MenuIds []int64) ([]*SysMenu, error) {
	args := make([]interface{}, len(MenuIds))
	for i, id := range MenuIds {
		args[i] = id
	}

	var resp []*SysMenu
	query := fmt.Sprintf(`select %s from %s where menu_id in (?`+strings.Repeat(",?", len(MenuIds)-1)+`)`, sysMenuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
