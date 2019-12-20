/* Box Drawing
│ └ ├ ─ ┬
*/

package util

import (
    "fmt"
)



func Error(err error) {
    fmt.Println("│ Error:", err)
}


func Warning(warning string) {
    fmt.Println("│ Warning:", warning)
}


func Log(message string) {
    fmt.Println("│ Log:", message)
}


func Help() {
    fmt.Println(`
│ Hello, I am Embedded Helper! Welcome!

│ Scoops requires at most two command line arguments to know what to do.
│ First one is a flag and it is optional in some cases.
│ Second one is a filename that must have a proper extention.

│ Let's look at some flags:
├───┬ [-a] -- (Assemble) -- Outputs *.scpa file with Scoops Assembly Code.
│   └──── For this flag, only *.scp files constitute valid input.
│
├───┬ [-c] -- (Compile)  -- Outputs *.scpb file with Scoops Bytecode.
│   └──── For this flag, *.scp and *.scpa files are accepted as valid input.
│
└───┬ [-e] -- (Execute)  -- Optional flag. Executes your code.
    └──── For this flag, any Scoops file is accepted as valid input.

│ Valid Scoops file formats:
├──── [.scp]  -- (Source)   -- Basic Scoops Source Code file.
├──── [.scpa] -- (Assembly) -- Scoops Assembly file.
└──── [.scpb] -- (Bytecode) -- Scoops Bytecode file.

│ Command line keywords.
│ To use command line keyword, type 'scoops <keyword>'.
└──── help -- Embedded helper. 
    `)
}
