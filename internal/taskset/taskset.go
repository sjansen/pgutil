package taskset

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsonnet "github.com/google/go-jsonnet"
)

type Config struct {
	Databases map[string]map[string]string
	Tasks     map[string]map[string]string
}

func Load(directory, filename string) (*Config, error) {
	pathname := filepath.Join(directory, filename)
	bytes, err := ioutil.ReadFile(pathname)
	if err != nil {
		return nil, err
	}

	vm := jsonnet.MakeVM()
	importer := &jsonnet.FileImporter{
		JPaths: []string{directory},
	}
	vm.Importer(importer)

	evaluated, err := vm.EvaluateSnippet(filename, string(bytes))
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(evaluated))
	dec.DisallowUnknownFields()
	cfg := &Config{}
	dec.Decode(cfg)

	return cfg, nil
}
