package raindrops

import (
    "fmt"
    "strings"
)

const testVersion = 3

func Convert(i int) string {

    s := make([]string, 3)
    s = s[:0]

    // need to factor here.

    if 0 == (i % 3) {
        s = append(s, "Pling")
    }

    if 0 == (i % 5) {
        s = append(s, "Plang")
    }

    if 0 == (i % 7) {
        s = append(s, "Plong")
    }


    if len(s) > 0 {
        return strings.Join( s, "" )
    } else {
        return fmt.Sprintf("%d", i);
    }

}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
