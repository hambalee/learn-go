package main

import "fmt"

func main() {
	var courseName = []string{"Java", "Python"}
	fmt.Println(courseName) // [Java Python]
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println(courseName) // [Java Python C C# HTML CSS JavaScript]
	courseWeb := courseName[4:7]
	fmt.Println(courseWeb) // [HTML CSS JavaScript]
	courseWeb2 := courseName[:4]
	fmt.Println(courseWeb2) // [Java Python C C#]
}
