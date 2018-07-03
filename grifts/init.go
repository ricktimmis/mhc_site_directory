package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/ricktimmis/mhclub/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
