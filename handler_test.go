// nolint: goconst
package bugsnag_test

import (
	"github.com/bugsnag/bugsnag-go"

	bugsnaghandler "handler.emperror.dev/bugsnag"
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
