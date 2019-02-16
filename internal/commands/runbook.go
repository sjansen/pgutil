package commands

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type runbookListCmd struct {
	file string
}

type runbookRunCmd struct {
	file string

	//dryRun      bool
	//interactive bool
	//listTasks   bool
	//tasks       []string
}

var runbookHelp = `Run books contain collections of tasks and their dependencies.

By default, all tasks in a run book are executed, and progress is reported
to stdout and stderr.
`

func (c *runbookListCmd) register(parent *kingpin.CmdClause) {
	cmd := parent.Command("ls", "List all tasks in a run book.").
		Action(c.run)
	cmd.Arg("FILE", "A run book filename").Required().
		ExistingFileVar(&c.file)
}

func (c *runbookRunCmd) register(parent *kingpin.CmdClause) {
	cmd := parent.Command("run", runbookHelp).
		Action(c.run).Default()
	cmd.Arg("FILE", "A run book filename").Required().
		ExistingFileVar(&c.file)
	/* TODO
	cmd.Flag("dry-run", "Print the tasks that would be run, but do not run them").
		Short('n').BoolVar(&c.dryRun)
	*/
	/* TODO
	cmd.Flag("interactive", "Request confirmation before running a task").
		Short('i').BoolVar(&c.interactive)
	*/
	/* TODO
	cmd.Flag("ls", "List the tasks in a run book").
		BoolVar(&c.listTasks)
	*/
	/* TODO
	cmd.Flag("task", "Run a specific task and its dependencies").
		Short('t').StringsVar(&c.tasks)
	*/
}

func (c *runbookRunCmd) run(pc *kingpin.ParseContext) error {
	fmt.Println("run:", c.file)
	return nil
}

func (c *runbookListCmd) run(pc *kingpin.ParseContext) error {
	fmt.Println("list:", c.file)
	return nil
}
