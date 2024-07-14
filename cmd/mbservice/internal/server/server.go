package server

import (
	"github.com/rkuprov/mbot/pkg/auth"
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

type MBot struct {
	db   *store.Store
	auth *auth.Auth
}

var _ mbotpbconnect.MBotServerServiceHandler = (*MBot)(nil)

func NewMBot(db *store.Store, auth *auth.Auth) *MBot {
	return &MBot{
		db:   db,
		auth: auth,
	}
}
