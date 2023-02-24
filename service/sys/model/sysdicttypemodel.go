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

var _ SysDictTypeModel = (*customSysDictTypeModel)(nil)

type (
	// SysDictTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictTypeModel.
	SysDictTypeModel interface {
		sysDictTypeModel
		FindAll(context.Context, int64, int64) ([]*SysDictType, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
	}

	customSysDictTypeModel struct {
		*defaultSysDictTypeModel
	}
)

// NewSysDictTypeModel returns a model for the database table.
func NewSysDictTypeModel(conn sqlx.SqlConn, c cache.CacheConf) SysDictTypeModel {
	return &customSysDictTypeModel{
		defaultSysDictTypeModel: newSysDictTypeModel(conn, c),
	}
}

func (m *customSysDictTypeModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysDictType, error) {
	var resp []*SysDictType
	query := "select sys_dict_type.* from sys_dict_type limit ?,?"
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

func (m *customSysDictTypeModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(dict_id) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}

func (m *customSysDictTypeModel) DeleteMulti(ctx context.Context, dictIds []int64) error {
	args := make([]interface{}, len(dictIds))
	for i, id := range dictIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where dict_id in (?`+strings.Repeat(",?", len(dictIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
