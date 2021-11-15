package setting

import (
	"fmt"
	"testing"
)

func TestSetup(t *testing.T) {
	Setup()
	fmt.Println(RedisSetting.Host)
}
