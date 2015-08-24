# contrib 
Show contributors for specified files.

## Installation

```
$ go get github.com/naoty/contrib
```

## Usage

```
$ contrib main.go
Naoto Kaneko
```

```
$ contrib main.go hello.go
Naoto Kaneko
```

`-n` option helps to specify the number of contributors.

```
$ contrib -n 3 main.go hello.go
Nobunaga Oda
Hideyoshi Toyotomi
Ieyasu Tokugawa
```

## Author

[naoty](https://github.com/naoty)

