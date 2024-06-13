package main

import "fmt"

func main() {
	v1 := [3]float64{1, 3, 4}
	var v2 [3]float64
	v2 = [3]float64{2, 4, 6}
	for v3, i := [...]float64{0, 0, 0}, 0; i < len(v3); i++ {
		v3[i] = v1[(i+1)%3]*v2[(i+2)%3] - v1[(i+2)%3]*v2[(i+1)%3]
		defer fmt.Printf("%t\n", v3)
		fmt.Printf("%v\n", v3[i])
	}

	vectors := []struct {
		x, y, z float64
	}{
		{1, 2, 3},
		{3.2, 4, 6},
		{4, 3, 1},
	}

	fmt.Printf("type %#T and value %v\n", vectors, vectors)
	vectors = append(vectors, struct{ x, y, z float64 }{7, 7, 7})
	fmt.Printf("type %#T and value %v\n", vectors[3:], vectors[3:])
	fmt.Printf("type %#T and value %v\n", vectors[3], vectors[3])

	for i, v := range vectors {
		fmt.Println(i, " : ", v)
	}
	numbers := make([]int, 10)
	fmt.Println(numbers)

	for i := range numbers {
		numbers[i] = i
	}

	fmt.Println(numbers)

	age := map[string]int{"max": 24, "tom": 28}
	fmt.Println("map: ", age)

	m := make(map[string]float64)
	m["E"] = 2.7
	m["Pi"] = 3.14
	m["Phi"] = 1.61

	for key, v := range m {
		fmt.Printf("Key: %v, Value: %v, Value: %v \n", key, v, m[key])
	}

	delete(m, "E")
	fmt.Println("len: ", len(m))
	fmt.Println("map: ", m)

	_, ok := m["E"]
	fmt.Println("ok: ", ok)
}
