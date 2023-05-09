package telegram

import (
	"context"
	"errors"
	"log"
	"password-storage-bot/internal/app/models"
	"password-storage-bot/internal/app/storage"
	e "password-storage-bot/pkg/lib"
	"strings"
)

const (
	SetCmd    = "/set"
	GetCmd    = "/get"
	DeleteCmd = "/delete"
	HelpCmd   = "/help"
	StartCmd  = "/start"
)

func (p *Processor) doCmd(text string, chatID int, userName string) error {
	text = strings.TrimSpace(text)

	command := command(text)
	splitedText := strings.Split(text, " ")

	//delete after
	log.Printf("got new command '%s' from '%s", text, userName)

	switch command {
	case SetCmd:
		if len(splitedText) == 4 {
			serviceName := splitedText[1]
			login := splitedText[2]
			password := splitedText[3]
			return p.saveService(chatID, userName, serviceName, login, password)
		}
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	case GetCmd:
		if len(splitedText) == 2 {
			serviceName := splitedText[1]
			return p.getService(chatID, userName, serviceName)
		}
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	case DeleteCmd:
		if len(splitedText) == 2 {
			serviceName := splitedText[1]
			return p.deleteService(chatID, userName, serviceName)
		}
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	case HelpCmd:
		if len(splitedText) == 1 {
			return p.sendHelp(chatID)
		}
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	case StartCmd:
		if len(splitedText) == 1 {
			return p.sendHello(chatID)
		}
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func command(text string) string {
	return strings.Split(text, " ")[0]
}

func (p *Processor) saveService(chatID int, userName, serviceName, login, password string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save service", err) }()

	service := &models.Service{
		UserName:    userName,
		ServiceName: serviceName,
		Login:       login,
		Password:    password,
	}

	isExists, _ := p.storage.IsExists(context.Background(), service)
	if err != nil {
		return err
	}

	if isExists {
		if err := p.storage.Update(context.Background(), service, login, password); err != nil {
			return err
		}
		if err := p.tg.SendMessage(chatID, msgUpdated); err != nil {
			return err
		}
		return nil
	}

	if err := p.storage.Set(context.Background(), service); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func (p *Processor) getService(chatID int, userName, serviceName string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: get service", err) }()
	service, err := p.storage.Get(context.Background(), userName, serviceName)

	if err != nil && !errors.Is(err, storage.ErrNoService) {
		return err
	}

	if errors.Is(err, storage.ErrNoService) {
		return p.tg.SendMessage(chatID, msgNoService)
	}

	msg := buildGetMessage(service.ServiceName, service.Login, service.Password)

	return p.tg.SendMessage(chatID, msg)
}

func (p *Processor) deleteService(chatID int, userName, serviceName string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: delete service", err) }()
	if err := p.storage.Delete(context.Background(), userName, serviceName); err != nil {
		return err
	}
	return p.tg.SendMessage(chatID, msgDeleted)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
