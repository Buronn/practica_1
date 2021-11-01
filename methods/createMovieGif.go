package methods

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"

	"gamificacion/functions"
	"gamificacion/structures"
)

func MovieInstagram(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("400: Bad request in SetAnswer.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if reqBody == nil {
		log.Println("204: No content in SetAnswer.")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var body structures.Body
	json.Unmarshal(reqBody, &body)

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	name, level := functions.ParseFlight(body.Key)
	var res [70][]byte
	url := "http://localhost:3000/en/" + name + ".html?lang=en&level=" + level + "&skin=0"
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate("window.localStorage.setItem('" + body.Key + "', '" + body.Value + "')").Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),

		chromedp.Reload(),
		chromedp.MouseClickXY(23, 474, chromedp.ButtonType("left")),
		chromedp.Sleep(1*time.Millisecond),
		chromedp.ActionFunc(func(c context.Context) error {
			for i := 0; i < len(res); i++ {
				log.Println("Sacando screenshot..", i)
				err := chromedp.Screenshot(`display`, &res[i], chromedp.NodeVisible, chromedp.ByID).Do(c)
				if err != nil {
					return err
				}
			}
			return nil
		}),
	)
	functions.GenerateGif(res)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
}
