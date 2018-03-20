package questionGetter

import (
    "time"
    "math/rand"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Round struct {
    Question string
    Answer string
    A string
    B string
    C string
    D string
}

type Prize struct {
    amount string
}

func (r Round) toString() string {
    return toJson(r)
}

func toJson(r interface{}) string {
    bytes, err := json.Marshal(r)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func getQuestions(difficulty string) []Round {
    raw, err := ioutil.ReadFile("./" + difficulty + ".json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Round
    json.Unmarshal(raw, &c)
    return c
}

func in_array(val Round, array [15]Round) (exists bool) {
    exists = false
    for _, v := range array {
        if val == v {
            exists = true
            return
        }
    }
    return
}

func BuildQuestionArray() [15]Round {
    var questionArray [15]Round
    easyQuestions := getQuestions("easyQuestions")
    mediumQuestions := getQuestions("mediumQuestions")
    hardQuestions := getQuestions("hardQuestions")
    for i := 0;  i<15; i++ {
        if i<5 {
            rand.Seed(time.Now().UnixNano())
            j := rand.Intn(15)
            for in_array(easyQuestions[j], questionArray) {
                j = rand.Intn(15)
            }
            questionArray[i] = easyQuestions[j]
            }
        if i>=5 && i<10 {
            rand.Seed(time.Now().UnixNano())
            j := rand.Intn(15)
            for in_array(mediumQuestions[j], questionArray) {
                j = rand.Intn(15)
            }
            questionArray[i] = mediumQuestions[j]
            }
        if i>=10 && i<15 {
            rand.Seed(time.Now().UnixNano())
            j := rand.Intn(15)
            for in_array(hardQuestions[j], questionArray) {
                j = rand.Intn(15)
            }
            questionArray[i] = hardQuestions[j]
        }
    }
    fmt.Println(questionArray)
    return questionArray
}
