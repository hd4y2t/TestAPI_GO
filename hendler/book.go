package hendler

import (
	"API/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHendler struct {
	bookService book.Service
}

func NewBookHendler(bookService book.Service) *bookHendler {
	return &bookHendler{bookService}
}
func (h *bookHendler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "tes API",
	})
}

func (h *bookHendler) HelloHandler(c *gin.Context) {
	text := c.Query("text")
	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}

func (h *bookHendler) GetBook(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *bookHendler) CreateBook(c *gin.Context) {
	var bookRequest book.BooksRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Errors on field : %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
