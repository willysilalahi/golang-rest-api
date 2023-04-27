package main

import (
	"golang-rest-api/handler"
	"golang-rest-api/model"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error!")
	}

	DB.AutoMigrate(&model.Book{})

	/*
		//creating book
		var book model.Book
		book.Title = "Seni Bodoh Amat"
		book.Description = "Sebuah buku yang mengajari orang untuk bodoh amat lah cok"
		book.Price = 140000
		book.Rating = 5
		DB.Create(&book)

		//Get Single Book
		var book model.Book
		err = DB.First(&book, 3).Error
		if err != nil {
			fmt.Println("=========================")
			fmt.Println("Error finding book record")
			fmt.Println("=========================")
		}
		fmt.Println("Title :", book.Title)
		fmt.Println("%v", book)


		//Get All Book
		var books []model.Book
		err = DB.Find(&books).Error
		if err != nil {
			fmt.Println("=========================")
			fmt.Println("Error get all book record")
			fmt.Println("=========================")
		}

		for i, value := range books {
			fmt.Println("Data ", i, " :", value.Title)
		}

		//Get All Book Query Search
		var books []model.Book
		search := "Naj"
		err = DB.Where("title LIKE ?", "%"+search+"%").Find(&books).Error
		if err != nil {
			fmt.Println("=========================")
			fmt.Println("Error get all book record")
			fmt.Println("=========================")
		}

		for i, value := range books {
			fmt.Println("Data ", i, " :", value.Title)
		}

		// Update book data by id
		var book model.Book
		DB.First(&book, 3)
		err = DB.Model(&book).Updates(model.Book{Title: "Catatan Ko Wilson", Price: 201000}).Error
		if err != nil {
			fmt.Println("=========================")
			fmt.Println("Error updated book data  ")
			fmt.Println("=========================")
		}

		// Delete Data
		var book model.Book
		err = DB.Delete(&book, 4).Error
		if err != nil {
			fmt.Println("=========================")
			fmt.Println("Error updated book data  ")
			fmt.Println("=========================")
		}
		fmt.Println("Data buku berhasil di hapus bro!")
	*/

	// Routing
	route := gin.Default()
	v1 := route.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/book", handler.CreateBook)

	route.Run()
}
