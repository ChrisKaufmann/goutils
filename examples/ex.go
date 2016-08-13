package main

import e "github.com/ChrisKaufmann/goutils"
import "github.com/msbranco/goconfig"

func main() {
	print("e.Evenodd(1):" + e.Evenodd(1) + "\n")
	print("e.Evenodd(0):" + e.Evenodd(0) + "\n")

	print("e.Tostr(32):" + e.Tostr(32) + "\n")
	print("e.Tostr(false):" + e.Tostr(false) + "\n")
	print("e.Tostr(true):" + e.Tostr(true) + "\n")

	print("e.Toint(\"0\"):" + e.Tostr(e.Toint("0")) + "\n")

	print("\n")

	print("arr := []int{1,2,3,4,5}\n")
	arr := []int{1, 2, 3, 4, 5}
	print("e.Shuffle(&arr):")
	e.Shuffle(&arr)
	for _, i := range arr {
		print(i)
		print(",")
	}
	print("\n")
	print("\n")

	print("ars := []string{\"a\",\"b\",\"c\",\"d\",\"e\"}\n")
	ars := []string{"a", "b", "c", "d", "e"}
	print("e.Shuffle(&ars):")
	e.Shuffle(&ars)
	for _, i := range ars {
		print(i)
		print(",")
	}
	print("\n")
	print("\n")
	print("e.GetHash(\"hello\"):" + e.GetHash("Hello") + "\n")
	print("e.RandomString(32):" + e.RandomString(32) + "\n")

	// use ew for error warning and ef for error fatal
	// e.g.
	_, err := goconfig.ReadConfigFile("config")
	print("e.Ew(err):")
	e.Ew(err)
	print("\n")
	//	e.Ef(err)
}
