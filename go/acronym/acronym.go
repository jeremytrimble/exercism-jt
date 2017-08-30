package acronym

import "strings"

const testVersion = 3


func Abbreviate(in string) string {
    fn := func( x rune ) bool {
        return x == ' ' || x == '-';
    }

    return strings.Join(
        strings.FieldsFunc( in, fn),
        "");
}
