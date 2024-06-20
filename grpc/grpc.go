package grpc

import "github.com/bayashi/qq/dictionary"

// Via https://github.com/googleapis/googleapis/blob/3597f7db2191c00b100400991ef96e52d62f5841/google/rpc/code.proto
var Ref = map[int]dictionary.Element{
	0: {
		Subject:  "OK",
		Description: "Not an error; returned on success.\nHTTP Mapping: 200 OK",
	},
	1: {
		Subject:  "Canceled",
		Description: "The operation was cancelled, typically by the caller.\nHTTP Mapping: 499 Client Closed Request",
	},
	2: {
		Subject:  "Unknown",
		Description: "Unknown error. For example, this error may be returned when a `Status` value received from another address space belongs to an error space that is not known in this address space.  Also errors raised by APIs that do not return enough error information may be converted to this error.\nHTTP Mapping: 500 Internal Server Error",
	},
	3: {
		Subject:  "InvalidArgument",
		Description: "The client specified an invalid argument.  Note that this differs from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments that are problematic regardless of the state of the system (e.g., a malformed file name).\nHTTP Mapping: 400 Bad Request",
	},
	4: {
		Subject:  "DeadlineExceeded",
		Description: "The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully.  For example, a successful response from a server could have been delayed long enough for the deadline to expire.\nHTTP Mapping: 504 Gateway Timeout",
	},
	5: {
		Subject:  "NotFound",
		Description: "Some requested entity (e.g., file or directory) was not found. Note to server developers: if a request is denied for an entire class of users, such as gradual feature rollout or undocumented allowlist, `NOT_FOUND` may be used. If a request is denied for some users within a class of users, such as user-based access control, `PERMISSION_DENIED` must be used.\nHTTP Mapping: 404 Not Found",
	},
	6: {
		Subject:  "AlreadyExists",
		Description: "The entity that a client attempted to create (e.g., file or directory) already exists.\nHTTP Mapping: 409 Conflict",
	},
	7: {
		Subject:  "PermissionDenied",
		Description: "The caller does not have permission to execute the specified operation. `PERMISSION_DENIED` must not be used for rejections caused by exhausting some resource (use `RESOURCE_EXHAUSTED` instead for those errors). `PERMISSION_DENIED` must not be used if the caller can not be identified (use `UNAUTHENTICATED` instead for those errors). This error code does not imply the request is valid or the requested entity exists or satisfies other pre-conditions.\nHTTP Mapping: 403 Forbidden",
	},
	8: {
		Subject:  "ResourceExhausted",
		Description: "Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space.\nHTTP Mapping: 429 Too Many Requests",
	},
	9: {
		Subject:  "FailedPrecondition",
		Description: "The operation was rejected because the system is not in a state required for the operation's execution.  For example, the directory to be deleted is non-empty, an rmdir operation is applied to a non-directory, etc.\nService implementors can use the following guidelines to decide between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:\n(a) Use `UNAVAILABLE` if the client can retry just the failing call.\n(b) Use `ABORTED` if the client should retry at a higher level. For\n    example, when a client-specified test-and-set fails, indicating the\n    client should restart a read-modify-write sequence.\n(c) Use `FAILED_PRECONDITION` if the client should not retry until\n    the system state has been explicitly fixed. For example, if an \"rmdir\"\n    fails because the directory is non-empty, `FAILED_PRECONDITION`\n    should be returned since the client should not retry unless\n    the files are deleted from the directory.\nHTTP Mapping: 400 Bad Request",
	},
	10: {
		Subject:  "Aborted",
		Description: "The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort.\nHTTP Mapping: 409 Conflict",
	},
	11: {
		Subject:  "OutOfRange",
		Description: "The operation was attempted past the valid range. E.g., seeking or reading past end-of-file.\nUnlike `INVALID_ARGUMENT`, this error indicates a problem that may be fixed if the system state changes. For example, a 32-bit file system will generate `INVALID_ARGUMENT` if asked to read at an offset that is not in the range [0,2^32-1], but it will generate `OUT_OF_RANGE` if asked to read from an offset past the current file size.\nThere is a fair bit of overlap between `FAILED_PRECONDITION` and `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific error) when it applies so that callers who are iterating through a space can easily look for an `OUT_OF_RANGE` error to detect when they are done.\nHTTP Mapping: 400 Bad Request",
	},
	12: {
		Subject:  "Unimplemented",
		Description: "The operation is not implemented or is not supported/enabled in this service.\nHTTP Mapping: 501 Not Implemented",
	},
	13: {
		Subject:  "Internal",
		Description: "Internal errors.  This means that some invariants expected by the underlying system have been broken.  This error code is reserved for serious errors.\nHTTP Mapping: 500 Internal Server Error",
	},
	14: {
		Subject:  "Unavailable",
		Description: "The service is currently unavailable.  This is most likely a transient condition, which can be corrected by retrying with a backoff. Note that it is not always safe to retry non-idempotent operations.\nHTTP Mapping: 503 Service Unavailable",
	},
	15: {
		Subject:  "DataLoss",
		Description: "Unrecoverable data loss or corruption.\nHTTP Mapping: 500 Internal Server Error",
	},
	16: {
		Subject:  "Unauthenticated",
		Description: "The request does not have valid authentication credentials for the operation.\nHTTP Mapping: 401 Unauthorized",
	},
}

func GetRef() map[int]dictionary.Element {
	return Ref
}
