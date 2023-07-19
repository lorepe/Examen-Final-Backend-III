package main

import (
	server "Final/cmd/server/handler"
	_ "Final/docs"
	"os"

	"Final/internal/odontologo"

	"Final/internal/paciente"

	"Final/internal/turno"

	"Final/pkg/middleware"

	"Final/pkg/store"
	_ "Final/pkg/web"
	"database/sql"
	"log"

	_ "Final/cmd/server/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Certified Tech Developer- Especialización Backend Golang
//	@version		1.0
//	@description	This API Handle a Dental Clinic in Go using Gin framework
//	@termsOfService	https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
//	@contact.name	API Support
//	@contact.url	https://developers.ctd.com.ar/support
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080

func main() {

	err := godotenv.Load("./cmd/server/.env")
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}
	puerto := os.Getenv("PORT")

	dbstring := os.Getenv("DATASOURCE")
	db, err := sql.Open("mysql", dbstring)
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

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//SI DESEA UTLIZAR EL LOGGER PERSONALIZADO DESCOMENTE ESTA LÍNEA
	// r.Use(middleware.Logger())

	//swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacienteHandler.GetAll())
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.GET(":id", pacienteHandler.GetById())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())

	}

	turnos := r.Group("/turnos")
	{
		turnos.GET("", turnoHandler.GetAll())
		turnos.POST("", middleware.Authentication(), turnoHandler.Post())
		turnos.GET(":id", turnoHandler.GetById())
		turnos.PUT(":id", middleware.Authentication(), turnoHandler.Put())
		turnos.PATCH(":id", middleware.Authentication(), turnoHandler.Patch())
		turnos.DELETE(":id", middleware.Authentication(), turnoHandler.Delete())
		turnos.POST("/turnoauxiliar", middleware.Authentication(), turnoHandler.PostDNIMat())
		turnos.GET("/paciente", turnoHandler.GetAllByDni())

	}

	r.Run(puerto)

}
