package main

import (
	"fmt"
	"net/http"
	"strings"
)

var input = func() []string {
	return []string{"www.abc.com",
		"www.testme.com",
		"www.abc.com",
		"www.abc.com",
		"www.testme.com",
		"www.test2.com",
		"www.testok.com",
		"www.testme.com",
		"www.abc.com"}

}
var responseFormat = func(domain string, position, status int) string {
	const msgBody = "%s - %d {status: %d, message: %s}"
	if status == http.StatusOK {
		return fmt.Sprintf(msgBody, domain, position, status, "OK")
	}
	return fmt.Sprintf(msgBody, domain, position, status, "Too many requests")
}

func main() {
	fmt.Println(strings.Join(solve(input()), "\n\n"))

}
func solve(requests []string) []string {
	type response struct {
		lastIndex int
		count     int
	}
	result := make([]string, len(requests))

	requestMap := make(map[string]response)
	maxRequestLimit := 2
	maxIntervalLimit := 5
	for i, r := range requests {
		_, found := requestMap[r]
		if !found {
			requestMap[r] = response{lastIndex: i, count: 1}
			result[i] = responseFormat(r, i+1, 200)
		} else {
			resp := requestMap[r]
			resp.count++
			if resp.count > maxRequestLimit && i-resp.lastIndex < maxIntervalLimit {
				result[i] = responseFormat(r, i+1, 429)
			} else {
				result[i] = responseFormat(r, i+1, 200)
			}

			resp.lastIndex = i
			requestMap[r] = resp

		}
	}
	return result
}

/*
0 - www.abc.com {status: 200, message: OK}
1 - www.testme.com {status: 200, message: OK}
2 - www.abc.com {status: 200, message: OK}
3 - www.abc.com {status: 429, message: Too many requests}
4 - www.testme.com {status: 200, message: OK}
5 - www.test2.com {status: 200, message: OK}
6 - www.testok.com {status: 200, message: OK}
7 - www.testme.com {status: 429, message: Too many requests}
8 - www.abc.com {status: 200, message: OK}
*/
