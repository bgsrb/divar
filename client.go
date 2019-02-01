package divar

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	apiBaseURL = "https://search.divar.ir/json/"
)

type Error struct {
	Code uint `json:"code"`
}

type Response struct {
	*Response
	Result Result `json:"result"`
}

type Result struct {
	*Result
	PostList []Post `json:"post_list"`
}

type Post struct {
	Category  string      `json:"c"`
	Category1 int         `json:"c1"`
	Category2 int         `json:"c2"`
	Category3 int         `json:"c3"`
	Category4 interface{} `json:"c4"`
	Ce        bool        `json:"ce"`
	D         int         `json:"d"`
	Desc      string      `json:"desc"`
	Hc        bool        `json:"hc"`
	Ic        int         `json:"ic"`
	Lm        int         `json:"lm"`
	P         int         `json:"p"`
	P1        interface{} `json:"p1"`
	P2        int         `json:"p2"`
	P3        interface{} `json:"p3"`
	P4        interface{} `json:"p4"`
	Title     string      `json:"title"`
	Token     string      `json:"token"`
	V01       int         `json:"v01"`
	V02       int         `json:"v02"`
	V03       int         `json:"v03"`
	V04       int         `json:"v04"`
	V05       int         `json:"v05"`
	V06       int         `json:"v06"`
	V07       int         `json:"v07"`
	V08       int         `json:"v08"`
	Price     int         `json:"v09"`
	V010      int         `json:"v010`
	V11       int         `json:"v11"`
	V12       int         `json:"v12"`
}

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type Client struct {
	BaseClient *http.Client
	BaseURL    *url.URL
}

func NewClient() *Client {
	baseURL, _ := url.Parse(apiBaseURL)
	c := &Client{
		BaseClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
	return c
}

func (c *Client) execute(request Request, v interface{}) error {
	var jsonBody, _ = json.Marshal(request)
	body := bytes.NewBuffer(jsonBody)
	req, _ := http.NewRequest("POST", apiBaseURL, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := c.BaseClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		//todo handel excpetion
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	return err
}
