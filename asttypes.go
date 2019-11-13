package main

import (
    "fmt"
    _ "github.com/timtadh/lexmachine"
)

//Define operation enum
type Op int
const (
    Op_MUL Op = 0
    Op_DIV Op = 1
    Op_PLUS Op = 2
    Op_MINUS Op = 3
    Op_EQUALS Op = 4
    Op_NEQ Op = 5
    Op_GT Op = 6
    Op_LT Op = 7
    Op_GTE Op = 8
    Op_LTE Op = 9
    Op_AND Op = 10
    Op_OR Op = 11
    Op_NEG Op = 12
)

type StringPrimitive struct {}
type VoidType struct {}
type Integer struct {
    expType    interface{}
    Number int
}

func NewInteger(number int) *Integer {
    return &Integer{
        Number: number,
    }
}

func NewIntegerPrimitive() *Integer {
    return &Integer{}
}


func (i *Integer) visit() string {
    return fmt.Sprintf("(int %d)", i.Number)
}

func (i *Integer) analyze(c *Context)  {
}

type InfixExpression struct {
    expType    interface{}
    OpType Op
    Left Node
    Right  Node
}

func NewInfixExpression(opType Op, left Node , right Node) *InfixExpression {
    return &InfixExpression{
        OpType: opType,
        Left: left,
        Right: right,
    }
}

func (ie *InfixExpression) visit() string {
    return fmt.Sprintf("(%s %v %v)", resolveOp(ie.OpType), string(ie.Left.Exp.visit()), ie.Right.Exp.visit())

}

func  (ie *InfixExpression) analyze(c *Context)  {
}

func resolveOp(opType Op) string {
    typeStr := ""
    switch opType {
        case 0:
            typeStr = "Op_MUL"
        case 1:
            typeStr = "Op_DIV"
        case 2:
            typeStr = "Op_PLUS"
        case 3:
            typeStr = "Op_MINUS"
        case 4:
            typeStr = "Op_EQUALS"
        case 5:
            typeStr = "Op_NEQ"
        case 6:
            typeStr = "Op_GT"
        case 7:
            typeStr = "Op_LT"
        case 8:
            typeStr = "Op_GTE"
        case 9:
            typeStr = "Op_LTE"
        case 10:
            typeStr = "Op_AND"
        case 11:
            typeStr = "Op_OR"
        case 12:
            typeStr = "Op_NEG"
    }
    return typeStr
}

type Negation struct {
    expType    interface{}
    Exp *Node
}

func NewNegation(exp *Node) *Negation {
    return &Negation{
        Exp: exp,
    }
}

func (ne *Negation) visit() string {
    var str string
    if(ne.Exp == nil) {
        fmt.Println("Exp is null")
        str = string(ne.Exp.Token.Lexeme)
    } else {
        str = fmt.Sprintf("(NEG %v)", ne.Exp.visit())
    }
    return str
}

func (ne *Negation) analyze(c *Context)  {
}


type SeqExpression struct {
    expType    interface{}
    Exps []Node
}

func NewSeqExpression(expressions []Node) *SeqExpression {
    // var copiedExpressionContents []Node
    // for _, e := range expressions {
        // copiedExpressionContents = append(copiedExpressionContents, *e)
    // }
    return &SeqExpression{
        Exps: expressions,
    }
}

func (se *SeqExpression) visit() string {
    str := "(seqexp "
    for _, n := range se.Exps {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\n%v ", n.Exp.visit())
        }
    }

    str += "\n)"
    return str
}

func (se *SeqExpression) analyze(c *Context)  {
}


type StringLiteral struct {
    expType    interface{}
    str string
}

func NewStringLiteral(s string) *StringLiteral {
    return &StringLiteral{
        str: s,
    }
}

func (sl *StringLiteral) visit() string {
    str := fmt.Sprintf("(strlit %s)", sl.str)
    return str
}

func (sl *StringLiteral) analyze(c *Context)  {
}



type Assignment struct {
    expType    interface{}
    lValue  Node
    exp     Node
}

func NewAssignment(lValue Node, exp Node) *Assignment {
    return &Assignment{
        lValue: lValue,
        exp:    exp,
    }
}

func (as *Assignment) visit() string {
    return fmt.Sprintf("(assignment lValue:%v exp:%v)", as.lValue.Exp.visit(), as.exp.Exp.visit())
}

func (as *Assignment) analyze(c *Context)  {
}



type Nil struct {
    expType interface{}
}

func NewNil() *Nil {
    return &Nil{}
}

func (ni *Nil) visit() string {
    return fmt.Sprintf("(nil)")
}

func (ni *Nil) analyze(c *Context)  {
}


type CallExpression struct {
    expType    interface{}
    name    string
    exps    []Node
}

func NewCallExpression(name string, exps []Node) *CallExpression {
    return &CallExpression{
        name: name,
        exps: exps,
    }
}

func (ce *CallExpression) visit() string {
    str := fmt.Sprintf("(callExp: %s", ce.name)
    for i, n := range ce.exps {
        if(n.Exp == nil) {
            str += string(n.Token.Lexeme)
        } else {
            str += fmt.Sprintf("\nparam %d: %v ", i+1, n.Exp.visit())
        }
    }
    str += "\n)"
    return str
}

func (ce *CallExpression) analyze(c *Context)  {
}


type TypeDeclaration struct {
    expType    interface{}
    id    string
    Exp    Node
}

func NewTypeDeclaration(identifier string, n *Node) *TypeDeclaration {
    return &TypeDeclaration{
        id: identifier,
        Exp: *n,
    }
}

func (td *TypeDeclaration) visit() string {
    return fmt.Sprintf("(tyDec: type:%s %s)", td.id, td.Exp.visit())
}

func (td *TypeDeclaration) analyze(c *Context)  {
    td.Exp.analyze(c)
}


type FuncDeclaration struct {
    expType    interface{}
    id     string
    id2     string
    decs   []Node
    exp    Node
}

func NewFuncDeclaration(id string, id2 string, declarations []Node, n Node) *FuncDeclaration {
    return &FuncDeclaration{
        id: id,
        id2: id2,
        decs: declarations,
        exp: n,
    }
}

func (fd *FuncDeclaration) visit() string {
    str := fmt.Sprintf("(funDec: id:%s id2:%s decs:", fd.id, fd.id2)
    for _, n := range fd.decs {
        str += fmt.Sprintf("(%v)\n", n.Exp.visit())
    }
    str += fmt.Sprintf("exp:%s)", fd.exp.Exp.visit())
    return str
}

func (fd *FuncDeclaration) analyze(c *Context)  {
}


type FieldDeclaration struct {
    expType    interface{}
    id    string
    fieldType   string
}

func NewFieldDeclaration(identifier1 string, fieldType string) *FieldDeclaration {
    return &FieldDeclaration{
        id: identifier1,
        fieldType: fieldType,
    }
}

func (fid *FieldDeclaration) visit() string {
    return fmt.Sprintf("fieldDec: (id:%s) (fieldType:%s)", fid.id, fid.fieldType)
}

func (fid *FieldDeclaration) analyze(c *Context)  {
}


type FieldExpression struct {
    expType    interface{}
    lValue    Node
    id        string
}

func NewFieldExpression(lValue Node, id string) *FieldExpression {
    return &FieldExpression{
        lValue: lValue,
        id: id,
    }
}

func (fe *FieldExpression) visit() string {
    return fmt.Sprintf("(fieldExp: (lValue:%v) (id:%s))", fe.lValue.Exp.visit(), fe.id)
}

func (fe *FieldExpression) analyze(c *Context)  {
}


type FieldCreate struct {
    expType    interface{}
    id    string
    exp   Node
}

func NewFieldCreate(identifier string, exp Node) *FieldCreate {
    return &FieldCreate{
        id: identifier,
        exp: exp,
    }
}

func (fc *FieldCreate) visit() string {
    return fmt.Sprintf("fieldCreate: id:%s exp:(%v)", fc.id, fc.exp.Exp.visit())
}

func (fc *FieldCreate) analyze(c *Context)  {
}



type VarDeclaration struct {
    expType    interface{}
    id    string
    typeId    string
    Exp    Node
}

func NewVarDeclaration(identifier1 string, typeId string, n *Node) *VarDeclaration {
    return &VarDeclaration{
        id: identifier1,
        typeId: typeId,
        Exp: *n,
    }
}

func (vd *VarDeclaration) visit() string {
    return fmt.Sprintf("(varDec: id:%s typeId:%s exp:%s)", vd.id, vd.typeId, vd.Exp.visit())
}

func (vd *VarDeclaration) analyze(c *Context)  {
    vd.Exp.analyze(c)
    if(vd.typeId != "") {//If type id is declared then we know the type from a lookup!
        vd.expType = c.lookup(vd.typeId)

        //Check assignable to ?
        isAssignable(vd.Exp.Exp, vd.expType)
    } else { // Inference type from init experssion:O
        vd.expType = vd.Exp.Exp
    }

    //add type to context
    c.add(vd)
}


type Identifier struct {
    expType    interface{}
    id    string
}

func NewIdentifier(identifier string) *Identifier {
    return &Identifier{
        id: identifier,
    }
}

func (id *Identifier) visit() string {
    return fmt.Sprintf("(ID: %s)", id.id)
}


func (id *Identifier) analyze(c *Context)  {
}


type Subscript struct {
    expType    interface{}
    id          string
    expId        *Node
    subscriptExp Node
}

func NewSubscriptExpression(id string, expId *Node, subscriptExp Node) *Subscript {
    return &Subscript{
        id: id,
        expId: expId,
        subscriptExp: subscriptExp,
    }
}

func (se *Subscript) visit() string {
    var str string
    if(se.id != "") {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", se.id, se.subscriptExp.Exp.visit())
    } else {
        str = fmt.Sprintf("(Subscript id:%s exp:%v)", se.expId.Exp.visit(), se.subscriptExp.Exp.visit())
    }
    return str
}

func (se *Subscript) analyze(c *Context)  {
}


type RecordType struct {
    expType    interface{}
    decs    []Node
}

func NewRecordType(decs []Node) *RecordType {
    return &RecordType{
        decs: decs,
    }
}

func (rt *RecordType) visit() string {
    str := fmt.Sprintf("(recTy: decs:(")
    for _, n := range rt.decs {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")"
    return str
}

func (rt *RecordType) analyze(c *Context)  {
}


type RecordCreate struct {
    expType    interface{}
    id string
    decs    []Node
}

func NewRecordCreate(id string, decs []Node) *RecordCreate {
    return &RecordCreate{
        id: id,
        decs: decs,
    }
}

func (rc *RecordCreate) visit() string {
    str := fmt.Sprintf("(recCreate: id:%s decs:(", rc.id)
    for _, n := range rc.decs {
        str += fmt.Sprintf("%v",n.Exp.visit())
    }
    str += ")\n"
    return str
}

func (rc *RecordCreate) analyze(c *Context)  {
}


type ArrayType struct {
    expType    interface{}
    id    string
}

func NewArrayType(identifier string) *ArrayType {
    return &ArrayType{
        id: identifier,
    }
}

func (at *ArrayType) visit() string {
    return fmt.Sprintf("(arrType: %s)", at.id)
}


func (at *ArrayType) analyze(c *Context)  {
     at.expType = c.lookup(at.id)
     fmt.Printf("Assigning type %T to arraytype\n", at.expType)
}


type ArrayCreate struct {
    expType  interface{}
    id    string
    exp1  Node
    exp2  Node
}

func NewArrayCreate(identifier string, exp1 Node, exp2 Node) *ArrayCreate {
    return &ArrayCreate{
        id: identifier,
        exp1: exp1,
        exp2: exp2,
    }
}

func (ac *ArrayCreate) visit() string {
    return fmt.Sprintf("(arrCreate: id:%s exp1:%v exp2:%v)", ac.id, ac.exp1.visit(), ac.exp2.visit())
}

func (ac *ArrayCreate) analyze(c *Context)  {
}



type LetExpression struct {
    expType    interface{}
    decs    []Node
    exps    []Node
}

func NewLetExpression(declarations []Node, expressions []Node) *LetExpression {
    return &LetExpression{
        decs: declarations,
        exps: expressions,
    }
}

func (le *LetExpression) visit() string {
    str := fmt.Sprintf("(letExp: decs:(")
    for _, n := range le.decs {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }
    str += fmt.Sprintf(")\n(exps: ")
    for _, n := range le.exps {
        str += fmt.Sprintf("\n %v",n.Exp.visit())
    }

    str += "))"
    return str
}

func (le *LetExpression) analyze(c *Context)  {
    newContext := c.createChildContextForBlock()
    for _, d := range le.decs {
        td, isTypeDec := d.Exp.(*TypeDeclaration)
        if(isTypeDec) { //If its a type declaration, add it to the new context
            newContext.add(td)
        }

        // typeDec, isTypeDec := d.Exp.(*FuncDeclaration)
        // if()
    }

    for _, d := range le.decs {
        d.analyze(newContext)
    }

    for _, d := range le.exps {
        d.analyze(newContext)
    }

    //If expressions has a body then take the type of the last element
    if(len(le.exps) > 0) {
        le.expType = le.exps[len(le.exps)-1].Exp
    } else {
        le.expType = VoidType{}
    }
}



type IfThenElseExpression struct {
    expType    interface{}
    exp1  Node
    exp2  Node
    exp3  Node
}

func NewIfThenElseExpression(exp1 Node, exp2 Node, exp3 Node) *IfThenElseExpression {
    return &IfThenElseExpression{
        exp1:exp1,
        exp2:exp2,
        exp3:exp3,
    }
}

func (itee *IfThenElseExpression) visit() string {
    return fmt.Sprintf("(ifThenElse if:%v then:%v else:%v)", itee.exp1.Exp.visit(), itee.exp2.Exp.visit(), itee.exp3.Exp.visit())
}

func (itee *IfThenElseExpression) analyze(c *Context)  {
}


type IfThenExpression struct {
    expType    interface{}
    exp1  Node
    exp2  Node
}

func NewIfThenExpression(exp1 Node, exp2 Node) *IfThenExpression {
    return &IfThenExpression{
        exp1:exp1,
        exp2:exp2,
    }
}

func (ite *IfThenExpression) visit() string {
    return fmt.Sprintf("(ifThen if:%v then:%v)", ite.exp1.Exp.visit(), ite.exp2.Exp.visit())
}

func (ite *IfThenExpression) analyze(c *Context)  {
}


type WhileExpression struct {
    expType    interface{}
    exp1  Node
    exp2  Node
}

func NewWhileExpression(exp1 Node, exp2 Node) *WhileExpression {
    return &WhileExpression{
        exp1:exp1,
        exp2:exp2,
    }
}

func (we *WhileExpression) visit() string {
    return fmt.Sprintf("(whileExp cond:%v do:%v)", we.exp1.Exp.visit(), we.exp2.Exp.visit())
}


func (we *WhileExpression) analyze(c *Context)  {
}


type ForExpression struct {
    expType    interface{}
    id    string
    exp1  Node
    exp2  Node
    exp3  Node
}

func NewForExpression(id string, exp1 Node, exp2 Node, exp3 Node) *ForExpression {
    return &ForExpression{
        id:id,
        exp1:exp1,
        exp2:exp2,
        exp3:exp3,
    }
}

func (fe *ForExpression) visit() string {
    return fmt.Sprintf("(forexp id:%s exp1:%v exp2:%v exp3:%v)", fe.id, fe.exp1.Exp.visit(), fe.exp2.Exp.visit(), fe.exp3.Exp.visit())
}

func (fe *ForExpression) analyze(c *Context)  {
}
