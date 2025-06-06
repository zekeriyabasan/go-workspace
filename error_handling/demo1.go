package errorhandling

import (
	"fmt"
	"io"
	"os"
)

func OpenFile(path string) {
	f, err := os.Open(path)

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Printf("HATA: '%s' dosyasına erişilemedi (%s işlemi)\n", pErr.Path, pErr.Op)
			return
		} else {
			fmt.Println("Error opening file:", err)
		}
	}

	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}
