# locale

Simple wrapper for [go-i18n](https://github.com/nicksnyder/go-i18n)

## Installation

Before use this validator, you need to install [go-i18n](https://github.com/nicksnyder/go-i18n) first, and then install this package.

```cmd
go get github.com/ramadani/locale
```

## Usage

Register the package

```go
locale.Register(func() *i18n.Localizer {
	bundle := &i18n.Bundle{DefaultLanguage: language.Indonesian}
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("resources/lang/id-ID.yml")

	return i18n.NewLocalizer(bundle, "id-ID")
})
```

And then, you can use it on your go project

```go
trans := locale.Locale()
subject := trans.Trans("password_reset_mail_subject")
// Password Reset Confirmation
```
