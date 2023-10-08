package main

import (
	"fmt"
	"strings"

	"github.com/fengde/gocommon/colorx"
)

func Show(prefix ...string) {
	conf, err := LoadYaml()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}
	if len(prefix) == 0 {
		for k, v := range *conf {
			_show(k, v)
		}
		return
	}

	pre := prefix[0]
	for k, v := range *conf {
		if strings.HasPrefix(k, pre) {
			_show(k, v)
		}
	}
}

func _show(cmd string, tips []Tip) {
	fmt.Println(colorx.WithColor(cmd, colorx.FgGreen))
	for _, tip := range tips {
		fmt.Printf("  %s: %s\n", colorx.WithColor(tip.Cmd, colorx.FgGreen), colorx.WithColor(tip.Desc, colorx.FgYellow))
	}
}
