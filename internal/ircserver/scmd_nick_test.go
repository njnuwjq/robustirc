package ircserver

import (
	"testing"

	"github.com/robustirc/robustirc/internal/robust"

	"gopkg.in/sorcix/irc.v2"
)

func TestServerNick(t *testing.T) {
	i, ids := stdIRCServerWithServices()

	i.ProcessMessage(&robust.Message{Session: ids["secure"]}, irc.ParseMessage("JOIN #test"))

	mustMatchIrcmsgs(t,
		i.ProcessMessage(&robust.Message{Session: ids["services"]}, irc.ParseMessage("NICK ChanServ 1 1422134861 services robustirc.net services.robustirc.net 0 :Operator Server")),
		[]*irc.Message{})

	mustMatchMsg(t,
		i.ProcessMessage(&robust.Message{Session: ids["services"]}, irc.ParseMessage("NICK ChanServ 1 1422134861 services robustirc.net services.robustirc.net 0 :Operator Server")),
		":robustirc.net 433 * ChanServ :Nickname is already in use")

	mustMatchMsg(t,
		i.ProcessMessage(&robust.Message{Session: ids["secure"]}, irc.ParseMessage("NICK NickServ")),
		":robustirc.net 433 sECuRE NickServ :Nickname is already in use")
}
