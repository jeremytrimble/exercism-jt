package hamming

const testVersion = 6

type HammingDistanceError struct {
    err_str string
};

func (hde HammingDistanceError) Error() string {
    return hde.err_str
}

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return -1, HammingDistanceError{ "String lengths are not equal" }
    }

    diffs := 0
    for i := range(a) {
        if a[i] != b[i] {
            diffs++
        }
    }

    return diffs, nil
}

