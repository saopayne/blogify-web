package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/saopayne/blogify-web/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
