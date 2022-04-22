package main

import (
	"log"
	"rest-api-go/book"
	"rest-api-go/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection Error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)

	v1.POST("/book", bookHandler.CreateBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)

	router.Run()
}

// layer |
// main
// handler
// service
// repository
// db
// mysql
