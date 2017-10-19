package http

import (
	"fmt"
	"testing"
)

var data = []string{
	"a=aaa&b=bbb",
	`{"key":"value"}`,
	`[1,2,3]`,

	`^a=@http_test.go&b=bbb`,
	fmt.Sprintf(`^name=@xxx.xxx$%x&b=bbb`, "no-if"),
}

func Test(t *testing.T) {

	resp := Get("http://httpbin.org/get", nil)
	fmt.Println("get", resp.Error, "\n", string(resp.Body))

	resp = Get("http://httpbin.org/get", map[string]string{"Cookie": "a=b"})
	fmt.Println("header", resp.Error, "\n", string(resp.Body))

	for _, data := range data {
		resp = Post("http://httpbin.org/post", data, nil)
		fmt.Println("post", data, resp.Error, "\n", string(resp.Body))
	}

}
