package tinify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type re struct {
	Input struct {
		Size int    `json:"size"`
		Type string `""json:"type"`
	} `json:"input"`
	Output struct {
		Size   int     `json:"size"`
		Type   string  `""json:"type"`
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Ratio  float64 `json:"ratio"`
		Url    string  `json:"url"`
	} `json:"output"`
}

var (
	KEY string
)

func Do(path string) (re re, err error) {
	var (
		data     []byte
		request  *http.Request
		response *http.Response
	)

	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	request, err = http.NewRequest("POST", "https://tinypng.com/site/shrink", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	request.SetBasicAuth("api", KEY)
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &re)
	return
}

// {
// 	"input":{
// 		"size":114936,
// 		"type":"image/jpeg"
// 	},
// 	"output":{
// 		"size":53193,
// 		"type":"image/jpeg",
// 		"width":710,
// 		"height":300,
// 		"ratio":0.4628,
// 		"url":"https://tinypng.com/site/output/7e87410cnfr96uf8.jpg"
// 	}
// }
