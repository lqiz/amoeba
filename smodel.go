package amoeba

type Type string

const (
	TypeInt     Type = "integer"
	TypeBoolean Type = "boolean"
	TypeFloat   Type = "float"
	TypeMap     Type = "map"
	TypeArray   Type = "array"
	TypeStruct  Type = "object"
	TypeString  Type = "string"
)

// Schema represents JSON schema.  Value and Key  are additional
type Schema struct {
	JsonType   string `json:"type"`
	Properties map[string]*Schema
	Value      string
	Key        string
	Default    interface{}
	Items      *Schema
	encoder    encoderFunc `json:"-"`
}

func (s *Schema) mustBe(expected Type) {
	// TODO
	//if Kind(f&flagKindMask) != expected {
	//	panic(&ValueError{methodName(), f.kind()})
	//}
}

func (s *Schema) kind() Type {
	return Type(s.JsonType)
}

func (s *Schema) IsNil() bool {
	// TODO
	return false
}
