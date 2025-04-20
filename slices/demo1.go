package slices

func Demo1() {
	names := make([]string, 3)
	names = append(names, "Mustafa")
	names = append(names, "Zekeriya")
	println(len(names))
	for i := 0; i < len(names); i++ {
		if i < 3 {
			names[i] = "DEFAULT"
		}
	}

	names = append(names, "Ali")
	names = append(names, "Veli")
	println(len(names))

	names2 := []string{"Celal"}
	println(names2[0])
}
