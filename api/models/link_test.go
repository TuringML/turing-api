package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const linkJSON = `
{
	"links": {
		"inputs": [{
			"from": {
				"node_id": 1,
				"field_id": 1
			},
			"to": {
				"node_id": 2,
				"field_id": 4
			}
		}],
		"output": {
			"from": {
				"node_id": 2,
				"field_id": 3
			},
			"to": {
				"node_id": 3,
				"field_id": 5
			}
		}
	}
}
`

func TestCreateLink(t *testing.T) {
	links := &Links{}
	err := json.Unmarshal([]byte(linkJSON), links)
	if err != nil {
		t.Errorf("TestCreateLink failed with error %v", err)
		return
	}

	l, err := CreateLink(testDb, links)
	if err != nil {
		t.Errorf("TestCreateLink failed with error %v", err)
		return
	}

	assert.Equal(t, 1, l.Link.Inputs[0].From.NodeID)
	assert.Equal(t, 1, l.Link.Inputs[0].From.FieldID)
	assert.Equal(t, 2, l.Link.Inputs[0].To.NodeID)
	assert.Equal(t, 4, l.Link.Inputs[0].To.FieldID)

	assert.Equal(t, 2, l.Link.Output.From.NodeID)
	assert.Equal(t, 3, l.Link.Output.From.FieldID)
	assert.Equal(t, 3, l.Link.Output.To.NodeID)
	assert.Equal(t, 5, l.Link.Output.To.FieldID)
}

func TestGetLinkInput(t *testing.T) {
	l, err := GetLink(testDb, 2)
	if err != nil {
		t.Errorf("TestGetLink failed with error %v", err)
	}
	assert.Equal(t, 3, l.Link.Inputs[0].From.NodeID)
	assert.Equal(t, 5, l.Link.Inputs[0].From.FieldID)
	assert.Equal(t, 1, l.Link.Inputs[0].To.NodeID)
	assert.Equal(t, 1, l.Link.Inputs[0].To.FieldID)
}

func TestGetLinkOutput(t *testing.T) {
	l, err := GetLink(testDb, 1)
	if err != nil {
		t.Errorf("TestGetLink failed with error %v", err)
	}
	assert.Equal(t, 4, l.Link.Output.From.NodeID)
	assert.Equal(t, 6, l.Link.Output.From.FieldID)
	assert.Equal(t, 2, l.Link.Output.To.NodeID)
	assert.Equal(t, 3, l.Link.Output.To.FieldID)
}

func TestGetLinkFail(t *testing.T) {
	l, err := GetLink(testDb, 100)
	if l != nil {
		t.Errorf("TestGetLinkFail failed. Expected nil, got %v", l)
	}
	assert.NotNil(t, err)
}

func TestGetLinks(t *testing.T) {
	ls, err := GetLinks(testDb, 1)
	if err != nil {
		t.Errorf("TestGetLinks failed with error %v", err)
	}
	assert.Equal(t, 3, len(ls))
}

func TestUpdateLink(t *testing.T) {
	l, err := GetLink(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdateLink failed. Expected nil, got %v", err)
	}

	l.Link.Output.From.FieldID = int(3)

	err = UpdateLink(testDb, l)
	if err != nil {
		t.Errorf("TestUpdateLink failed. Expected nil, got %v", err)
	}

	assert.Equal(t, 4, l.Link.Output.From.NodeID)
	assert.Equal(t, 3, l.Link.Output.From.FieldID)
	assert.Equal(t, 2, l.Link.Output.To.NodeID)
	assert.Equal(t, 3, l.Link.Output.To.FieldID)
}

func TestDeleteLink(t *testing.T) {
	err := DeleteLink(testDb, 1)
	if err != nil {
		t.Errorf("TestDeleteLink failed. Expected nil, got %v", err)
	}
}
