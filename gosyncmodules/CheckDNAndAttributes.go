package gosyncmodules

import (
	"github.com/nohupped/ldap" //using a forked version that includes custom methods to retrieve and edit *AddRequest struct.
	//"fmt"
)

func IfDNExists(checkfor *ldap.AddRequest , checkin []*ldap.AddRequest ) bool {
	for _, i := range checkin {
		if checkfor.DN == i.DN {
			return true
		}

	}
	return false
}

/*func (attributes *ldap.Attribute)()  {
	
}*/