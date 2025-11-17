package main

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string) {

//		return "golang", "javascript", "c"
//	}
//
//	func processIt(fn func(a int) int) {
//		fn(1)
//	}
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)
	fn := processIt()
	fn(6)
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3	)
	// fmt.Println(getLanguages())
	// result := add(3, 5)
	// fmt.Println(result)
}
