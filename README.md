# ansible-local-exec

[![Go Reference](https://pkg.go.dev/badge/github.com/jay-babu/ansible-local-exec.svg)](https://pkg.go.dev/github.com/jay-babu/ansible-local-exec)

## Completions

Run `ansible-local-exec completion {fish|bash|zsh|powershell}` to generate completions.

Run `ansible-local-exec help completion` for more information.

## Command

`ansible-local-exec --source apt.yml`.

Runs `ansible-playbook --connection local --inventory 127.0.0.1, -K SOURCE` under the hood.
