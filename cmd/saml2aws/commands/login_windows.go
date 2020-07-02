package commands

import (
	"github.com/anoop2811/saml2aws/v2/helper/credentials"
	"github.com/anoop2811/saml2aws/v2/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}
