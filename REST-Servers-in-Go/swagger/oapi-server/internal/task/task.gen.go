package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// Task defines model for Task.
type Task struct {
	Due  *time.Time `json:"due,omitempty"`
	Id   *int       `json:"id,omitempty"`
	Tags *[]string  `json:"tags,omitempty"`
	Text *string    `json:"text,omitempty"`
}

// PostTaskJSONBody defines parameters for PostTask.
type PostTaskJSONBody struct {
	Due  *time.Time `json:"due,omitempty"`
	Tags *[]string  `json:"tags,omitempty"`
	Text *string    `json:"text,omitempty"`
}

// PostTaskJSONRequestBody defines body for PostTask for application/json ContentType.
type PostTaskJSONRequestBody PostTaskJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get tasks with given due date
	// (GET /due/{year}/{month}/{day})
	GetDueYearMonthDay(ctx echo.Context, year int, month int, day int) error
	// Get tasks with given tag name
	// (GET /tag/{tagname})
	GetTagTagname(ctx echo.Context, tagname string) error
	// Returns a list of all tasks
	// (GET /task)
	GetTask(ctx echo.Context) error
	// Create a task
	// (POST /task)
	PostTask(ctx echo.Context) error
	// Delete task with specific id
	// (DELETE /task/{id})
	DeleteTaskId(ctx echo.Context, id int) error
	// Delete all tasks
	// (DELETE /task/
	DeleteAllTasks(ctx echo.Context) error
	// Get task with specific id
	// (GET /task/{id})
	GetTaskId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDueYearMonthDay converts echo context to params.
func (w *ServerInterfaceWrapper) GetDueYearMonthDay(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "year" -------------
	var year int

	err = runtime.BindStyledParameter("simple", false, "year", ctx.Param("year"), &year)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter year: %s", err))
	}

	// ------------- Path parameter "month" -------------
	var month int

	err = runtime.BindStyledParameter("simple", false, "month", ctx.Param("month"), &month)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter month: %s", err))
	}

	// ------------- Path parameter "day" -------------
	var day int

	err = runtime.BindStyledParameter("simple", false, "day", ctx.Param("day"), &day)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter day: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDueYearMonthDay(ctx, year, month, day)
	return err
}

// GetTagTagname converts echo context to params.
func (w *ServerInterfaceWrapper) GetTagTagname(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tagname" -------------
	var tagname string

	err = runtime.BindStyledParameter("simple", false, "tagname", ctx.Param("tagname"), &tagname)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tagname: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTagTagname(ctx, tagname)
	return err
}

// GetTask converts echo context to params.
func (w *ServerInterfaceWrapper) GetTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTask(ctx)
	return err
}

// PostTask converts echo context to params.
func (w *ServerInterfaceWrapper) PostTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostTask(ctx)
	return err
}

// DeleteTaskId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTaskId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTaskId(ctx, id)
	return err
}

func (w *ServerInterfaceWrapper) DeleteAllTasks(ctx echo.Context) error {
	return w.Handler.DeleteAllTasks(ctx)
}

// GetTaskId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTaskId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTaskId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/due/:year/:month/:day", wrapper.GetDueYearMonthDay)
	router.GET(baseURL+"/tag/:tagname", wrapper.GetTagTagname)
	router.GET(baseURL+"/task", wrapper.GetTask)
	router.POST(baseURL+"/task", wrapper.PostTask)
	router.DELETE(baseURL+"/task/:id", wrapper.DeleteTaskId)
	router.DELETE(baseURL+"/task", wrapper.DeleteAllTasks)
	router.GET(baseURL+"/task/:id", wrapper.GetTaskId)

}
