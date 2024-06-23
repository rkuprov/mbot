package server

import (
	"github.com/rkuprov/mbot/pkg/gen/mbotpb/mbotpbconnect"
	"github.com/rkuprov/mbot/pkg/store"
)

type MBot struct {
	// mbotpbconnect.UnimplementedMBotServerServiceHandler
	db *store.Store
}

var _ mbotpbconnect.MBotServerServiceHandler = (*MBot)(nil)

func NewMBot(db *store.Store) *MBot {
	return &MBot{
		db: db,
	}
}
