package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	tk := time.NewTicker(1 * time.Second)
	for range tk.C {
		request()
	}
}

// {"content-type":"text/xml","data":"PHhtbD48VG9Vc2VyTmFtZT48IVtDREFUQVt3dzQ4ZmIyMWVhYjVjYzg4MDJdXT48L1RvVXNlck5hbWU+PEVuY3J5cHQ+PCFbQ0RBVEFbcXNvT1c5WEZJSFdoSENab3JpTXlKZzQ2UDFlTUgvV3Raclo1VlRPNngyWHc2bS9zdDMwOUdnMmY1TUtnV1crcUtWbFAxelU2Y3RlMlI5c1hZV0p5YXVhMEFJZ01NYWVhS0c3aEZ6b3RIcGNSRTBxcVl5cjBMcWhCQ0dCeUlESmJKQVl1aUh2VHpDL0pDc1pUK1lmNHFCNk0wK3V0eXozNkQ2TExrcG9ncjIxTHJzU1Z0d2ZMNldsRjh5SGlIWHFOMlBVYi83dmR5NXd6dXN3bC83OTlSNlR3amVvdVNDTTgvVVpkMG4xSFFvMEJrUXJrVWRudWY1TTh6MzZUYnVaVjAvNVRnVk9DMXBXTVAxT2ZGNjdqVzJLNjFJQXlIVlVaTlE3eTVXanY5ZGRvbGQ3T1FQYk1BeExUTSs2YVVOWEUyajZFY2VYL1FHOGhaeGdUYlY3YjZuWkF6N2wyMUNJV1YzTzFJSFBBWnN4d29IZ3JOQ1FUMkFENGRid1BzRG5rRFdzdHJZdTRVUUhnMlJPUFJFSTVMUUhsMEdQWHRVcTAybEdsbGQ0WWV6S0dDYjlRSjFMQ1N5dlIwZFJzUmJ2cWxUZWtPaUIzT2JsWVVCbXNDWjZQTlhJMjR3Nk1HSFJiOUUrUVBOU0wydGk1L2RDcGQySnBodW5VZTdWeWRJOVdHeklBWkVkQzVwblE1dExnNkgxZFJvNHhrUFhNblU2SHdYMGZlaWErL244PV1dPjwvRW5jcnlwdD48QWdlbnRJRD48IVtDREFUQVtdXT48L0FnZW50SUQ+PC94bWw+","method":"POST"}
func request() {
	req, _ := http.NewRequest(http.MethodGet, "http://wxwork.clearcode.cn/pull", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var dict = make(map[string]string)
	json.Unmarshal(data, &dict)

	a, _ := http.NewRequest(dict["method"], "http://localhost:8080/callback/addressbook", strings.NewReader(dict["data"]))
	a.Header.Add("Content-Type", dict["content-type"])
	aa, _ := http.DefaultClient.Do(a)
	if aa != nil {
		aa.Body.Close()
	}
}
