package kweb

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/g/assert"
	"github.com/kooksee/g/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var if_ = utils.If
var ipAddress = utils.IpAddress

var assertBool = assert.Bool
var assertErr = assert.Err
