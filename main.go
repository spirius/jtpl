package main

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	var (
		templateFile string
		vars         = make(map[string]interface{})
		err          error
		tpl          *template.Template
		decoder      *json.Decoder
		content      []byte
    funcMap      map[string]interface{}
	)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <template>\nVariables should be sent as JSON from STDIN\n", os.Args[0])
		os.Exit(1)
	}

	templateFile = os.Args[1]

	if _, err = os.Stat(templateFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s not found\n", templateFile)
		os.Exit(1)
	}

	content, err = ioutil.ReadFile(templateFile)

	if err != nil {
		panic(err)
	}

  funcMap = sprig.TxtFuncMap()
  funcMap["acell"] = NewCell

	tpl, err = template.New("").Funcs(funcMap).Parse(string(content))

	if err != nil {
		panic(err)
	}

	decoder = json.NewDecoder(os.Stdin)

	if err = decoder.Decode(&vars); err != nil {
		fmt.Fprintf(os.Stderr, "error reading parsing stdin: %s", err)
		os.Exit(1)
	}

	err = tpl.Execute(os.Stdout, vars)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v", templateFile, err)
		os.Exit(1)
	}
}

type Cell struct{ v interface{} }

func NewCell(v ...interface{}) (*Cell, error) {
  switch len(v) {
  case 0:
    return new(Cell), nil
  case 1:
    return &Cell{v[0]}, nil
  default:
    return nil, fmt.Errorf("wrong number of args: want 0 or 1, got %v", len(v))
  }
}

func (c *Cell) Set(v interface{}) *Cell { c.v = v; return c }
func (c *Cell) Get() interface{}        { return c.v }

