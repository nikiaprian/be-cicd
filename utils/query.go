package utils

import (
	"kel15/models"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetPaginationDefault(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 15
	}

	search := c.Query("search")

	pagination := models.Pagination{
		PageActive: page,
		Limit:      limit,
		Search:     "%" + search + "%",
	}
	c.Set("pagination", pagination)
}

func SetPaginationNew(c *gin.Context, pagination models.Pagination) {
	c.Set("pagination", pagination)
}

func SetTotalPagePagination(c *gin.Context, pagination models.Pagination) models.Pagination {
	var totalPage float64 = float64(pagination.Count) / float64(pagination.Limit)
	var totalPageInt int = int(math.Ceil(totalPage))
	pagination.TotalPage = totalPageInt
	return pagination
}

func GetPagination(c *gin.Context) models.Pagination {
	pagination := c.MustGet("pagination").(models.Pagination)
	return pagination
}
