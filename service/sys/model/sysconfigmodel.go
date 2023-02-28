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

var _ SysConfigModel = (*customSysConfigModel)(nil)

type (
	// SysConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysConfigModel.
	SysConfigModel interface {
		sysConfigModel
		FindAll(context.Context) ([]*SysConfig, error)
		FindByKey(context.Context, string) ([]*SysConfig, error)
		FindAll1(context.Context, int64, int64) ([]*SysConfig, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
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

func (m *customSysConfigModel) FindByKey(ctx context.Context, configKey string) ([]*SysConfig, error) {
	var resp []*SysConfig
	query := fmt.Sprintf("select %s from %s where config_key = ?", sysConfigRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, configKey)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysConfigModel) FindAll1(ctx context.Context, Current int64, PageSize int64) ([]*SysConfig, error) {
	var resp []*SysConfig
	query := "select sys_config.* from sys_config limit ?,?"
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

func (m *customSysConfigModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(id) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}

func (m *customSysConfigModel) DeleteMulti(ctx context.Context, dictIds []int64) error {
	args := make([]interface{}, len(dictIds))
	for i, id := range dictIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where id in (?`+strings.Repeat(",?", len(dictIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
