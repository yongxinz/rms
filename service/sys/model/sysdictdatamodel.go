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

var _ SysDictDataModel = (*customSysDictDataModel)(nil)

type (
	// SysDictDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictDataModel.
	SysDictDataModel interface {
		sysDictDataModel
		FindByDictType(context.Context, string) ([]*SysDictData, error)
		FindAll(context.Context, string, int64, int64) ([]*SysDictData, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
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

func (m *customSysDictDataModel) FindAll(ctx context.Context, DictType string, Current int64, PageSize int64) ([]*SysDictData, error) {
	var resp []*SysDictData
	query := "select sys_dict_data.* from sys_dict_data where dict_type = ? limit ?,?"
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, DictType, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysDictDataModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(dict_code) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}

func (m *customSysDictDataModel) DeleteMulti(ctx context.Context, dictIds []int64) error {
	args := make([]interface{}, len(dictIds))
	for i, id := range dictIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where dict_code in (?`+strings.Repeat(",?", len(dictIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
