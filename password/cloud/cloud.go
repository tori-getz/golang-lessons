package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	panic("to do")

	// return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) error {
	panic("to do")

	// return []byte{}, nil
}
