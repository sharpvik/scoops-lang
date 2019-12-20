package bytes

import (
    "bufio"
//  "fmt"
    "io"
    "os"
    "github.com/sharpvik/scoops/Package/Shared"
)



func ReadAllBytes(rdr *bufio.Reader) ([]byte, error) {
    var buf []byte
    for {
        b, err := rdr.ReadByte()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        buf = append(buf, b)
    }
    return buf, nil
}


func Read(filename string) ([]*shared.Instruction, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    rdr := bufio.NewReader(file)

    buf, err := ReadAllBytes(rdr)
    if err != nil {
        return nil, err
    }

    var code []*shared.Instruction

//  fmt.Println(buf)
    for i := 0; i < len(buf); i += 2 {
        code = append( code, shared.NewInstruction(buf[i], buf[i + 1]) )
    }

    return code, nil
}
