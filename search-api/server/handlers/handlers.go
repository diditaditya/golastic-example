package handlers

import (
	schemas "app/schemas"
	search "app/search"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDocuments(c *gin.Context) {
	var docs []schemas.DocumentRequest
	if err := c.BindJSON(&docs); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}
	err := search.CreateDocuments(c, docs)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func SearchDocuments(c *gin.Context) {
	query := schemas.DocumentQuery{
		Query: c.Query("query"),
		Skip: 0,
		Take: 10,
	}
	if query.Query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		query.Skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		query.Take = i
	}

	res, err := search.SearchDocuments(c, query)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, res)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}