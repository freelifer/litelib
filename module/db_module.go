package module

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/go-xorm/xorm"

	. "github.com/freelifer/litelib/public"
)

var (
	DatabaseEngine struct {
		Tables []interface{}
	}
)

type DatabaseModule struct {
}

func NewDatabaseModule() Module {
	module := DatabaseModule{}
	return &module
}

func (module *DatabaseModule) Setup(config map[string]string) error {
	log.Println("database module setup")
	var err error
	DB.Engine, err = getEngine(config)
	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}
	DB.Engine.ShowSQL(true)
	// x.Logger().SetLevel(core.LOG_DEBUG)

	if err = DB.Engine.StoreEngine("InnoDB").Sync2(DatabaseEngine.Tables...); err != nil {
		log.Fatalf("sync database struct error: %v\n", err)
	}
	return nil
}

func init() {
	Register("database", NewDatabaseModule)
}

func DropTables() error {
	// return x.DropTables(new(WxUser), new(PasswdInfo), new(IconInfo))
	return nil
}

func getEngine(config map[string]string) (*xorm.Engine, error) {
	connStr := ""
	// var Param string = "?"
	// if strings.Contains(settings.DatabaseCfg.HostName, Param) {
	// 	Param = "&"
	// }
	switch config["type"] {
	case "mysql":
		if config["host"][0] == '/' {
			connStr = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8",
				config["user"], config["passwd"], config["host"], config["name"])
		} else {
			connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
				config["user"], config["passwd"], config["host"], config["name"])
		}
		// var engineParams = map[string]string{"rowFormat": "DYNAMIC"}
		// return xorm.NewEngineWithParams(settings.DatabaseCfg.Type, connStr, engineParams)
	case "sqlite3":
		// if !EnableSQLite3 {
		// 	return nil, errors.New("This binary version does not build support for SQLite3.")
		// }
		if err := os.MkdirAll(path.Dir(config["path"]), os.ModePerm); err != nil {
			return nil, fmt.Errorf("Fail to create directories: %v", err)
		}
		connStr = "file:" + config["path"] + "?cache=shared&mode=rwc"
	default:
		return nil, fmt.Errorf("Unknown database type: %s", config["type"])
	}
	fmt.Println(connStr)
	return xorm.NewEngine(config["type"], connStr)
}
