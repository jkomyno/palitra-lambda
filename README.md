# palitra

[![Build Status](https://travis-ci.org/jkomyno/palitra-lambda.svg?branch=master)](https://travis-ci.org/jkomyno/palitra-lambda) [![GoDoc](https://godoc.org/github.com/jkomyno/palitra-lambda?status.svg)](https://godoc.org/github.com/jkomyno/palitra-lambda) [![Go Report Card](https://goreportcard.com/badge/github.com/jkomyno/palitra-lambda)](https://goreportcard.com/report/github.com/jkomyno/palitra-lambda)

## Description

This repository hosts the code for `palitra-lambda`, an AWS Lambda function that given a picture and a certain limit `N`, returns the `N` most frequent colors in a picture, approximated to CSS3 color names.

```bash
.
├── LICENSE                     <-- MIT License file
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── palitra                     <-- Source code for lambda function + test
├── Gopkg.lock                  <-- Go dep dependencies exact versions
├── Gopkg.toml                  <-- Go dependencies requirements
└── template.yaml
```

## Available Scripts

This scripts are accessible in the Makefile.

- `install`: Installs the function's dependencies
- `format`: Formats the code according to the Go standard
- `build`: Builds an optimized binary for 64-bit Linux in `./bin/palitra`
- `zip`: Packages the binary into a zip file ready to be uploaded to AWS
- `clean`: Removes the compiled binary and the zip file
- `clean-all`: Removes the Go dependency folder created by `dep`
- `test`: Runs the Go integration tests

## Related packages

- [palitra](https://github.com/jkomyno/palitra): this is the original Golang CLI utility that has been used as the base of this Lambda function.

## Make an AWS Lambda function yourself

How can you generate a Golang lambda function similar to this one?
Just type in `sam init -r go1.x -n awesome-unicorns` and there you go.

It requires you to have AWS SAM cli installed. You can installed via the following:

`python3 -m pip install aws-sam-cli`

Check that the script is installed properly by typing `which sam`.
