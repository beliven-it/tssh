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
  brew install tssh
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
```

## Usage

To see available commands and options, run: `tssh`, `tssh help`, `tssh --help` or `tssh -h`.

## Development

Clone the repository and run inside the folder:

- `go mod init tssh`
- `go mod vendor`
- `go build -ldflags="-X tssh/cmd.Version=1.0.0"`

Run `./tssh` inside the folder to test the CLI.

## Have found a bug?

Please open a new issue [here](https://github.com/beliven-it/tssh/issues).

## License

Licensed under [MIT](./LICENSE)
