package server

import (
	"net/http"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/amenities"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/city"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/friendshipevent"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/member"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/office"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/table"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/workspace"
	"github.com/goware/cors"
	"github.com/julienschmidt/httprouter"
)

func Router() http.Handler {
	router := httprouter.New()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	cityHandler := city.Handler{}
	router.GET("/cities", cityHandler.Retrieve)

	amenitiesHandler := amenities.Handler{}
	router.GET("/amenities", amenitiesHandler.Retrieve)

	officeHandler := office.Handler{}
	router.GET("/offices", officeHandler.Retrieve)
	router.GET("/maps", officeHandler.Maps)

	tableHandler := table.Handler{}
	router.GET("/tables", tableHandler.Retrieve)
	router.PATCH("/tables", tableHandler.ChangeStatus)
	router.DELETE("/tables", tableHandler.Delete)
	router.POST("/tables", tableHandler.Create)
	router.PUT("/tables", tableHandler.Update)

	memberHandler := member.Handler{}
	router.GET("/members", memberHandler.Retrieve)
	router.DELETE("/members", memberHandler.Delete)
	router.POST("/members", memberHandler.Create)

	workspaceHandler := workspace.Handler{}
	router.PATCH("/workspaces", workspaceHandler.Update)
	router.DELETE("/workspaces", workspaceHandler.Delete)
	router.POST("/workspaces", workspaceHandler.Assign)

	friendshipeventHandler := friendshipevent.Handler{}
	router.POST("/friendship_events", friendshipeventHandler.Create)
	router.GET("/friendship_events", friendshipeventHandler.Retrieve)

	return cors.Handler(router)
}
