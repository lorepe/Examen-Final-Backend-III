package main

import (
	server "Final/cmd/server/handler"
	"Final/internal/odontologo"
	"Final/internal/paciente"
	"Final/pkg/store"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//repo/service para odontologo
	//FIXME pasar a env
	db, err := sql.Open("mysql", "root:root@/clinica")
	//user1:secret_password@/my_db"
	if err != nil {
		log.Fatal(err)
	}
	storage := store.NewSqlStore(db)

	repoOdontologo := odontologo.NewRepository(storage)
	serviceOdontologo := odontologo.NewService(repoOdontologo)
	odontologoHandler := server.NewOdontologoHandler(serviceOdontologo)
	//repo/service para odontologo
	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := server.NewPacienteHandler(servicePaciente)

	//FIXME pasar a new
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("", func(c *gin.Context) { c.String(200, "Bienvenido a la Clínica Odontológica") })

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("", odontologoHandler.GetAll())
		odontologos.POST("", odontologoHandler.Post())
		odontologos.GET(":id", odontologoHandler.GetById())

	}

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacienteHandler.GetAll())
	}

	//TODO hacer una variable del puerto
	r.Run(":8080")

}
