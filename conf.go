package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fengde/gocommon/filex"
	"github.com/fengde/gocommon/httpx"
	"github.com/fengde/gocommon/timex"
	"gopkg.in/yaml.v3"
)

const YamlPath = "./conf.yaml"
const YamlUrl = "https://raw.githubusercontent.com/fengde/tip/main/cmd.yaml"
const YamlTmp = "/var/tmp/tip"

type Tip struct {
	Cmd  string
	Desc string
}

type Conf map[string][]Tip

func LoadYaml() (*Conf, error) {
	content, err := os.ReadFile(YamlPath)
	if err != nil {
		return nil, err
	}
	var conf Conf
	if err := yaml.Unmarshal(content, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func UpdateYaml() {
	lastUpdate, _ := filex.ReadFileToString(YamlTmp)
	if lastUpdate != "" {
		if timex.NowUnix()-timex.String2Unix(lastUpdate) > 86400 {
			resp, err := httpx.Get(YamlUrl, nil, nil, time.Second*3)
			if err != nil {
				fmt.Println(err)
				return
			}
			if resp.StatusCode() == 200 {
				if err := filex.WriteStringToFile(YamlPath, resp.String(), false); err != nil {
					fmt.Println(err)
					return
				}
				filex.WriteStringToFile(YamlTmp, fmt.Sprintf("%d", timex.NowUnix()), false)
			}
		}
	}
}
