package spider

import (
	"fmt"
	"log"
	"net/http"
	"smh-api/models"

	"github.com/PuerkitoBio/goquery"
)

type DSD struct {
	Href     string
	Name     string
	Cover    string
	Years    string
	Region   string
	Genre    string
	Director string
	Actor    string
}

//https://www.diaosidao.net/search.php?searchword=%E8%AF%AF%E6%9D%80
var url = "https://www.diaosidao.net/search.php?searchword="

func Spider(key string) []*models.Movie {

	fmt.Println(url)
	resp, err := http.Get(url + key)
	if err != nil {
		log.Println(err)
	}
	// b, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}
	serachResult := make([]*DSD, 0)
	doc.Find(".col-lg-wide-75 .stui-pannel_bd ul li").Each(func(i int, s *goquery.Selection) {

		s.Find("div .detail").Each(func(j int, s1 *goquery.Selection) {

			fmt.Println(s1.ChildrenFiltered(`h3`).Text())
			s1.Find("p").Each(func(k int, s2 *goquery.Selection) {
				fmt.Println(s2.Text())
				fmt.Println(s1.Find("a(0)").AttrOr("href", ""))
			})

			serachResult = append(serachResult, &DSD{
				Name: s.ChildrenFiltered(`.h3`).Text(),
			})
		})

	})
	return nil
}
