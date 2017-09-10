package main

import (
	"fmt"
	"github.com/dsparling/go-apache-log-parser"
	"log"
	"regexp"
	"sort"
	"strings"
)

// A data structure to hold a key/value pair.
type pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by Value.
type pairList []pair

func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	uniqueUrls := make(map[string]int)
	lines, err := apachelogparser.Parse("logs/access_log")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		re, err := regexp.Compile(`([^?=&]+)(=([^&]*))?`)
		if err != nil {
			log.Fatal(err)
		}
		res := re.FindAllStringSubmatch(line.Url, -1)
		if len(res) > 0 {
			url := res[0][0]
			if skipURL(url) {
				continue
			}
			uniqueUrls[url]++
		} else {
			continue
		}
	}
	fmt.Println(len(uniqueUrls))
	fmt.Println("\n##### Sorted by link #####")
	// Store the keys in slice in sorted order
	var keys []string
	for k := range uniqueUrls {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// Print in keys alpha order
	for _, k := range keys {
		fmt.Println(k, "(", uniqueUrls[k], ")")
	}

	fmt.Println("\n##### Sorted by value #####")
	sortedUniqueUrlsByValue := sortMapByValue(uniqueUrls)
	for _, v := range sortedUniqueUrlsByValue {
		fmt.Println(v.Key, "(", v.Value, ")")
	}
}

// A function to turn a map into a pairList, then sort and return it.
// Andrew Gerrand: https://groups.google.com/d/msg/golang-nuts/FT7cjmcL7gw/Gj4_aEsE_IsJ
func sortMapByValue(m map[string]int) pairList {
	p := make(pairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func skipURL(url string) bool {
	// Do any filtering you need here
	if strings.HasSuffix(url, ".js") ||
		strings.HasSuffix(url, ".css") ||
		strings.HasSuffix(url, ".html") ||
		strings.HasSuffix(url, ".txt") ||
		strings.HasSuffix(url, ".php") ||
		strings.HasSuffix(url, ".asp") ||
		strings.HasSuffix(url, ".cgi") ||
		strings.HasSuffix(url, ".xml") ||
		strings.HasSuffix(url, ".com") ||
		strings.HasSuffix(url, ".gif") ||
		strings.HasSuffix(url, ".png") ||
		strings.HasSuffix(url, ".jpg") ||
		strings.HasSuffix(url, ".ico") {
		return true
	}
	return false
}
