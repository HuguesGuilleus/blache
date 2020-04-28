#!/bin/bash

# Requirement:
# - Typescript compiler: tsc
# - github.com/go-bindata/go-bindata/
# - https://github.com/tdewolff/minify/

cd web/ts/
tsc --outFile /dev/stdout | minify --type=js > ../app.js
cd ../..

go-bindata -pkg=webData -o web/webData/data.go web/{index.html,style.css,*.js}
