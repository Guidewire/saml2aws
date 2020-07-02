package commands

import (
	"github.com/anoop2811/saml2aws/v2/helper/credentials"
	"github.com/anoop2811/saml2aws/v2/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}
