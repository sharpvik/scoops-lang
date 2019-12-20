package assembly

import (
    "github.com/sharpvik/scoops/Package/Shared"
    "testing"
    //"fmt"
)


func TestGetIntegerAndBase(t *testing.T) {
    cases := []string{
        "b101",
        "xFF",
        "42",
    }
    answers := []int{
        2, 16, 10,
    }
    for i, c := range cases {
        _, base := GetIntegerAndBase(c)
        if base != answers[i] {
            t.Errorf("Invalid base %d for %s.", base, c)
        }
    }
}


func TestAssembleLine(t *testing.T) {
    cases := []string{
        "MAKE_BYTE b101010",
        "MAKE_BYTE x2A",
        "MAKE_BYTE xFF",
        "MAKE_BYTE 42",
        "MAKE_BYTE MAKE_BYTE",
        "END 0",
    }
    answers := []*shared.Instruction{
        shared.NewInstruction(1, 42),
        shared.NewInstruction(1, 42),
        shared.NewInstruction(1, 255),
        shared.NewInstruction(1, 42),
        shared.NewInstruction(1, 1),
        shared.NewInstruction(0, 0),
    }
    for i, c := range cases {
        instruction, err := AssembleLine(c)
        if err != nil {
            t.Error(err)
        }
        if *instruction != *answers[i] {
            t.Errorf("Cannot properly assemble line %d.", i + 1)
        }
    }
    cases = []string{
        "PUSH_BYTE 1",
        "MAKE_BYTE 11 42",
        "GO_BROKE",
    }
    for i, c := range cases {
        _, err := AssembleLine(c)
        if err == nil {
            t.Errorf("Cannot properly detect errors (line %d).", i + 1)
        }
    }
}