package wx

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"regexp"
)

func sign(parameter map[string]string) (code string) {
	var (
		parameter_url = url.Values{}
		key           = parameter["key"]
	)
	delete(parameter, "key")

	for k, v := range parameter {
		parameter_url.Set(k, v)
	}

	code, _ = url.QueryUnescape(parameter_url.Encode())

	return fmt.Sprintf("%X", md5.Sum([]byte(fmt.Sprintf("%s&key=%s", code, key))))
}

func Xml2map(_xml string) (_map map[string]string) {
	_map = make(map[string]string)

	for _, ele := range regexp.MustCompile(`<(\w+)><!\[CDATA\[(.+)\]\]></\w+>`).FindAllStringSubmatch(_xml, -1) {
		_map[ele[1]] = ele[2]
	}

	return
}

func map2json(_map map[string]string) (json_string string) {
	var first = ""
	json_string = "{"
	for k, v := range _map {
		json_string += fmt.Sprintf(`%s"%s":"%s"`, first, k, v)
		if first == "" {
			first = ","
		}
	}
	return json_string + "}"
}

func map2xml(_map map[string]string) (xml_data string) {
	xml_data = `<xml>`
	for k, v := range _map {
		xml_data += fmt.Sprintf(`<%s>%s</%s>`, k, v, k)
	}
	return xml_data + "</xml>"
}
