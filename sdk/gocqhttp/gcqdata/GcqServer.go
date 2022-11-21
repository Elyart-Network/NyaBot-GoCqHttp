package gcqdata

import jsoniter "github.com/json-iterator/go"

var Handler CqHandler
var CqHttpHost = "http://127.0.0.1:4000/"
var ServerPort = "3000"
var DebugMode = false
var FileLogger = false
var JsonLib = jsoniter.ConfigCompatibleWithStandardLibrary
