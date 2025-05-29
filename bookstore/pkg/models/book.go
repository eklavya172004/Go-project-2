package models

import(
	"github.com/jinzhu/gorm"
	"github.com/eklavya/go-bookstore/pkg/config"
)

var db *gorm.DB
//no inheritance in go : no super and parent class
type Book struct {
	gorm.Model
	// Name string `gorm:""json:"name"`
	//what is grom:""json:"name"`?
	// gorm:""json:"name"` is a struct tag in Go that specifies how the field should be handled by the GORM ORM and JSON encoding/decoding.
	//here `gorm:""` indicates that the field should be included in the database schema, and `json:"name"` specifies that when the struct is converted to JSON, this field will be represented as "name".
	// Author string `json:"author"`
	// Publication string `json:"publication"`

	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect() // Connect to the database using the config package
	db = config.GetDB()
	// Migrate the Book struct to create the corresponding table in the database
	db.AutoMigrate(&Book{})
}

//this fuction is used to create a new book record in the database
// CreateBook creates a new book record in the database
// It uses the GORM ORM to handle database operations.
// The function returns a pointer to the Book instance that was created.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	// Create a new record in the database
	// The NewRecord method checks if the record is new (not yet saved to the database).
	db.Create(&b);
	// returns the created book record
	return b
}

func GetallBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBooksById(Id int64) (*Book,*gorm.DB){

	//(*Book, *gorm.DB): The function returns two values:

// A pointer to a Book struct (*Book) — this will contain the book's data if found.

// A pointer to a gorm.DB object — this includes the result of the database operation (e.g., errors, rows affected, etc.).

	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

func DeleteBookById(Id int64) Book{
	var book Book

	// First fetch the book by ID
	result := db.First(&book, Id)
	if result.Error != nil {
		// Optionally handle not found or return zero-value book
		return Book{}
	}

	// Then delete it
	db.Delete(&book)

	// Return the deleted book
	return book
}