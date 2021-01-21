package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/tombuildsstuff/tfdev/models"
)


func main() {
	var args []string
	if len(os.Args) > 0 {
		args = os.Args[1:]
	}

	if len(args) != 2 {
		log.Printf("Usage: `tfdev [provider] [dev|prod]")
		os.Exit(1)
		return
	}

	if err := run(args); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
		return
	}

	os.Exit(0)
}


func run(args []string) error {
	path, err := configFilePath()
	if err != nil {
		return fmt.Errorf("retrieving file path: %+v", err)
	}
	config, err := models.ConfigFromPath(*path)
	if err != nil {
		return fmt.Errorf("parsing config: %+v", err)
	}

	providerName := args[0]
	mode := args[1]
	var provider *models.Provider
	for _, p := range config.Providers {
		if p.Alias == providerName {
			provider = &p
			break
		}
	}
	if provider == nil {
		return fmt.Errorf("missing alias for %q", providerName)
	}

	switch mode {
	case "dev":
		log.Printf("Templating %q to use %q..", providerName, provider.Dev)
		return templateProvider("providers.tf", providerName, provider.Dev)
		break

	case "prod":
		log.Printf("Templating %q to use %q..", providerName, provider.Prod)
		return templateProvider("providers.tf", providerName, provider.Prod)
		break
	default:
		return fmt.Errorf("unknown mode %q", mode)
	}

	return nil
}

func templateProvider(fileName, providerName, source string) error {
	providerConfig := fmt.Sprintf(`
terraform {
  required_providers {
    %s = {
      source = "%s"
    }
  }
}
`, providerName, source)
	providerConfig = strings.TrimSpace(providerConfig)

	_ = os.Remove(fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, _ = file.WriteString(providerConfig)
	return nil
}

func configFilePath() (*string, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	out := path.Join(home, ".tfdev.hcl")
	return &out, nil
}