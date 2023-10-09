package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	PdamCd              PdamCd
	PdamAdmin           PdamAdmin
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RedisUrl            string        `mapstructure:"REDIS_URL"`
	RedisDB             string        `mapstructure:"REDIS_DB"`
	RedisPassword       string        `mapstructure:"REDIS_PASSWORD"`
	AllowDeposit        string        `mapstructure:"ALLOW_DEPOSIT"`
	KilatUrl            string        `mapstructure:"KILAT_URL"`
	KilatSecret         string        `mapstructure:"KILAT_SECRET"`
	KilatUser           string        `mapstructure:"KILAT_USER"`
	KilatMerchant       string        `mapstructure:"KILAT_MERCHANT"`
	DigiSellerUser      string        `mapstructure:"DIGI_SELLER_USERNAME"`
	DigiApiKey          string        `mapstructure:"DIGI_API_KEY"`
	WhiteListIP         string        `mapstructure:"WHITELIST_API"`
	MongoUrl            string        `mapstructure:"MONGO_URL"`
	MongoIsPassword     string        `mapstructure:"MONGO_IS_PASSWORD"`
	MongoUser           string        `mapstructure:"MONGO_USER"`
	MongoPassword       string        `mapstructure:"MONGO_PASSWORD"`
	MongoDbName         string        `mapstructure:"MONGO_DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
