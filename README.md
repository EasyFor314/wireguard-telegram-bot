# wireguard-telegram-bot


Простой Telegram-бот для генерации конфигурации VPN WireGuard

## Функциональность

- `/menu` — список доступных команд
- `/newkeys` — создайте новый конфигурационный файл и qr-код для новой сгенерированной пары ключей
- `/pubkey` — создайте новый шаблон конфигурационного файла для предоставленного вами открытого ключа
- `/help` — распечатайте это сообщение

## Публичный телеграмм-бот Wireguard

[Install](https://www.wireguard.com/install/) Клиент WireGuard для вашего устройства и импортируйте сгенерированный файл или отсканируйте qr-код

## Настройте своего собственного телеграмм-бота WireGuard

- Go to [@BotFather](https://t.me/BotFather), send him `/newbot`, choose a bot's name and username, and receive Telegram Bot API Token
- Go to AWS, GCP, whatever ☁️ and setup your remote server in desired region
  - You need to open corresponding port (e.g. `udp:51820`)
- Install `go`, `wireguard` and `wireguard-tools` on your server
  - Someday, we hope there will be a handy Dockerfile for it 🐳
- Generate Wireguard key pair for your server, create appropriate config file (e.g. `wg0.conf`) and run Wireguard
  - You're all big boys, you'll handle it
- `git clone git@github.com:skoret/wireguard-telegram-bot.git`
- `cd wireguard-telegram-bot && cp .env.example .env`
- Set env variables in `.env` file:

  | Variable              | Content | Notes |
    | :-------------------- | ------- | ----- |
  | `TELEGRAM_APITOKEN`   | your Telegram Bot API token from [@BotFather](https://t.me/BotFather) | keep it in _secret_! |
  | `ADMIN_USERNAMES`     | list of Telegram usernames, separated by commas, who are allowed to access this bot | leave variable _empty_ for public access |
  | `DNS_IPS`             | list of DNS ip addresses, separated by commas | e.g. `8.8.8.8,8.8.4.4` |
  | `SERVER_ENDPOINT`     | `<your_machine's_external_ip:open_port>` | |
  | `WIREGUARD_INTERFACE` | new Wireguard interface name | e.g. `wg0` |
  | `TEMPLATES_FOLDER`    | path to configuration template files | probably, you don't wanna change it |
  | `DEV_MODE`            | `false` for common uses<br />`true` for mocked internal wireguard client | dev mode suitable for manual bot ui tests |
- `sudo go run cmd/bot/main.go`
- 🎉 🍻 🥳
