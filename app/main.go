package main

import "github.com/kataras/iris"

func main() {
		// iris.Favicon("./favicon.ico")
		iris.Config.IsDevelopment = true
		iris.Get("/", func(ctx *iris.Context) {
		    ctx.Render("index.html",  struct{ Name string }{Name: "iris"})
				// ctx.Render("text/plain", "hello world")
		})

		// iris.Get("/login", func(ctx *iris.Context) {
		//     ctx.Render("login.html", iris.Map{"Title": "Login Page"})
		// })

		// iris.Post("/login", func(ctx *iris.Context) {
		//     secret := ctx.PostValue("secret")
		//     ctx.Session().Set("secret", secret)

		//     ctx.Redirect("/user")
		// })

		iris.Listen(":8080")
}