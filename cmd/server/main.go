package main

import (
	server "Final/cmd/server/handler"
	"Final/internal/odontologo"
	"Final/internal/paciente"
	"Final/internal/turno"
	"Final/pkg/middleware"
	"Final/pkg/store"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./cmd/server/.env")
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	//FIXME pasar a env
	db, err := sql.Open("mysql", "root:root@/clinica")
	//user1:secret_password@/my_db"
	if err != nil {
		log.Fatal(err)
	}
	storage := store.NewSqlStore(db)

	//repo/service para odontologo
	repoOdontologo := odontologo.NewRepository(storage)
	serviceOdontologo := odontologo.NewService(repoOdontologo)
	odontologoHandler := server.NewOdontologoHandler(serviceOdontologo)

	//repo/service para paciente
	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := server.NewPacienteHandler(servicePaciente)

	//repo/service para turno
	repoTurno := turno.NewRepository(storage)
	serviceTurno := turno.NewService(repoTurno)
	turnoHandler := server.NewTurnoHandler(serviceTurno)

	//FIXME pasar a new
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("", func(c *gin.Context) { c.String(200, "Bienvenido a la Clínica Odontológica") })

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("", odontologoHandler.GetAll())
		odontologos.POST("", middleware.Authentication(), odontologoHandler.Post())
		odontologos.GET(":id", odontologoHandler.GetById())
		odontologos.PUT(":id", middleware.Authentication(), odontologoHandler.Put())
		odontologos.PATCH(":id", middleware.Authentication(), odontologoHandler.Patch())
		odontologos.DELETE(":id", middleware.Authentication(), odontologoHandler.Delete())

	}
	//TODO AUTENTICACION
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacienteHandler.GetAll())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.GET(":id", pacienteHandler.GetById())
		pacientes.PUT(":id", pacienteHandler.Put())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.DELETE(":id", pacienteHandler.Delete())

	}
	//TODO AUTENTICACION

	turnos := r.Group("/turnos")
	{
		turnos.GET("", turnoHandler.GetAll())
		turnos.POST("", turnoHandler.Post())
		turnos.GET(":id", turnoHandler.GetById())
		turnos.PUT(":id", turnoHandler.Put())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.POST("/turnoauxiliar", turnoHandler.PostDNIMat())
		turnos.GET("/paciente", turnoHandler.GetAllByDni())

	}

	//TODO hacer una variable del puerto
	r.Run(":8080")

}
