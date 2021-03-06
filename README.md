# genstruct
![build](https://github.com/fifsky/genstruct/workflows/build/badge.svg)

Golang struct generator from mysql schema

[![asciicast](https://asciinema.org/a/12i6QmbaUCQgPZ4o2rz5QmPVE.png)](https://asciinema.org/a/12i6QmbaUCQgPZ4o2rz5QmPVE)

## Install

```
go install github.com/fifsky/genstruct
```

## Usage

```
genstruct -h localhost -u root -p 3306 -P 123456
```

* `-h` default `localhost`
* `-u` default `root`
* `-p` default `3306`

## online

https://go.fifsky.com/

## gosql

The structure can be applied to a [gosql](https://github.com/ilibs/gosql) package

## License

The source code is available under the MIT [License](/LICENSE).