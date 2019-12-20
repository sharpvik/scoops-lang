package assembly

import (
    "os"
    "testing"
)



func Join(arr []string) string {
    var output string
    for _, line := range arr {
        output += line
    }
    return output
}



func TestRead(t *testing.T) {
    sampleFilename := "test.scpb"
    file, err := os.Create(sampleFilename)
    if err != nil {
        t.Error("System Error: Cannot create sample file.")
    }

    cases := 
`MAKE_BYTE 'c'
PRINT_OBJ 0

PRINT_NEWLINE 0
THE_END 0
# the end
`
    
    file.WriteString(cases)
    file.Close()
    
    _, err = Read(sampleFilename)

    if err != nil {
        t.Error("Cannot read sample file.")
    }

    os.Remove(sampleFilename)
}