package comm

import (
	"fmt"

	"github.com/ameamezhou/xiawuyue/xconfig"
)

// Settings 配置
var (
	Config *xconfig.WeConfig
)

// LoadSettings 初始化配置
func LoadSettings(configFile string) error {
	c, err := xconfig.LoadConfig(configFile)
	if err != nil {
		fmt.Printf("LoadSettings error: %v\n", err)
		return err
	}
	Config = c
	return nil
}
