package shared

import "bufio"

/*
 * It is important to note that the 'Object' interface type is satisfied by all
 * of the pointer types in the DataTypes folder.
 *
 * In other words, *primitives.Rune can be retrieved from an Object type using
 * type assertion as *primitives.Rune perfectly satisfies the Object interface:
 *     var a Object = interpreter.scope.data.Pop()
 *     var b *primitives.Rune = a.(*primitives.Rune)
 *
 * However, primitives.Rune cannot. It doesn't satisfy the Object interface.
 * Here's why... We declare the Print and Type functions as follows:
 *     func (r *Rune) Print() { ... }
 * so type *Rune satisfies the Object interface but type Rune doesn't.
 * This goes for every type in the DataTypes folder. Similar logic applies to
 * the Collection interface.
 */
type Object interface {
    Clone() Object
    Print(w *bufio.Writer)
    Type() string
}


/* Meanwhile, Collection type has all the methods of the Object type plus its
   own Size method. */
type Collection interface {
    Object
    Size() uint64
}


type Instruction struct {
    Opcode, Operand byte
}

func NewInstruction(opcode, operand byte) *Instruction {
    return &Instruction{opcode, operand}
}
