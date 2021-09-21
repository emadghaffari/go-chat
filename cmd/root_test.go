package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "a",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Runner.RootCmd()
		})
	}

	rootCmd.Run(&rootCmd, []string{})
}
