package jsonrpc

import "github.com/bayashi/qq/dictionary"

// http://www.jsonrpc.org/specification#error_object
var Ref = map[int]dictionary.Element{
	-32700: {
		Subject:     "Parse error",
		Description: "Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.",
	},
	-32600: {
		Subject:     "Invalid Request",
		Description: "The JSON sent is not a valid Request object.",
	},
	-32601: {
		Subject:     "Method not found",
		Description: "The method does not exist / is not available.",
	},
	-32602: {
		Subject:     "Invalid params",
		Description: "Invalid method parameter(s).",
	},
	-32603: {
		Subject:     "Internal error",
		Description: "Internal JSON-RPC error.",
	},
}

func GetRef() map[int]dictionary.Element {
	for i := 0; i < 100; i++ {
		Ref[(32000+i)*-1] = dictionary.Element{
			Subject:     "Server error",
			Description: "Reserved for implementation-defined server-errors.",
		}
	}

	return Ref
}
