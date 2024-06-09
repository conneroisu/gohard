# hardgo

design hardware with golang

Compiles golang to verilog and vhdl, but allows simulation of the golang code.

## Installation

Get the cli.

```bash
go install github.com/conneroisu/hardgo/cmd/hardgo@latest
```

Get the package.

```bash
go get github.com/conneroisu/hardgo
```

## Usage

```bash
hardgo -h
```

## Example

```bash
hardgo -o test.v -i test.go
```
