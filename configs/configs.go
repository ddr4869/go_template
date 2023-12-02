package configs

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/kafka-go/ent"
)

const AppConfigPath = "./configs/.env"

var ServerConfig *Config
var GlobalTest *GlobalConfigTest

func GlobalInit() *GlobalConfigTest {
	GlobalTest = &GlobalConfigTest{
		Foo: "a",
		Bar: "b",
	}
	return GlobalTest
}

type GlobalConfigTest struct {
	Foo string
	Bar string
}

type Config struct {
	Backend    *BackendConfig
	DB         *DBConf
	Redis      *RedisConf
	EntClient  *ent.Client
	RedisCient *redis.Client
}

type BackendConfig struct {
	ServerPort        string
	ServerPortOpenApi string
}

type RedisConf struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Password   string `yaml:"password"`
	SessionKey string `yaml:"session-key"`
}

type DBConf struct {
	Address  string
	Port     string
	Username string
	Password string
	DBName   string
	Driver   string
	Tables   string
	Logging  string
}

func Init() *Config {
	backend := GetBackendConfig()
	db := GetDBConfig()
	redis := GetRedisConf()
	GlobalInit()

	ServerConfig = &Config{
		Backend: backend,
		DB:      db,
		Redis:   redis,
	}
	// ServerConfig.Backend = backend
	// ServerConfig.DB = db
	// ServerConfig.Redis = redis
	// return ServerConfig
	return ServerConfig
}

func GetBackendConfig() *BackendConfig {
	return &BackendConfig{
		ServerPort:        os.Getenv("SERVER_PORT"),
		ServerPortOpenApi: os.Getenv("SERVER_PORT_OPENAPI"),
	}
}

func GetDBConfig() *DBConf {
	return &DBConf{
		Address:  os.Getenv("DB_ADDRESS"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Tables:   os.Getenv("DB_TABLES"),
		Logging:  os.Getenv("DB_LOGGING"),
	}
}

func GetRedisConf() *RedisConf {
	return &RedisConf{
		Host:       os.Getenv("REDIS_ADDRESS"),
		Port:       os.Getenv("REDIS_PORT"),
		Password:   os.Getenv("REDIS_PASSWORD"),
		SessionKey: os.Getenv("REDIS_SESSIONKEY"),
	}
}
