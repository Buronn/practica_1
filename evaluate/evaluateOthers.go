package evaluate

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"gamificacion/functions"
	"gamificacion/structures"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func Others(body structures.Body) (correct bool, bleft int) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	name, level := functions.ParseFlight(body.Key)
	log.Println("Esperando resultados de " + name + "" + level)
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
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
	var blocks string
	url := "http://localhost:3000/en/" + name + ".html?lang=en&level=" + level + "&skin=0"
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate("window.localStorage.setItem('" + body.Key + "', '" + body.Value + "')").Do(ctx)
			log.Println("LINEA 44 AnctionFunc")
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Reload(),
		haveBlockLimiter(name, level, &blocks),
		chromedp.Click(`runButton`, chromedp.NodeVisible),
		chromedp.MouseClickXY(164, 519, chromedp.ButtonType("left")),
		chromedp.WaitVisible(`//*[@id="dialogHeader"]`),
		chromedp.Text(`dialogDone`, &example, chromedp.NodeVisible),
	)
	if err != nil {
		log.Println("ERROR LINEA 61")
		return false, -1
	}

	i, _ := strconv.Atoi(blocks)
	if strings.Contains(example, "Congratulations!") {
		log.Println("Return LINEA 67")
		return true, i
	} else {
		log.Println("Return LINEA 70")
		return false, i
	}
}
func haveBlockLimiter(name string, level string, blocks *string) chromedp.Tasks {
	if name == "maze" && level != "1" && level != "2" {
		log.Println("LINEA 73 haveBlockLimiter")
		return chromedp.Tasks{
			chromedp.Text(`capacityNumber`, blocks, chromedp.NodeVisible),
		}
	} else {
		log.Println("LINEA 77 haveBlockLimiter")
		return chromedp.Tasks{}
	}
}
