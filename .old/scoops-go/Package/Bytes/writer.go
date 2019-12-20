package bytes

import (
    "bufio"
    "os"
    "github.com/sharpvik/scoops/Package/Shared"
)

func Write(byteCode []*shared.Instruction, filename string) error {
    // Convert []*shared.Instruction to []byte
    var digest []byte
    for _, i := range byteCode {
        digest = append(digest, i.Opcode, i.Operand)
    }
    
    // Create file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    
    // Write to file
    wrtr := bufio.NewWriter(file)
    _, err = wrtr.Write(digest)
    if err != nil {
        return err
    }
    wrtr.Flush()
    file.Close()
    
    return nil
}
