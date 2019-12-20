package bytes

import (
    "os"
    "testing"
)



func TestRead(t *testing.T) {
    sampleFilename := "test.scpb"
    file, err := os.Create(sampleFilename)
    if err != nil {
        t.Error("Cannot create sample file.")
    }

    //

    file.Close()
    os.Remove(sampleFilename)
}

