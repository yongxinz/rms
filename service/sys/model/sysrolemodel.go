package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		FindAll(context.Context, int64, int64) ([]*SysRole, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn, c),
	}
}

func (m *customSysRoleModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysRole, error) {
	var resp []*SysRole
	query := "select sys_role.* from sys_role limit ?,?"
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

func (m *customSysRoleModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(role_id) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}

func (m *customSysRoleModel) DeleteMulti(ctx context.Context, roleIds []int64) error {
	args := make([]interface{}, len(roleIds))
	for i, id := range roleIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where role_id in (?`+strings.Repeat(",?", len(roleIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
