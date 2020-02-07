package spider

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func Spider(url string) {
	client := http.DefaultClient

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("user-agent", "Mobile")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, err := client.Do(req)
	fmt.Println(resp.Request.Header)
	if err != nil {
		log.Println(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}
	defer resp.Body.Close()

}
