package telegram

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler func(b *Bot, chatID int64, arg string) (responses, error)

type command struct {
	tgbotapi.BotCommand
	text     string
	keyboard *tgbotapi.InlineKeyboardMarkup
	handler  handler
}

var (
	MenuCmd = command{
		BotCommand: tgbotapi.BotCommand{
			Command:     "menu",
			Description: "меню бота",
		},
		text: "Итак, чего ты хочешь?",
	}
	ConfigForNewKeysCmd = command{
		BotCommand: tgbotapi.BotCommand{
			Command:     "newkeys",
			Description: "Новая конфигурация для новой пары ключей",
		},
		text: "",
	}
	ConfigForPublicKeyCmd = command{
		BotCommand: tgbotapi.BotCommand{
			Command:     "pubkey",
			Description: "Новая конфигурация для вашего открытого ключа",
		},
		text: "пришлите мне свой открытый ключ wire guard, вот так:\n" +
			"`/pubkey <ваш ключ в base64>`",
	}
	HelpCmd = command{
		BotCommand: tgbotapi.BotCommand{
			Command:     "help",
			Description: "Показать функциональность бота с описанием",
		},
		text: "Привет, я телеграмм-бот wireguard\n" +
			"Я могу создать для вас новые файлы конфигурации WireGuard VPN\n\n" +
			"/menu — доступные команды\n" +
			"/newkeys — новая конфигурация для новой пары ключей\n" +
			"/pubkey — новая конфигурация для вашего открытого ключа\n" +
			"/help — это сообщение",
	}
)

var commands = map[string]*command{
	MenuCmd.Command:               &MenuCmd,
	ConfigForNewKeysCmd.Command:   &ConfigForNewKeysCmd,
	ConfigForPublicKeyCmd.Command: &ConfigForPublicKeyCmd,
	HelpCmd.Command:               &HelpCmd,
}

// setMyCommands is adapted method from unreleased v5.0.1
// https://github.com/go-telegram-bot-api/telegram-bot-api/commit/4a2c8c4547a868841c1ec088302b23b59443de2b
func setMyCommands(api *tgbotapi.BotAPI) error {
	params := make(tgbotapi.Params)
	data, err := json.Marshal([]command{
		MenuCmd,
		ConfigForNewKeysCmd,
		ConfigForPublicKeyCmd,
		HelpCmd,
	})
	if err != nil {
		return err
	}
	params.AddNonEmpty("commands", string(data))
	_, err = api.MakeRequest("setMyCommands", params)
	if err != nil {
		return err
	}
	return nil
}
