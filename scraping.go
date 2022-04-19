package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	url := "https://aws.amazon.com/certification/certified-cloud-practitioner/"

	response, error := http.Get(url)
	defer response.Body.Close()
	check(error)

	if response.StatusCode == 200 {
		fmt.Println("Stauts code :", response.StatusCode)
	}
	by, err := ioutil.ReadAll(response.Body)
	check(err)
	log.Println(string(by))
	//s := string(by)
	//log.Println(s)

	//cont, error := s.Find("id.video-title").Html()
	//check(error)
	//log.Println((cont))

	doc, error := goquery.NewDocumentFromReader(response.Body)
	check(error)
	filter := func(i int, s *goquery.Selection) bool {
		title, _ := s.Attr("href")
		return strings.HasPrefix(title, "/watch")
	}
	var titletext string
	doc.Find("div lb-col lb-tiny-24 lb-mid-18").FilterFunction(filter).Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("href")
		titletext = s.Text()
		//log.Println(titletext)
		_ = title
	})

	i := strings.Index(titletext, `"refinements":[`)
	log.Println(titletext[i:])

	//contant := doc.Find("div.section-hero-home-alt-inner-text text-center").Size()

	//log.Println(contant)
}
