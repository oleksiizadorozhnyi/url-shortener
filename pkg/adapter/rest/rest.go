package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"perviymoiserver/pkg/model"
	"perviymoiserver/pkg/shortener"
	"perviymoiserver/pkg/utils"
)

type Rest struct {
	router    *gin.Engine
	shortener *shortener.Shortener
}

func New(shortener *shortener.Shortener) *Rest {
	router := gin.Default()
	return &Rest{
		router:    router,
		shortener: shortener,
	}
}

func (r *Rest) Start() error {
	r.CreateHandlers()
	return r.router.Run(":8080")
}

func (r *Rest) CreateHandlers() {
	r.router.POST("/makeShortLink", r.NewLinkHandler)
	r.router.GET("/:shortLink", r.RedirectToLongHandler)
}

func (r *Rest) NewLinkHandler(c *gin.Context) {
	var newLink model.Page
	newLink.LongUrl = c.PostForm("longLink")
	if !utils.UrlValidator(newLink.LongUrl) {
		c.String(http.StatusBadRequest, "It's not a link")
		return
	}
	err := r.shortener.NewLink(&newLink)
	if err != nil {
		c.String(http.StatusInternalServerError, "Technical error")
		log.Print("NewLinkHandler() Link creation error ", err)
		return
	}
	c.String(http.StatusOK, newLink.ShortUrl)
}

func (r *Rest) RedirectToLongHandler(c *gin.Context) {
	var link model.Page
	link.ShortUrl = c.Param("shortLink")
	err := r.shortener.RedirectToLong(&link)
	if err != nil {
		c.String(http.StatusBadRequest, "There is no such short link")
		return
	}
	c.Redirect(http.StatusPermanentRedirect, link.LongUrl)
}
