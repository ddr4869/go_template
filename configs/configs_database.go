package configs

import (
	"context"
	"fmt"

	"github.com/go-board/ent"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/the-medium-tech/platform-externals/log"
)

func (c *Config) SetEntRepo() *ent.Client {
	dbCfg := c.DB

	dbType := dbCfg.Driver
	var dataSourceFmt string
	switch dbType {
	case "mariadb":
		fallthrough
	case "mysql":
		dataSourceFmt = fmt.Sprintf("%s:%s@tcp(%s:%s)/",
			dbCfg.Username,
			dbCfg.Password,
			dbCfg.Address,
			dbCfg.Port,
		) + "%s?parseTime=true"
	}

	clientManager, err := ent.Open(dbType, fmt.Sprintf(dataSourceFmt, dbCfg.DBName))
	if err != nil {
		panic("could not open manager db: " + err.Error())
	}
	clientManager = clientManager.Debug()
	ctx := context.Background()
	err = clientManager.Schema.Create(ctx)

	if err != nil {
		panic("could not create 'mdl_manager' schema:" + err.Error())
	}

	return clientManager
}

func (c *Config) SetRedisRepo() *redis.Client {

	redisCfg := c.Redis
	log.Infof("Addr:%s, Port:%s", redisCfg.Host, redisCfg.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password, // no password set
		DB:       0,                 // use default DB
	})

	c.RedisCient = rdb
	ctx := context.Background()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	return rdb
}

func (c *Config) GetRedisClient() *redis.Client {
	return c.RedisCient
}

func (c *Config) RedisTest() {

	log.Info("RedisTest Call !!!")
	ctx := context.Background()
	err := c.RedisCient.Set(ctx, "key3", "value", 0).Err()
	if err != nil {
		log.Info("RedisTest Error !!!")
		panic(err)
	}
}
