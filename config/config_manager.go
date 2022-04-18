package config

import (
	"cmd.unisource.icu/util"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// AllConfigAtDir 获取一个地址下所有的配置
func AllConfigAtDir(dir string) ([]*UniConfig, error) {
	var configs []*UniConfig
	exist, err := util.PathExists(dir)
	if exist {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			if !info.IsDir() {
				name := info.Name()
				if strings.HasSuffix(name, ".uni") {
					conf, err := UniConfigForJsonFile(path)
					if conf != nil {
						configs = append(configs, conf)
					} else {
						fmt.Printf("uni文件读取错误 %v", err)
					}
				}
			}
			return nil
		})
	} else {
		if err == nil {
			fmt.Printf("路径不存在")
		} else {
			fmt.Printf("错误: %v", dir)
		}

	}
	return configs, err
}
