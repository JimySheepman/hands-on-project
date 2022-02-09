package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"routing/gorilla"
)

const port = 8080

func main() {
	if len(os.Args) < 2 || routers[os.Args[1]] == nil {
		fmt.Fprintf(os.Stderr, "usage: go-routing router\n\n")
		fmt.Fprintf(os.Stderr, "router is one of: %s\n", strings.Join(routerNames, ", "))
		os.Exit(1)
	}
	routerName := os.Args[1]
	router := routers[routerName]

	fmt.Printf("listening on port %d using %s router\n", port, routerName)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

var routers = map[string]http.Handler{

	"gorilla": gorilla.Serve,
}

var routerNames = func() []string {
	routerNames := []string{}
	for k := range routers {
		routerNames = append(routerNames, k)
	}
	sort.Strings(routerNames)
	return routerNames
}()
