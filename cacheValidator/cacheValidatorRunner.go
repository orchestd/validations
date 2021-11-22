package cacheValidator

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	. "bitbucket.org/HeilaSystems/servicereply"
	. "bitbucket.org/HeilaSystems/validations"
	"context"
	"sort"
	"sync"
)

func NewCacheValidatorRunner(cache cache.CacheStorageGetter) ValidatorRunner {
	return &cacheValidatorRunner{cache: cache}
}

type cacheValidatorRunner struct {
	cache           cache.CacheStorageGetter
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
		err := cvg.cache.GetAll(c, "validators", "1", validators)
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
		cvg.cacheValidators = cacheValidators
	}
	var validators []Validator
	for _, cv := range cvg.cacheValidators {
		for _, iv := range initValidators {
			if iv.GetId() == cv.Id {
				validators = append(validators, iv)
			}
		}
	}
	return validators, nil
}
