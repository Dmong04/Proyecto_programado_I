package main

import (
	"database/sql"
	"log"
	"project/api"
	"project/dto"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver  = "mysql"
	dbSource  = "root:@tcp(localhost:3306)/coco_tours_db?charset=utf8mb4&parseTime=True&loc=Local"
	serverUrl = "127.0.0.1:8080"
)

func main() {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("No se pudo establecer la conexión", err)
	}
	dbtx := dto.NewDbTransaction(connection)
	server, err := api.NewServer(dbtx)
	if err != nil {
		log.Fatal("No se ha podido iniciar el servidor", err)
	}
	err = server.Start(serverUrl)
	if err != nil {
		log.Fatal("No se ha podido iniciar el servidor", err)
	}
}
