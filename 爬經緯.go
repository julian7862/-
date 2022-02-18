package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	port = 8080
)

type restaurant struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
	Rate    string `json:"rate,omitempty"`
}

func main() {

	opts := []selenium.ServiceOption{
		// Enable fake XWindow session.
		// selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr), // Output debug information to STDERR
	}

	// Enable debug info.
	// selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService("c:/Users/julia/Desktop/portfolio/爬經緯/chromedriver.exe", port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	wd.Get("https://www.google.com.tw/maps/")
	newfile, err := os.OpenFile("./LatLong.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer newfile.Close()
	w := csv.NewWriter(newfile)

	file, err := os.Open("restaurant0820_2.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(record[1]) // record has the type []string

		//找到输入框id
		we, err := wd.FindElement(selenium.ByID, "searchboxinput")
		if err != nil {
			panic(err)
		}
		//清空
		err = we.Clear()
		if err != nil {
			panic(err)
		}
		//向输入框发送“”
		err = we.SendKeys(record[1])
		if err != nil {
			panic(err)
		}
		//找到提交按钮id
		wclick, err := wd.FindElement(selenium.ByID, "searchbox-searchbutton")
		if err != nil {
			panic(err)
		}
		//点击提交
		err = wclick.Click()
		if err != nil {
			panic(err)
		}
		time.Sleep(8 * time.Second)
		wurl, err := wd.CurrentURL()
		r, err := regexp.Compile("@([0-9]+).([0-9]+),([0-9]+).([0-9]+)")
		rFind := r.FindString(wurl)
		println(record[0] + " " + rFind)

		var newrecord []string
		newrecord = append(newrecord, record[0])
		newrecord = append(newrecord, record[1])
		newrecord = append(newrecord, strings.Replace(rFind, "@", "", 1))
		w.Write(newrecord)
		w.Flush()
	}

	// langs := []string{
	// 	"西子灣",
	// 	"台北101",
	// }

	// for _, obj := range langs {

	// }

	time.Sleep(10 * time.Second)

}
