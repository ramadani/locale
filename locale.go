package locale

import (
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localable interface {
	Trans(key string, datas ...TemplateData) string
}

type TemplateData map[string]interface{}

type locale struct {
	localizer *i18n.Localizer
}

func (l *locale) Trans(key string, datas ...TemplateData) string {
	config := &i18n.LocalizeConfig{}
	config.MessageID = key

	if len(datas) == 1 {
		config.TemplateData = datas[0]
	}

	return l.localizer.MustLocalize(config)
}

var (
	l          *locale
	localeOnce sync.Once
)

type RegisterFunc func() *i18n.Localizer

func Register(regFunc RegisterFunc) {
	if l == nil {
		localeOnce.Do(func() {
			localizer := regFunc()
			l = &locale{localizer}
		})
	}
}

func Locale() Localable {
	return l
}
