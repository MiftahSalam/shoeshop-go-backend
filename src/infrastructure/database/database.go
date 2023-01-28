package database

import (
	"context"

	"gorm.io/gorm"
)

// common implementation related to database goes here

const (
	MySQL      string = "mysql"
	SQLite     string = "sqlite"
	PostgreSQL string = "postgresql"
)

type ORM interface {
	Error() error
	Close() error
	Begin() ORM
	Commit() error
	Rollback() error
	Offset(offset int64) ORM
	Limit(limit int64) ORM
	First(object interface{}) error
	Last(object interface{}) error
	Find(object interface{}) error
	Model(value interface{}) ORM
	Select(query interface{}, args ...interface{}) ORM
	OmitAssoc() ORM
	Where(query interface{}, args ...interface{}) ORM
	Order(value interface{}) ORM
	Create(args interface{}) error
	Update(args interface{}) error
	UpdateColumns(args interface{}) error
	Delete(model interface{}, args ...interface{}) error
	WithContext(ctx context.Context) ORM
	Raw(query string, args ...interface{}) ORM
	Exec(query string, args ...interface{}) ORM
	Scan(object interface{}) error
	Preload(assoc string) ORM
	Joins(assoc string) ORM
	GetGormInstance() *gorm.DB
	Count(count *int64) error
	Association(column string) ORMAssociation
	Or(query interface{}, args ...interface{}) ORM
	Save(data interface{}) error
	Migrate(data interface{}) error
}

type ORMAssociation interface {
	Replace(values ...interface{}) error
	Find(out interface{}, conds ...interface{}) error
	Clear() error
}
