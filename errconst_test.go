package errconst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	errconst "github.com/gomatic/yze-go-errconst"
)

func TestDisallowedErrorConstructionIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), errconst.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, errconst.Registration.Validate())
	assert.Equal(t, "yze/errconst", errconst.Registration.RuleID())
	assert.Same(t, errconst.Analyzer, errconst.Registration.Analyzer)
}
