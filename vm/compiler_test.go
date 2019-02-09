package vm

import (
	"testing"

	"github.com/actgardner/gogen-avro/schema"
	"github.com/stretchr/testify/assert"
)

// The compiler handles missing and re-ordered fields
func TestCompilePrimitive(t *testing.T) {
	reader := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "one", "type":"string"},
    {"name": "two", "type":"int"}
  ]
}
`

	writer := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "two", "type":"int"},
    {"name": "three", "type":"string"},
    {"name": "one", "type":"string"}
  ]
}
`

	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(reader))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	writerNs := schema.NewNamespace(false)
	writerType, err := writerNs.TypeForSchema([]byte(writer))
	assert.Nil(t, err)

	err = writerType.ResolveReferences(writerNs)
	assert.Nil(t, err)

	program, err := Compile(writerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

// The compiler handles nested record types
func TestCompileNested(t *testing.T) {
	schemaStr := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": "string"},
    {
      "name": "one", 
      "type": {
        "name": "nested",
        "type": "record",
        "fields": [
           {"name": "one", "type": "int"},
           {"name": "two", "type": "string"}
        ]
      }
    },
    {"name": "two", "type":"int"}
  ]
}
`

	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(schemaStr))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	program, err := Compile(readerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 2},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

// The compiler handles removing nested record types
func TestCompileNestedRemoved(t *testing.T) {
	writer := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": "string"},
    {
      "name": "one", 
      "type": {
        "name": "nested",
        "type": "record",
        "fields": [
           {"name": "one", "type": "int"},
           {"name": "two", "type": "string"}
        ]
      }
    },
    {"name": "two", "type":"int"}
  ]
}
`

	reader := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": "string"},
    {"name": "two", "type":"int"}
  ]
}
`

	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(reader))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	writerNs := schema.NewNamespace(false)
	writerType, err := writerNs.TypeForSchema([]byte(writer))
	assert.Nil(t, err)

	err = writerType.ResolveReferences(writerNs)
	assert.Nil(t, err)

	program, err := Compile(writerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

func TestCompileMap(t *testing.T) {
	schemaStr := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": "string"},
    {
      "name": "map", 
      "type": {
        "type": "map",
        "values": "string"
      }
    },
    {"name": "two", "type": "int"}
  ]
}
`
	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(schemaStr))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	program, err := Compile(readerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: BlockStart, Type: Unused, Field: 65535},
		Instruction{Op: Read, Type: MapKey, Field: 65535},
		Instruction{Op: Append, Type: Unused, Field: 65535},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: BlockEnd, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 2},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

func TestCompileArray(t *testing.T) {
	schemaStr := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": "string"},
    {
      "name": "array", 
      "type": {
        "type": "array",
        "items": "int"
      }
    },
    {"name": "two", "type": "int"}
  ]
}
`
	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(schemaStr))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	program, err := Compile(readerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: BlockStart, Type: Unused, Field: 65535},
		Instruction{Op: Append, Type: Unused, Field: 65535},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: BlockEnd, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Enter, Type: Unused, Field: 2},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

func TestCompilePrimitiveUnion(t *testing.T) {
	schemaStr := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {"name": "first", "type": ["string", "int"]}
  ]
}
`
	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(schemaStr))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	program, err := Compile(readerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: UnionElem, Field: 65535},
		Instruction{Op: SwitchStart, Type: Unused, Field: 65535},
		Instruction{Op: SwitchCase, Type: Unused, Field: 0},
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: SwitchCase, Type: Unused, Field: 1},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: SwitchEnd, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}

func TestCompileRecordUnion(t *testing.T) {
	schemaStr := `
{
  "type": "record",
  "name": "test",
  "fields": [
    {
      "name": "first", 
      "type": [
        "string",
        {"type": "record", "name": "innertest", "fields": [{"name": "second", "type":"int"}]}
      ]
    }
  ]
}
`
	readerNs := schema.NewNamespace(false)
	readerType, err := readerNs.TypeForSchema([]byte(schemaStr))
	assert.Nil(t, err)

	err = readerType.ResolveReferences(readerNs)
	assert.Nil(t, err)

	program, err := Compile(readerType, readerType)

	expectedProgram := []Instruction{
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: UnionElem, Field: 65535},
		Instruction{Op: SwitchStart, Type: Unused, Field: 65535},
		Instruction{Op: SwitchCase, Type: Unused, Field: 0},
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: String, Field: 65535},
		Instruction{Op: Set, Type: String, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: SwitchCase, Type: Unused, Field: 1},
		Instruction{Op: Enter, Type: Unused, Field: 1},
		Instruction{Op: Enter, Type: Unused, Field: 0},
		Instruction{Op: Read, Type: Int, Field: 65535},
		Instruction{Op: Set, Type: Int, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
		Instruction{Op: SwitchEnd, Type: Unused, Field: 65535},
		Instruction{Op: Exit, Type: Unused, Field: 65535},
	}
	assert.Equal(t, expectedProgram, program)
	assert.Nil(t, err)
}