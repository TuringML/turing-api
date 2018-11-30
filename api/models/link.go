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
func GetLinks(db *gorm.DB, ID int) ([]Link, error) {
	var links []Link
	if err := db.Where("from_node_id = ? OR to_node_id = ?", ID, ID).Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

// GetLink returns a single link of a given node
func GetLink(db *gorm.DB, ID int) (*Link, error) {
	var link Link
	if err := db.Where("id = ?", ID).Find(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// CreateLink creates a new Link
func CreateLink(db *gorm.DB, l Link) (*Link, error) {
	if err := db.Create(&l).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

// UpdateLink updates a link
func UpdateLink(db *gorm.DB, ID int, l Link) error {
	if err := db.Where("id = ?", ID).Save(&l).Error; err != nil {
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
