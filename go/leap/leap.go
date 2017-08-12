package leap

const testVersion = 3

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(y int) bool {
    if y % 4 == 0 {
        if y % 100 == 0 {
            if y % 400 == 0 {
                return true;
            } else {
                return false;
            }
        }
        return true;
    }
    return false;
}
