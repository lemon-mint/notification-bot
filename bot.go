package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/lemon-mint/godotenv"
	"github.com/lemon-mint/notification-bot/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

var addr string = ":19071"
var telegramAPIServerPrefix string

func main() {
	godotenv.Load()

	client, err := ent.Open("sqlite3", "file:test.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Error: %v", err)
	}

	r := router.New()
	telegramAPIServerPrefix = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_API_KEY")
	var telegramWebhookEndPoint string = base64Rand(32)

	r.POST("/tghook/"+telegramWebhookEndPoint, telegramWebhook)

	r.GET("/", func(c *fasthttp.RequestCtx) {
		c.Response.SetBodyString("# notification-bot\nEasy-to-use notification bot\n")
	})
	log.Println("Starting Server On", addr)
	PUBLIC_URL := os.Getenv("PUBLIC_URL")
	var WEBHOOK_URL string
	if len(PUBLIC_URL) == 0 {
		log.Fatalln("PUBLIC_URL must be set")
	}
	if PUBLIC_URL[len(PUBLIC_URL)-1] == '/' {
		WEBHOOK_URL = PUBLIC_URL + "tghook/" + telegramWebhookEndPoint
	} else {
		WEBHOOK_URL = PUBLIC_URL + "/tghook/" + telegramWebhookEndPoint
	}
	err = tgSetWebhook(WEBHOOK_URL, []string{"message"})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Webhook created.")
		log.Println("Webhook URL:", WEBHOOK_URL)
	}
	log.Fatalln(fasthttp.ListenAndServe(addr, r.Handler))
}

func telegramWebhook(c *fasthttp.RequestCtx) {
	log.Printf("%s", c.Request.Body())
	e := new(Update)
	err := json.Unmarshal(c.Request.Body(), e)
	if err != nil {
		log.Println("Error :", err, "Body :", string(c.Request.Body()))
		return
	}
	ExecuteUserCommand(e, c.ConnID())
}

func base64Rand(size int) string {
	buf := make([]byte, size)
	rand.Read(buf)
	return base64.RawURLEncoding.EncodeToString(buf)
}

func RequestJSONPost(URL string, payload interface{}) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req.SetBody(data)
	req.Header.Add("Content-Type", "application/json")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(URL)
	resp := fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)
	return resp, err
}

func ExecuteUserCommand(e *Update, ID uint64) {
	if e.Message != nil {
		if e.Message.Text == "/ping" {
			r, err := RequestJSONPost(telegramAPIServerPrefix+"/sendMessage", SendMessage{
				Method: "sendMessage",
				ChatID: e.Message.Chat.ID,
				Text:   "pong!",
			})
			defer fasthttp.ReleaseResponse(r)
			if err != nil {
				log.Println("Error :", err, "Event :", e)
				return
			} else {
				log.Println(string(r.Body()))
			}
		}
	}
}
