package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindAll(context.Context, int64, int64) ([]*SysUser, error)
		FindAll1(context.Context, int64, int64) ([]*SysUserList, error)
		Count(context.Context) (int64, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}

	SysUserList struct {
		SysUser
		SysDept
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}

func (m *defaultSysUserModel) FindAll1(ctx context.Context, Current int64, PageSize int64) ([]*SysUserList, error) {
	var resp []*SysUserList
	query := "select sys_user.*, sd.* from sys_user left join sys_dept sd on sys_user.dept_id = sd.dept_id limit ?,?"
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

func (m *customSysUserModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysUser, error) {
	var resp []*SysUser
	query := "select sys_user.*, sys_dept.* from sys_user left join sys_dept sd on sys_user.dept_id = sd.dept_id limit ?,?"
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

func (m *customSysUserModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(user_id) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}
