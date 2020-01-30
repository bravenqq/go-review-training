package main

func main() {

}

func Deduplicate(s []string) []string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	i := 1
	for j := 1; j < len(s); j++ {
		if s[j] == s[j-1] {
			continue
		}
		s[i] = s[j]
		i++
	}
	s = s[0:i]
	return s
}
