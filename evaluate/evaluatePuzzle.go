package evaluate

import (
	"context"
	"log"
	"strings"
	"time"

	"gamificacion/functions"
	"gamificacion/structures"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func Puzzle(body structures.Body) (correct bool) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	name, level := functions.ParseFlight(body.Key)
	log.Println("Esperando resultados de " + name + "" + level)
	var example string
	err := chromedp.Run(ctx,
		Entra(),
		chromedp.Navigate(`http://localhost:3000/en/puzzle.html?lang=en`),
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
		chromedp.Click(`secondary`, chromedp.NodeVisible),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Click(`checkButton`, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Text(`answerMessage`, &example, chromedp.NodeVisible),
	)
	if err != nil {
		log.Println("no entra a chromedp")
		return false
	}

	if strings.Contains(example, "Perfect! All 16 blocks are correct.") {
		return true
	} else {
		return false
	}

}
func Entra() chromedp.Tasks {
	log.Println("Entra al chromedp")
	return chromedp.Tasks{}

}
