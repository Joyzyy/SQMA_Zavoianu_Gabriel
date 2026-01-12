package gotests

func AddElement(arr []int, element int) []int {
	return append(arr, element)
}

func RemoveElement(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func GetElement(arr []int, index int) (int, bool) {
	if index < 0 || index >= len(arr) {
		return 0, false
	}
	return arr[index], true
}
