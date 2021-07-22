package initialize

import (
	"bufio"
	"encoding/json"
	"gin/config"
	"os"
)

var _cfg *config.Config
func ParseAllConfig(path string) (*config.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err2 := decoder.Decode(&_cfg)
	if err2 != nil {
		return nil, err
	}
	return _cfg, nil
}
