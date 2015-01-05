/*
syncGoogleHosts

for the acccess to google against the Chinese wall

@author JaSonS (JaSonS_toy@foxmail.com)

*/

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	resAddr   = "http://www.360kb.com/kb/2_122.html"
	startStr  = "#google hosts 2015 by 360kb.com"
	endStr    = "#google hosts 2015 end"
	hostsPath = "C:/Windows/System32/drivers/etc/hosts"
)

func main() {
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
