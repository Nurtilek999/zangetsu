package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"

	"log"
	"net/http"
	"zangetsu/internal/domain/entity"
	"zangetsu/internal/domain/service"
	"zangetsu/pkg/config"
	"zangetsu/pkg/response"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	var userHandler = UserHandler{}
	userHandler.userService = userService
	return &userHandler
}

func (h *UserHandler) LoginGmail(c *gin.Context) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate") //oauth2.AccessTypeOffline
	c.Redirect(http.StatusFound, url)
}

func (h *UserHandler) CallbackGmail(c *gin.Context) {
	code := c.Query("code")

	googleConfig := config.SetupConfig()
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(c.Writer, "User data fetch failed")
		return
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(c.Writer, "Json parsing failed")
		return
	}

	var user entity.UserRegistrationModel
	err = json.Unmarshal(userData, &user)
	if err != nil {
		fmt.Fprintln(c.Writer, "Encoding failed")
	}
	err = h.userService.RegistrationByGmail(&user)
	if err != nil {
		fmt.Fprintln(c.Writer, "Error in service layer: "+err.Error())
	}
	fmt.Fprintln(c.Writer, string(userData))
	response.ResponseOKWithData(c, userData)
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var user entity.UserViewModel
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.userService.SignUp(&user)
	
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response.ResponseOKWithData(c, user)
}
