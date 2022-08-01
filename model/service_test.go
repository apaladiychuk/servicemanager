package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartCommand(t *testing.T) {

	s, err := StartCommand("cmd", "/C", "dir")
	assert.NoError(t, err)
	assert.NotNil(t, s.PID)
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

}
