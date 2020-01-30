// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"iohttps.com/nqq/go-review-training/exercise/ch4/github"
)

//!+
func main() {
	var recentMonthIss []*github.Issue
	var recentYearIss []*github.Issue
	var overYearIss []*github.Issue
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if IsRecentMonth(item.CreatedAt) {
			recentMonthIss = append(recentMonthIss, item)
			continue
		}
		if IsRecentYear(item.CreatedAt) {
			recentYearIss = append(recentYearIss, item)
			continue
		}
		overYearIss = append(overYearIss, item)
	}
	for _, item := range recentMonthIss {
		fmt.Println("近一个月的提问")
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	for _, item := range recentYearIss {
		fmt.Println("近一年的提问")
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	for _, item := range overYearIss {
		fmt.Println("超过一年的提问")
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

}

func IsRecentMonth(publish time.Time) bool {
	if publish.Month() < 12 {
		afterMonth := time.Date(publish.Year(), publish.Month()+1, publish.Day(), publish.Hour(), publish.Minute(), publish.Second(), 0, publish.Location())
		return time.Now().Before(afterMonth)
	}

	afterMonth := time.Date(publish.Year()+1, 1, publish.Day(), publish.Hour(), publish.Minute(), publish.Second(), 0, publish.Location())
	return time.Now().Before(afterMonth)
}

func IsRecentYear(publish time.Time) bool {
	afterYear := time.Date(publish.Year()+1, publish.Month(), publish.Day(), publish.Hour(), publish.Minute(), publish.Second(), 0, publish.Location())
	return time.Now().Before(afterYear)
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
