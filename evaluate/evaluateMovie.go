package evaluate

import (
	"context"
	"log"
	"time"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"strings"
	"gamificacion/structures"
	"gamificacion/functions"
)

func Movie(body structures.Body) (correct bool) {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	name, level := functions.ParseFlight(body.Key)
	log.Println("Esperando resultados de " + name + "" + level)

	if !strings.Contains(body.Value, "<xml xmlns=") {
		body.Value = strings.Replace(body.Value, "\n", "~ ", -1)
		r := strings.Split(body.Value, "~ ")
		a := ""
		for i := 0; i < len(r); i++ {
			if !strings.Contains(r[i], "while(") && !strings.Contains(r[i], "for(") && !strings.Contains(r[i], "if(") && !strings.Contains(r[i], "{") && !strings.Contains(r[i], ";") && !strings.Contains(r[i], "else") {
				r[i] = r[i] + ";"
			}
			a = a + r[i]
		}
		body.Value = a
	}

	var example string

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
		chromedp.WaitVisible(`//*[@id="dialogHeader"]`),
		chromedp.Text(`dialogDone`, &example, chromedp.NodeVisible),
	)
	if err != nil {
		return false
	}

	if strings.Contains(example, "Congratulations!") {
		return true
	} else {
		return false
	}
}
