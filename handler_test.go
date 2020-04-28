// nolint: goconst
package bugsnag_test

import (
	"github.com/bugsnag/bugsnag-go"

	bugsnaghandler "emperror.dev/handler/bugsnag"
)

func ExampleNew() {
	apiKey := "key"

	_ = bugsnaghandler.New(apiKey)

	// Output:
}

func ExampleNewFromNotifier() {
	apiKey := "key"

	_ = bugsnaghandler.NewFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey: apiKey,
	}))

	// Output:
}
