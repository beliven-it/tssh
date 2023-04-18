<br>
<p align="center"><img src="./assets/tssh.svg" /></p>
<br>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/beliven-it/tssh?color=512fc9&style=for-the-badge" />
<img src="https://img.shields.io/github/v/release/beliven-it/tssh?color=512fc9&style=for-the-badge" />
<img src="https://img.shields.io/github/license/beliven-it/tssh?color=512fc9&style=for-the-badge" />
</p>
<p align="center">
<img src="https://img.shields.io/github/issues-pr/beliven-it/tssh?color=512fc9&style=for-the-badge" />
<img src="https://img.shields.io/github/issues/beliven-it/tssh?color=512fc9&style=for-the-badge" />
<img src="https://img.shields.io/github/contributors/beliven-it/tssh?color=512fc9&style=for-the-badge" />
</p>

A CLI to easily list, search and connect to SSH hosts via GoTeleport service.
This CLI is just a wrapper around `tsh` command provided by Goteleport application.

## Install

Add Homebrew Beliven tap with:

```bash
  brew tap beliven-it/tap
```

Then install `tssh` CLI with:

```bash
  brew install beliven-it/tap/tssh
```

## Configuration

Run `tssh init` to generate config file inside `~/.config/tssh/config.yml` (works only if not exists yet) or let the CLI creating it automatically on first run (every command).

### fzf options

See the man page (`man fzf`) for the full list of available options and add the desired ones to the `fzf_options` string inside `~/.config/tssh/config.yml`. See more about the fzf options in the [official repository](https://github.com/junegunn/fzf#options).

### Config file example

This is a complete config file example with two providers:

```yaml
# TSSH configuration file
fzf_options: "-i"

# Admin users can be identifier by a role.
# For security reasons this role is default set to empty.
# If you are a admin user, set these values below.
# 
# When a user have the role declared the system force the 
# user to connect with a specific privileged user defined by 
# the admin_user variable.
#
# example:
#
# admin_role: "sysadmin"
# admin_user: "root"
#
# launching tssh c the results are a list of the host the user can connect
# and instead to use the role convension host.user, the user will be replaced
# by root value.
admin_role: "<YOUR PRIVILEGED ROLE IDENTIFIER>"

# The privileged user. Usually is root
admin_user: "root"

# TSSH Easily allow to login or logout into the cluster.
# compile the following values for handle login actions.
teleport_proxy: "teleport.domain.com" 
teleport_user: "my_user"

# If you need to use the passwordless feature by default you can enable with this setting
teleport_passwordless: true
```

## Usage

To see available commands and options, run: `tssh`, `tssh help`, `tssh --help` or `tssh -h`.

## Development

Clone the repository and run inside the folder:

- `go mod init tssh`
- `go mod vendor`
- `go build -ldflags="-X tssh/cmd.Version=1.0.0"`

Run `./tssh` inside the folder to test the CLI.

## Use as normal SFTP, SSH and other services

You can also use as normal behavior. 

## Have found a bug?

Please open a new issue [here](https://github.com/beliven-it/tssh/issues).

## License

Licensed under [MIT](./LICENSE)
