package ds_

import "github.com/darksubmarine/versatile/failure"

type Error = *failure.ExtendedError

func NewError(err error) *failure.ExtendedError {
	return failure.NewExtendedError(err)
}
