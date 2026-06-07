#!/usr/bin/env bash
set -euo pipefail

mkdir -p ./build

build() {

	local goos="$1"
	local goarch="$2"

	local output="antimiasma_${goos}_${goarch}";

	if [[ "${goos}" == "windows" ]]; then
		output="${output}.exe";
	fi

	echo "Building ${goos}/${goarch} to ./build/${output}";

	env CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" go build -o "./build/${output}" ./cmds/antimiasma;

}

build linux amd64;
build linux arm64;
build darwin amd64;
build darwin arm64;
build windows amd64;

