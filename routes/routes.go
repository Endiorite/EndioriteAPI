package routes

import (
	"EndioriteAPI/controllers"
	"EndioriteAPI/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userLink := r.Group("/userLink")
	{
		userLink.Use(middleware.KeyAuth())

		userLink.GET("/check/:userId", controllers.CheckUserLink)
		userLink.POST("/link", controllers.LinkUser)
	}

	playersStats := r.Group("/playersStats")
	{
		playersStats.GET("/getAll", controllers.GetAllPlayersStats)
		playersStats.GET("/get", controllers.GetPlayerStats)
		playersStatsTop := playersStats.Group("/top")
		{
			playersStatsTop.GET("/kills", controllers.GetPlayersStatsTopKills)
			playersStatsTop.GET("/deaths", controllers.GetPlayersStatsTopDeaths)
			playersStatsTop.GET("/killStreak", controllers.GetPlayersStatsTopKillStreak)
			playersStatsTop.GET("/bestKillStreak", controllers.GetPlayersStatsTopBestKillStreak)
			playersStatsTop.GET("/playingTime", controllers.GetPlayersStatsTopPlayingTime)
		}
	}
}
