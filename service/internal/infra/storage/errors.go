package storage

import (
	appErrors "github.com/m11ano/neurochar-experiments-2/service/internal/app/errors"
)

var ErrBucketAlreadyExists = appErrors.ErrConflict.Extend("bucket already exists")
