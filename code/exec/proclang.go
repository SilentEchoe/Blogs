package main

import (
	"regexp"
	"strings"
)

var rubyModule = regexp.MustCompile(`^(.*/)?ruby[\d.]*$`)
var pythonModule = regexp.MustCompile(`^(.*/)?python[\d.]*$`)

func instrumentableFromModuleMap(moduleName string) InstrumentableType {
	if strings.Contains(moduleName, "libcoreclr.so") {
		return InstrumentableDotnet
	}
	if strings.Contains(moduleName, "libjvm.so") {
		return InstrumentableJava
	}
	if strings.HasSuffix(moduleName, "/node") || moduleName == "node" {
		return InstrumentableNodejs
	}
	// 通过正则表达
	if rubyModule.MatchString(moduleName) {
		return InstrumentableRuby
	}
	if pythonModule.MatchString(moduleName) {
		return InstrumentablePython
	}

	return InstrumentableGeneric
}
