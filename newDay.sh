#!/bin/bash

touch "./data/day$1"

echo -e "package main\n\nfunc day$1() Result{}" >> day$1.go