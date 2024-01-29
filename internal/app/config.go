package app

import (
	"os"
	"strings"

	"github.com/knadh/koanf/providers/confmap"

	"github.com/joomcode/errorx"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const EnvPrefix = "VERSIONER_"

var (
	ErrFailedReadConfig = errorx.NewType(errorx.CommonErrors, "ErrFailedReadConfig")
)

type Config struct {
	// Files - list of files that contain version
	Files []File `mapstructure:"files"`
	// Debug - whether to print debug messages
	Debug bool `mapstructure:"debug"`
}

type File struct {
	// Name - the name of the file that contains version
	Name string `mapstructure:"name"`
	// Path - path to the version in structured documents, like JSON, YAML, XML
	Path string `mapstructure:"path"`
	// Regexp - expression to find version in unstructured documents like MD, TXT
	Regexp string `mapstructure:"regexp"`
}

func ReadConfig() (*Config, error) {
	k := koanf.New(".")

	err := k.Load(
		confmap.Provider(
			map[string]any{
				"debug": "false",
			}, ".",
		), nil,
	)
	if err != nil {
		return nil, ErrFailedReadConfig.Wrap(err, "Failed to load configuration from map")
	}

	parser := yaml.Parser()
	fileNames := []string{"versioner.yml", "versioner.yaml"}
	for _, fileName := range fileNames {
		err := k.Load(file.Provider(fileName), parser)
		if err != nil && !os.IsNotExist(err) {
			return nil, ErrFailedReadConfig.Wrap(err, "Failed to load configuration from file %v", fileName)
		}
	}

	err = k.Load(
		env.Provider(
			EnvPrefix, "_", func(name string) string {
				name = strings.TrimPrefix(name, EnvPrefix)
				name = strings.ToLower(name)
				name = strings.ReplaceAll(name, "_", ".")
				return name
			},
		), nil,
	)
	if err != nil {
		return nil, ErrFailedReadConfig.Wrap(err, "Failed to load environment variables")
	}

	config := &Config{}
	err = k.Unmarshal("", config)
	if err != nil {
		return nil, ErrFailedReadConfig.Wrap(err, "Failed to unmarshall config")
	}

	return config, nil
}
