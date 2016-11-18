#!/bin/bash

GIT_VER=`git describe --tags`
if [ -z "${GIT_VER}" ]; then
    GIT_VER="[version unknown]"
fi
cat <<EOF > version.go
package main

var version string = \`${GIT_VER}\`

EOF
