package database

import (
	"fmt"

	db "shoeshop-backend/src/infrastructure/database"
	"shoeshop-backend/src/infrastructure/logger"
)

type Database struct {
	db.ORM
}

func Setup(config Option, logger *logger.Logger) Database {
	fmt.Println("Try Setup Database ...")

	if logger == nil {
		panic("DB setup failed. logger is nil")
	}

	var (
		dbConnectionString = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable search_path=%s",
			config.Username,
			config.Password,
			config.Name,
			config.Host,
			config.Port,
			config.Schema)

		opts = &db.PostgreSqlOption{
			ConnectionString:      dbConnectionString,
			MaxOpenConnection:     config.MaxOpenConnections,
			MaxIdleConnection:     config.MinIdleConnections,
			MaxLifeTimeConnection: config.ConnMaxLifetime,
		}
	)

	if config.DebugMode {
		opts.Logger = *logger
	}

	db, err := db.NewPostgreSql(opts)
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}

	return Database{db}
}
