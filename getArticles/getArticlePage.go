package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getUrl(page int) []string {

	var urls []string

	fileInfos, _ := ioutil.ReadFile("./out/lists/p" + strconv.Itoa(page) + ".html")
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		fmt.Print("url scarapping failed")
	}

	// リストから記事ページURLを取得する
	doc.Find(".list > li >a ").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		// fmt.Println(url)
		urls = append(urls, url)
	})
	return urls

}

func getArticle(url string) {

	// html取得
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	// res, err := doc.Find("body").Html()
	// if err != nil {
	// 	fmt.Print("dom get failed")
	// }
	// fmt.Println("dom get success")
	// doc, err := goquery.NewDocumentFromReader(res)
	// ioutil.WriteFile("./out/article/1", []byte(res), os.ModePerm)
	doc.Find(".headline > .headlineTxt ").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Find(".newsTitle").Text())
		fmt.Println(s.Find(".hbody").Text())

	})
	// fmt.Println([]byte(res))
}

func main() {

	// 記事ページのurlリスト作成
	// var urls []string
	// for i := 10; i <= 100; i++ {
	// 	urls = append(urls, getUrl(i)...)
	// }
	// fmt.Println(urls)

	// 1ページごと、スクレイピングする
	getArticle("http://news.yahoo.co.jp/pickup/6203247")
}
