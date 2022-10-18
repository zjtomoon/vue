package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetRequest(url string) string {
	resp,err := http.Get(url)
	if err != nil {
		log.Println("err = ",err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err = ",err)
	}
	return string(body)
}

func PostRequest(url string,requestBody map[string]interface{}) string {
	requestStr := ""
	count := 0
	for key,value := range requestBody {
		count++
		if len(requestBody) == count {
			requestStr += key + "=" + cast.ToString(value)
		}else {
			requestStr += key + "=" + cast.ToString(value) + "&"
		}
	}

	resp,err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(requestStr))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err = ",err)
	}
	return string(body)
}

func GetUrlParam(url string,key string) string {
	temp := strings.Split(url,"&")
	for _,n := range temp {
		arr := strings.Split(n,"=")
		if arr[0] == key {
			return arr[1]
		}
	}
	return ""
}