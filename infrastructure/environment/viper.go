package environment

import (
	"github.com/spf13/viper"
	"time"
)

type Environment struct {
	// ============== APPLICATION ==================//
	Env      string `mapstructure:"ENV"`
	HTTPPort string `mapstructure:"HTTP_PORT"`
	GRPCPort int    `mapstructure:"GRPC_PORT"`

	// ============== DATABASE ==================//
	DBDriver      string        `mapstructure:"DB_DRIVER"`
	DBHost        string        `mapstructure:"DB_HOST"`
	DBPort        string        `mapstructure:"DB_PORT"`
	DBUsername    string        `mapstructure:"DB_USERNAME"`
	DBPassword    string        `mapstructure:"DB_PASSWORD"`
	DBName        string        `mapstructure:"DB_NAME"`
	DBSslMode     string        `mapstructure:"DB_SSL_MODE"`
	DBMaxOpenConn int           `mapstructure:"DB_MAX_OPEN_CONN"`
	DBMaxIdleConn int           `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxLifeTime time.Duration `mapstructure:"DB_MAX_LIFE_TIME"`
	DBMaxIdleTime time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
}

func Load(fileType, filePath string) (*Environment, error) {
	viper.SetConfigType(fileType)
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return &Environment{}, err
	}

	env := Environment{}
	err = viper.Unmarshal(&env)

	return &env, err
}
