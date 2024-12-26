package main


func main() {
	var numbers []int

	_, err := getElement(numbers, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getElement(slice []int, index int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	if index < 0 || index >= len(slice) {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	return slice[index], nil
}


func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://YOUR_SENTRY_DSN@example.com/1",
	})
	if err != nil {
		zap.L().Fatal("Error initializing Sentry", zap.Error(err))
	}
	defer sentry.Flush(2)

	var numbers []int

	// Attempting to access a nil slice
	if numbers != nil {
		zap.L().Info("Accessing element at index 0", zap.Int("index", 0), zap.Ints("slice", numbers))
	} else {
		zap.L().Error("Slice is nil")
		sentry.CaptureMessage("Slice is nil")
	}
}  


func TestGetElement(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		index    int
		expected int
		wantErr  bool
	}{
		{
			name:     "Empty Slice",
			slice:    []int{},
			index:    0,
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Valid Index",
			slice:    []int{1, 2, 3},
			index:    1,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Out-of-Bounds Index",
			slice:    []int{1, 2, 3},
			index:    3,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getElement(tc.slice, tc.index)
			if tc.wantErr != (err != nil) {
				t.Errorf("Expected error: %v, Got: %v", tc.wantErr, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result: %d, Got: %d", tc.expected, result)