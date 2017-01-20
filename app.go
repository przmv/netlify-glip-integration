package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	n := &NetlifyWebhook{}

	ctx := appengine.NewContext(r)

	if err := decoder.Decode(&n); err != nil {
		log.Errorf(ctx, "Error decoding Netlify Webhook data: %s", err.Error())
		return
	}

	g := &GlipWebhook{
		Icon:  "https://www.netlify.com/img/global/favicon/apple-touch-icon.png",
		Title: n.Name,
	}

	switch n.State {
	case "ready":
		g.Activity = "Deploy succeeded"
		g.Body = fmt.Sprintf("Deployed from branch `%s` to %s", n.Branch, n.DeployURL)
	case "error":
		g.Activity = "Deploy failed"
		g.Body = n.Error
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(g)

	client := urlfetch.Client(ctx)
	res, err := client.Post(os.Getenv("GLIP_HOOK_URL"), "application/json; charset=utf-8", b)
	if err != nil {
		log.Errorf(ctx, "Error sending request to Glip: %s", err.Error())
		return
	}
	defer res.Body.Close()
}
