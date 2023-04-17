package handler

type PackageHandler interface {
}

type packageHandler struct {
}

func NewPackageHandler() PackageHandler {
	return &packageHandler{}
}
