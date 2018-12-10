package models

import (
	"database/sql/driver"
	"errors"

	"github.com/jinzhu/gorm"
)

// ChannelType specifies the type of Entrance. It transforms it in enum in SQL
type ChannelType string

const (
	// InputChannelType specify the input enum
	InputChannelType ChannelType = "INPUT"
	// OutputChannelType specify the output enum
	OutputChannelType ChannelType = "OUTPUT"
)

// Scan scans the values in the current class
func (c *ChannelType) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	*c = ChannelType(string(asBytes))
	return nil
}

// Value returns the string representation of the dirver value
func (c ChannelType) Value() (driver.Value, error) {
	return string(c), nil
}

// LinkGorm is used to store the information into the database
type LinkGorm struct {
	gorm.Model
	FromNode    Node        `gorm:"foreignkey:NodeID;auto_preload"`
	FromNodeID  uint        `json:"from_node_id"`
	FromField   Field       `gorm:"foreignkey:FieldID;auto_preload"`
	FromFieldID uint        `json:"from_field_id"`
	ToNode      Node        `gorm:"foreignkey:NodeID;auto_preload"`
	ToNodeID    uint        `json:"to_node_id"`
	ToField     Field       `gorm:"foreignkey:FieldID;auto_preload"`
	ToFieldID   uint        `json:"to_field_id"`
	Type        ChannelType `json:"type" sql:"not null;type:ENUM('INPUT', 'OUTPUT')"`
}

// TableName exports the name links into SQL
func (LinkGorm) TableName() string {
	return "links"
}

// Links is the overall object htat incapsulate inputs and output
type Links struct {
	Link Link `json:"links"`
}

// Link creates the actual link between two nodes. Used to unmarshall the requests
type Link struct {
	Inputs []Input `json:"inputs"`
	Output Output  `json:"output"`
}

// Input describes the input node and field
type Input struct {
	From from `json:"from"`
	To   to   `json:"to"`
}

// Output describes the output node and field
type Output struct {
	From from `json:"from"`
	To   to   `json:"to"`
}

type from struct {
	NodeID  int `json:"node_id"`
	FieldID int `json:"field_id"`
}

type to struct {
	NodeID  int `json:"node_id"`
	FieldID int `json:"field_id"`
}

// GetLinks returns all the links from/to a node
func GetLinks(db *gorm.DB, nodeID int) ([]Links, error) {
	var links []LinkGorm
	if err := db.Where("from_node_id = ? OR to_node_id = ?", nodeID, nodeID).Find(&links).Error; err != nil {
		return nil, err
	}

	var ls []Links
	for _, l := range createLinkObjs(links) {
		ls = append(ls, Links{Link: l})
	}

	return ls, nil
}

// GetLink returns a single link of a given node
func GetLink(db *gorm.DB, ID int) (*Links, error) {
	var link LinkGorm
	if err := db.Where("id = ?", ID).First(&link).Error; err != nil {
		return nil, err
	}

	links := &Links{Link: *createLinkObj(link)}

	return links, nil
}

// CreateLink creates a new Link
func CreateLink(db *gorm.DB, links *Links) (*Links, error) {
	l := links.Link

	// Create input links
	for _, input := range l.Inputs {
		inputLink := LinkGorm{
			FromFieldID: uint(input.From.FieldID),
			FromNodeID:  uint(input.From.NodeID),
			ToFieldID:   uint(input.To.FieldID),
			ToNodeID:    uint(input.To.NodeID),
			Type:        InputChannelType,
		}

		if err := db.Create(&inputLink).Error; err != nil {
			return nil, err
		}
	}

	// Create output link
	outputLink := LinkGorm{
		FromFieldID: uint(l.Output.From.FieldID),
		FromNodeID:  uint(l.Output.From.NodeID),
		ToFieldID:   uint(l.Output.To.FieldID),
		ToNodeID:    uint(l.Output.To.NodeID),
		Type:        OutputChannelType,
	}
	if err := db.Create(&outputLink).Error; err != nil {
		return nil, err
	}

	return links, nil
}

// UpdateLink updates a link
func UpdateLink(db *gorm.DB, links *Links) error {
	l := links.Link

	// Update input links
	for _, input := range l.Inputs {
		inputLink := LinkGorm{
			FromFieldID: uint(input.From.FieldID),
			FromNodeID:  uint(input.From.NodeID),
			ToFieldID:   uint(input.To.FieldID),
			ToNodeID:    uint(input.To.NodeID),
			Type:        InputChannelType,
		}

		if err := db.Save(&inputLink).Error; err != nil {
			return err
		}
	}

	// Update output link
	outputLink := LinkGorm{
		FromFieldID: uint(l.Output.From.FieldID),
		FromNodeID:  uint(l.Output.From.NodeID),
		ToFieldID:   uint(l.Output.To.FieldID),
		ToNodeID:    uint(l.Output.To.NodeID),
		Type:        OutputChannelType,
	}
	if err := db.Save(&outputLink).Error; err != nil {
		return err
	}

	return nil
}

// DeleteLink deletes a link
func DeleteLink(db *gorm.DB, ID int) error {
	if err := db.Where("id = ?", ID).Delete(Link{}).Error; err != nil {
		return err
	}
	return nil
}

func createLinkObj(l LinkGorm) *Link {
	link := &Link{}
	if l.Type == OutputChannelType {
		link.Output = Output{
			From: from{
				FieldID: int(l.FromFieldID),
				NodeID:  int(l.FromNodeID),
			},
			To: to{
				FieldID: int(l.ToFieldID),
				NodeID:  int(l.ToNodeID),
			},
		}
	} else {
		if link.Inputs == nil {
			link.Inputs = []Input{}
		}
		i := Input{
			From: from{
				NodeID:  int(l.FromNodeID),
				FieldID: int(l.FromFieldID),
			},
			To: to{
				NodeID:  int(l.ToNodeID),
				FieldID: int(l.ToFieldID),
			},
		}
		link.Inputs = append(link.Inputs, i)
	}
	return link
}

func createLinkObjs(ls []LinkGorm) []Link {
	var links []Link

	for i := 0; i < len(ls); i++ {
		link := Link{}
		if ls[i].Type == OutputChannelType {
			link.Output = Output{
				From: from{
					FieldID: int(ls[i].FromFieldID),
					NodeID:  int(ls[i].FromNodeID),
				},
				To: to{
					FieldID: int(ls[i].ToFieldID),
					NodeID:  int(ls[i].ToNodeID),
				},
			}
		} else {
			if link.Inputs == nil {
				link.Inputs = []Input{}
			}

			for j := i + 1; j < len(ls); j++ {
				if ls[j].Type == InputChannelType && ls[i].ToNodeID == ls[j].ToNodeID {
					input := Input{
						From: from{
							NodeID:  int(ls[j].FromNodeID),
							FieldID: int(ls[j].FromFieldID),
						},
						To: to{
							NodeID:  int(ls[j].ToNodeID),
							FieldID: int(ls[j].ToFieldID),
						},
					}

					link.Inputs = append(link.Inputs, input)
				}
			}

			// Add current node
			input := Input{
				From: from{
					NodeID:  int(ls[i].FromNodeID),
					FieldID: int(ls[i].FromFieldID),
				},
				To: to{
					NodeID:  int(ls[i].ToNodeID),
					FieldID: int(ls[i].ToFieldID),
				},
			}

			link.Inputs = append(link.Inputs, input)
		}
		links = append(links, link)
	}

	return links
}
