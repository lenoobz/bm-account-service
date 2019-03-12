// Package user provides abstraction layer to User core business logic
package user

import (
	"github.com/gin-gonic/gin"
)

// AppRoute defines application's route structure
type AppRoute struct {
	Group  string
	Routes []Route
}

// Route defines a single route, e.g. a human readable name, HTTP method and
// the pattern, the function that will execute when the route is called
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

// InitUserRouter returns a pointer to a mux.Router we can use as a handler
func InitUserRouter(uc *Control) *gin.Engine {
	engine := gin.Default()
	appRoutes := generateRoutes(uc)

	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, appRoute := range appRoutes {
		groupRoute := engine.Group(appRoute.Group)

		for _, route := range appRoute.Routes {
			switch route.Method {
			case "POST":
				groupRoute.POST(route.Pattern, route.HandlerFunc)
			case "GET":
				groupRoute.GET(route.Pattern, route.HandlerFunc)
			}

		}
	}

	return engine
}

func generateRoutes(uc *Control) []AppRoute {
	return []AppRoute{
		AppRoute{
			"/api/v1",
			[]Route{
				Route{
					"POST",
					"/users",
					uc.CreateUser,
				},
				Route{
					"GET",
					"/users",
					uc.GetUsers,
				},
			},
		},
	}
}
