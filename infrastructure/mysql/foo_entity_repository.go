package persistence

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hobord/golang-poc-rest/domain/entity"
	"github.com/hobord/golang-poc-rest/domain/repository"
)

// FooRepository Implements repository.FooRepository
type FooRepository struct {
	conn *sql.DB
}

// NewFooMysqlRepository returns initialized FooRepositoryImpl
func NewFooMysqlRepository(conn *sql.DB) repository.FooRepository {
	return &FooRepository{conn: conn}
}

func (r *FooRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func (r *FooRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *FooRepository) GetByID(ctx context.Context, id string) (*entity.Foo, error) {
	row, err := r.queryRow(ctx, "SELECT id, title FROM foo WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	entity := &entity.Foo{}
	err = row.Scan(&entity.ID, &entity.Title)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *FooRepository) GetAll(ctx context.Context) ([]*entity.Foo, error) {
	rows, err := r.query(ctx, "SELECT id, title FROM foo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*entity.Foo, 0)

	for rows.Next() {
		entity := &entity.Foo{}
		err = rows.Scan(&entity.ID, &entity.Title)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}

	return results, nil
}

func (r *FooRepository) Save(ctx context.Context, entity *entity.Foo) error {
	var query string
	isExists, err := r.GetByID(ctx, entity.ID)
	if isExists == nil {
		query = "INSERT INTO foo (id, title) VALUES (?, ?)"
	} else {
		query = "UPDATE foo SET title = ? WHERE ID = ?"
	}
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if isExists == nil {
		_, err = stmt.ExecContext(ctx, entity.ID, entity.Title)
	} else {
		_, err = stmt.ExecContext(ctx, entity.Title, entity.ID)
	}
	return err
}

func (r *FooRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.conn.Prepare("DELETE FROM entity WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
