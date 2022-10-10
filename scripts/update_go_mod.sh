#!/bin/bash

GITUSERREPOCMD=$(git config --get-regexp user.name)
GITUSERREPO=${GITUSERREPOCMD:(10)} #It's cutting the "user.name" from output

go mod init github.com/$GITUSERREPO/gconcat;
go mod tidy;
