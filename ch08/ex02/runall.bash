#!/bin/bash
cd `dirname $0`

go build -o FTPServer
sudo ./FTPServer &