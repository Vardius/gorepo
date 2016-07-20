package gorepo

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/vardius/goquery"
)

type mysqlRepository struct {
	db      *sql.DB
	builder goquery.Builder
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
	length := len(ids)
	if length < 1 {
		return nil, errors.New("sql gorepo: not enought arguments")
	}
	s := make([]interface{}, length)
	marks := "?"
	for i, v := range ids {
		if i > 0 {
			marks += ",?"
		}
		s[i] = v
	}
	return repo.builder.Reset().Delete().Where("id IN (" + marks + ")").SetParameters(s...).GetQuery().Execute(repo.db)
}

func newRepository(db *sql.DB, t reflect.Type) Repository {
	return &mysqlRepository{db, goquery.New(t)}
}
