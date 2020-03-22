package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartCommand(t *testing.T) {

	s, err := StartCommand("cmd", "/C", "dir")
	assert.NoError(t, err)
	assert.NotNil(t, s.PID)
	b, _ := json.Marshal(s)
	fmt.Println(string(b))

}
