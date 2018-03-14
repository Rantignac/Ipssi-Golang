package model

import (
	"net/http"
)

// Route represents one route.
type Route struct {
	Name    string
	Method  string
	Pattern string
	Func    http.HandlerFunc
}

// Routes is a an array of route.
type Routes []Route
