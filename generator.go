//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type entry struct {
	name        string
	Default     interface{}             `yaml:"default,omitempty"`
	Description string                  `yaml:"description"`
	Required    bool                    `yaml:"required"`
	Type        interface{}             `yaml:"type"`
	Keys        *map[string]interface{} `yaml:"keys,omitempty"`
}

// the begining and end tags for the configuration in the markdown
const (
	cfs = "{% configuration %}"
	cfe = "{% endconfiguration %}"
)

const discTemplate = `package discovery

type {{.Name | convertKey}} struct {
	{{range $key, $value := .Data}}
	{{$value.Description | comment}}
	// Default: {{$value.Default}}
	{{$key | convertKey}} {{$value | getType}} ` + "`json:\"{{$key}}{{if not $value.Required}},omitempty{{end}}\"`" + `
	{{end}}
}
`

func inList(s string, l []interface{}) bool {
	for _, li := range l {
		if v, ok := li.(string); ok && s == v {
			return true
		}
	}
	return false
}

func convertKey(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")

	return s
}

func comment(s string) string {
	substrs := strings.Split(s, "\n")
	ns := ""
	for _, substr := range substrs {
		ns += "// " + substr + "\n"
	}
	return ns[:len(ns)-2]
}

func getType(n string, ty interface{}) string {
	switch t := ty.(type) {
	case string:
		switch t {
		case "integer":
			return "int"
		case "boolean":
			return "bool"
		case "template", "icon", "device_class":
			return "string"
		// special case for availability
		case "list":
			if n == "availability" {
				return "[]Availability"
			}
		// special case for device
		case "map":
			if n == "device" {
				return "*Device"
			}
		}
	case []interface{}:
		switch {
		// handle probable typo
		case len(t) == 1:
			return getType(n, t[0])
		case len(t) == 2 && inList("string", t) && inList("list", t):
			return "[]string"
		}
		log.Fatalf("%+v - Unhandled Type: %T, %q", n, t, t)
	default:
		log.Fatalf("%+v - Unhandled Type: %T", n, t)
	}
	return "string"
}

func getTypeFromEntry(e entry) string {
	return getType(e.name, e.Type)
}

func getBytes(src string) ([]byte, error) {
	prefix := strings.SplitN(src, ":", 2)[0]

	if prefix == "http" || prefix == "https" {
		resp, err := http.Get(src)
		if err != nil {
			return nil, fmt.Errorf("could not get url: %v", err)
		}

		bs, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("could not read body: %v", err)
		}

		return bs, nil
	}

	return ioutil.ReadFile(src)
}

func readYAML(url string) error {
	bs, err := getBytes(url)
	if err != nil {
		return err
	}

	// find the begining and end of the configuration sections
	m := make(map[string]*entry)

	start := bytes.Index(bs, []byte(cfs)) + len(cfs)
	end := bytes.Index(bs, []byte(cfe))
	err = yaml.Unmarshal(bs[start:end], m)
	if err != nil {
		fmt.Printf("%s\n", bs[start:end])
		return fmt.Errorf("could not unmarshal bytes: %v", err)
	}

	for k := range m {
		m[k].name = k
	}

	su := strings.Split(url, "/")
	fn := su[len(su)-1]
	name := strings.TrimSuffix(fn, ".mqtt.markdown")
	name = convertKey(name)

	s := struct {
		Name string
		Data map[string]*entry
	}{Name: name, Data: m}

	funcMap := template.FuncMap{
		"convertKey": convertKey,
		"getType":    getTypeFromEntry,
		"comment":    comment,
	}

	t, err := template.New("discoverable").Funcs(funcMap).Parse(discTemplate)
	if err != nil {
		return fmt.Errorf("could not create template: %v", err)
	}

	dbs := &bytes.Buffer{}
	err = t.Execute(dbs, s)
	if err != nil {
		return fmt.Errorf("could not execute template: %v", err)
	}

	fbs, err := format.Source(dbs.Bytes())
	if err != nil {
		fmt.Printf("%s\n", dbs)
		return fmt.Errorf("could not format output: %v", err)
	}

	filename := strings.ToLower(name) + ".go"
	err = ioutil.WriteFile(filename, fbs, 0664)
	if err != nil {
		return fmt.Errorf("could not write file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <markdown_url>\n", os.Args[0])
		os.Exit(1)
	}

	err := readYAML(os.Args[1])
	if err != nil {
		log.Fatalf("could not read yaml: %v", err)
	}
}
