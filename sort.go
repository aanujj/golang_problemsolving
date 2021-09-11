import "fmt"

//without using sort library

func BubbleSort(elements []int) []int {
	keeprunning := true

	for keeprunning {
		keeprunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keeprunning = true
			}
		}
	}
	return elements
}

func main() {
	elements := []int{101, 20, 90, 29, 1}
	a := BubbleSort(elements)
	fmt.Println(a)
}
