package envvar

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// LoadConfig will load files optionally from the json file stored at path, then
// will override those values based on the envconfig struct tags. The envPrefix is how
// we prefix our environment variables.
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from file")
}

// LoadFile unmarshalls a json file into a config struct
