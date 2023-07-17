package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"zangetsu/internal/domain/entity"
	"zangetsu/pkg/config"
	"zangetsu/pkg/response"
)

//type UserHandler struct {
//	userService service.IUserService
//	logger      logging.Logger
//}

//func NewUserHandler(userService service.IUserService, logger logging.Logger) *UserHandler {
//	var userHandler = UserHandler{}
//	userHandler.userService = userService
//	userHandler.logger = logger
//	return &userHandler
//}

func (h *Handler) LoginGmail(c *gin.Context) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate") //oauth2.AccessTypeOffline
	c.Redirect(http.StatusFound, url)
}

func (h *Handler) CallbackGmail(c *gin.Context) {
	code := c.Query("code")

	googleConfig := config.SetupConfig()
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		h.logger.Errorf(err.Error())
		log.Fatal(err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		h.logger.Errorf(err.Error())
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.logger.Errorf(err.Error())
		return
	}

	var user entity.UserRegistrationModel
	err = json.Unmarshal(userData, &user)
	if err != nil {
		h.logger.Errorf(err.Error())
	}
	err = h.services.RegistrationByGmail(&user)
	//err = h.userService.RegistrationByGmail(&user)
	if err != nil {
		h.logger.Errorf(err.Error())
	}
	response.ResponseOKWithData(c, userData)
}

func (h *Handler) SignUp(c *gin.Context) {
	var user entity.UserViewModel
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		h.logger.Errorf(err.Error())
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.services.SignUp(&user)
	//err = h.userService.SignUp(&user)

	if err != nil {
		h.logger.Errorf(err.Error())
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response.ResponseOKWithData(c, user)
}
