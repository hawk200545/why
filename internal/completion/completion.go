package completion

import (
	_ "embed"
)

// embedding the files at compiletime

//go:embed bash-completions.bash
var bashCompletion string
//go:embed zsh-completions.zsh
var zshCompletion string
//go:embed fish-completions.fish
var fishCompletion string

func CompleteSubCommands() []string {
	return []string{
		"curl",
		"dns",
		"http",
		"tcp",
		"ssh",
		"tls",
		"completion",
	}
}

func CompleteFlags() []string {
	return []string{
		"--deep",
		"--ai",
		"--explain",
		"--json",
		"--verbose",
		"--version",
		"--timeout",
		"-v",
		"-i",
		"-p",
		"-X",
		"-H",
	}
}

func Complete(shell string) string {
	switch shell {
		case "bash":
				return bashCompletion
		case "zsh":
				return zshCompletion
		case "fish":
				return fishCompletion
		default:
				return "echo Invalid or unsupported shell completion is sourced for why completions"
		}

}
