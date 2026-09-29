complete -c why -f

complete -c why \
    -n "not __fish_seen_subcommand_from (why list-commands | string split ' ')" \
    -a "(why list-commands | string split ' ')"

complete -c why \
    -n "string match -q -- '-*' (commandline -ct)" \
    -a "(why list-flags | string split ' ')"
