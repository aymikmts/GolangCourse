#!/bin/bash
# TODO: 値のチェックしてない。有理数か・正の値か、など

cd lengthconv
go test

cd ../weightconv
go test

