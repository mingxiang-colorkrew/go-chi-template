package config

import "github.com/gookit/validate"

func setupGlobalValidation() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
	})
}
