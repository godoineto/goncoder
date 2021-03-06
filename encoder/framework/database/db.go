package database

import (
	"encoder/domain"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstace := NewDb()
	dbInstace.Env = "Test"
	dbInstace.DbTypeTest = "sqlite3"
	dbInstace.DsnTest = ":memory:"
	dbInstace.AutoMigrateDb = true
	dbInstace.Debug = true

	connection, err := dbInstace.Connect()

	if err != nil {
		log.Fatal("Test db error: %v\n", err)
	}
	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "Test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}
	if err != nil {
		return nil, err
	}
	if d.Debug {
		d.Db.LogMode(true)
	}
	if d.AutoMigrateDb {
		d.Db.AutoMigrateDb(&domain.Video{}, &domain.Job{})
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}
	return d.Db, nil
}
