package main

import (
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	r := routing.New()

	r.Get("/", func(c *routing.Context) error {
		c.Response.SetBodyString("# notification-bot\nEasy-to-use notification bot\n")
		return nil
	})

	log.Fatalln(fasthttp.ListenAndServe(":19071", r.HandleRequest))
}
