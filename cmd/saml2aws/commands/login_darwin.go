package commands

import (
	"github.com/guidewire/saml2aws/v2/helper/credentials"
	"github.com/guidewire/saml2aws/v2/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}
