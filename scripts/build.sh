#!/bin/bash -e

go_mains=$(find ./cmd -name "main.go")

version=$(git describe --always)

for go_main in ${go_mains}; do
  bin_name="${go_main//\//_}"
  bin_name="${bin_name//_main.go/}"
  bin_name="${bin_name//._cmd_/}"
  echo "compiling ${bin_name}"
  CGO_ENABLED=0 go build -ldflags "-X github.com/VJftw/docker-registry-proxy/pkg/cmd.BuildVersion=${version}" -o "dist/${bin_name}" "${go_main}"
done
