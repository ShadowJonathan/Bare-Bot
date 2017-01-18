package Belt

import (
	fmt "fmt"

	"github.com/bwmarrin/discordgo"
)

type Version struct {
	Major               byte
	Minor               byte
	Build               byte
	Experimental        bool
	ExperimentalVersion byte
}

type Bot struct {
	dg      *discordgo.Session
	version Version
}

// Vars after this

var bbb *Bot
var err error

// Functions after this

func Initialize(Token string) {
	bb = &Bot{
		version: Version{0, 0, 1, true, 1},
	}

	bb.dg, err = discordgo.New("BOT " + Token)
	if err != nil {
		fmt.Println("Discord Session error, check token, error message: " + err.Error())
		return
	}
}
