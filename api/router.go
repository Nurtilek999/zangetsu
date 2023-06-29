package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"zangetsu/internal/domain/handler"
	"zangetsu/internal/domain/repository"
	"zangetsu/internal/domain/service"
	"zangetsu/pkg/logging"
)

func SetupRouter(pgdb *sql.DB, esdb *elastic.Client, logger logging.Logger) *gin.Engine {
	router := gin.Default()

	animeRepo := repository.NewAnimeRepository(pgdb)
	animeEsRepo := repository.NewElasticsearchAnimeRepository(esdb, "anime")
	animeService := service.NewAnimeService(animeRepo, animeEsRepo, logger)
	animeHandler := handler.NewAnimeHandler(animeService, logger)

	anime := router.Group("v1/anime")
	{
		anime.POST("/save", animeHandler.Save)
		anime.GET("/search", animeHandler.SearchAnime)
	}

	userRepo := repository.NewUserRepository(pgdb)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handler.NewUserHandler(userService, logger)

	user := router.Group("v1/user")
	{
		user.GET("loginGmail", userHandler.LoginGmail)
		user.GET("callback", userHandler.CallbackGmail)
		user.POST("/signup", userHandler.SignUp)
	}

	return router
}
