package api

func (a *api) Start(port int) error {
	return a.App.Listen(port)
}
