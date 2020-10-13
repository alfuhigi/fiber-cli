package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

func Test_Root_Execute(t *testing.T) {
	setupOsExit()
	defer teardownOsExit()

	b := &bytes.Buffer{}
	rootCmd.SetErr(b)
	rootCmd.SetOut(b)
	rootCmd.RunE = func(_ *cobra.Command, _ []string) error {
		return errors.New("fake error")
	}

	Execute()
}

func Test_Root_FiberCmd(t *testing.T) {
	assert.Equal(t, rootCmd, FiberCmd())
}