package db

// Test struct
type Test struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Data string `json:"data"`
}

// AddTest add test data
func AddTest(data string) (*Test, error) {
	test := Test{
		Data: data,
	}

	if err := db.Create(&test).Error; err != nil {
		return nil, err
	}

	return &test, nil
}

// GetTest get test data
func GetTest() ([]*Test, error) {
	var test []*Test

	if err := db.Find(&test).Error; err != nil {
		return nil, err
	}

	return test, nil
}
