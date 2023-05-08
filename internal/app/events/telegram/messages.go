package telegram

import "fmt"

const msgHelp = `I can store your logins and passwords for different services.
If you save login and password for current service just send me /set <service> <login> <password>.
If you wanna get login and password for current service send me /get <service>.
If you wanna delete login and password for current service send me /delete <service>.
CAUTION! BOT IS WORKING ON A TEST MODE(DATA IS NOT ENCRYPTED)`

const msgHello = "Hi there! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤔"
	msgNoService      = "Yo haven't data for this service🙊"
	msgSaved          = "Saved! 👌"
	msgUpdated        = "Updated! 👌"
	msgDeleted        = "Deleted! 👌"
	msgRewriteService = "Login and password are rewritten 👌"
)

func buildGetMessage(serviceName, login, password string) string {
	return fmt.Sprintf("Service: %s \n Login: %s \n Password: %s", serviceName, login, password)
}
