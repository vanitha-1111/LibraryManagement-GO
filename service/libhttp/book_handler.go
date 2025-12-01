package libhttp

import (
	"net/http"
	"strconv"

	"library/service/handler"
	"library/service/models"

	"github.com/gin-gonic/gin"
)

type BookHTTPHandler struct {
	service *handler.BookService
}

func NewBookHTTPHandler(service *handler.BookService) *BookHTTPHandler {
	return &BookHTTPHandler{service: service}
}

// Post/books

func (h *BookHTTPHandler) CreateBook(c *gin.Context) {
	var req models.Book
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	id, err := h.service.CreateBook(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"book_id": id})
}

// Get/books
func (h *BookHTTPHandler) ListBooks(c *gin.Context) {
	books, err := h.service.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)

}

// Get/books/:id
func (h *BookHTTPHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	book, err := h.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func (h *BookHTTPHandler) GetBooksByCategoryName(c *gin.Context) {
	name := c.Param("name")

	books, err := h.service.ListBooksByCategoryName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
