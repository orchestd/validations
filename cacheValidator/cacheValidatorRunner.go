package cacheValidator

import (
	"context"
	"github.com/orchestd/dependencybundler/interfaces/cache"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	. "github.com/orchestd/servicereply"
	. "github.com/orchestd/validations"
	"sort"
	"sync"
)

func NewCacheValidatorRunner(cache cache.CacheStorageGetter, conf configuration.Config) ValidatorRunner {
	return &cacheValidatorRunner{cache: cache, conf: conf}
}

type cacheValidatorRunner struct {
	cache           cache.CacheStorageGetter
	conf            configuration.Config
	cacheValidators []CacheValidator
	sync.Mutex
}

func (cvg *cacheValidatorRunner) Validate(c context.Context, validatorCont ValidatorCont) ServiceReply {
	sortedValidators, err := cvg.getSortedValidators(c, validatorCont.Validators)
	if err != nil {
		return NewInternalServiceError(err)
	}
	var notRejectedReply ServiceReply
	for _, v := range sortedValidators {
		rep := v.Validate()
		if rep != nil {
			if !rep.IsSuccess() {
				return rep
			} else if notRejectedReply == nil {
				notRejectedReply = rep
			}
		}
	}
	return notRejectedReply
}

func (cvg *cacheValidatorRunner) getSortedValidators(c context.Context, initValidators []Validator) ([]Validator, error) {
	cvg.Lock()
	defer cvg.Unlock()
	if cvg.cacheValidators == nil {
		validators := make(map[string]CacheValidator)
		serviceName, err := cvg.conf.GetServiceName()
		if err != nil {
			return nil, err
		}
		err = cvg.cache.GetAll(c, "validators", serviceName, validators)
		if err != nil {
			return nil, err
		}
		var cacheValidators []CacheValidator
		cacheValidators = make([]CacheValidator, 0, 0)
		for _, v := range validators {
			cacheValidators = append(cacheValidators, v)
		}
		sort.Slice(cacheValidators, func(i, j int) bool {
			return cacheValidators[i].SortOrder < cacheValidators[j].SortOrder
		})
		cvg.cacheValidators = cacheValidators
	}
	var validators []Validator
	for _, iv := range initValidators {
		foundInCache := false
		for _, cv := range cvg.cacheValidators {
			if iv.GetId() == cv.Id {
				foundInCache = true
				if cv.Enabled {
					validators = append(validators, iv)
				}
			}
		}
		if !foundInCache && iv.GetIsEnabledByDefault() {
			validators = append(validators, iv)
		}
	}
	return validators, nil
}
