package bob

import "fmt"

const testVersion = 3

func any_words_all_caps( runes []rune ) (bool,bool) {

    any_all_caps  := false
    word_all_caps := true
    total_caps    := true

    for _, ru := range runes {
        up := int(ru)

        if !(int('A') <= up && up <= int('Z')) {
            word_all_caps = false
            total_caps    = false
        }

        if ru == ' ' {
            if word_all_caps {
                any_all_caps = true
            }
            word_all_caps = true
        }
    }

    return any_all_caps, total_caps
}

// func Hey simulates interaction with a stereotyped millenial.
func Hey( in string ) string {

    fmt.Println("in: ", in)

    runes := []rune(in)

    any_all_caps, total_caps  := any_words_all_caps( runes )
    ends_with_exc := runes[ len(runes)-1 ] == '!'

    fmt.Println("any_all_caps", any_all_caps, "total_caps", total_caps)

    ends_with_que := runes[ len(runes)-1 ] == '?'
    
    if (any_all_caps && ends_with_exc) || total_caps {
        return "Whoa, chill out!"
    }

    if ends_with_que {
        return "Sure."
    }

    return "Whatever."
}

