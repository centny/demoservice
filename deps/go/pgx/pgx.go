package pgx

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func C() *Pool {
	return Shared
}

var Shared *Pool

func Bootstrap(connString string) (err error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err == nil {
		Shared = (*Pool)(pool)
	}
	return
}

func Exec(sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return Shared.Exec(sql, args...)
}

func QueryRow(sql string, args ...interface{}) *Row {
	return Shared.QueryRow(sql, args...)
}

func Query(sql string, args ...interface{}) (rows *Rows, err error) {
	return Shared.Query(sql, args...)
}

func CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return Shared.CopyFrom(tableName, columnNames, rowSrc)
}

func SendBatch(b *pgx.Batch) *BatchResults {
	return Shared.SendBatch(b)
}

func Begin() (tx *Tx, err error) {
	return Shared.Begin()
}

var ErrNoRows = pgx.ErrNoRows
var ErrTxClosed = pgx.ErrTxClosed
var ErrTxCommitRollback = pgx.ErrTxCommitRollback

type Row struct {
	SQL string
	pgx.Row
}

func (r Row) Scan(dest ...interface{}) (err error) {
	defer func() {
		xerr := r.Row.Scan(dest...)
		if err == nil {
			err = xerr
		}
	}()
	err = mockerCheck("Row.Scan", r.SQL)
	return
}

type Rows struct {
	SQL string
	pgx.Rows
}

func (r *Rows) Scan(dest ...interface{}) error {
	if err := mockerCheck("Rows.Scan", r.SQL); err != nil {
		return err
	}
	return r.Rows.Scan(dest...)
}

func (r *Rows) Values() ([]interface{}, error) {
	if err := mockerCheck("Rows.Values", r.SQL); err != nil {
		return nil, err
	}
	return r.Rows.Values()
}

type BatchResults struct {
	pgx.BatchResults
}

func (b *BatchResults) Exec() (pgconn.CommandTag, error) {
	if err := mockerCheck("BatchResult.Exec", ""); err != nil {
		return nil, err
	}
	return b.BatchResults.Exec()
}

func (b *BatchResults) Query() (rows *Rows, err error) {
	if err := mockerCheck("BatchResult.Query", ""); err != nil {
		return nil, err
	}
	raw, err := b.BatchResults.Query()
	if err == nil {
		rows = &Rows{Rows: raw}
	}
	return
}

func (b *BatchResults) QueryRow() *Row {
	return &Row{Row: b.BatchResults.QueryRow()}
}

func (b *BatchResults) Close() error {
	if err := mockerCheck("BatchResult.Close", ""); err != nil {
		return err
	}
	return b.BatchResults.Close()
}

type Tx struct {
	pgx.Tx
}

// Begin starts a pseudo nested transaction.
func (t *Tx) Begin() (tx *Tx, err error) {
	if err := mockerCheck("Tx.Begin", ""); err != nil {
		return nil, err
	}
	raw, err := t.Tx.Begin(context.Background())
	if err == nil {
		tx = &Tx{Tx: raw}
	}
	return
}

func (t *Tx) Commit() error {
	if err := mockerCheck("Tx.Commit", ""); err != nil {
		t.Tx.Rollback(context.Background())
		return err
	}
	return t.Tx.Commit(context.Background())
}

func (t *Tx) Rollback() error {
	if err := mockerCheck("Tx.Rollback", ""); err != nil {
		t.Tx.Rollback(context.Background())
		return err
	}
	return t.Tx.Rollback(context.Background())
}

func (t *Tx) CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	if err := mockerCheck("Tx.CopyFrom", ""); err != nil {
		return 0, err
	}
	return t.Tx.CopyFrom(context.Background(), tableName, columnNames, rowSrc)
}

func (t *Tx) SendBatch(b *pgx.Batch) *BatchResults {
	return &BatchResults{
		BatchResults: t.Tx.SendBatch(context.Background(), b),
	}
}

func (t *Tx) Prepare(name, sql string) (*pgconn.StatementDescription, error) {
	if err := mockerCheck("Tx.Prepare", sql); err != nil {
		return nil, err
	}
	return t.Tx.Prepare(context.Background(), name, sql)
}

func (t *Tx) Exec(sql string, args ...interface{}) (commandTag pgconn.CommandTag, err error) {
	if err := mockerCheck("Tx.Exec", sql); err != nil {
		return nil, err
	}
	return t.Tx.Exec(context.Background(), sql, args...)
}

func (t *Tx) ExecRow(sql string, args ...interface{}) (err error) {
	res, err := t.Exec(sql, args...)
	if err == nil && res.RowsAffected() < 1 {
		err = pgx.ErrNoRows
	}
	return
}

func (t *Tx) Query(sql string, args ...interface{}) (rows *Rows, err error) {
	if err := mockerCheck("Tx.Query", sql); err != nil {
		return nil, err
	}
	raw, err := t.Tx.Query(context.Background(), sql, args...)
	if err == nil {
		rows = &Rows{SQL: sql, Rows: raw}
	}
	return
}
func (t *Tx) QueryRow(sql string, args ...interface{}) *Row {
	return &Row{
		SQL: sql,
		Row: t.Tx.QueryRow(context.Background(), sql, args...),
	}
}

type Pool pgxpool.Pool

func (p *Pool) Exec(sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if err := mockerCheck("Pool.Exec", sql); err != nil {
		return nil, err
	}
	pool := (*pgxpool.Pool)(p)
	return pool.Exec(context.Background(), sql, args...)
}

func (p *Pool) ExecRow(sql string, args ...interface{}) (err error) {
	if err := mockerCheck("Pool.ExecRow", sql); err != nil {
		return err
	}
	res, err := p.Exec(sql, args...)
	if err == nil && res.RowsAffected() < 1 {
		err = pgx.ErrNoRows
	}
	return
}

func (p *Pool) QueryRow(sql string, args ...interface{}) *Row {
	pool := (*pgxpool.Pool)(p)
	return &Row{
		SQL: sql,
		Row: pool.QueryRow(context.Background(), sql, args...),
	}
}

func (p *Pool) Query(sql string, args ...interface{}) (rows *Rows, err error) {
	if err := mockerCheck("Pool.Query", sql); err != nil {
		return nil, err
	}
	pool := (*pgxpool.Pool)(p)
	raw, err := pool.Query(context.Background(), sql, args...)
	if err == nil {
		rows = &Rows{SQL: sql, Rows: raw}
	}
	return
}

func (p *Pool) CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	if err := mockerCheck("Pool.CopyFrom", ""); err != nil {
		return 0, err
	}
	pool := (*pgxpool.Pool)(p)
	return pool.CopyFrom(context.Background(), tableName, columnNames, rowSrc)
}

func (p *Pool) SendBatch(b *pgx.Batch) *BatchResults {
	pool := (*pgxpool.Pool)(p)
	return &BatchResults{
		BatchResults: pool.SendBatch(context.Background(), b),
	}
}

func (p *Pool) Begin() (tx *Tx, err error) {
	if err := mockerCheck("Pool.Begin", ""); err != nil {
		return nil, err
	}
	pool := (*pgxpool.Pool)(p)
	raw, err := pool.Begin(context.Background())
	if err == nil {
		tx = &Tx{Tx: raw}
	}
	return
}

type PoolF func() *Pool
type QueryF func(sql string, args ...interface{}) (rows *Rows, err error)
type QueryRowF func(sql string, args ...interface{}) *Row
type ExecF func(sql string, args ...interface{}) (pgconn.CommandTag, error)

//ExecRowTx will exec and check result must be at least update one row
func ExecRowTx(tx *Tx, sql string, args ...interface{}) (err error) {
	res, err := tx.Exec(sql, args...)
	if err == nil && res.RowsAffected() < 1 {
		err = pgx.ErrNoRows
	}
	return
}

//ExecRow will exec and check result must be at least update one row
func ExecRow(pool *Pool, sql string, args ...interface{}) (err error) {
	res, err := pool.Exec(sql, args...)
	if err == nil && res.RowsAffected() < 1 {
		err = pgx.ErrNoRows
	}
	return
}

type Caller interface {
	Exec(sql string, args ...interface{}) (pgconn.CommandTag, error)
	ExecRow(sql string, args ...interface{}) (err error)
	QueryRow(sql string, args ...interface{}) *Row
	Query(sql string, args ...interface{}) (rows *Rows, err error)
}
