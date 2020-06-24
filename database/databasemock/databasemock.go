package databasemock

//go:generate moq -pkg databasemock -out database.go .. Database
//go:generate moq -pkg databasemock -out rows.go .. Rows
//go:generate moq -pkg databasemock -out row.go .. Row
//go:generate moq -pkg databasemock -out transaction.go .. Transaction

// DefaultRowsMock returns a quick-to-use basic instance of RowsMock.
func DefaultRowsMock() *RowsMock {
	return &RowsMock{
		ScanFunc: func(dest ...interface{}) error {
			return nil
		},
		NextFunc: func() bool {
			return false
		},
		CloseFunc: func() error {
			return nil
		},
		ErrFunc: func() error {
			return nil
		},
	}
}

// DefaultRowMock returns a quick-to-use basic instance of RowMock.
func DefaultRowMock() *RowMock {
	return &RowMock{
		ScanFunc: func(dest ...interface{}) error {
			return nil
		},
	}
}
