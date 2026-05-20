package main

import (
	"fmt"
	"strings"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMIddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello",c.HelloCount),"World!")
}

func main(){
	router := web.New(AppContext{}).

	// Create your router

	Middleware(web.LoggerMiddleware).

	// Use some included middleware


}
