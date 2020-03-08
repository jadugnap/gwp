package main

import (
	"github.com/go-martini/martini"
	"github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-martini/post"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/post/:id", post.Retrieve, post.HandleGet)
	m.Post("/post", binding.Json(post.Post{}), post.HandlePost)
	m.Put("/post/:id", post.Retrieve, post.HandlePut)
	m.Delete("/post/:id", post.Retrieve, post.HandleDelete)
	m.Run()

}
