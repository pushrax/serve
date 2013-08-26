# serve

Start serving a directory over HTTP (à la `python -m SimpleHTTPServer`).
Defaults to the current working directory and port 8080.

## Installation

[Install Go], and do `go install github.com/pushrax/serve`. Ensure your Go `bin` folder is in your path.

[Install Go]: http://golang.org/doc/install

## Usage

```
$ serve -help
Usage of serve:
  -path=".": Directory to serve
  -port=8080: Port to listen on
```

