package main

import (
	"flag"
	"fmt"
	þ "github.com/4thel00z/gojsonschema"
	λ "github.com/4thel00z/lambda/v1"
	"os"
)

var (
	schemaPath = flag.String("schema", "", "path to a json schema")
	load       = þ.NewStringLoader
)

func main() {
	flag.Parse()
	jsonContent := λ.Slurp(os.Stdin).Catch(λ.Die).UnwrapString()
	schemaContent := λ.Open(*schemaPath).Catch(λ.Die).Slurp().UnwrapString()
	loader := þ.NewSchemaLoader()

	s, err := loader.Compile(load(schemaContent))
	if err != nil {
		λ.Die(err)
	}
	validate, err := s.Validate(load(jsonContent))
	if err != nil {
		λ.Die(err)
	}
	if validate.Valid() {
		os.Exit(0)
	}
	if len(validate.Errors()) > 0 {
		for _, err := range validate.Errors() {
			fmt.Println(err.Description())
		}
	}
}
