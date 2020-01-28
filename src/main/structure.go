package main

cmd = newCommand(
	"version",
	&version,
	"show version",
)
func newCommand(name string, varref *int, comment string) *Command {
    return &Command{
        Name:    name,
        Var:     varref,
        Comment: comment,
    }
}

func main()  {
	
}

