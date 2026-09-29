#!/bin/env zsh

#compdef why

_why() {
    # get the input from the user
    local input=$words[$CURRENT]

    # subcommands completion
    if (( CURRENT == 2 )); then
        local -a commands
        commands=(
          ${(s: :)"$(why list-commands)"}
        )

        _describe 'command' commands
    fi

    # flags completion
    if [[ "$input" == -* ]]; then
        local -a flags
        flags=(
            ${(s: :)"$(why list-flags)"}
        )

        _describe 'option' flags
    fi
}

# registering the function
compdef _why why
