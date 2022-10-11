#!/bin/bash

echo -ne "\e[36m###########################\e[0m\n"
echo -ne "\e[36m#                         #\e[0m\n"
echo -ne "\e[36m#\e[0m   GENERATING COVERAGE   \e[36m#\e[0m\n"
echo -ne "\e[36m#                         #\e[0m\n"
echo -ne "\e[36m###########################\e[0m\n"

go test -coverprofile=coverage.out;
go tool cover -html=coverage.out;


echo -ne "\e[36m#######################\e[0m\n"
echo -ne "\e[36m#                     #\e[0m\n"
echo -ne "\e[36m#\e[0m   COVER GENERATED   \e[36m#\e[0m\n"
echo -ne "\e[36m#                     #\e[0m\n"
echo -ne "\e[36m#######################\e[0m\n"

