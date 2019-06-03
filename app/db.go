package app

import (
	"errors"
	"sync"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var db *xorm.Engine
var once sync.Once

func NewDB(connectionString string) *xorm.Engine {

	once.Do(func() {
		engine, err := xorm.NewEngine("mysql", connectionString)

		if err != nil {
			log.Panic(err)
		}

		if err = engine.Ping(); err != nil {
			log.Panic(err)
		}

		// Use an adapter to the logrus Standard logger
		logger := &logrusAdapter{Logger: log.StandardLogger()}
		engine.SetLogger(logger)
		engine.ShowSQL(true)
		// TODO: Switch to GonicMapper to handle ID field to column names better
		//engine.SetColumnMapper(core.GonicMapper{})

		results, err := engine.QueryString("SELECT VERSION() as v")

		log.Infof("Connected to SQL server:%s", results[0]["v"])

		db = engine
	})

	return db

}

//GetDB returns the already initialised DB otherwise error will be returned.
func GetDB() (*xorm.Engine, error) {
	if db == nil {
		return nil, errors.New("Expected DB instance to be in initialised state but found instance as nil")
	}

	return db, nil
}

type logrusAdapter struct {
	*log.Logger
	showSQL bool
}

//
//func (logger *logrusAdapter) Infof(format string, v ...interface{}) {
//	logger.Logger.Infof(format, v)
//}

func (logger *logrusAdapter) Level() core.LogLevel {
	// Convert between logrus and xorm log levels
	logint := -1 * (int(logger.Logger.Level) - 5)

	return core.LogLevel(logint)
}

func (logger *logrusAdapter) SetLevel(logLevel core.LogLevel) {
	logint := -1 * (logLevel - 5)
	logger.Logger.SetLevel(log.Level(logint))
}

func (logger *logrusAdapter) ShowSQL(show ...bool) {
	if len(show) > 0 {
		logger.showSQL = show[0]
	} else {
		logger.showSQL = true
	}
}

func (logger *logrusAdapter) IsShowSQL() bool {
	return logger.showSQL
}
