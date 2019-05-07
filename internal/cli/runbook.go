package cli

import "github.com/sjansen/pgutil/internal/commands"

var runbookHelp = `Run books contain collections of tasks and their dependencies.

By default, all tasks in a run book are executed, and progress is reported
to stdout and stderr.
`

func registerRunbook(p *ArgParser) {
	parent := p.addParent("runbook")

	// == ls ==
	list := &commands.RunBookListCmd{}
	cmd := p.addSubCommand(parent, list, "list", "List all tasks in a run book").
		Alias("ls")
	cmd.Arg("FILE", "A run book filename").Required().
		ExistingFileVar(&list.File)

	// == dot ==
	dot := &commands.RunBookDotCmd{}
	cmd = p.addSubCommand(parent, dot, "dot",
		"Generate a GraphViz compatible description of a runbook's tasks",
	)
	cmd.Arg("FILE", "A run book filename").Required().
		ExistingFileVar(&dot.File)
	cmd.Flag("output", "Write to FILENAME instead of stdout").
		Short('o').PlaceHolder("FILENAME").StringVar(&dot.Output)
	cmd.Flag("", "Choose output filename based on the input filename").
		Short('O').BoolVar(&dot.AutoOut)
	cmd.Flag("splines", "Control how edges are represented").
		EnumVar(&dot.Splines, "curved", "line", "ortho")

	// == run ==
	run := &commands.RunBookRunCmd{}
	cmd = p.addSubCommand(parent, run, "run", runbookHelp).
		Default()
	cmd.Arg("FILE", "A run book filename").Required().
		ExistingFileVar(&run.File)
	/* TODO
	cmd.Flag("dry-run", "Print the tasks that would be run, but do not run them").
		Short('n').BoolVar(&run.DryRun)
	*/
	/* TODO
	cmd.Flag("interactive", "Request confirmation before running a task").
		Short('i').BoolVar(&run.Interactive)
	*/
	/* TODO
	cmd.Flag("task", "Run a specific task and its dependencies").
		Short('t').StringsVar(&run.Tasks)
	*/
}
