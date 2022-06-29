package main

import (
	"errors"
	"fmt"
	"github.com/analogj/justvanish/data"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {

	//parse schema
	compiler := jsonschema.NewCompiler()
	schemaFileReader, err := os.Open("organization-schema.json")
	if err := compiler.AddResource("schema.json", schemaFileReader); err != nil {
		panic(err)
	}
	schema, err := compiler.Compile("schema.json")
	if err != nil {
		panic(err)
	}

	//loop through and process each organization yaml file
	orgConfigList, err := data.Organizations.ReadDir("organizations")
	if err != nil {
		panic(err)
	}

	failedValidation := map[string]error{}
	for _, orgConfigFileName := range orgConfigList {
		orgConfig, err := data.Organizations.ReadFile(fmt.Sprintf("organizations/%s", orgConfigFileName.Name()))
		if err != nil {
			panic(err)
		}

		var m interface{}
		err = yaml.Unmarshal(orgConfig, &m)
		if err != nil {
			panic(err)
		}
		m, err = toStringKeys(m)
		if err != nil {
			panic(err)
		}
		if err := schema.Validate(m); err != nil {
			failedValidation[orgConfigFileName.Name()] = err
			fmt.Printf("%s failed validation: %v", orgConfigFileName, err)
		}
	}

}

func toStringKeys(val interface{}) (interface{}, error) {
	var err error
	switch val := val.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range val {
			k, ok := k.(string)
			if !ok {
				return nil, errors.New("found non-string key")
			}
			m[k], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return m, nil
	case []interface{}:
		var l = make([]interface{}, len(val))
		for i, v := range val {
			l[i], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		return val, nil
	}
}
