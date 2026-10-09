package telegram

import (
	"golang-lod/lib/e"
	"golang-lod/storage"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:
	}
}

func (p *Processor) Savepage(chatid int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("cant do command: save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}
	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if isExists {
		return p.tg.SendMessage(chatid, msgAlreadyExists)
	}
	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatid, msgSaved); err != nil {
		return err
	}
	return nil
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
