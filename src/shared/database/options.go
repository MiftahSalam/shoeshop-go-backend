package database

import "time"

type Option struct {
	Username           string        `json:"username" mapstructure:"username"`
	Password           string        `json:"password" mapstructure:"password"`
	Name               string        `json:"name" mapstructure:"name"`
	Schema             string        `json:"schema" mapstructure:"schema"`
	Host               string        `json:"host" mapstructure:"host"`
	Port               int           `json:"port" mapstructure:"port"`
	MinIdleConnections int           `json:"min_idle_connections" mapstructure:"min_idle_connections"`
	MaxOpenConnections int           `json:"max_open_connections" mapstructure:"max_open_connections"`
	ConnMaxLifetime    time.Duration `json:"conn_max_lifetime" json:"conn_max_lifetime"`
	DebugMode          bool          `json:"debug_mode" mapstructure:"debug_mode"`
}
