package main

import (
	"backend-mytodo/migration"
	"backend-mytodo/repository"
	"backend-mytodo/resthandler"
	"backend-mytodo/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/mytodo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// migration
	migration.RunMigration(db)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoResthandler := resthandler.NewTodoResthandler(todoService)

	router := gin.Default()
	// cors middleware
	config :=  cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	api := router.Group("/api/v1")
	api.POST("/todo", todoResthandler.Create)
	api.GET("/todo", todoResthandler.FindAll)

	router.Run()
}