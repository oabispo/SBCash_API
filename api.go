package main

import (
	"log"
	"net/http"

	"sbcash_api/endpoint"

	mysql "github.com/oabispo/BAMySQLHelper"

	router "sbcash_api/BARoutingHelper"

	"github.com/oabispo/BAIniHandler"
)

//import "test" //test.TestaTudo( db );

func main() {
	ini, _ := BAIniHandler.NewBAIniHandler("api.ini")
	ini.Save(true)
	host := ini.ReadString("DBSettings", "host", "127.0.0.1:3306")
	database := ini.ReadString("DBSettings", "database", "database")
	userName := ini.ReadString("DBSettings", "user", "root")
	password := ini.ReadString("DBSettings", "passwd", "password")
	portNum := ini.ReadInteger("APISettings", "portNum", 8080)

	api := router.NewBARoutHelper(router.NewBARoutConfig(portNum, onStart, onBeforeHandling, onAfterHandling))
	db := mysql.NewSQLConnection(host, database, userName, password)
	db.Ping()

	api.StartRouting(endpoint.NewRouterHome(db))
}

func onStart(portNum int) {
	log.Printf("Server starting on port %v\n", portNum)
}

func onBeforeHandling(r *http.Request) {
	log.Printf("Chamarei %v, method: %v\n", r.URL.String(), r.Method)
}

func onAfterHandling(path string, method string) {
	log.Printf("Chamei %v\n", path)
}
