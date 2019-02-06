## About

`lambda2pi` is a toy program that converts a lambda-calculus term into a pi-calculus term.

## Usage

```sh
lambda2pi <input file> [-o <output file>]
```

The input file contains a lambda term, and the output file will contain a pi term.

Also, you can run lambda2pi in REPL mode when no argument is given:

```sh-session
$ lambda2pi
>> 
```

Look at `lambda2pi --help` for detailed syntax for lambda-calculus and pi-calculus.

## Example

Let's convert a very simple lambda term `\x. x` (`example/simpleAbs.lambda`):

```sh-session
$ lambda2pi ./example/simpleAbs.lambda -o out.pi
$ cat out.pi
pp!yy0.(*yy0?ww.ww?x.(new pp in pp!yy1.(*yy1?ww.x!ww.O)))
```

## Installation

### Prerequisites

* Go >= 1.11
* goyacc: Install by running `go install golang.org/x/tools/cmd/goyacc`

### Manual Build

```sh-session
$ go get -u github.com/nekketsuuu/lambda2pi
$ # cd to $GOPATH/src/github.com/nekketsuuu/lambda2pi
$ go generate
$ go build
$ go install
```

## References

* Robin Milner, "Functions as processes", 1990

## Author

nekketsuuu

## License

The MIT License. See `LICENSE` for details.
