package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"reflect"
	"strings"
)

var dbInstance = &gormInstances{}

type dbConfig struct {
	Name     string `toml:"Name"`
	Strategy string `toml:"Strategy"`
	Database struct {
		Username        string `toml:"Username"`
		Password        string `toml:"Password"`
		ConnMaxLifeTime int    `toml:"ConnMaxLifeTime"`
		ConnMaxIdleTime int    `toml:"ConnMaxIdleTime"`
		MaxIdleConns    int    `toml:"MaxIdleConns"`
		DBDriver        string `toml:"DBDriver"`
		DBName          string `toml:"DBName"`
		DSNParams       string `toml:"DSNParams"`
		Manual          []struct {
			Name string `toml:"Name"`
			Host string `toml:"Host"`
			Port int    `toml:"Port"`
		} `toml:"Manual"`
	} `toml:"Database"`
}

type gormInstances struct {
	conf    string
	Default *gorm.DB `database:"default"`
}

func initGORM(_ context.Context, cf string) error {
	//set config dir
	dbInstance.conf = cf

	var dbsTyp = reflect.TypeOf(*dbInstance)
	var numFields = dbsTyp.NumField()
	for {
		numFields--
		if numFields < 0 {
			break
		}
		var field = dbsTyp.Field(numFields)
		if field.Tag.Get("database") == "" {
			continue
		}
		var db, err = getInstance(dbInstance.conf, field.Tag.Get("database"))
		if err != nil {
			return err
		}
		reflect.ValueOf(dbInstance).Elem().FieldByName(field.Name).Set(reflect.ValueOf(db))
	}
	return nil
}

func getInstance(cf, srv string) (db *gorm.DB, err error) {
	var dc dbConfig
	//write config for database sources
	if dc, err = loadConfig(cf, srv, "write"); err == nil {
		var dsnList = getDSNList(dc)
		if db, err = gorm.Open(mysql.Open(dsnList[0]), &gorm.Config{DryRun: false, Logger: nil}); err != nil {
			return
		}
		var sources, replicas []gorm.Dialector
		for _, v := range dsnList[1:] {
			sources = append(sources, mysql.Open(v))
		}

		//read config for database replicas
		if dc, err = loadConfig(cf, srv, "read"); err != nil {
			return
		}
		for _, v := range dsnList {
			replicas = append(replicas, mysql.Open(v))
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:           sources,
			Replicas:          replicas,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}))
	}
	return
}

func loadConfig(cf, srv, tpl string) (dc dbConfig, err error) {
	if _, err = toml.DecodeFile(fmt.Sprintf("%s/database/%s_%s.toml", cf, srv, tpl), &dc); err != nil {
		return
	}
	if len(dc.Database.Manual) <= 0 {
		err = errors.New(fmt.Sprintf("database config (%s) has empty manual", srv))
	}
	if strings.TrimSpace(dc.Database.DBDriver) == "" {
		err = errors.New("db driver is empty")
	}
	if strings.ToLower(dc.Database.DBDriver) != "mysql" {
		err = errors.New(fmt.Sprintf("db driver (%s) is unsupport", dc.Database.DBDriver))
	}
	if len(dc.Database.Manual) <= 0 {
		err = errors.New(fmt.Sprintf("invalid database manual srv (%s)", srv))
		return
	}
	return
}

func getDSNList(c dbConfig) (dsn []string) {
	var s = c.Database
	for _, v := range s.Manual {
		dsn = append(dsn, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", s.Username, s.Password,
			v.Host,
			v.Port,
			s.DBName,
			s.DSNParams))
	}
	return
}

func GetGORM() *gormInstances {
	return dbInstance
}
