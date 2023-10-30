package crm_core

import (
	"github.com/spf13/viper"
	"log"
)

type (
	// Configuration -.
	Configuration struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		Gin  `yaml:"gin"`
		DB   `yaml:"db"`
		//Auth `yaml:"auth"`
		Jwt `yaml:"jwt"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port                   string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		DefaultReadTimeout     int64  `env-required:"true" yaml:"default_read_timeout" env:"DEFAULT_READ_TIMEOUT"`
		DefaultWriteTimeout    int64  `env-required:"true" yaml:"default_write_timeout" env:"DEFAULT_WRITE_TIMEOUT"`
		DefaultShutdownTimeout int64  `env-required:"true" yaml:"default_shutdown_timeout" env:"DEFAULT_SHUTDOWN_TIMEOUT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Gin struct {
		Mode string `env-required:"true" yaml:"mode" env:"GIN_MODE"`
	}

	DB struct {
		PoolMax  int64  `env-required:"true" yaml:"pool_max" env:"DB_POOL_MAX"`
		Host     string `env-required:"true" yaml:"host" env:"DB_HOST"`
		User     string `env-required:"true" yaml:"user" env:"DB_USER"`
		Password string `env-required:"true" yaml:"password" env:"DB_PASSWORD"`
		Name     string `env-required:"true" yaml:"name" env:"DB_NAME"`
		Port     int64  `env-required:"true" yaml:"port" env:"DB_PORT"`
	}

	//Auth struct {
	//	Login    string `mapstructure:"login"`
	//	Password string `mapstructure:"pass"`
	//}

	Jwt struct {
		AccessPrivateKey     string `mapstructure:"access_private_key"`
		AccessPublicKey      string `mapstructure:"access_public_key"`
		AccessTokenExpiredIn int64  `mapstructure:"access_token_expired_in"`
		accessTokenMaxAge    int64  `mapstructure:"access_token_max_age"`

		RefreshPrivateKey     string `mapstructure:"refresh_private_key"`
		RefreshPublicKey      string `mapstructure:"refresh_public_key"`
		RefreshTokenExpiredIn int64  `mapstructure:"refresh_token_expired_in"`
		refreshTokenMaxAge    int64  `mapstructure:"refresh_token_max_age"`
	}
)

func NewConfig() *Configuration {
	var config Configuration
	viper.SetConfigFile("config/crm_core/config.yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
