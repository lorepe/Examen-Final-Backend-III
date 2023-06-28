package main

import (
	server "Final/cmd/server/odontologoHandler"
	"Final/internal/odontologo"
	"Final/pkg/store"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//repo/service para odontologo
	//FIXME pasar a env
	db, err := sql.Open("mysql", "root:root@tcp(serverPort:3306)/clinica")
	if err != nil {
		panic(err)
	}
	storage := store.NewSqlStore(db)

	repoOdontologo := odontologo.NewRepository(storage)
	serviceOdontologo := odontologo.NewService(repoOdontologo)
	odontologoHandler := server.NewOdontologoHandler(serviceOdontologo)

	//FIXME pasar a new
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("", func(c *gin.Context) { c.String(200, "Bienvenido a la Clínica Odontológica") })

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("", odontologoHandler.GetAll())
	}

	//TODO hacer una variable del puerto
	r.Run(":8080")

}
