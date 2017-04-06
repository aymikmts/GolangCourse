#!/bin/bash

cd echo2    
go test -bench=.

cd ../echo3
go test -bench=.