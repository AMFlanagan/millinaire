package gameEngine

import (
    "../questionGetter"
    "../graphics"
    "fmt"
)

func marker (selection string, answer string) (result bool) {
    if (selection == answer) {
        result = true
    } else {
        result = false
    }
    return
}

func printQuestion (i int, question string, rightAnswer string, options ...string) (answer string, correctAnswer string) {
    graphics.Printer(question, options[0], options[1], options[2], options[3], i)
    var input string
    fmt.Scanln(&input)
    answer = input
    correctAnswer = rightAnswer
    return
}

func askQuestion (i int, question questionGetter.Round) (pass bool) {
    answer, correctAnswer := printQuestion(
        i,
        question.Question,
        question.Answer,
        question.A,
        question.B,
        question.C,
        question.D,
    )
    result := marker(answer, correctAnswer)
    pass = result
    fmt.Println(result)
    return
}

func RunGame() {
    questions := questionGetter.BuildQuestionArray()
    for i, r := range questions {
        if askQuestion(i, r) == false {
            fmt.Println("Ooooooooh noooooo you lost!")
            return
        }
    }
}
