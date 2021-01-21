package models

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Providers []Provider `hcl:"provider,block"`
}

type Provider struct {
	Alias string `hcl:"alias,label"`
	Dev string `hcl:"dev"`
	Prod string `hcl:"prod"`
}

func ConfigFromPath(fileName string) (*Config, error) {
	var config Config
	if err := hclsimple.DecodeFile(fileName, nil, &config); err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", fileName, err)
	}
	return &config, nil
}