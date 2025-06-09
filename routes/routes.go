package routes

import (
	"EndioriteAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
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
