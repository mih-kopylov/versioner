package pathfinder

import (
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
		offset := specificNode.Token.Position.Offset - 1
		value := specificNode.Token.Value
		return &Result{offset - len(value), offset, value}, nil
	}
	if specificNode, ok := node.(*ast.FloatNode); ok {
		offset := specificNode.Token.Position.Offset - 1
		value := specificNode.Token.Value
		return &Result{offset - len(value), offset, value}, nil
	}
	if specificNode, ok := node.(*ast.StringNode); ok {
		offset := specificNode.Token.Position.Offset - 1
		value := specificNode.Token.Value
		return &Result{offset - len(value), offset, value}, nil
	}
	return nil, errorx.AssertionFailed.New("Unknown node type '%+v'", node)
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
