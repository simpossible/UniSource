package config

import (
	"cmd.unisource.icu/util"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type UniConfig struct {
	Source     string   `json:"source"`
	ImgItems   []string `json:"img_items"`
	VapItems   []string `json:"vap_items"`
	VideoItems []string `json:"video_items"`
	LottyItems []string `json:"lotty_items"`
}

func UniConfigForJsonFile(filePath string) (*UniConfig, error) {
	exist, _ := util.PathExists(filePath)
	if exist {
		data, err := ioutil.ReadFile(filePath)
		if err == nil {
			config := UniConfig{}
			//m := new(map[string]interface{})
			//err := json.Unmarshal(data, m)
			err := json.Unmarshal(data, &config)
			return &config, err
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("文件不存在 %v", filePath))
	}
}
