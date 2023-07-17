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

	repos := repository.NewRepository(pgdb, esdb, logger, "anime")
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services, logger)

	//animeRepo := repository.NewAnimeRepository(pgdb, logger)
	//animeEsRepo := repository.NewElasticsearchAnimeRepository(esdb, "anime", logger)
	//animeService := service.NewAnimeService(animeRepo, animeEsRepo, logger)
	//animeHandler := handler.NewAnimeHandler(animeService, logger)

	anime := router.Group("v1/anime")
	{
		//anime.POST("/save", animeHandler.Save)
		//anime.GET("/search", animeHandler.SearchAnime)
		anime.POST("/save", handlers.Save)
		anime.GET("/search", handlers.SearchAnime)
	}

	//userRepo := repository.NewUserRepository(pgdb, logger)
	//userService := service.NewUserService(userRepo, logger)
	//userHandler := handler.NewUserHandler(userService, logger)

	user := router.Group("v1/user")
	{
		user.GET("loginGmail", handlers.LoginGmail)
		user.GET("callback", handlers.CallbackGmail)
		user.POST("/signup", handlers.SignUp)
	}

	return router
}
