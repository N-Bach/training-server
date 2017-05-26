package requestModel

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValid_Test1(t *testing.T) {
	req := RequestFeedback{
		Title: "Test title",
		Description: "Test description",
	}

	result := req.IsValid()

	assert.Nil(t, result)
}

func TestIsValid_Test2(t *testing.T) {
	req := RequestFeedback{
		Title: "",
		Description: "Test description",
	}

	result := req.IsValid()

	assert.EqualError(t, result, "Empty title")

}
func TestIsValid_Test3(t *testing.T) {
	req := RequestFeedback{
		Title: "Test title",
		Description: "",
	}
	result := req.IsValid()

	assert.Nil(t, result)
}

func TestIsValid_Test4(t *testing.T) {
	req := RequestFeedback{
		Title: "",
		Description: "",
	}
	result := req.IsValid()

	assert.EqualError(t, result, "Empty title")
}
