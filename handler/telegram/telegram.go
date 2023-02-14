package telegram

import (
	"strings"

	"github.com/homeinchina/chatgpt-wx/openai"
	log "github.com/sirupsen/logrus"
)

func Handle(msg string) *string {
	requestText := strings.TrimSpace(msg)
	reply, err := openai.Completions(requestText)
	if err != nil {
		log.Println(err)
	}
	return reply
}
