package assembly

import (
    "errors"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
    "regexp"
    "strconv"
)



func SyntaxCheck(line string) error {
    /* 
     * b[01]+           -- BOOLEAN
     * x[\dA-F]+        -- HEXADECIMAL
     * \d+              -- DECIMAL
     * '[\x00-\xFF]'    -- ASCII CHAR
     * [A-Z_]+          -- OPERATOR
     */
    validInstruction := regexp.MustCompile(
        `^[A-Z_]+( b[01]+| x[\dA-F]+| \d+| '[\x00-\xFF]'| [A-Z_]+)\s*$`)
    if !validInstruction.MatchString(line) {
        return fmt.Errorf("Syntactically incorrect line detected:\n\t%s", line)
    }
    return nil
}


func FindOpcode(line string) string {
	validOpcode := regexp.MustCompile(`[A-Z_]+`)
	return validOpcode.FindString(line)
}


func FindOperand(line string) string {
    validOperandByte := regexp.MustCompile(
        `b[01]+|x[\dA-F]+|\d+|'[\x00-\xFF]'|[A-Z_]+`)
    return validOperandByte.FindString(line)
}


func OperandCheck(operand string) bool {
    prefix := operand[0]
    switch prefix {
    case 'b':
        return len(operand) < 10

    case 'x':
        _, e := strconv.ParseUint(operand[1:], 16, 8)
        return e == nil

    case '\'':
        // ASCII characters are 1 byte long by definition and their validity is
        // checked syntactically.
        return true

    default: // it's either integer or opcode at this point
        validOpcode := regexp.MustCompile(`[A-Z_]+`)
        if validOpcode.MatchString(operand) {
            return true
        }
        _, e := strconv.ParseUint(operand, 10, 8)
        return e == nil
    }
}


func SemanticsCheck(line string) error {
    opcodeString := FindOpcode(line)
    _, exists := OpcodeMap[opcodeString]
    if !exists {
        return errors.New("Opcode does not exist.")
    }

    // Check that operand is made of values that can be stored in a byte
    /* 
     * ASCII characters and opcodes are by definition a byte long.
     * Must only check:
     *     1. Booleans
     *     2. Decimals
     *     3. Hexadecimals
     */
    opcodeLen := len(opcodeString)
    operand := FindOperand(line[opcodeLen:])

    // Check integers
    if !OperandCheck(operand) {
        return errors.New("Operand contains byte value out of range.")
    }

    return nil
}


func GetIntegerAndBase(integer string) (string, int) {
    /* 
     * [A-Z_]+       -- OPERATOR
     * '[\x00-\x7F]' -- ASCII CHAR CONVERSION OPERAND
     * [0-9]+        -- DECIMAL OPERAND
     * x[0-9A-F]+    -- HEXADECIMAL OPERAND
     * b[01]{1,8}    -- BINARY OPERAND
     */
    prefix := integer[0]

    switch prefix {
    case 'b':
        return integer[1:], 2
    case 'x':
        return integer[1:], 16
    default:
        return integer, 10
    }
}


func AssembleLine(line string) (instruction *shared.Instruction, err error) {
    err = SyntaxCheck(line)
    if err != nil {
        return
    }
    err = SemanticsCheck(line)
    if err != nil {
        return
    }

    opcodeString := FindOpcode(line)
    opcodeLen := len(opcodeString)
    operandString := FindOperand(line[opcodeLen:])
    opcode := OpcodeMap[opcodeString]
    validOpcode := regexp.MustCompile(`^[A-Z_]+$`)
    var operand byte
    if operandString[0] == '\'' { // char
        operand = operandString[1]
    } else if validOpcode.MatchString(operandString) {
        operand = OpcodeMap[operandString]
    } else {
        // If this byte is not a char, it must be an integer with base
        // 2 / 10 / 16. Let's detect its base and parse it.
        integer, integerBase := GetIntegerAndBase(operandString)
        operand64, err := strconv.ParseUint(integer, integerBase, 8)
        if err != nil {
            return nil, err
        }
        operand = byte(operand64)
    }

    instruction = shared.NewInstruction(opcode, operand)
    return
}


/*
 * This function should work in a linear manner, checking line by line,
 * returning error whenever it finds a line that is either syntactically or
 * semantically incorrect.
 */
func Assemble(assemblyCode []string) (byteCode []*shared.Instruction,
                                      err error) {
    for _, line := range assemblyCode {
        instruction, err := AssembleLine(line)
        if err != nil {
            return nil, err
        }
        byteCode = append(byteCode, instruction)
    }
    return
}
