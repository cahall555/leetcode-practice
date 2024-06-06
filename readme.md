# Leet code practice

This repository is to help practice and implement leet code.

## Javascript

Each file is a leetcode problem, tests have been written in jest.

to test run: `pnpn test` or `pnpm test [FILENAME].test.js`

## php

Each file has has a test written in PHPUnit

to test run: `./tools/phpunit.phar tests/unit/[FILENAME].php`
or run: `./tools/phpunit.phar --testsuite unit` to test everything

## Go

Each file has a test written in go. Tests in go are in files ending with `_test.go`

to test run: `go test` in the folder of the file you want to test


# Big O Notation Notes

better                                                           worse

<---------------------------------------------------------------->

O(1)      O(log(n))    O(n)      O(n^c)      O(c^n)        O(n!)

constant  logarithmic  linear    polynomial  exponential   factorial
