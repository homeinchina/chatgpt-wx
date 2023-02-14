package bootstrap

import (
	"os"

	"github.com/eatmoreapple/openwechat"
	"github.com/homeinchina/chatgpt-wx/handler/wechat"
	log "github.com/sirupsen/logrus"
)

func StartWebChat() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = wechat.Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewJsonFileHotReloadStorage("token.json")
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		err := os.Remove("token.json")
		if err != nil {
			return
		}
		reloadStorage = openwechat.NewJsonFileHotReloadStorage("token.json")
		err = bot.HotLogin(reloadStorage)
		if err != nil {
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	friends, err := self.Friends()

	for i, friend := range friends {
		log.Println(i, friend)
	}
	groups, err := self.Groups()
	for i, group := range groups {
		log.Println(i, group)
	}
	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}
}
