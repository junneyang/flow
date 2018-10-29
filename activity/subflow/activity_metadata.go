package subflow

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "flogo-subflow",
  "type": "flogo:activity",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/subflow",
  "version": "0.0.1",
  "title": "Start a SubFlow",
  "description": "Simple SubFlow Activity",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/activity/subflow",
  "dynamicIO": true,
  "settings": [
    {
      "name": "flowURI",
      "type": "string",
      "required": true
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
