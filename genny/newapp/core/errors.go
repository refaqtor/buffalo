package core

import "github.com/pkg/errors"

// ErrGoModulesWithDep is thrown when trying to use both dep and go modules
var ErrGoModulesWithDep = errors.New("dep and modules can not be used at the same time")

// ErrNotInGoPath is thrown when not using go modules outside of GOPATH
var ErrNotInGoPath = errors.New("currently not in a $GOPATH")
