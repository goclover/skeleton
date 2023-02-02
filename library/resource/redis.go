package resource

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/redis/go-redis/v9"
	"reflect"
)

var redisInstance = &redisInstances{}

type redisConfig struct {
	Name  string `toml:"Name"`
	Redis struct {
		Host     string `toml:"Host"`
		Port     int    `toml:"Port"`
		Password string `toml:"Password"`
		Database int    `toml:"Database"`
	} `toml:"Redis"`
}

type redisInstances struct {
	conf    string
	Default *redis.Client `redis:"default"`
}

func initRedis(ctx context.Context, cd string) (err error) {
	//config path
	redisInstance.conf = cd

	var reTyp = reflect.TypeOf(*redisInstance)
	var numFields = reTyp.NumField()
	for {
		numFields--
		if numFields < 0 {
			break
		}
		var field = reTyp.Field(numFields)
		if field.Tag.Get("redis") == "" {
			continue
		}
		var ins, err = getRedisInstance(dbInstance.conf, field.Tag.Get("redis"))
		if err != nil {
			return err
		}
		reflect.ValueOf(redisInstance).Elem().FieldByName(field.Name).Set(reflect.ValueOf(ins))
	}
	return
}

func getRedisInstance(cd, srv string) (c *redis.Client, err error) {
	var rc = redisConfig{}
	if _, err = toml.DecodeFile(fmt.Sprintf("%s/redis/%s.toml", cd, srv), &rc); err != nil {
		return
	}
	c = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rc.Redis.Host, rc.Redis.Port),
		Password: rc.Redis.Password,
		DB:       rc.Redis.Database,
	})
	return
}

func GetRedis() *redisInstances {
	return redisInstance
}
