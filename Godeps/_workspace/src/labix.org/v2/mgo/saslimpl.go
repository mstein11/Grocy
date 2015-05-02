//+build sasl

package mgo

import (
	"local/Grocy/Godeps/_workspace/src/labix.org/v2/mgo/sasl"
)

func saslNew(cred Credential, host string) (saslStepper, error) {
	return sasl.New(cred.Username, cred.Password, cred.Mechanism, cred.Service, host)
}
