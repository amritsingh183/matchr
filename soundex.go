package matchr

import "strings"

func Soundex(s1 string) (string){
    if len(s1) == 0 {
        return ""
    }

    // we should work with all uppercase
    s1 = strings.ToUpper(s1)

    // the encoded value
    enc := s1[0:1]

    i := 0
    c := ""
    prev := ""
    hw := false

    for _, v := range(s1) {
        switch v {
        case 'B', 'F', 'P', 'V':
            c = "1"
        case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
            c = "2"
        case 'D', 'T':
            c = "3"
        case 'L':
            c = "4"
        case 'M', 'N':
            c = "5"
        case 'R':
            c = "6";
        case 'H', 'W':
            hw = true
        default:
            c = ""
        }

        // don't encode the first position, but we need its code value
        // to prevent repeats
        if c != "" && c != prev && i > 0 {
            // if the next encoded digit is different, we can add it right away
            // if it is the same, though, it must not have been preceded
            // by an 'H' or a 'W'
            if enc[len(enc)-1:len(enc)] != c || !hw {
                enc = enc + c
            }

            // we're done when we reach four encoded characters
            if len(enc) == 4 {
                break
            }
        }

        prev = c
        hw = false
        i++
    }

    // if we've fallen short of 4 "real" encoded characters,
    // it gets padded with zeros
    for i := len(enc); i < 4; i++ {
        enc = enc + "0"
    }

    return enc
}