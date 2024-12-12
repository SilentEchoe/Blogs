package main

type InstrumentableType int

const (
	InstrumentableGolang InstrumentableType = iota + 1
	InstrumentableJava
	InstrumentableDotnet
	InstrumentablePython
	InstrumentableRuby
	InstrumentableNodejs
	InstrumentableRust
	InstrumentableGeneric
	InstrumentablePHP
)
