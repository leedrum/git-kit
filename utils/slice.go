package utils

import "fmt"

func RemoveSliceElm(iptSlice []string, txt string) []string {
	var newSlice []string
	for _, v := range iptSlice {
		if v == txt {
			continue
		}
		newSlice = append(newSlice, v)
	}

	return newSlice
}

func IsSliceContain(iptSlice []string, txt string) bool {
	for _, v := range iptSlice {
		if v == txt {
			return true
		}
	}

	return false
}

func PrintSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
