/*    */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/* yahoo news のスポーツのリストページ */
const yahoo_sports_news_list_url string = "http://news.yahoo.co.jp/list/?c=sports&p="

func getList(p int) {

	fmt.Println(yahoo_sports_news_list_url + strconv.Itoa(p))

	doc, err := goquery.NewDocument(yahoo_sports_news_list_url + strconv.Itoa(p))
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	res, err := doc.Find("body").Html()
	if err != nil {
		fmt.Print("dom get failed")
	}
	fmt.Println("dom get success")
	ioutil.WriteFile("./out/lists/p"+strconv.Itoa(p)+".html", []byte(res), os.ModePerm)

}

func doPaging() {

	for i := 10; i <= 100; i++ {
		getList(i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {

	doPaging()

}
