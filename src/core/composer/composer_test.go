package composer_test

import (
	"testing"

	"github.com/anasmohammad611/compose-cli/core/composer"
	"github.com/anasmohammad611/compose-cli/core/errors"
	"github.com/anasmohammad611/compose-cli/core/types"
)

func Test_AddConfig(t *testing.T) {
	got, err := composer.AddConfig("", types.Template{})

	if err != errors.ErrInvalidYaml {
		t.Errorf("expected %v; got: %v", errors.ErrInvalidYaml, err)
	}

	if got != "" {
		t.Errorf("expected \"\"; got: %v", got)
	}

}
