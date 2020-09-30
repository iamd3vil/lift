<a href="https://zerodha.tech"><img src="https://zerodha.tech/static/images/github-badge.svg" align="right" /></a>

# lift

A CLI Client for Liftbridge

## Usage

```
$ lift -h
NAME:
   lift - A CLI Client for Liftbridge

USAGE:
   lift [global options] command [command options] [arguments...]

AUTHOR:
   Sarat Chandra <me@saratchandra.in>

COMMANDS:
   stream, s  Commands about creating/deleting/consuming from streams
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose      Enable verbose logging (default: false)
   --config value, -c value  Configuration Path
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Configuration

Configuration path can be passed to `lift` by using `--config` flag. If nothing is given `lift` expects the path to be at `$HOME/.config/lift/config.toml`.

A sample config can be generated by running `lift init` which generates the following config at the given path:

```toml
[servers]
addresses = ["localhost:9292"]
```
