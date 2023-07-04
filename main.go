package main

import (
	"backend-mytodo/migration"
	"backend-mytodo/repository"
	"backend-mytodo/resthandler"
	"backend-mytodo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	api := router.Group("/api/v1")
	api.POST("/todo", todoResthandler.Create)
	api.GET("/todo", todoResthandler.FindAll)

	router.Run()
}