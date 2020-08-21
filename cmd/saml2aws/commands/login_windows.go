package commands

import (
	"github.com/guidewire/saml2aws/v2/helper/credentials"
	"github.com/guidewire/saml2aws/v2/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}
