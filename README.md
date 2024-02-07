# Simple Bank

## Description

This is a simple bank application that allows users to create accounts, deposit money, withdraw money, and check their balance.

## How to Run?

To run the application make sure you have docker installed and run the following commands:

```bash
make start
```

this will expose the application on port 3000.

the application runs on a docker container, so it is possible to stop the application with the following command:

```bash
make stop
```

## How to Test?

To test the application run the following command:

```bash
make test
```

It is also possible to run the tests with coverage:

```bash
make test.cover
```

The coverage report will be available in the `coverage.html` file, where it is possible to open it in the browser to get a better understanding of the coverage.

