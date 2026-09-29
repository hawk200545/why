package completion

import (
	"slices"
	"testing"
)


func TestCompleteSelectShell(t *testing.T){
	list := [...]string {"bash", "zsh", "fish"}
	
	for _, shell := range list{
		completionScript := Complete(shell)
		if(len(completionScript) == 0){
			t.Errorf("Error in %s completion script", shell)
		}
	}
}

func TestListCommands(t *testing.T){
	subCommands := [...]string{"ssh", "curl", "tcp", "tls", "dns", "http", "completion"}

	got:= CompleteSubCommands()
	
	for _, cmd := range subCommands{
		if(!slices.Contains(got, cmd)){
			t.Errorf("Error: subcommand %s is missing in list-commands", cmd)
		}
	}
}

func TestListFlags(t *testing.T){
	flags := [...]string{ "-i", "-p", "-X", "-v", "-H", "--ai", "--deep", "--explain", "--timeout", "--json", "--version", "--verbose"}
	
	got := CompleteFlags()

	for _, flag := range flags{
		if(!slices.Contains(got, flag)){
			t.Errorf("Error: flag %s is missing in list-flags", flag)
		}
	}
}

