package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fengde/gocommon/colorx"
)

func Show(prefix ...string) {
	conf, err := LoadYaml()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}

	var commands []string
	for command := range *conf {
		commands = append(commands, command)
	}

	sort.Strings(commands)

	for _, command := range commands {
		if len(prefix) == 0 || strings.HasPrefix(command, prefix[0]) {
			colorPrint(command, (*conf)[command])
		}

	}
}

func colorPrint(command string, tips []Tip) {
	fmt.Println(colorx.WithColor(command, colorx.FgGreen))
	for _, tip := range tips {
		fmt.Printf("  %s: %s\n", colorx.WithColor(tip.Cmd, colorx.FgGreen), colorx.WithColor(tip.Desc, colorx.FgYellow))
	}
}
