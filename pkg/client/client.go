package client

import (
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"time"
)

func Status(shopNumber, orderNumber string) (data string, err error) {
	client := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   180 * time.Second,
	}

	values := url.Values{}
	values.Add("com.livinglogic.cms.apps.search.model.SearchState.search_submit", "true")
	values.Add("searchDescription", shopNumber)
	values.Add("searchDescription2", orderNumber)
	values.Add("searchDescription3", "")
	resp, err := client.PostForm("https://bestellstatus.apollo.de/bestellstatus.htm", values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	parser, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	statusNode := find("div", "bestellstatus ", parser)
	statusText := find("h4", "desc", statusNode)

	data, err = flatten(statusText)
	return
}
