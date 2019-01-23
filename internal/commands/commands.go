package commands

import kingpin "gopkg.in/alecthomas/kingpin.v2"

func Register(app *kingpin.Application, version string) {
	(&pingCmd{}).register(app)
	(&versionCmd{}).register(app, version)
}
