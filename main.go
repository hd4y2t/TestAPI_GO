package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"API/book"
	"API/hendler"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/api"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("koneksi db gagal")
		// fmt.Println("koneksi db gagal")
	}

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHendler := hendler.NewBookHendler(bookService)
	// books, err := bookRepository.FindByID(2)
	// if err != nil {
	// 	log.Fatal("error")
	// }
	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// }
	// fmt.Println("Title : ", books.Title)
	// migratoin
	// db.AutoMigrate(book.Book{})

	// CRUD

	// CREATE
	// book := book.Book{}
	// book.Title = "sdasdas"
	// book.Price = 800000
	// book.Discoud = 5
	// book.Rating = 5
	// book.Desciption = "tesetttt"

	// db.Create(&book)

	// READ

	// READ seluruh data menggunakan find dengan deklarasi book dalam array
	// READ satu data menggunakan find dengan &books ditambah , dan variabel yang di cari
	// var books []book.Book
	// err = db.Find(&books).Error
	// if err != nil {
	// 	fmt.Println("gagal read data")
	// 	// fmt.Println("koneksi db gagal")
	// }

	// for _, b := range books {
	// 	fmt.Println("title : ", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	//Update data menggunakan where
	// var book book.Book
	// err = db.Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("gagal read data")
	// 	// fmt.Println("koneksi db gagal")
	// }

	// book.Title = "Coba coba"
	// db.Save(&book)

	// Delete Data
	// var book book.Book
	// err = db.Where("id = ?", 3).First(&book).Error
	// if err != nil {
	// 	fmt.Println("gagal read data")
	// 	// fmt.Println("koneksi db gagal")
	// }

	// book.Title = "Coba coba"
	// db.Delete(&book)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", bookHendler.RootHandler)
	v1.GET("/hello", bookHendler.HelloHandler)
	v1.GET("/books/:id", bookHendler.GetBook)
	v1.POST("/books", bookHendler.CreateBook)

	router.Run(":8081")

	// main
	// hendler
	// service
	// repository
	// db
	// mysql
}
