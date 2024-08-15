package object

type ObjectType string

type Objecet interface {
    Type()      ObjectType
    Inspect()   string
}
