package stringfunctions

import (
	"fmt"
	s "strings"
)

var name = "Zekeriya"
var chars = []string{"z", "e", "k"}

func Demo1() {
	cloneName := s.Clone(name)                // clonela
	nameChars := s.Split(s.ToLower(name), "") // empty e göre ayır array yap
	zekName := s.Trim(cloneName, "r")         // r a kadar olan kısmı kes al
	zekName2 := s.Join(chars, "")             // birleştir arayı araların empty koy
	isEqual := s.Compare(zekName, zekName2)
	isEqual2 := s.Compare(s.ToLower(zekName), zekName2) // eşit mi
	hasPref := s.HasPrefix(name, zekName)               // ile mi başlıyor
	hasSuffix := s.HasSuffix(name, "ya")                // ile mi bitiyor
	contains := s.Contains(name, "ek")
	countSubstr := s.Count(name, "e")
	replacedName := s.Replace(name, "", "*", 4)            // "" olan yerleri * yap 4 tane yap
	replacedName2 := s.Replace(replacedName, "*", "-", -1) // - 1 tümünü demek
	repeatName := s.Repeat(zekName2, 5)

	fmt.Println("Clone:", cloneName)
	fmt.Println("Split", nameChars)
	fmt.Println("Join:", zekName2)
	fmt.Println("Compare", isEqual)
	fmt.Println("Compare2", isEqual2)
	fmt.Println("HasPrefix", hasPref)
	fmt.Println("HasSuffix", hasSuffix)
	fmt.Println("Contains", contains)
	fmt.Println("Count", countSubstr)
	fmt.Println("Replace", replacedName)
	fmt.Println("Replace2", replacedName2)
	fmt.Println("Repeat", repeatName)

}
