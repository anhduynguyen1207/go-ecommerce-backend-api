package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )
	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	// }

	assert.Equal(t, AddOne(2), 4, "AddOne(2) should return 3")
}

// // Hàm này sẽ dừng lại ở dòng require.Equal vì nó fail
// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3)
// 	fmt.Println("Not executing TestRequire")
// }

// Hàm này sẽ chạy hết dù có fail
// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	fmt.Println("executing")
// }
