package app

type MiddlewareServicePort interface {
	GetIndex(param string) (respApp string, err error)
}

type MiddlewareDataStorePort interface {
}
