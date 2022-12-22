package main

import (
    "fmt"
    "math/rand"
    "time"
)

    var guess int
    var trialcount int = 1

    func main(){
        source := rand.NewSource(time.Now().UnixNano())
        randomizer := rand.New(source)
        secretNumber := randomizer.Intn(10)

        fmt.Println("Adivinhe o número!")

        for{
            fmt.Println("Qual é o número?")
            fmt.Scan(&guess)
            if trialcount > 2 && guess != secretNumber{
                fmt.Println("Acabaram as tentativas!")
                fmt.Println("O número era: ", secretNumber)
                break
            }else{
                if guess > secretNumber{
                    fmt.Println("Maior!")
                }else if guess < secretNumber{
                    fmt.Println("Menor!")
                }else {
                    fmt.Println("Acertou!")
                    break
                }
            }
            trialcount ++
        }
    }
