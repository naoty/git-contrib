# git-contrib
Show contributors for specified files.

## Installation

```
$ go get github.com/naoty/git-contrib
```

## Usage

```
$ git contrib main.go
Naoto Kaneko
```

```
$ git contrib main.go hello.go
Naoto Kaneko
```

`-n` option helps to specify the number of contributors.

```
$ git contrib -n 3 main.go hello.go
Nobunaga Oda
Hideyoshi Toyotomi
Ieyasu Tokugawa
```

## Author

[naoty](https://github.com/naoty)

