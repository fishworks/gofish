package rig

import "errors"

var (
	// ErrExists indicates that a rig already exists
	ErrExists = errors.New("rig already exists")
	// ErrDoesNotExist indicates that a rig does not exist
	ErrDoesNotExist = errors.New("rig does not exist")
	// ErrHomeMissing indicates that the directory expected to contain rigs does not exist
	ErrHomeMissing = errors.New(`rig home "$(gofish home)/Rigs" does not exist`)
	// ErrMissingSource indicates that information about the source of the rig was not found
	ErrMissingSource = errors.New("cannot get information about the source of this rig")
	// ErrRepoDirty indicates that the rig repo was modified
	ErrRepoDirty = errors.New("rig repo is in a dirty git tree state so we cannot update. Try removing and adding this rig back")
	// ErrVersionDoesNotExist indicates that the request version does not exist
	ErrVersionDoesNotExist = errors.New("requested version does not exist")
)
