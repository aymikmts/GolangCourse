#!/bin/bash

cd popcountroop
go test -bench=.

cd ../popcount
go test -bench=.