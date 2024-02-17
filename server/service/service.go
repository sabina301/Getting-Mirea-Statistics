package service

import (
	"context"
	"fmt"
	"getting-statistics-mirea/server/entity"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"time"
)

var countMinus uint8
var countPlus uint8
var countN uint8

type Service struct {
}

func (s *Service) GetResultService(email string, password string, number string) *entity.Statistics {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var url string

	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://attendance-app.mirea.ru/"),
		chromedp.Click(".ant-btn.css-1qhpsh8.ant-btn-primary"),
		chromedp.Sleep(2*time.Second),
		chromedp.Location(&url),
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Идет загрузка ... ")

	var nodes []*cdp.Node

	if err2 := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
		chromedp.SetValue(`input[type="email"]`, email, chromedp.ByQuery),
		chromedp.SetValue(`input[type="password"]`, password, chromedp.ByQuery),
		chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Location(&url),
		chromedp.Click(".ant-btn.css-1qhpsh8.ant-btn-primary"),
		chromedp.Sleep(2*time.Second),
		chromedp.Location(&url),
		chromedp.Click(".ant-space.css-1qhpsh8.ant-space-horizontal.ant-space-align-center.ant-space-gap-row-small.ant-space-gap-col-small .ant-space-item:nth-child(2)"),
		chromedp.Sleep(1*time.Second),
		chromedp.Nodes(".ant-spin-container li a a", &nodes),
	); err2 != nil {
		log.Fatal(err2)
	}

	var userNodes []*cdp.Node
	countPlus, countMinus, countN = 0, 0, 0
	for _, node := range nodes {
		hr, _ := node.Attribute("href")

		if err3 := chromedp.Run(ctx,
			chromedp.Navigate("https://attendance-app.mirea.ru"+hr),
			chromedp.Nodes(".ant-table-row.ant-table-row-level-0", &userNodes),
		); err3 != nil {
			log.Fatal(err3)
		}
		num, _ := strconv.Atoi(number)

		key, _ := userNodes[num].Attribute("data-row-key")

		var tdTexts []string
		if err := chromedp.Run(ctx,
			chromedp.Evaluate(fmt.Sprintf(`Array.from(document.querySelectorAll('tr.ant-table-row.ant-table-row-level-0[data-row-key="%s"] .ant-table-cell[style="text-align: center;"]')).map(td => td.textContent)`, key), &tdTexts),
		); err != nil {
			log.Fatal(err)
		}

		for _, tdText := range tdTexts {
			if tdText == "-" {
				countMinus++
			} else if tdText == "+" {
				countPlus++
			} else if tdText == "Н" {
				countN++
			}
		}

	}
	log.Println("Посещений = ", countPlus, "Пропусков = ", countMinus, "По уважительной = ", countN)
	stat := entity.Statistics{
		Plus:  countPlus,
		Minus: countMinus,
		N:     countN,
	}
	return &stat

}
func NewService() *Service {
	return &Service{}
}
