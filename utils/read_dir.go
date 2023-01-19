package utils

import (
	"io/ioutil"
	"strings"

	c "github.com/chaewonkong/go-json-server/config"
)

func ReadDirAndGetRouteConfig(targetDir string) (configs []c.RouteConfig, err error) {
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return
	}

	for _, file := range files {
		var b []byte
		filePath := targetDir + "/" + file.Name()
		urlPath := strings.Split(file.Name(), ".")[0]

		b, err = ioutil.ReadFile(filePath)
		if err != nil {
			return
		}

		routeConfig := c.NewRouteConfig(urlPath, b)
		configs = append(configs, routeConfig)
	}

	return
}
