# hardgo

Design hardware with golang.

## Introduction

Compiles golang to verilog and vhdl, but allows simulation of the golang code.

Introduces a new language called `hardgo` (with extension `.hgo`) that compiles to verilog and vhdl.
Similar to [templ](https://templ.guide/) but for hardware rather than html.

## Installation

Get the cli.

```bash
go install github.com/conneroisu/gohard/cmd/gohard@latest
```

Get the package.

```bash
go get github.com/conneroisu/gohard
```

## Usage

```bash
gohard -h
```

## Example

```bash
gohard -o test.v -i test.go
```
