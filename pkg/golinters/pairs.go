package golinters

import (
	"github.com/ZipRecruiter/splinter/pairs"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewPairs(settings config.Settings) *goanalysis.Linter {
	// we configure the analyzer here because the goanalysis.Linter.configureAnalyzer
	// method assumes that flags only ever get set once.

	a := pairs.NewAnalyzer()
	for k, values := range settings {
		for _, v := range values {
			if err := a.Flags.Set(k, v); err != nil {
				// this panics because NewPairs, and more
				// importantly the caller,
				// lintersdb.Manager.GetAllSuppportedLinterConfigs,
				// has no error return.
				panic(err.Error())
			}
		}
	}
	return goanalysis.NewLinter(
		"pairs",
		"Validates key/value paired functions",
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeWholeProgram)
}
