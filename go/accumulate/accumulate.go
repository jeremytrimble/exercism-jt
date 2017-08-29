package accumulate

const testVersion = 1

func Accumulate( in []string, fun func(string) string) (out []string)  {

    out = make( []string, len(in) )

    for i := 0; i < len(in); i++ {
        out[i] = fun(in[i])
    }

    return

}
