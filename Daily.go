package main

import (
	"github.com/bmob/bmob-go-sdk"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"time"
	"log"
)

var (
	appConfig = bmob.RestConfig{"913a4de2d6622184c2f2167742c7a73a",
				    "4e99a1565b33ebd1887ac0450df50853"}
)

type IcibaSentence struct {
	Sid         string `json:"sid"`
	Tts         string `json:"tts"`
	Content     string `json:"content"`
	Note        string `json:"note"`
	Love        string `json:"love"`
	Translation string `json:"translation"`
	Picture     string `json:"picture"`
	Picture2    string `json:"picture2"`
	Caption     string `json:"caption"`
	Dateline    string `json:"dateline"`
	FenxiangImg string `json:"fenxiang_img"`
}

type Sentence struct {
	Sid         string `json:"sid"`
	Tts         string `json:"tts"`
	Content     string `json:"content"`
	Note        string `json:"note"`
	Love        string `json:"love"`
	Translation string `json:"translation"`
	Picture     string `json:"picture"`
	Picture2    string `json:"picture2"`
	Caption     string `json:"caption"`
	Dateline    string `json:"dateline"`
	ShareImg    string `json:"shareImg"`
}

func httpGet() {
	date := time.Now().Local().Format("2006-01-02")
	resp, err := http.Get("http://open.iciba.com/dsapi/?date=" + date)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	icibaSentence := &IcibaSentence{}
	json.Unmarshal(body, &icibaSentence)
	fmt.Println(icibaSentence.FenxiangImg)
	sentence := &Sentence{}
	sentence.Sid = icibaSentence.Sid
	sentence.Tts = icibaSentence.Tts
	sentence.Content = icibaSentence.Content
	sentence.Note = icibaSentence.Note
	sentence.Love = icibaSentence.Love
	sentence.Translation = icibaSentence.Translation
	sentence.Picture = icibaSentence.Picture
	sentence.Picture2 = icibaSentence.Picture2
	sentence.Caption = icibaSentence.Caption
	sentence.Dateline = icibaSentence.Dateline
	sentence.ShareImg = icibaSentence.FenxiangImg
	fmt.Println(sentence.Content)
	data, err := json.Marshal(sentence)
	if err != nil {
		return
	}
	header, err := bmob.DoRestReq(appConfig, bmob.RestRequest{
		BaseReq: bmob.BaseReq{
			Method: "POST",
			//Path:   bmob.ApiRestURL("DailySentence"),
			Path:  bmob.ApiRestURL("DailySentenceNet") + "/",
			Token: ""},
		Type: "application/json",
		Body: data}, &sentence)
	if err != nil {
		log.Panic(err)
		log.Println("error data:", date)
	} else {
		log.Println(header)
		log.Println("success data:", date)
	}
}

func main() {
	//next := time.Now().Add(time.Second * 5)
	next := time.Now().Add(time.Hour * 9)
	fmt.Println(next)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	fmt.Println(next)
	timer := time.NewTimer(next.Sub(time.Now()))
	<-timer.C
	httpGet()
	//ticker := time.NewTicker(5 * time.Second)
	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			go httpGet()
		}
	}
}
