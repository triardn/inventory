package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/triardn/inventory/driver"
)

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
}

type Cache struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	DialConnectTimeout time.Duration `mapstructure:"dial_connect_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	MaxIdle            int           `mapstructure:"max_idle"`
	MaxActive          int           `mapstructure:"max_active"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout"`
	Wait               bool          `mapstructure:"wait"`
	MaxConnLifetime    time.Duration `mapstructure:"max_conn_lifetime"`
	Namespace          string        `mapstructure:"namespace"`
	Password           string        `mapstructure:"password"`
	Items              Items         `mapstructure:"items"`
}

type CacheItem struct {
	Prefix string        `mapstructure:"prefix" validate:"required"`
	TTL    time.Duration `mapstructure:"ttl" validate:"required"`
}

type Items struct {
	GetOneCampaign CacheItem `mapstructure:"get_one_campaign"`
	GetOneUser     CacheItem `mapstructure:"get_one_user"`
}

type DatabaseSqlite struct {
	Path string `mapstructure:"path_to_db"`
}

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

type Config struct {
	App      App
	Database DatabaseSqlite
	Cache    Cache
}

var defaultConfig *viper.Viper

func NewAppConfig() (conf *Config, err error) {
	err = viper.Unmarshal(&conf)
	return
}

func (c *Config) GetCacheOption() driver.CacheOption {
	return driver.CacheOption{
		Host:               c.Cache.Host,
		Port:               c.Cache.Port,
		DialConnectTimeout: c.Cache.DialConnectTimeout,
		WriteTimeout:       c.Cache.WriteTimeout,
		ReadTimeout:        c.Cache.ReadTimeout,
		Namespace:          c.Cache.Namespace,
		Password:           c.Cache.Password,
		MaxIdle:            c.Cache.MaxIdle,
		MaxActive:          c.Cache.MaxActive,
		IdleTimeout:        c.Cache.IdleTimeout,
		Wait:               c.Cache.Wait,
		MaxConnLifetime:    c.Cache.MaxConnLifetime,
	}
}

func (c *Config) GetPathToDBSqlite() string {
	return c.Database.Path
}
