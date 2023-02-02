package bvalidator

import (
	"container/list"
	"github.com/orchestd/validations"
)

type validatorConfig struct {
	customValidations []validations.CustomValidation
}

type validatorBuilder struct {
	ll *list.List
}

func Builder() validations.Builder {
	return &validatorBuilder{
		ll: list.New(),
	}
}

func (v *validatorBuilder) AddCustomValidations(validations ...validations.CustomValidation) validations.Builder {
	v.ll.PushBack(func(cfg *validatorConfig) {
		cfg.customValidations = validations
	})
	return v
}

func (v *validatorBuilder) Build() (validations.Validations, error) {
	vCfg := &validatorConfig{}
	for e := v.ll.Front(); e != nil; e = e.Next() {
		f := e.Value.(func(cfg *validatorConfig))
		f(vCfg)
	}
	return NewValidator(vCfg)
}
