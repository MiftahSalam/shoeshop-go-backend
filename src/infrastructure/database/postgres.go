package database

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "shoeshop-backend/src/infrastructure/logger"
)

type (
	postgresqldb struct {
		db  *gorm.DB
		err error
	}

	PostgreSqlOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
		Logger                               log.Logger
	}
)

func (d *postgresqldb) Error() error {
	return d.err
}

func (d *postgresqldb) Close() error {
	sql, err := d.db.DB()

	if err != nil {
		return err
	}

	if err := sql.Close(); err != nil {
		return err
	}
	return nil
}

func (d *postgresqldb) Begin() ORM {
	var (
		db  = d.db.Begin()
		err = db.Error
	)
	return &postgresqldb{db, err}
}

func (d *postgresqldb) Migrate(object interface{}) error {
	return d.db.AutoMigrate(object)
}

func (d *postgresqldb) Commit() error {
	return d.db.Commit().Error
}

func (d *postgresqldb) Rollback() error {
	return d.db.Rollback().Error
}

func (d *postgresqldb) Offset(offset int64) ORM {
	var (
		db  = d.db.Offset(int(offset))
		err = d.db.Error
	)
	return &postgresqldb{db, err}
}

func (d *postgresqldb) Limit(limit int64) ORM {
	var (
		db  = d.db.Limit(int(limit))
		err = d.db.Error
	)
	return &postgresqldb{db, err}
}

func (d *postgresqldb) First(object interface{}) error {
	var (
		res = d.db.First(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *postgresqldb) Last(object interface{}) error {
	var (
		res = d.db.Last(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *postgresqldb) Find(object interface{}) error {
	var (
		res = d.db.Find(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *postgresqldb) Model(value interface{}) ORM {
	var (
		db  = d.db.Model(value)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) OmitAssoc() ORM {
	var (
		db  = d.db.Omit(clause.Associations)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Select(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Select(query, args...)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Where(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Where(query, args...)
		err = db.Error
	)
	return &postgresqldb{db, err}
}

func (d *postgresqldb) Order(value interface{}) ORM {
	var (
		db  = d.db.Order(value)
		err = d.db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Create(args interface{}) error {
	return d.db.Create(args).Error
}

func (d *postgresqldb) Update(args interface{}) error {
	return d.db.Updates(args).Error
}

func (d *postgresqldb) UpdateColumns(args interface{}) error {
	return d.db.UpdateColumns(args).Error
}

func (d *postgresqldb) Delete(model interface{}, args ...interface{}) error {
	return d.db.Delete(model, args...).Error
}

func (d *postgresqldb) WithContext(ctx context.Context) ORM {
	var (
		db = d.db.WithContext(ctx)
	)

	return &postgresqldb{db: db, err: nil}
}

func (d *postgresqldb) Raw(query string, args ...interface{}) ORM {
	var (
		db  = d.db.Raw(query, args...)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Exec(query string, args ...interface{}) ORM {
	var (
		db  = d.db.Exec(query, args...)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Scan(object interface{}) error {
	var (
		db = d.db.Scan(object)
	)

	return db.Error
}

func (d *postgresqldb) Preload(assoc string) ORM {
	var (
		db  = d.db.Preload(assoc)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Joins(assoc string) ORM {
	var (
		db  = d.db.Joins(assoc)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) GetGormInstance() *gorm.DB {
	return d.db
}

func (d *postgresqldb) Count(count *int64) error {
	var (
		res = d.db.Count(count)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *postgresqldb) Association(column string) ORMAssociation {
	return d.db.Association(column)
}

func (d *postgresqldb) Or(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Or(query, args...)
		err = db.Error
	)

	return &postgresqldb{db, err}
}

func (d *postgresqldb) Save(data interface{}) error {
	var (
		db  = d.db.Save(data)
		err = db.Error
	)

	return err
}

func NewPostgreSql(option *PostgreSqlOption) (ORM, error) {
	var (
		opts = &gorm.Config{}
	)

	if option.Logger != nil {
		opts.Logger = logger.New(option.Logger, logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
		})
	}

	db, err := gorm.Open(postgres.Open(option.ConnectionString), opts)

	if err != nil {
		return nil, err
	}

	sql, err := db.DB()

	if err != nil {
		return nil, err
	}

	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)

	return &postgresqldb{db: db}, nil
}
