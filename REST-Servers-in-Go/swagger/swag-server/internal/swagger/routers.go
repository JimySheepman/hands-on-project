package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"DueYearMonthDayGet",
		strings.ToUpper("Get"),
		"/due/{year}/{month}/{day}/",
		DueYearMonthDayGet,
	},

	Route{
		"TagTagnameGet",
		strings.ToUpper("Get"),
		"/tag/{tagname}/",
		TagTagnameGet,
	},

	Route{
		"TaskGet",
		strings.ToUpper("Get"),
		"/task/",
		TaskGet,
	},

	Route{
		"TaskIdDelete",
		strings.ToUpper("Delete"),
		"/task/{id}/",
		TaskIdDelete,
	},

	Route{
		"TaskIdGet",
		strings.ToUpper("Get"),
		"/task/{id}/",
		TaskIdGet,
	},

	Route{
		"TaskPost",
		strings.ToUpper("Post"),
		"/task/",
		TaskPost,
	},

	Route{
		"TaskDeleteAll",
		strings.ToUpper("Delete"),
		"/task/",
		TaskDeleteAll,
	},
}
