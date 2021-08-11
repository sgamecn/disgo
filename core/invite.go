package core

import (
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/rest"
	"github.com/DisgoOrg/disgo/rest/route"
)

type Invite struct {
	discord.Invite
	Disgo      Disgo
	Inviter    *User
	TargetUser *User
}

// URL returns the invite URL in format like https://discord.gg/{code}
func (i *Invite) URL() string {
	url, err := route.InviteURL.Compile(nil, i.Code)
	if err != nil {
		return ""
	}
	return url.Route()
}

func (i *Invite) Delete() rest.Error {
	return i.Disgo.RestServices().InviteService().DeleteInvite(i.Code)
}
