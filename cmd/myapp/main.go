package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"student-management/config"
	"student-management/pkg/controller"
	"student-management/pkg/repository"
	"student-management/pkg/repository/inmemory"
	"student-management/pkg/repository/mysqldb"
	"student-management/pkg/service"
	"student-management/routes"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Only log the debug severity or above.
	log.SetLevel(log.DebugLevel)

	log.Info("Welcome to the student management application...")
	log.Info("starting the server")
	defer log.Warn("Exiting the server..")

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	r := gin.Default()

	studentController, courseController := setupDependancies(cfg)
	routes.InitRoutes(r, studentController, courseController)

	server := http.Server{
		Addr:        cfg.HTTPServer.Port,
		Handler:     r,
		ReadTimeout: time.Duration(cfg.HTTPServer.ReadTimeout) * time.Second,
	}
	// err := r.Run(":8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Panic("error occured while starting the server : " + err.Error())
	}
}

func setupDependancies(cfg *config.Config) (*controller.StudentController, *controller.CourseController) {
	var studentRepo repository.StudentRepository
	var courseRepo repository.CourseRepository
	var err error
	if cfg.Env == "production" {
		db := newMySQLDBConnection(cfg)
		studentRepo, err = mysqldb.NewMysqlStudentRepository(db)
		if err != nil {
			log.Fatal("error occured while initializing student repository :" + err.Error())
		}
		courseRepo, err = mysqldb.NewMysqlCourseRepository(db)
		if err != nil {
			log.Fatal("error occured while initializing course repository" + err.Error())
		}
	} else {
		// Use in-memory repository for development
		studentRepo = inmemory.NewStudentRepository()
		courseRepo = inmemory.NewCourseRepository()
	}

	studentService := service.NewStudentService(studentRepo)
	courseService := service.NewCourseService(courseRepo)

	studentController := controller.NewStudentController(studentService)
	courseController := controller.NewCourseController(courseService)

	return studentController, courseController
}

func newMySQLDBConnection(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", cfg.MysqlConfig.Username, cfg.MysqlConfig.Password,
		cfg.MysqlConfig.HostName, cfg.MysqlConfig.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("error occured while initializing database : ", err.Error())
	}

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.MysqlConfig.Database)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("error occured while creating database : ", err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("USE %s", cfg.MysqlConfig.Database))
	if err != nil {
		log.Fatal("error occured while connecting with database : ", err.Error())
	}

	return db
}
