package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"toy/excelutil"
	"toy/pageutil"

	"golang.org/x/net/html"
)

func fetch(url string) {
	//设置请求，10秒超时
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//解析网页
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	htmlBody := string(body)
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("开始解析网页---")
	profile, link := pageutil.FindProfile(doc)
	title := pageutil.FindTitle(doc)
	tags := pageutil.FindVideoTags(doc)

	for _, profileContent := range profile {
		fmt.Println(profileContent)
	}
	if link != nil {
		fmt.Println(link)
	}
	if title != "" {
		fmt.Println("标题: " + title)
	}
	if len(tags) > 0 {
		fmt.Println(tags)
	}
	fmt.Println("解析结束---")

}

func main() {
	reqUrl := "https://www.bilibili.com/video/BV1MH4y1g77V"
	fetch(reqUrl)
	path := ReadJsonConfig()
	excelutil.CreateExcel(path)
	
}
