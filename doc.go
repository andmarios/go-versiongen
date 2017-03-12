// Copyright 2017 Marios Andreopoulos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
  Package versiongen provides an opinionated method to add version information
  (version number and git commit hash) to go projects via 'go generate'. It
  produces a go file that includes a vgVersion constant with the version of
  your software and a vgHash constant with the SHA1 of the HEAD at build time.

  It assumes:

  1. your go project uses git for source control

  2. you have at least one tagged commit (e.g v0.1)

  The file it creates (version.go by default) looks like:
    package main
    const (
            vgVersion   = "v0.1-2-g7665642+"
            vgHash      = "7665642a17cecca2693d48276b07c28bd29cda06"
            vgClean     = false
    )

  Look into the example subdirectory to see how you can use it:
    git clone github.com/andmarios/go-versiongen
    cd go-versiongen/example
    ls
    go generate
    go run main.go version.go
*/
package versiongen
