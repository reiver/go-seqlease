package seqleasedriver

import (
	"sync"
)

type internalRegistrar struct {
	mutex sync.RWMutex
	lessorCreators map[string]LessorCreator
}

func (receiver *internalRegistrar) Lessor(lessorCreatorName string, params ...interface{}) (Lessor, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	lessorCreators := receiver.lessorCreators
	if nil == lessorCreators {
		return nil, errNotFound
	}

	lessorCreator, ok := receiver.lessorCreators[lessorCreatorName]
	if !ok {
		return nil, errNotFound
	}
	if nil == lessorCreator {
		return nil, errInternalError
	}

	lessor, err := lessorCreator.Lessor(params...)
	if nil != err {
		return nil, err
	}

	return lessor, nil
}

func (receiver *internalRegistrar) Register(lessorCreatorName string, lessorCreator LessorCreator) error {
	if nil == receiver {
		return errNilReceiver
	}


	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == lessorCreator {
		return errBadRequest
	}

	if nil == receiver.lessorCreators {
		receiver.lessorCreators = map[string]LessorCreator{}
	}
	if _, ok := receiver.lessorCreators[lessorCreatorName]; ok {
		return errFound
	}

	receiver.lessorCreators[lessorCreatorName] = lessorCreator

	return nil
}
