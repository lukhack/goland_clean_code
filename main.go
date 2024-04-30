package main

import (
	"app/src/infraestructure/server"
	task "app/src/modules/task"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(task.ModuleProviders())

	app.Up()

}
