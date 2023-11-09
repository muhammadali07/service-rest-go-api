package app

import "fmt"

func (app *MiddlewareApplication) GetIndex(params string) (respApp string, err error) {
	respApp = fmt.Sprintf("Hello, world : %v", params)
	return
}
