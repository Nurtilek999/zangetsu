package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"zangetsu/internal/domain/entity"
	"zangetsu/internal/domain/service"
	"zangetsu/pkg/logging"
	"zangetsu/pkg/response"
)

type AnimeHandler struct {
	animeService service.IAnimeService
	logger       logging.Logger
}

func NewAnimeHandler(animeService service.IAnimeService, logger logging.Logger) *AnimeHandler {
	var animeHandler = AnimeHandler{}
	animeHandler.animeService = animeService
	animeHandler.logger = logger
	return &animeHandler
}

func (h *AnimeHandler) SearchAnime(c *gin.Context) {
	query := c.Query("query")
	animes, err := h.animeService.SearchAnime(query)
	if err != nil {
		h.logger.Errorf(err.Error())
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response.ResponseOKWithData(c, animes)
}

func (h *AnimeHandler) Save(c *gin.Context) {
	var anime entity.AnimeViewModel
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &anime)

	if err != nil {
		h.logger.Errorf(err.Error())
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.animeService.SaveAnime(&anime)
	if err != nil {
		h.logger.Errorf(err.Error())
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response.ResponseOK(c, "successfully saved")
	return
}
