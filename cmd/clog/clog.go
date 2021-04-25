package clog

import (
	"log"

	"github.com/i582/cfmt/cmd/cfmt"
)

// Printf is the same as log.Printf.
func Printf(format string, a ...interface{}) {
	log.Print(cfmt.Sprintf(format, a...))
}

// Print is the same as log.Print.
func Print(a ...interface{}) {
	log.Print(cfmt.Sprint(a...))
}

// Println is the same as log.Println.
func Println(a ...interface{}) {
	log.Println(cfmt.Sprint(a...))
}

// Fatalf is the same as log.Fatalf.
func Fatalf(format string, a ...interface{}) {
	log.Fatal(cfmt.Sprintf(format, a...))
}

// Fatal is the same as log.Fatal.
func Fatal(a ...interface{}) {
	log.Fatal(cfmt.Sprint(a...))
}

// Fatalln is the same as log.Fatalln.
func Fatalln(a ...interface{}) {
	log.Fatalln(cfmt.Sprint(a...))
}

// Panicf is the same as log.Panicf.
func Panicf(format string, a ...interface{}) {
	log.Panic(cfmt.Sprintf(format, a...))
}

// Panic is the same as log.Panic.
func Panic(a ...interface{}) {
	log.Panic(cfmt.Sprint(a...))
}

// Panicln is the same as log.Panicln.
func Panicln(a ...interface{}) {
	log.Panicln(cfmt.Sprint(a...))
}
