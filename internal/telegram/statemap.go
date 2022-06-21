package telegram

import (
	tm "github.com/and3rson/telemux/v2"
)

func (t *Telegram) initMux() {
	stateMap := tm.StateMap{
		stateDefault: {
			tm.NewHandler(tm.IsCommandMessage(cmdTextStart), t.startCmd),
			tm.NewHandler(tm.IsCommandMessage(cmdTextAdd), t.addCmd),
		},
		stateLocation: {
			tm.NewHandler(tm.HasLocation(), t.newLocation),
			tm.NewHandler(tm.IsCommandMessage(cmdTextCancel), t.cancel),
		},
		stateComment: {
			tm.NewHandler(tm.HasText(), t.newComment),
			tm.NewHandler(tm.IsCommandMessage(cmdTextCancel), t.cancel),
		},
		statePhoto: {
			tm.NewHandler(tm.HasPhoto(), t.newPhoto),
			tm.NewHandler(tm.IsCommandMessage(cmdTextSkip), t.skip),
			tm.NewHandler(tm.IsCommandMessage(cmdTextCancel), t.cancel),
		},
	}

	def := []*tm.Handler{
		tm.NewHandler(tm.IsCommandMessage(cmdTextCancel)),
	}

	cmdHandler := []*tm.Handler{
		tm.NewHandler(tm.And(tm.IsPrivate(), tm.IsCommandMessage("start")), t.startCmd),
	}

	t.cmds = tm.NewMux().
		AddHandler(tm.NewConversationHandler("menu", tm.NewLocalPersistence(), stateMap, def)).
		AddHandler(cmdHandler...).
		SetRecover(t.onError)
}
