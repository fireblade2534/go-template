#!/bin/bash
go build -buildmode=c-shared  -o bind/template.so go_template/bind/*.go