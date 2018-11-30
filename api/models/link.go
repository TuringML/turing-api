package models

import "github.com/jinzhu/gorm"

// Link creates the actual link between two nodes
type Link struct {
	gorm.Model
	FromNode    Node  `gorm:"foreignkey:NodeID;auto_preload"`
	FromNodeID  uint  `json:"from_node_id"`
	FromField   Field `gorm:"foreignkey:FieldID;auto_preload"`
	FromFieldID uint  `json:"from_field_id"`
	ToNode      Node  `gorm:"foreignkey:NodeID;auto_preload"`
	ToNodeID    uint  `json:"to_node_id"`
	ToField     Field `gorm:"foreignkey:FieldID;auto_preload"`
	ToFieldID   uint  `json:"to_field_id"`
}

// GetLinks returns all the links from/to a node
func GetLinks(db *gorm.DB, nodeID int) ([]Node, error) {
	return nil, nil
}

// GetLink returns a single link of a given node
func GetLink(db *gorm.DB, linkID int) (*Node, error) {
	return nil, nil
}

// CreateLink creates a new Link
func CreateLink(db *gorm.DB, l Link) (*Node, error) {
	return nil, nil
}

// UpdateLink updates a link
func UpdateLink(db *gorm.DB, linkID int, l Link) error {
	return nil
}

// DeleteLink deletes a link
func DeleteLink(db *gorm.DB, linkID int) error {
	return nil
}
