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

func askQuestion (i int, question questionGetter.Round) (pass bool, take bool) {
    take = false

    answer, correctAnswer := printQuestion(
        i,
        question.Question,
        question.Answer,
        question.A,
        question.B,
        question.C,
        question.D,
    )

    if answer == "t" {
        take = true
    }

    result := marker(answer, correctAnswer)
    pass = result
    fmt.Println(result)
    return
}

func returnPrize (i int, take bool) {
    amount := ""
    if take {
        switch i {
            case 0: amount = "£0"
            case 1: amount = "£100"
            case 2: amount = "£200"
            case 3: amount = "£300"
            case 4: amount = "£500"
            case 5: amount = "£1,000"
            case 6: amount = "£2,000"
            case 7: amount = "£4,000"
            case 8: amount = "£8,000"
            case 9: amount = "£16,000"
            case 10: amount = "£32,000"
            case 11: amount = "£64,000"
            case 12: amount = "£125,000"
            case 13: amount = "£250,000"
            case 14: amount = "£500,000"
            case 15: amount = "£1,000,000"
        }
    } else {
        switch {
            case i < 5: amount = "£0"
            case i >= 5 && i < 10: amount = "£1,000"
            case i >= 10 && i < 14: amount = "£32,000"
            case i == 14: amount = "£1,000,000"
        }
    }

    fmt.Println("You go away with " + amount)
}

func RunGame() {
    questions := questionGetter.BuildQuestionArray()
    for i, r := range questions {
        fmt.Println(i)
        pass, take := askQuestion(i, r)
        if take == true {
            returnPrize(i, true)
            return
        }
        if pass == false {
            returnPrize(i, false)
            return
        }
        if pass == true && i ==14 {
            returnPrize(i, false)
            return
        }
    }
}
