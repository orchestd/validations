package validations

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"bitbucket.org/HeilaSystems/servicereply"
	"context"
	"sort"
	"sync"
)

type CustomValidation interface {
	GetValidationTag() string
	GetValidationFunction() interface{}
}

type Builder interface {
	AddCustomValidations(...CustomValidation) Builder
	Build() (Validations, error)
}
type Validations interface {
	Validate(req interface{}) servicereply.ServiceReply
}

type Validator interface {
	Validate() error
	GetId() string
	GetName() string
	SetLevel(string)
	SetMessage(string)
}

type CacheValidator struct {
	Id      string
	Level   string
	Message string
	SortOrder    int
}

func NewCacheValidatorGetter(cache cache.CacheStorageGetter) ValidatorGetter {
	return &cacheValidatorGetter{cache: cache}
}

type ValidatorGetter interface {
	GetValidators(c context.Context, initValidators []Validator) ([]Validator, error)
}

type cacheValidatorGetter struct {
	cache           cache.CacheStorageGetter
	cacheValidators []CacheValidator
	sync.Mutex
}

func (v *cacheValidatorGetter) GetValidators(c context.Context, initValidators []Validator) ([]Validator, error) {
	v.Lock()
	defer v.Unlock()
	if v.cacheValidators == nil {
		validators := make(map[string]CacheValidator)
		err := v.cache.GetAll(c, "validators", "1", validators)
		if err != nil {
			return nil, err
		}
		var cacheValidators []CacheValidator
		for _, v := range validators {
			cacheValidators = append(cacheValidators, v)
		}
		sort.Slice(cacheValidators, func(i, j int) bool {
			return cacheValidators[i].SortOrder < cacheValidators[j].SortOrder
		})
		v.cacheValidators = cacheValidators
	}
	var validators []Validator
	for _, cv := range v.cacheValidators {
		for _, iv := range initValidators {
			if iv.GetId() == cv.Id {
				iv.SetLevel(cv.Level)
				iv.SetMessage(cv.Message)
				validators = append(validators, iv)
			}
		}
	}
	return validators, nil
}
