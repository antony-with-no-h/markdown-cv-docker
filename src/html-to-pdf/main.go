// the example: github.com/chromedp/examples/blob/master/pdf/main.go was
// exactly what I wanted so all credit to them

package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)

	// create context
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	fileIn := printEnv("CV_FILE", "file:///cv/cv.html")
	fileOut := printEnv("CV_PDF", "/cv/cv.pdf")

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(fileIn, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fileOut, buf, 0o644); err != nil {
		log.Fatal(err)
	}
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func printEnv(name string, valDefault string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return valDefault
	}
	return val
}
