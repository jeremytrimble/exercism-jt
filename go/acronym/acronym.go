package acronym

import "strings"
import "bytes"

const testVersion = 3


func Abbreviate(in string) string {

    var prev rune
    first := true

    var out bytes.Buffer

    for _, cur := range in {

        beg := (prev == ' ') || (prev == '-')

        if first || beg {
            out.WriteString( strings.ToUpper( string(cur) ) )
        }

        prev = cur

        first = false
    }

    return out.String()
}
