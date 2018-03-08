package main

import (
    "fmt"
)

func Marker (selection string, answer string) (result bool) {
    if (selection == answer) {
        result = true
    } else {
        result = false
    }
    return
}

func Title () {
    fmt.Println("                     oOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOo")
    fmt.Println("             _ _ _  _ _   _    _ _ _   _   _  _  ___  __   ___   _    ___  ___    _  ")
    fmt.Println("            | | | || U | / \\  | | | | / \\ | \\| ||_ _|/ _| |_ _| / \\  | o )| __|  / \\ ")
    fmt.Println("            | V V ||   |( o ) | V V || o || \\\\ | | | \\_ \\  | | ( o ) | o \\| _|  | o |")
    fmt.Println("             \\_n_/ |_n_| \\_/   \\_n_/ |_n_||_|\\_| |_| |__/  |_|  \\_/  |___/|___| |_n_|")
    fmt.Println("           oOO                                                                    OOo")
    fmt.Println("         oOO                                                                        OOo")
    fmt.Println("       oOO                                                                            OOo")
    fmt.Println("      oOO                                                                              OOo")
    fmt.Println("     oOO                                                                                OOo")
    fmt.Println("    oOO                                                                                  OOo")
    fmt.Println("   oOO                                                                                    OOo")
    fmt.Println("  oOO                                                                                      OOo")
    fmt.Println(" oOO                                                                                        OOo")
    fmt.Println(" oO ███    ███  ██  ██       ██       ██   ██████   ███    ██   █████   ██  ██████   ███████ Oo")
    fmt.Println(" oO ████  ████  ██  ██       ██       ██  ██    ██  ████   ██  ██   ██  ██  ██  ██   ██      Oo")
    fmt.Println(" oO ██ ████ ██  ██  ██       ██       ██  ██    ██  ██ ██  ██  ███████  ██  ██████   █████   Oo")
    fmt.Println(" oO ██  ██  ██  ██  ██       ██       ██  ██    ██  ██  ██ ██  ██   ██  ██  ██  ██   ██      Oo")
    fmt.Println(" oO ██      ██  ██  ███████  ███████  ██   ██████   ██   ████  ██   ██  ██  ██   ██  ███████ Oo")
    fmt.Println(" oOO                                                                                        OOo")
    fmt.Println("  oOO                                                                                      OOo")
    fmt.Println("   oOO                                                                                    OOo")
    fmt.Println("    oOO                                                                                  OOo")
    fmt.Println("     oOO                                                                                OOo")
    fmt.Println("      oOO                                                                              OOo")
    fmt.Println("       oOO                                                                            OOo")
    fmt.Println("         oOO                                                                        OOo")
    fmt.Println("           oOO                                                                    OOo")
    fmt.Println("             _ _ _  _ _   _    _ _ _   _   _  _  ___  __   ___   _    ___  ___    _  ")
    fmt.Println("            | | | || U | / \\  | | | | / \\ | \\| ||_ _|/ _| |_ _| / \\  | o )| __|  / \\ ")
    fmt.Println("            | V V ||   |( o ) | V V || o || \\\\ | | | \\_ \\  | | ( o ) | o \\| _|  | o |")
    fmt.Println("             \\_n_/ |_n_| \\_/   \\_n_/ |_n_||_|\\_| |_| |__/  |_|  \\_/  |___/|___| |_n_|")
    fmt.Println("                        oOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOo")
}

func PrintQuestion (question string, rightAnswer string, options ...string) (answer string, correctAnswer string) {
    fmt.Println("<<  " + question + "  >>")
    fmt.Println("< (a) " + options[0] + " >" + "       " + "< (b) " + options[1] + " >")
    fmt.Println("< (c) " + options[2] + " >" + "       " + "< (d) " + options[3] + " >")
    var input string
    fmt.Scanln(&input)
    answer = input
    correctAnswer = rightAnswer
    return
}

func main() {
    Title()
    answer, correctAnswer := PrintQuestion("What is the letter a?", "a", "a", "b", "c", "d")
    result := Marker(answer, correctAnswer)
    fmt.Println(result)
}
