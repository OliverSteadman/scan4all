package hydra

import (
	"fmt"
	"github.com/hktalent/scan4all/pkg"
	"strings"
)

func init() {
	InitDefaultAuthMap()
	var a1, a2 []string
	HydraUser := pkg.GetVal4File("HydraUser", "")
	if "" != HydraUser {
		a1 = strings.Split(HydraUser, "\n")
	}

	HydraPass := pkg.GetVal4File("HydraPass", "")
	if "" != HydraPass {
		a2 = strings.Split(HydraPass, "\n")
	}
	//加载自定义字典
	InitCustomAuthMap(a1, a2)
}

// 密码破解
func Start(IPAddr string, Port int, Protocol string) {
	authInfo := NewAuthInfo(IPAddr, Port, Protocol)
	crack := NewCracker(authInfo, false, 16)
	fmt.Printf("[hydra]->开始对%v:%v[%v]进行暴力破解，字典长度为：%d", IPAddr, Port, Protocol, crack.Length())
	go crack.Run()
}
