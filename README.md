[![Build Status](https://github.com/harshadnawathe/geektrust-ride-sharing/actions/workflows/go-build.yml/badge.svg)](https://github.com/harshadnawathe/geektrust-ride-sharing/actions/workflows/go-build.yml)

# Geektrust Rise Sharing

This repository contains the solution to the
[Geektrust - Rise Sharing](https://www.geektrust.com/coding/detailed/ride-sharing)
challenge.

> [!NOTE]
> The repository structure is not as per the
> [guidelines](https://github.com/geektrust/coding-problem-artefacts/tree/master/Go)
> provided by the geektrust.

## Requirements

go@1.17+

## How to build ?

Run command below to build the code from the project directory.

```sh
go build .
```

## How to run tests?

Run command below to run tests from the project directory.

```sh
go test ./...
```

## How to run application ?

Once you build the code
you will find executable file named 'geektrust' in the current directory.

> Sample inputs
>
> Two sample inputs are provided in the [sample_input](./sample_input/) directory.

Run following command in the current directory to run the code with sample inputs.

```sh
./geektrust ./sample_input/input1.txt
```
