package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

const (
	newsPg    = "newsPg"
	crawlerPg = "crawlerPg"
)

func initNewsPg() {
	connStr := fmt.Sprintf(
		"dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		beego.AppConfig.String(newsPg+"::Db"),
		beego.AppConfig.String(newsPg+"::Host"),
		beego.AppConfig.String(newsPg+"::User"),
		beego.AppConfig.String(newsPg+"::Pwd"),
		beego.AppConfig.String(newsPg+"::Port"),
		beego.AppConfig.String(newsPg+"::SSLMode"),
	)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", connStr, beego.AppConfig.DefaultInt("ConnPoolSize", 20))
}

func initCrawlerPg() {
	connStr := fmt.Sprintf(
		"dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		beego.AppConfig.String(crawlerPg+"::Db"),
		beego.AppConfig.String(crawlerPg+"::Host"),
		beego.AppConfig.String(crawlerPg+"::User"),
		beego.AppConfig.String(crawlerPg+"::Pwd"),
		beego.AppConfig.String(crawlerPg+"::Port"),
		beego.AppConfig.String(crawlerPg+"::SSLMode"),
	)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("crawler", "postgres", connStr, beego.AppConfig.DefaultInt("ConnPoolSize", 30))
}

func setDefault() {
	orm.SetMaxIdleConns("default", 20)
	orm.SetMaxOpenConns("default", 20)
	orm.SetMaxIdleConns("crawler", 30)
	orm.SetMaxOpenConns("crawler", 30)
}

func initDb() {
	initNewsPg()
	initCrawlerPg()
	setDefault()
}
