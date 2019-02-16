package commands

import kingpin "gopkg.in/alecthomas/kingpin.v2"

func Register(app *kingpin.Application, version string) {
	(&pingCmd{}).register(app)
	parent := app.Command("runbook", "")
	(&runbookListCmd{}).register(parent)
	(&runbookRunCmd{}).register(parent)
	(&versionCmd{}).register(app, version)
}
