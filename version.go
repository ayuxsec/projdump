package main

import "runtime/debug"

var Version = "dev"

func VersionString() string {
	if Version != "dev" {
		return Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}
	return Version
}
