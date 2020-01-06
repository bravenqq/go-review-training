package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: ":Mar", 4: "Apr", 5: "May", 6: "June", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov.", 12: "Dec."}
	summer := make([]string, 3, 7)
	summer[0] = months[6]
	summer[1] = months[7]
	summer[2] = months[8]
	fmt.Printf("summer len:%d cap:%d\n", len(summer), cap(summer))
	fmt.Println(summer)
	summer = summer[:5]
	summer[3] = months[9]
	summer[4] = months[10]
	fmt.Printf("summer len:%d cap:%d\n", len(summer), cap(summer))
	fmt.Println(summer)
}
