package bytes

import (
	"encoding/binary"
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/DataTypes/Queue"
    "github.com/sharpvik/scoops/Package/DataTypes/Scoop"
    "github.com/sharpvik/scoops/Package/DataTypes/Slice"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
    "github.com/sharpvik/scoops/Package/DataTypes/String"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
    "math"
)



func (interpreter *Interpreter) Evaluate() {
    instruction := *(interpreter.scope.code[interpreter.scope.ip])

    //fmt.Println(instruction) // debug

    switch instruction.Opcode {

    case shared.END:
        interpreter.running = false

    case shared.MAKE_BYTE:
        interpreter.scope.data.Push( primitives.NewByte(instruction.Operand) )

    case shared.MAKE_BLN:
        b := instruction.Operand != 0
        interpreter.scope.data.Push( primitives.FromBoolean(b) )

    /*
     * It's important to remember that the sequence of bytes that are pushed
     * onto the data stack to become a flt or an int in the future must be
     * generated using the BigEndian byte order. Stack reverses them and here
     * interpreter uses the LittleEndian to instantiate primitives.
     */
    case shared.MAKE_FLT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        bits := binary.LittleEndian.Uint64(buffer)
        float := math.Float64frombits(bits)
        interpreter.scope.data.Push( primitives.NewFloat(float) )

    case shared.MAKE_INT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        n := int64( binary.LittleEndian.Uint64(buffer) )
        interpreter.scope.data.Push( primitives.NewInteger(n) )

    case shared.MAKE_NIL:
        interpreter.scope.data.Push( primitives.NewNil() )

    case shared.MAKE_ASCII_RUNE:
        interpreter.scope.data.Push(
            primitives.NewRune([]byte{instruction.Operand}),
        )

    case shared.MAKE_RUNE:
        var buffer []byte
        c := int(instruction.Operand)
        for i := 0; i < c; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        interpreter.scope.data.Push( primitives.NewRune(buffer) )

    case shared.MAKE_STRING:
        c := interpreter.scope.data.Pop().(*primitives.Integer).Value
        var i int64
        var buffer []*primitives.Rune
        for i = 0; i < c; i++ {
            r := interpreter.scope.data.Pop().(*primitives.Rune)
            buffer = append(buffer, r)
        }
        interpreter.scope.data.Push( _string.New(buffer) )

    case shared.STRING_CONCATENATE:
        b := interpreter.scope.data.Pop().(*_string.String)
        a := interpreter.scope.data.Pop().(*_string.String)
        c := _string.Concatenate(a, b)
        interpreter.scope.data.Push(c)

    case shared.MAKE_ERROR:
        //

    case shared.SCREECH:
        if instruction.Operand > 0 {
            err := interpreter.scope.data.Pop().(*primitives.Error)
            interpreter.err = err
        } else {
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                fmt.Sprintf(
                    "Instruction #%d: Retarded Screeching is not a good sign.",
                    interpreter.scope.ip,
                ),
            )
        }

    case shared.MAKE_SLICE:
        interpreter.scope.data.Push( slice.New() )

    case shared.SLICE_APPEND:
        obj := interpreter.scope.data.Pop()
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        slice.Append(obj)

    case shared.SLICE_GET_ITEM_BY_INDEX:
        i := interpreter.scope.data.Pop().(*primitives.Integer).Value
        index := uint64(i)
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        obj := slice.GetItemByIndex(index)
        interpreter.scope.data.Push(obj)

    case shared.SLICE_POP:
        i := interpreter.scope.data.Pop().(*primitives.Integer).Value
        index := uint64(i)
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        obj := slice.Pop(index)
        interpreter.scope.data.Push(obj)

    case shared.MAKE_STACK:
        interpreter.scope.data.Push( stack.New() )

    case shared.STACK_PUSH:
        obj := interpreter.scope.data.Pop()
        interpreter.scope.data.Peek().(*stack.Stack).Push(obj)

    case shared.STACK_POP:
        obj := interpreter.scope.data.Peek().(*stack.Stack).Pop()
        interpreter.scope.data.Push(obj)

    case shared.STACK_CLEAR:
        interpreter.scope.data.Peek().(*stack.Stack).Clear()

    case shared.STACK_EMPTY:
        bln := interpreter.scope.data.Peek().(*stack.Stack).Empty()
        interpreter.scope.data.Push( primitives.FromBoolean(bln) )

    case shared.STACK_PEEK:
        obj := interpreter.scope.data.Peek().(*stack.Stack).Peek()
        interpreter.scope.data.Push(obj)

    case shared.MAKE_QUEUE:
        interpreter.scope.data.Push( queue.New() )

    case shared.QUEUE_PUSH:
        obj := interpreter.scope.data.Pop()
        interpreter.scope.data.Peek().(*queue.Queue).Push(obj)

    case shared.QUEUE_POP:
        obj := interpreter.scope.data.Peek().(*queue.Queue).Pop()
        interpreter.scope.data.Push(obj)

    case shared.QUEUE_CLEAR:
        interpreter.scope.data.Peek().(*queue.Queue).Clear()

    case shared.QUEUE_EMPTY:
        bln := interpreter.scope.data.Peek().(*queue.Queue).Empty()
        interpreter.scope.data.Push( primitives.FromBoolean(bln) )

    case shared.QUEUE_PEEK:
        obj := interpreter.scope.data.Peek().(*queue.Queue).Peek()
        interpreter.scope.data.Push(obj)

    case shared.CLONE_OBJ:
        interpreter.scope.data.Push(
            interpreter.scope.data.Peek().Clone(),
        )

    case shared.PRINT_OBJ:
        interpreter.scope.data.Peek().Print(interpreter.scope.writer)

    case shared.GET_TYPE:
        interpreter.scope.data.Push(
            _string.FromString( interpreter.scope.data.Peek().Type() ),
        )

    case shared.GET_SIZE:
        collection := interpreter.scope.data.Peek().(shared.Collection)
        size := int64( collection.Size() )
        interpreter.scope.data.Push( primitives.NewInteger(size) )

    case shared.MAKE_SCOOP:
        name := interpreter.scope.data.Pop().(*_string.String).ToGoString()
        size := uint64(
            interpreter.scope.data.Pop().(*primitives.Integer).Value,
        )
        var code []*shared.Instruction
        final := interpreter.scope.ip + size + 1
        for interpreter.scope.ip++;
            interpreter.scope.ip < final;
            interpreter.scope.ip++ {
                scoopInstruction := interpreter.scope.code[interpreter.scope.ip]
                code = append(code, scoopInstruction)
        }
        interpreter.scope.data.Push( scoop.New(name, code) )
        return

    case shared.INLINE_SCOOP:
        interpreter.scope = NewInlineEnvironment(
            "inline",
            interpreter.scope.ip,
            interpreter.scope.code,
            interpreter.scope.global,
            interpreter.scope,
        )

    case shared.SCOOP_CALL:
        s := interpreter.scope.data.Pop().(*scoop.Scoop)
        interpreter.scope = NewEnvironment(
            s.Name, s.Code, interpreter.consts,
            interpreter.scope.global, interpreter.scope,
            interpreter.scope.writer,
        )
        /*
         * Here, we have to use 'return' in order to prevent instruction
         * pointer 'ip' from being incremented as it must be 0 at the start of
         * the new Scoop.
         */
        return

    case shared.SCOOP_ACCEPT_ARG:
        interpreter.scope.data.Push( interpreter.scope.prev.data.Pop() )

    case shared.SCOOP_RETURN:
        for !interpreter.scope.data.Empty() {
            interpreter.scope.prev.data.Push( interpreter.scope.data.Pop() )
        }
        ip := interpreter.scope.ip
        name := interpreter.scope.name
        interpreter.scope = interpreter.scope.prev
        if name == "inline" {
            interpreter.scope.ip = ip
        }

    case shared.STORE_CONST:
        interpreter.consts = append(
            interpreter.consts,
            interpreter.scope.data.Peek(),
        )

    case shared.LOAD_CONST:
        id := uint64( interpreter.scope.data.Pop().(*primitives.Integer).Value )
        interpreter.scope.data.Push( interpreter.consts[id] )

    case shared.STORE_VAR:
        storeMode := instruction.Operand
        if storeMode == 'N' {           // new variable created
            interpreter.scope.vars = append(
                interpreter.scope.vars,
                interpreter.scope.data.Peek(),
            )
        } else if storeMode == 'R' {    // variable reassigned
            id := uint64(
                interpreter.scope.data.Pop().(*primitives.Integer).Value,
            )
            interpreter.scope.vars[id] = interpreter.scope.data.Peek()
        } else {
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                fmt.Sprintf(
                    "Unknown store mode '%c' in call to STORE_VAR.",
                    storeMode,
                ),
            )
        }

    case shared.LOAD_VAR:
        id := uint64( interpreter.scope.data.Pop().(*primitives.Integer).Value )
        interpreter.scope.data.Push( interpreter.scope.vars[id] )

    case shared.UNARY_NOT:
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.NotBoolean(b) )

    case shared.BINARY_AND:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.AndBoolean(a, b) )

    case shared.BINARY_OR:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.OrBoolean(a, b) )

    case shared.BINARY_XOR:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.XorBoolean(a, b) )

    case shared.BINARY_ADD:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "byte":
            a := x.(*primitives.Byte)
            b := y.(*primitives.Byte)
            interpreter.scope.data.Push( primitives.AddByte(a, b) )
        case "int":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.AddInteger(a, b) )
        case "flt":
            a := x.(*primitives.Float)
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.AddFloat(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_SUB:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "byte":
            a := x.(*primitives.Byte)
            b := y.(*primitives.Byte)
            interpreter.scope.data.Push( primitives.SubByte(a, b) )
        case "int":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.SubInteger(a, b) )
        case "flt":
            a := x.(*primitives.Float)
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.SubFloat(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_MUL:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "byte":
            a := x.(*primitives.Byte)
            b := y.(*primitives.Byte)
            interpreter.scope.data.Push( primitives.MulByte(a, b) )
        case "int":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.MulInteger(a, b) )
        case "flt":
            a := x.(*primitives.Float)
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.MulFloat(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_DIV:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "byte":
            a := x.(*primitives.Byte)
            b := y.(*primitives.Byte)
            interpreter.scope.data.Push( primitives.DivByte(a, b) )
        case "int":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.DivInteger(a, b) )
        case "flt":
            a := x.(*primitives.Float)
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.DivFloat(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_MOD:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "byte":
            a := x.(*primitives.Byte)
            b := y.(*primitives.Byte)
            interpreter.scope.data.Push( primitives.ModByte(a, b) )
        case "int":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.ModInteger(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_POW:
        y := interpreter.scope.data.Pop()
        x := interpreter.scope.data.Pop()
        _type := x.Type()
        switch _type {
        case "int":
            a := x.(*primitives.Integer).ToFloat()
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.PowFloat(a, b) )
        case "flt":
            a := x.(*primitives.Float)
            b := y.(*primitives.Float)
            interpreter.scope.data.Push( primitives.PowFloat(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.PRINT_NEWLINE:
        interpreter.scope.writer.WriteString("\n")
        interpreter.scope.writer.Flush()

    case shared.POP:
        interpreter.scope.data.Pop()

    case shared.ABSOLUTE_JUMP:
        dst := uint64(
            interpreter.scope.data.Pop().(*primitives.Integer).Value,
        )
        interpreter.scope.ip = dst
        return

    case shared.RELATIVE_JUMP:
        diff := uint64(
            interpreter.scope.data.Pop().(*primitives.Integer).Value,
        )
        jumpMode := instruction.Operand

        if jumpMode == '-' {
            interpreter.scope.ip -= diff
        } else if jumpMode == '+' {
            interpreter.scope.ip += diff
        } else {
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                fmt.Sprintf(
                    "Unknown jump mode '%c' in call to RELATIVE_JUMP.",
                    jumpMode,
                ),
            )
        }
        return

    case shared.ABSOLUTE_JUMP_IF_FALSE:
        dst := uint64(
            interpreter.scope.data.Pop().(*primitives.Integer).Value,
        )
        condition := interpreter.scope.data.Pop().(*primitives.Boolean).Value

        if !condition {
            interpreter.scope.ip = dst
            return
        }

    case shared.RELATIVE_JUMP_IF_FALSE:
        diff := uint64(
            interpreter.scope.data.Pop().(*primitives.Integer).Value,
        )
        condition := interpreter.scope.data.Pop().(*primitives.Boolean).Value

        if !condition {
            interpreter.scope.ip += diff
            return
        }

    default:
        interpreter.err = primitives.NewError(
            shared.OpcodeError,
            fmt.Sprintf(
                "Instruction #%d: Unknown opcode %d.",
                interpreter.scope.ip,
                instruction.Opcode,
            ),
        )

    }

    /* Instruction pointer is incremented by default. Use 'return' inside the
       case body to force it not to. */
    interpreter.scope.ip++

    //interpreter.scope.data.Print(interpreter.scope.writer) // debug
}


func (interpreter *Interpreter) Execute() {
    for interpreter.running && interpreter.err == nil {
        interpreter.Evaluate()
    }
    if interpreter.err != nil {
        interpreter.err.Print(interpreter.scope.writer)
        util.Log("Interpreter exited with non-zero return value.")
    }
}
