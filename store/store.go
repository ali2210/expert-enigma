package store

import (
	"errors"
	"reflect"
)

// Request mode specified normal or internal routing
// 0 => default
// 1 => internal
type request_mode int

// Render Route specified web routing information
// Router Handle Name  => "home"
// RLink => "/home"
// tmode => normal or internal

type RenderRoutes struct {
	Route_Handler_Name string
	RLink              string
	tmode              request_mode
}

// webroute routing continous memory
var webroute []RenderRoutes

// continous memory allocation size
var SIZE_OF int64

// Routing object created
func SetRouteParams(handleRoute, rrlink string) RenderRoutes {
	return RenderRoutes{Route_Handler_Name: handleRoute, RLink: rrlink, tmode: 0}
}

// routing objects
func SetRoute(rr []RenderRoutes) {
	webroute = make([]RenderRoutes, SIZE_OF)
	webroute = append(webroute, rr...)
}

// get routing object
func GetRoute() []RenderRoutes { return webroute }

// map routing return route rules
func MapRoute(route string) RenderRoutes {

	for r := range GetRoute() {
		if reflect.DeepEqual(route, GetRoute()[r].Route_Handler_Name) {
			return GetRoute()[r]
		}
	}
	return RenderRoutes{}
}

func (r RenderRoutes) Validate() error {

	for i := range GetRoute() {
		if reflect.DeepEqual(r, GetRoute()[i]) {
			return nil
		}
	}
	return errors.New("invalid route")
}
