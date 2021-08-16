package fynex

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
)

var _ fyne.Theme = (*Theme)(nil)

func DarkTheme(opt ...Option) fyne.Theme {
	t, err := NewTheme(theme.DarkTheme(), opt...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return t
}

func LightTheme(opt ...Option) fyne.Theme {
	t, err := NewTheme(theme.LightTheme(), opt...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return t
}

type Theme struct {
	fyne.Theme
	font fyne.Resource
}

type Options struct {
	font string
}

type Option func(*Options)

func WithFont(font string) Option {
	return func(opt *Options) {
		opt.font = font
	}
}

func NewTheme(base fyne.Theme, opts ...Option) (t *Theme, err error) {
	t = new(Theme)
	t.Theme = base
	var o Options
	for _, v := range opts {
		v(&o)
	}
	fPath, err := findfont.Find(o.font)
	if err != nil {
		return
	}
	t.font, err = fyne.LoadResourceFromPath(fPath)
	return
}

func (t *Theme) Font(fyne.TextStyle) fyne.Resource {
	return t.font
}
