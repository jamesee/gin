package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamessee/gin.git/newfeed/httpd/platform/newsfeed"
)

type newsFeedPostRequest struct {
	Title string `json :"title"`
	Post  string `json :"post"`
}

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
