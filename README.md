# gfm-serve

A simple [serve] like utility to list files and render GitHub Flavored Markdown.

## Installation

```sh
$ go install github.com/fa7ad/gfm-serve/cmd/gfm-serve
```

Or, clone the repo and build the binary.

```sh
$ git clone https://github.com/fa7ad/gfm-serve.git
$ cd gfm-serve
$ make
```

The binary will be available in the `out` directory.

## Usage

```sh
$ gfm-serve --help
NAME:
   gfm-serve - Serve GitHub Flavored Markdown files

USAGE:
   gfm-serve [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value, -P value       Port number to listen on (default: 8080)
   --directory value, -d value, -D value  Path to serve files from (default: ".")
   --addr value, -a value, -A value       Hostname/Address to listen on (default: "localhost")
   --help, -h                             show help
```

[serve]: https://npm.im/serve
