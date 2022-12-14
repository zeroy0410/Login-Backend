package config

// Cache is the configuration of the redis cache server.
import (
    "errors"

    "github.com/spf13/viper"
)

type Cache struct {
    Path string `mapstructure:"path" json:"path" yaml:"path"`
}

var CacheConfig Cache

func init() {
    Register(&CacheConfig)
}

func (c *Cache) Save() error {
    viper.Set("cache", c)
    err := viper.WriteConfig()
    if err != nil {
        return err
    }
    return nil
}

func (c *Cache) Load() error {
    configReader := viper.Sub("cache")
    if configReader == nil {
        return errors.New("could not read cache config")
    }
    err := configReader.Unmarshal(&c)
    if err != nil {
        return err
    }
    return nil
}
