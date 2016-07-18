package gorepo

import (
	"database/sql"
	"reflect"

	"github.com/vardius/query"
)

type mysqlRepository struct {
	db      *sql.DB
	builder query.Builder
}

var NewSQL = newRepository

func (repo *mysqlRepository) GetAll() (interface{}, error) {
	return repo.builder.Reset().Select().GetQuery().GetResults(repo.db)
}

func (repo *mysqlRepository) Get(id int64) (interface{}, error) {
	return repo.builder.Reset().Select().Where("id = ?").SetParameters(id).GetQuery().GetResult(repo.db)
}

func (repo *mysqlRepository) Save(v interface{}) error {
	_, err := repo.builder.Reset().Save(v).GetQuery().Execute(repo.db)
	return err
}

func (repo *mysqlRepository) Remove(ids ...int64) (interface{}, error) {
	return repo.builder.Reset().Delete().Where("id IN (?)").SetParameters(ids).GetQuery().Execute(repo.db)
}

func newRepository(db *sql.DB, t reflect.Type) Repository {
	return &mysqlRepository{db, query.New(t)}
}
