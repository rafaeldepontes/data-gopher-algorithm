package main

func main() {
	capacity := 3
	hashMap := make(map[int]string, capacity)

	hashMap[0] = "A"
	hashMap[1] = "B"
	hashMap[2] = "C"

	for index, value := range hashMap {
		println("Index:", index, "- Value", value)
	}
}
