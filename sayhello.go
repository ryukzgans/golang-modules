package belajar_golang

import "strconv"

func SayHello() string {
	return "Hello"
}

func SayHelloWithName(name string) string {
	return "Hello " + name
}

func SayHelloWithNameAndAge(name string, age int) string {
	return "Hello " + name + ", Age " + strconv.Itoa(age)
}
