package utils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultBrushSize    float64 `json:"default_brush_size"`
	DefaultMaxBrushSize float64 `json:"default_max_brush_size"`
	DefaultBrushChange  float64 `json:"default_brush_change_rate"`
	DefaultMaskPreffix  string  `json:"default_mask_preffix"`
	DefaultMaskDir      string  `json:"default_mask_dir"`
	DefaultResourceDir  string  `json:"default_resource_dir"`
}

func LoadConfig(file string) (*Config, error) {
	cwd, _ := os.Getwd()
	dir := path.Join(cwd, file)
	viper.SetConfigFile(dir)
	viper.SetConfigType("json")
	var config Config
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	return &config, err
}
