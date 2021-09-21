package app

var (
	Base            Application = &App{}
)

// Application interface for start application
type Application interface {
	StartApplication()
}

type App struct{}

// StartApplication func
func (a *App) StartApplication() {}