package test

import (
	"fmt"
	"github.com/smartystreets/assertions"
	"path/filepath"
	"testing"
)

func TestExtensionUtil(t *testing.T) {
	ext := filepath.Ext("test.png")
	fmt.Print(ext)
	assertions.ShouldEqual(".png", ext)
}
