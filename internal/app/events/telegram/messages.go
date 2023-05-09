package telegram

import "fmt"

const msgHelp = `I can store your logins and passwords for different services.

✅ To save login and password for service send me /set <service> <login> <password>

✅ To get login and password for service send me /get <service>.

✅ To delete login and password for service send me /delete <service>.

❗ CAUTION! BOT IS WORKING ON A TEST MODE (DATA IS NOT ENCRYPTED)`

const msgHello = "Hi there! 😉\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤔"
	msgNoService      = "You haven't login and password for this service 🙊"
	msgSaved          = "Saved! 👌"
	msgUpdated        = "Updated! 👌"
	msgDeleted        = "Deleted! 👌"
)

func buildGetMessage(serviceName, login, password string) string {
	return fmt.Sprintf("Service: %s \nLogin: %s \nPassword: %s", serviceName, login, password)
}
