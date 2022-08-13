package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AccessTokenExpiration       time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRATION"`
	TokenSymmetricKey           string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	DriverName                  string        `mapstructure:"DRIVER_NAME"`
	DataSourceName              string        `mapstructure:"DATA_SOURCE_NAME"`
	SeverPort                   uint          `mapstructure:"SEVER_PORT"`
	MigrationSource             string        `mapstructure:"MIGRATION_SOURCE"`
	VerificationTokenExpiration time.Duration `mapstructure:"VERIFICATION_TOKEN_EXPIRATION"`
	RedisAddress                string        `mapstructure:"REDIS_ADDRESS"`
	RedisPassword               string        `mapstructure:"REDIS_PASSWORD"`
	SMTPServer                  string        `mapstructure:"SMTP_SERVER"`
	SMTPPort                    int           `mapstructure:"SMTP_PORT"`
	SMTPUser                    string        `mapstructure:"SMTP_USER"`
	SMTPPassword                string        `mapstructure:"SMTP_PASSWORD"`
	SenderIdentity              string        `mapstructure:"SENDER_IDENTITY"`
	SenderEmail                 string        `mapstructure:"SENDER_EMAIL"`
	SiteURL                     string        `mapstructure:"SITE_URL"`
	SiteActivateEndpoint        string        `mapstructure:"SITE_ACTIVATE_ENDPOINT"`
	SiteResetPasswordEndpoint   string        `mapstructure:"SITE_RESET_PASSWORD_ENDPOINT"`
	SystemName                  string        `mapstructure:"SYSTEM_NAME"`
	LogoURL                     string        `mapstructure:"LOGO_URL"`
	MYSQLPort                   string        `mapstructure:"MYSQL_PORT"`
	MYSQLUser                   string        `mapstructure:"MYSQL_USER"`
	MYSQLPass                   string        `mapstructure:"MYSQL_PASSWORD"`
	MYSQLDBName                 string        `mapstructure:"MYSQL_DB_NAME"`
	MYSQLHostName               string        `mapstructure:"MYSQL_HOST_NAME"`
}

const (
	defaultTokenExpiration             = time.Hour * 24
	defaultVerificationTokenExpiration = time.Minute * 30
	defaultServerPort                  = 7100
	defaultRedisAddress                = "localhost:6379"
)

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	setDefaultConfigValues(config)
	return config, nil
}

func setDefaultConfigValues(config *Config) {
	if config.AccessTokenExpiration == 0 {
		config.AccessTokenExpiration = defaultTokenExpiration
	}
	if config.VerificationTokenExpiration == 0 {
		config.VerificationTokenExpiration = defaultVerificationTokenExpiration
	}
	if config.RedisAddress == "" {
		config.RedisAddress = defaultRedisAddress
	}
	if config.SeverPort == 0 {
		config.SeverPort = defaultServerPort
	}
}
