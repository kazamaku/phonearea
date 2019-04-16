package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ReJson struct {
	Code string `json:"code"`
	Data ResultData `json:"data"`
	Error string `json:"error"`
}

type ResultData struct {
	Province string `json:"province"`
	City string `json:"city"`
}

func selectPhone(w http.ResponseWriter, r *http.Request) {
	result := initReJson()

	//if r.FormValue("key") != "d502d8e789c6f63d26e2f04e385f3b11" {
	//	result.failure("公钥错误")
	//	result.returnJson(&w)
	//	return
	//}
	
	resp, err := http.Get("http://m.ip138.com/mobile.asp?mobile="+strings.Trim(r.FormValue("phone")," "))

	if err != nil {
		result.failure(err.Error())
		result.returnJson(&w)
		return
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.failure(err.Error())
		result.returnJson(&w)
		return
	}

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil{
		result.failure(err.Error())
		result.returnJson(&w)
		return
	}

	span := dom.Find(".table").Find("span")
	area := span.Eq(1).Text()

	area_arr := strings.Split(string(area), " ")

	if len(area_arr) == 0 || area_arr[0] == "" {
		result.failure("验证手机号有误")
		result.returnJson(&w)
		return
	}

	if len(area_arr) == 1 {
		area_arr = append(area_arr, area_arr[0])
	}

	fmt.Printf("%s  %s \n", area_arr[0], area_arr[1])

	result.success(ResultData{Province: string(area_arr[0]), City: string(area_arr[1])})
	result.returnJson(&w)
}

// 初始化返回的json
func initReJson() *ReJson {
	var result ReJson
	return &result
}

// 返回json
func (result *ReJson) returnJson(w *http.ResponseWriter) {
	jsons, errs := json.Marshal(result) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	(*(w)).Write(jsons)
}

// 处理成功json
func (r *ReJson) success(rd ResultData) {
	r.Code = "200"
	r.Data = rd
	r.Error = ""
}

// 处理失败json
func (r *ReJson) failure(err string) {
	r.Code = "400"
	r.Data = ResultData{}
	r.Error = err
}

func main() {
	port := ":8088"

	if len(os.Args) > 1 {
		port = ":"+os.Args[1]
	}
	
	http.HandleFunc("/phone", selectPhone)
	http.ListenAndServe(port, nil)
}
