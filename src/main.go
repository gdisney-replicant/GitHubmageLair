package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "math/rand"
    "time"
)

var (
    answer1 string
    answer2 string
    answer3 string
    answer4 string
    answer5 string
    answer6 string
    answer7 string
    answer8 string
    answer9 string
    answer10 string
    answer11 string
    answer12 string
    answer13 string
    answer14 string
    answer15 string
    flag string
)

type Riddle struct {
    Question string
    Answer   *string
}

func main() {
    fmt.Println("Welcome to the Mage's Riddle Challenge!")
    fmt.Println("Answer three riddles correctly to uncover the secret.")
    fmt.Println()

    riddles := []Riddle{
        {"What is the most secure way to store your dragon's treasure in the cloud?", &answer1},
        {"What incantation deflects the hordes from storming your cloud castle?", &answer2},
        {"I'm a mythical creature that guards the entrance to your cloud fortress. What am I?", &answer3},
        {"What charm wards off malevolent wizards trying to hex your cloud data?", &answer4},
        {"In the realm of cloud, what potion do you drink to vanish from the enemy’s sight?", &answer5},
        {"What is the name of the magical artifact that can summon a storm of packets to overwhelm your enemies?", &answer6},
        {"What is the name of the spell that can reveal the true form of a disguised hacker?", &answer7},
        {"What is the name of the magical creature that can detect and alert you of intruders in your cloud kingdom?", &answer8},
        {"What ancient artifact grants you the power to audit all who enter your cloud lair?", &answer9},
        {"Which mystical tome holds the secrets to your cloud kingdom's logs?", &answer10},
        {"What spell alerts you to intruders breaching your cloud defenses?", &answer11},
        {"What rune do you invoke to ensure your dragon’s hoard is only accessed by the rightful owner?", &answer12},
        {"What is the sacred barrier that keeps out all but the most trusted warriors of your cloud realm?", &answer13},
        {"What relic allows you to scale your cloud defenses in an instant?", &answer14},
        {"What scroll binds attackers in a net of rules and denies them passage?", &answer15},
    }

    // Shuffle the riddles
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(riddles), func(i, j int) {
        riddles[i], riddles[j] = riddles[j], riddles[i]
    })
    if askRiddles(riddles, 3) {
        fmt.Println("Congratulations! You've solved three riddles! ", flag)
    } else {
        fmt.Println("Try the riddles again to uncover the secret.")
    }
}
func askRiddles(riddles []Riddle, numCorrect int) bool {
    reader := bufio.NewReader(os.Stdin)
    correctAnswers := 0
    answeredRiddles := make(map[int]bool)

    for correctAnswers < numCorrect {
        for i, riddle := range riddles {
            if answeredRiddles[i] {
                continue
            }

            fmt.Println(riddle.Question)
            fmt.Print("Your answer: ")
            answer, _ := reader.ReadString('\n')
            answer = strings.ToLower(strings.TrimSpace(answer))

            correctAnswer := strings.ToLower(strings.ReplaceAll(*riddle.Answer, "_", " "))
            words := strings.Fields(correctAnswer)

            if answer == correctAnswer {
                correctAnswers++
                fmt.Println("Correct!")
                fmt.Println()
                answeredRiddles[i] = true
            } else {
                for _, word := range words {
                    if word == answer {
                        correctAnswers++
                        fmt.Println("Good enough!")
                        fmt.Println()
                        answeredRiddles[i] = true
                        break
                    }
                }
                if !answeredRiddles[i] {
                    fmt.Println("Incorrect. Let's move to the next riddle.")
                    fmt.Println()
                }
            }
            if correctAnswers >= numCorrect {
                break
            }
        }
    }
    return correctAnswers >= numCorrect
}
