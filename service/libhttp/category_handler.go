package libhttp

import (
	"net/http"
	"strconv"

	"library/service/handler"

	"github.com/gin-gonic/gin"
)

type CategoryHTTPHandler struct {
	service *handler.CategoryService
}

func NewCategoryHTTPHandler(service *handler.CategoryService) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{service: service}
}

// Post
func (h *CategoryHTTPHandler) CreateCategory(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	category, err := h.service.CreateCategory(body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHTTPHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)

}

func (h *CategoryHTTPHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
