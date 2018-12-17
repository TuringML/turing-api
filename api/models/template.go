package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cbroglie/mustache"
	uuid "github.com/satori/go.uuid"
)

// Template struct holds the XML full template of the playground
type Template struct {
	Workflow string // The whole NiFi template
	GroupID  string // The root group ID
	Snippet  string // The string containing all the Group Processors
}

// NewTemplate return a pointer to a new Template object
func NewTemplate(name string) *Template {
	groupID := uuid.Must(uuid.NewV4()).String()

	context := map[string]string{
		"description": "Root Template",
		"groupId":     groupID,
		"name":        name,
		"timestamp":   time.Now().String(),
	}

	pwd, _ := os.Getwd()
	fmt.Print(pwd)

	t, err := mustache.RenderFile("../models/templates/skeleton.xml", context)
	if err != nil {
		return nil
	}

	return &Template{
		Workflow: t,
		GroupID:  groupID,
		Snippet:  "",
	}
}

// AddFlow will add the XML sub-template for a specific type and subtype
func (t *Template) AddFlow(node *NodeGraph, tp Type, sbt SubType) error {
	switch tp {
	case CollectorType:
		return t.doCollector(node, sbt)
	case EnricherType:
		return t.doEnricher(node, sbt)
	case JoinerType:
		return t.doJoiner(node, sbt)
	case OperatorType:
		return t.doOperator(node, sbt)
	case IntellectorType:
		return t.doIntellector(node, sbt)
	case StorerType:
		return t.doStorer(node, sbt)
	default:
		return errors.New("Not a correct type")
	}
}

// Components

func (t *Template) doCollector(node *NodeGraph, sbt SubType) error {
	switch sbt {
	case S3SubType:
		return t.doS3Collector(node)
	case GCSSubType:
		return t.doGCSCollector(node)
	case ADLSubType:
		return t.doADLCollector(node)
	case KafkaSubType:
		return t.doKafkaCollector(node)
	default:
		return errors.New("Not a correct sub type for Collector")
	}
}

func (t *Template) doEnricher(node *NodeGraph, sbt SubType) error {
	return nil
}

func (t *Template) doJoiner(node *NodeGraph, sbt SubType) error {
	return nil
}

func (t *Template) doOperator(node *NodeGraph, sbt SubType) error {
	return nil
}

func (t *Template) doIntellector(node *NodeGraph, sbt SubType) error {
	return nil
}

func (t *Template) doStorer(node *NodeGraph, sbt SubType) error {
	switch sbt {
	case S3SubType:
		return t.doS3Storer(node)
	case GCSSubType:
		return t.doGCSStorer(node)
	case ADLSubType:
		return t.doADLStorer(node)
	case KafkaSubType:
		return t.doKafkaStorer(node)
	default:
		return errors.New("Not a correct sub type for Storer")
	}
}

// Collectors sub types

/*
	These set of functions will do the following:
	1. Read the appropriate NiFi XML file
	2. Inject eventual credentials if needed
	3. Add the newly created XML to the Main template file
*/
func (t *Template) doS3Collector(node *NodeGraph) error {
	return nil
}

func (t *Template) doGCSCollector(node *NodeGraph) error {
	return nil
}

func (t *Template) doADLCollector(node *NodeGraph) error {
	return nil
}

func (t *Template) doKafkaCollector(node *NodeGraph) error {
	return nil
}

// Enricher sub types

// Joiner sub types

// Operator sub types

// Intellector sub types

// Storer sub types

func (t *Template) doS3Storer(node *NodeGraph) error {
	return nil
}

func (t *Template) doGCSStorer(node *NodeGraph) error {
	return nil
}

func (t *Template) doADLStorer(node *NodeGraph) error {
	return nil
}

func (t *Template) doKafkaStorer(node *NodeGraph) error {
	return nil
}
