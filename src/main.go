package main

import (
    "fmt"
)
// func main() {
//     fmt.Print("Enter text: ")
//     var input string
//     // fmt.Scanln(&input)
//     fmt.Print(input)
// }
func Marker (selection string, answer string) (result bool) {
    if (selection == answer) {
        result = true
    } else {
        result = false
    }
    return
}

func PrintQuestion (question string, rightAnswer string, options ...string) (answer string, correctAnswer string) {
    fmt.Println("<<" + question + ">>")
    fmt.Println("<" + options[0] + ">" + " " + "<" + options[1] + ">")
    fmt.Println("<" + options[2] + ">" + " " + "<" + options[3] + ">")
    var input string
    fmt.Scanln(&input)
    answer = input
    correctAnswer = rightAnswer
    return
}

func main() {
    answer, correctAnswer := PrintQuestion("What is the letter a?", "a", "a", "b", "c", "d")
    result := Marker(answer, correctAnswer)
    fmt.Println(result)
}
