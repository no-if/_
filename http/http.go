package http

import (
	"bytes"
	"encoding/hex"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	. "net/url"
	"os"
	"strings"
)

var (
	DefaultTransport = NewTransport()
	Multipart_first  = "^"
	File_name        = "@"
	File_date        = "$"
)

type Response struct {
	StatusCode int
	Location   string
	Cookie     map[string]string
	Cookies    string
	Header     map[string]string
	Body       []byte
	Error      error
}

func (self *transport) do(method, url, data string, header map[string]string) (response Response) {
	var (
		body = new(bytes.Buffer)
		ct   = "text/html" //get
	)

	if method == http.MethodPost && data != "" {
		var query Values

		switch string(data[0]) {
		default:
			if query, response.Error = ParseQuery(data); response.Error != nil {
				return
			}
			body.WriteString(query.Encode())
			ct = "application/x-www-form-urlencoded"
		case "<":
			body.WriteString(data)
			ct = "text/xml"
		case "[", "{":
			body.WriteString(data)
			ct = "application/json"
		case Multipart_first:
			//form-data
			var mw = multipart.NewWriter(body)

			if query, response.Error = ParseQuery(data[1:]); response.Error != nil {
				return
			}

			for field, values := range query {
				for _, value := range values {
					if string(value[0]) == File_name {
						var (
							file_tab []string
							fw       io.Writer
							data     []byte
						)

						file_tab = strings.Split(value[1:], File_date)

						if fw, response.Error = mw.CreateFormFile(field, file_tab[0]); response.Error != nil {
							return
						}
						if len(file_tab) == 1 {
							//read file
							var fr *os.File
							if fr, response.Error = os.Open(file_tab[0]); response.Error != nil {
								return
							}
							defer fr.Close()

							if _, response.Error = io.Copy(fw, fr); response.Error != nil {
								return
							}
						} else if data, response.Error = hex.DecodeString(file_tab[1]); response.Error == nil {
							//write data
							if _, response.Error = fw.Write(data); response.Error != nil {
								return
							}
						} else {
							return
						}
					} else {
						mw.WriteField(field, value)
					}
				}
			}
			if response.Error = mw.Close(); response.Error != nil {
				return
			}
			ct = mw.FormDataContentType()
		}
	}

	var request *http.Request

	if request, response.Error = http.NewRequest(method, url, body); response.Error != nil {
		return
	}

	if ct != "" {
		request.Header.Set("Content-Type", ct)
	}

	for name, value := range header {
		request.Header.Set(name, value)
	}

	var resp *http.Response
	//http.DefaultClient.Do(request)
	if resp, response.Error = self.RoundTrip(request); response.Error != nil {
		return
	}

	if response.Body, response.Error = ioutil.ReadAll(resp.Body); response.Error != nil {
		return
	}

	resp.Body.Close()

	response.StatusCode = resp.StatusCode
	// response.ContentLength = resp.ContentLength

	response.Header = make(map[string]string)
	for name, value := range resp.Header {
		response.Header[name] = strings.Join(value, "; ")
	}

	response.Location = response.Header["Location"]

	response.Cookie = make(map[string]string)
	var cookies []string
	for _, cookie := range resp.Cookies() {
		cookies = append(cookies, cookie.Name+"="+cookie.Value)
		response.Cookie[cookie.Name] = cookie.Value
	}

	response.Cookies = strings.Join(cookies, "; ")
	return
}

func Get(url string, header map[string]string) Response {
	return DefaultTransport.do(http.MethodGet, url, "", header)
}

func (self *transport) Get(url string, header map[string]string) Response {
	return self.do(http.MethodGet, url, "", header)
}
func Post(url, data string, header map[string]string) Response {
	return DefaultTransport.do(http.MethodPost, url, data, header)
}

func (self *transport) Post(url, data string, header map[string]string) Response {
	return self.do(http.MethodPost, url, data, header)
}
