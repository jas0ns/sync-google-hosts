/*
syncGoogleHosts

for the acccess to google against the Chinese wall

@author JaSonS (JaSonS_toy@foxmail.com)

*/

package main

import (
	"bytes"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	hostsPath = "C:/Windows/System32/drivers/etc/hosts"
	config    = "config.ini"
)

var resAddr string
var startStr string
var endStr string

func GetConfig() {
	c, err := goconfig.LoadConfigFile(config)
	if err != nil {
		log.Fatal(err)
	}

	resAddr, err = c.GetValue("hosts", "url")
	if err != nil {
		log.Fatal(err)
	}

	startStr, err = c.GetValue("hosts", "start")
	if err != nil {
		log.Fatal(err)
	}

	endStr, err = c.GetValue("hosts", "end")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	GetConfig()

	res, err := http.Get(resAddr)
	if err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//remove the head and the tail
	starti := bytes.Index(buff, []byte(startStr))
	endi := bytes.Index(buff, []byte(endStr))
	stepone := buff[starti : endi+len(endStr)]

	//remove all of the <> tags
	var steptwo []byte
	isTag := false
	for _, v := range stepone {
		if isTag {
			if v == '>' {
				isTag = false
			}
			continue
		} else {
			if v == '<' {
				isTag = true
				continue
			}
			steptwo = append(steptwo, v)
		}
	}

	//remove the &nbsp;
	result := bytes.Replace(steptwo, []byte("&nbsp;"), []byte(""), -1)

	err = ioutil.WriteFile(hostsPath, result, 0722)
	if err != nil {
		log.Fatal(err)
	}

}
