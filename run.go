package main

func Run(args []string) error {
	cmd := NewBcryptGenerateCmd()
	return cmd.Execute(args)
}
