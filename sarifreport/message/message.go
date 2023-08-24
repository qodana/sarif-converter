package message

import (
	"github.com/owenrumney/go-sarif/v2/sarif"
	"sarif-converter/sarifreport/rule"
)

func GetTextMessage(result *sarif.Result, rules rule.Wrappers) *string {
	message := result.Message

	if message.Text != nil {
		return message.Text
	}

	r := rules.Find(result)
	return format(r.TextMessage(result.Message), message.Arguments)
}
