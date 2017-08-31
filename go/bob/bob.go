package bob

//import "fmt"
import "strings"
import "unicode"

const testVersion = 3

func count_yelling( words []string ) (int, int) {

    total_words := 0
    total_yells := 0

    for _, word := range words {

        is_word     := false
        is_all_caps := true

        for _, ru := range word {

            if unicode.IsLetter(ru) {
                is_word = true
                if !unicode.IsUpper(ru) {
                    is_all_caps = false
                }
            }

        }

        if is_word {
            total_words++
        }

        if is_word && is_all_caps {
            total_yells++
        }

    }

    return total_yells, total_words
}

// func Hey simulates interaction with a stereotyped millenial.
func Hey( in string ) string {

    //fmt.Println("in: ", in)

    words := strings.Fields(in)

    if len(words) == 0 {
        return "Fine. Be that way!"
    }

    total_yells, total_words := count_yelling( words )

    ends_with_que := strings.HasSuffix(words[len(words)-1], "?")

    mostly_yelling := total_yells > 0 && float32(total_yells) >= (float32(total_words)/float32(2.0))

    //fmt.Println("total_yells", total_yells, "total_words", total_words)

    if mostly_yelling {
        return "Whoa, chill out!"
    }

    if ends_with_que {
        return "Sure."
    }

    return "Whatever."
}

