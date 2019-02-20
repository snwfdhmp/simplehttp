# Quickly serve a local directory over http

## Getting started

Install with

```
go get github.com/snwfdhmp/stupidhttp
```

## Usage

With no arguments, stupidhttp starts serving files under ./ over port 8080.

```
$ stupidhttp
INFO[0000] Serving ./ over 0.0.0.0:8080... Stop with ^C 
```

With `-d` arg, specify the directory to be served.

```
$ stupidhttp -d ./templates/
INFO[0000] Serving ./templates/ over 0.0.0.0:8080... Stop with ^C 
```


With `-p` arg, specify the port to serve on.

```
$ stupidhttp -p 31415
INFO[0000] Serving ./ over 0.0.0.0:31415... Stop with ^C 
```

## Feedback

Feel free to open an issue for any feedback or suggestion.

I fix bugs quickly.

## Contributions

PR are accepted as soon as they follow Golang common standards.
For more information: https://golang.org/doc/effective_go.html
