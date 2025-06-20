package main

import (
	"errors"   //build in
	"net/http" //build in

	"github.com/gin-gonic/gin" //gin framework
)

//capital = public(can export), lowercase = private
type Book struct {
	ID 			string	`json:"id"` 
	Title 		string	`json:"title"`
	Author 		string	`json:"author"`
	Quantity	int		`json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell",Quantity: 5},
	{ID: "2", Title: "To Kill a Mockingbird", Author:"Harper Lee", Quantity: 3},
	{ID: "3", Title: "The Great Gatsby", Author:"F. Scott Fitzgerald", Quantity: 2},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(id string) (*Book, error) {
	for i := range books {
		if books[i].ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookByIDHandler(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book data"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id,ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID in query"})
		return
	}

	book, err := getBookByID(id)
	// need to use index for else queantity would not be updated
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book is not available"})
		return
	}
	book.Quantity -= 1 
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID in query"})
		return
	}
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	//GET all books
	router.GET("/books", getBooks)
	//GET book by ID
	router.GET("/books/:id", bookByIDHandler)
	//POST create a new book
	router.POST("/books", createBook)
	//POST checkout a book
	router.PATCH("/checkout", checkoutBook)

	router.PATCH("/return",returnBook)

	router.Run("localhost:8080")
}

