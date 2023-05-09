package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	netUrl "net/url"
	"os"

	"github.com/gookit/color"
	"github.com/thoas/go-funk"
)

var green = color.FgGreen.Render

const TPL_EXPORT = `
{{- range $i, $a := .Items }}
export {{ $i }}="{{ $a }}"
{{- end }}
`

func main() {
	data, err := httpGet(
		"https://raw.githubusercontent.com/cnpm/binary-mirror-config/master/package.json",
		map[string]string{},
	)

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	packageJson := make(map[string]interface{})
	json.Unmarshal(data, &packageJson)

	envs := funk.Get(packageJson, "mirrors.china.ENVS")
	// log.Printf("%s %#v\n", green("[response.body]"), envs)

	t := template.Must(template.New("npm-binary-export-tpl").Parse(TPL_EXPORT))
	t.Execute(os.Stdout, map[string]interface{}{
		"Items": envs,
	})
}

func httpGet(url string, params map[string]string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	// add params

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}

	body, _ := netUrl.QueryUnescape(req.URL.String())

	log.Printf("%s %s URL : %s \n", green("[http.request]"), http.MethodGet, body)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
