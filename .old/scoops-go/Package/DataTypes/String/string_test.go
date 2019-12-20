package _string

import (
    "testing"
)



func TestClone(t *testing.T) {
    a := FromString("hello")
    b := a.Clone().(*String)
    for i, r := range a.value {
        if r.Value[0] != b.value[i].Value[0] {
            t.Error("Function 'Clone' works incorrectly.")
        }
    }
}