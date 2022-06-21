package math_test

import (
	. "GOLANG-SIMPLE-UNIT-TEST"
	"fmt"
	"testing"
	"time"
)

// Eps 4. Test Table Driven (Golang Testing Series)
func TestAdd(t *testing.T) {
	fmt.Println("MASUK TEST ADD")
	testTable := []struct {
		a, b			int
		expectedOutcome	int
	}{
		{a: 1, b: 2, expectedOutcome: 3},
		{a: -1, b: 2, expectedOutcome: 1},
		{a: 1, b: -2, expectedOutcome: -1},
		{a: -1, b: -2, expectedOutcome: -3},
		{a: 0, b: 0, expectedOutcome: 0},
	}

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("Got: %d But expect %d", result, test.expectedOutcome)
			t.Fail()
		}
	}
}

// Eps 5. Hierarchical Test Dengan Subtest (Golang Testing Series)
// func TestAddWithHierarchical(t *testing.T) {
// 	t.Run("a=positive", func(t *testing.T) {
// 		a := 10	
// 		t.Run("b=positive", func(t *testing.T) {
// 			result := Add(a, 5)
// 			if result != 15 {
// 				t.Logf("Got: %d But expect %d", result, 15)
// 				t.Fail()
// 			}
// 		})

// 		t.Run("b=negative", func(t *testing.T) {
// 			result := Add(a, -5)
// 			if result != 5 {
// 				t.Logf("Got: %d But expect %d", result, 5)
// 				t.Fail()
// 			}
// 		})
// 	})
// }


func TestNeedsToBeSkip(t *testing.T) {
	t.Skip("this test will be skipped")
	fmt.Println("MASUK TEST SKIP")
}

func TestCallToDb(t *testing.T) {
	// RUNNING THIS QUERY WITH "go test -v -short", "go test -v -short"
	if testing.Short() {
		t.Skip("SKIP BECAUSE CALLING DB IS WAY TO LONG")
	}
	<-time.After(3 * time.Second)
}