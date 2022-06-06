// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userId int64) error
		FindOneByName(ctx context.Context, userName string) (*User, error)
		FindOneByNameAndPassword(ctx context.Context, userName, password string) (*User, error)
		TransUpdateFollow(ctx context.Context, session sqlx.Session, userId int64, isFollow bool) error
		TransUpdateFollower(ctx context.Context, session sqlx.Session, userId int64, isFollow bool) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		UserId        int64  `db:"user_id"`
		UserName      string `db:"user_name"`
		Password      string `db:"password"`
		FollowCount   int64  `db:"follow_count"`
		FollowerCount int64  `db:"follower_count"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Password, data.FollowCount, data.FollowerCount)
	return ret, err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userId int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, userName string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByNameAndPassword(ctx context.Context, userName, password string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_name` = ? and `password` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, userName, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Password, data.FollowCount, data.FollowerCount, data.UserId)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserModel) TransUpdateFollow(ctx context.Context, session sqlx.Session, userId int64, isFollow bool) error {
	query := fmt.Sprintf("update %s set follow_count = follow_count + 1 where `user_id` = ?", m.table)
	if !isFollow {
		query = fmt.Sprintf("update %s set follow_count = follow_count - 1 where `user_id` = ?", m.table)
	}
	_, err := session.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserModel) TransUpdateFollower(ctx context.Context, session sqlx.Session, userId int64, isFollow bool) error {
	query := fmt.Sprintf("update %s set follower_count = follower_count + 1 where `user_id` = ?", m.table)
	if !isFollow {
		query = fmt.Sprintf("update %s set follower_count = follower_count - 1 where `user_id` = ?", m.table)
	}
	_, err := session.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
