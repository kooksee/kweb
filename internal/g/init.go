package g

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/g/assert"
	"github.com/kooksee/g/try"
	"github.com/kooksee/g/utils"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

var If = utils.If
var IpAddress = utils.IpAddress

var AssertBool = assert.Bool
var AssertErr = assert.Err
var Assert = assert.MustNotError

var Try = try.Try
