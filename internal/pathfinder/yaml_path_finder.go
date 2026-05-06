package pathfinder

import (
	"unicode/utf8"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/parser"
	"github.com/joomcode/errorx"
)

var (
	ErrFailedParseYaml = errorx.NewType(errorx.CommonErrors, "ErrFailedParseYaml")
)

type YamlPathFinder struct {
}

func (f YamlPathFinder) Find(content string, parts []string) (*Result, error) {
	tokens := lexer.Tokenize(content)
	file, err := parser.Parse(tokens, 0)
	if err != nil {
		return nil, ErrFailedParseYaml.WrapWithNoMessage(err)
	}
	if len(file.Docs) != 1 {
		return nil, errorx.AssertionFailed.New("It's expected to have a single document in yaml file")
	}
	documentNode := file.Docs[0]
	node := documentNode.Body
	for index := range parts {
		mappingValueNode, err := resolveCurrentLevelMappingValueNode(node, parts, index)
		if err != nil {
			return nil, err
		}
		node = mappingValueNode.Value
	}

	if specificNode, ok := node.(*ast.IntegerNode); ok {
		value := specificNode.Token.Value
		byteOffset := runeOffsetToByteOffset(content, specificNode.Token.Position.Offset)
		return &Result{byteOffset, byteOffset + len(value), value}, nil
	}
	if specificNode, ok := node.(*ast.FloatNode); ok {
		value := specificNode.Token.Value
		byteOffset := runeOffsetToByteOffset(content, specificNode.Token.Position.Offset)
		return &Result{byteOffset, byteOffset + len(value), value}, nil
	}
	if specificNode, ok := node.(*ast.StringNode); ok {
		value := specificNode.Token.Value
		byteOffset := runeOffsetToByteOffset(content, specificNode.Token.Position.Offset)
		return &Result{byteOffset, byteOffset + len(value), value}, nil
	}
	return nil, errorx.AssertionFailed.New("Unknown node type '%+v'", node)
}

// runeOffsetToByteOffset converts a rune offset to a byte offset in the content string.
func runeOffsetToByteOffset(content string, runeOffset int) int {
	byteOffset := 0
	runeIndex := 0
	for _, r := range content {
		// Token.Position.Offset is 1-indexed (starts from 1), so we subtract 1 to get 0-indexed rune position.
		if runeIndex >= runeOffset-1 {
			break
		}
		byteOffset += utf8.RuneLen(r)
		runeIndex++
	}
	return byteOffset
}

func resolveCurrentLevelMappingValueNode(node ast.Node, parts []string, partIndex int) (*ast.MappingValueNode, error) {
	nodeName := parts[partIndex]
	if mappingValueNode, ok := node.(*ast.MappingValueNode); ok {
		key, err := getMappingValueNodeKey(mappingValueNode)
		if err != nil {
			return nil, err
		}
		if key == nodeName {
			return mappingValueNode, nil
		} else {
			return nil, errorx.AssertionFailed.New("Can't find path '%v'", parts[:partIndex+1])
		}
	}
	if mappingNode, ok := node.(*ast.MappingNode); ok {
		for _, mappingValueNode := range mappingNode.Values {
			key, err := getMappingValueNodeKey(mappingValueNode)
			if err != nil {
				return nil, err
			}
			if key == nodeName {
				return mappingValueNode, nil
			}
		}
		return nil, errorx.AssertionFailed.New("Can't find path '%v'", parts[:partIndex+1])
	}
	return nil, errorx.AssertionFailed.New("Unknown node type '%v'", node)
}

func getMappingValueNodeKey(node *ast.MappingValueNode) (string, error) {
	if text, ok := node.Key.(*ast.StringNode); ok {
		return text.Value, nil
	}
	return "", errorx.AssertionFailed.New("Node key is expected to be a string, but was '%+v'", node.Key)
}
