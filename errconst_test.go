package errconst_test

import (
	"testing"

	errconst "github.com/gomatic/yze-go-errconst"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDisallowedErrorConstructionIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), errconst.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, errconst.Registration.Validate())
	assert.Equal(t, "yze/go/errconst", errconst.Registration.RuleID())
	assert.Same(t, errconst.Analyzer, errconst.Registration.Analyzer)
}
