import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var numbers int
    flag.IntVar(&numbers, "n", 0, "Number of digits in the PW")

    var lowercase int
    flag.IntVar(&lowercase, "l", 0, "Number of lowercase chars in the PW")

    var uppercase int
    flag.IntVar(&uppercase, "u", 0, "Number of uppercase chars in the PW")

    var specialChars int
    flag.IntVar(&specialChars, "s", 0, "Number of special chars in the PW")

    var totalLength int
    flag.IntVar(&totalLength, "t", 0, "The total password length. If passed, it will ignore -n, -l, -u and -s, and generate completely random passwords with the specified length")

    var amount int
    flag.IntVar(&amount, "a", 1, "")

    var outputFile string
    flag.StringVar(&outputFile, "o", "", "")

    flag.Parse()

    passwords := make([]string, amount)

    for i := 0; i < amount; i++ {
        if totalLength > 0 {
            password := make([]byte, totalLength)
            for j := 0; j < totalLength; j++ {
                password[j] = byte(rand.Intn(94) + 33)
            }
            passwords[i] = string(password)
        } else {
            password := make([]byte, numbers+lowercase+uppercase+specialChars)
            for j := 0; j < numbers; j++ {
                password[j] = byte(rand.Intn(10) + 48)
            }
            for j := 0; j < lowercase; j++ {
                password[numbers+j] = byte(rand.Intn(26) + 97)
            }
            for j := 0; j < uppercase; j++ {
                password[numbers+lowercase+j] = byte(rand.Intn(26) + 65)
            }
            for j := 0; j < specialChars; j++ {
                password[numbers+lowercase+uppercase+j] = byte(rand.Intn(15) + 33)
            }
            rand.Shuffle(len(password), func(i, j int) {
                password[i], password[j] = password[j], password[i]
            })
            passwords[i] = string(password)
        }
    }

    if outputFile != "" {
        file, err := os.Create(outputFile)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        for _, password := range passwords {
            fmt.Fprintln(file, password)
        }
    } else {
        for _, password := range passwords {
            fmt.Println(password)
        }
    }
}
