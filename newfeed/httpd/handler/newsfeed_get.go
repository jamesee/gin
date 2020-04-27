package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamessee/gin.git/newfeed/httpd/platform/newsfeed"
)

func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
