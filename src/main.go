package main

import "github.com/nsheremet/banjo"

func main() {
	app := banjo.Create(banjo.DefaultConfig())

	app.Get("/foo", func(ctx *banjo.Context) {
		ctx.HTML("<h1>Hello from BONjO!</h1>")
	})

	app.Run()
}
