package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (cmd command) button() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(cmd.Description, cmd.Command)
}

var (
	menuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(ConfigForNewKeysCmd.button()),
		tgbotapi.NewInlineKeyboardRow(ConfigForPublicKeyCmd.button()),
		tgbotapi.NewInlineKeyboardRow(HelpCmd.button()),
	)

	goToMenuButton = tgbotapi.NewInlineKeyboardButtonData("вернуться в меню", MenuCmd.Command)

	configForPublicKeyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(
				"Как сгенерировать пару ключей WireGuard?",
				"https://www.wireguard.com/quickstart/#key-generation",
			),
		),
		tgbotapi.NewInlineKeyboardRow(goToMenuButton),
	)

	helpKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(goToMenuButton),
	)
)

func init() {
	MenuCmd.keyboard = &menuKeyboard
	ConfigForPublicKeyCmd.keyboard = &configForPublicKeyKeyboard
	HelpCmd.keyboard = &helpKeyboard
}
