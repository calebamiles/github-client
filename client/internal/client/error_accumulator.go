package client

import (
	"fmt"
	"strings"
	"sync"
)

type errorAccumulator struct {
	sync.Mutex

	errorList []error
}

func (e *errorAccumulator) Add(err error) {
	e.Lock()
	defer e.Unlock()

	e.errorList = append(e.errorList, err)
}

func (e *errorAccumulator) IsNil() bool {
	e.Lock()
	defer e.Unlock()

	return (len(e.errorList) == 0)
}

func (e *errorAccumulator) Error() string {
	e.Lock()
	defer e.Unlock()

	if len(e.errorList) == 0 {
		return ""
	}

	errorListStrings := []string{}
	for i := range e.errorList {
		errorListStrings = append(errorListStrings, e.errorList[i].Error())
	}

	return fmt.Sprintf("the following errors occurred: %s", strings.Join(errorListStrings, "\n"))
}
