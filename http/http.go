package http

import "github.com/bayashi/qq/dictionary"

// Via https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
var Ref = map[int]dictionary.Element{
	100: {
		Subject:     "Continue",
		Description: "The 100 (Continue) status code indicates that the initial part of a request has been received and has not yet been rejected by the server. The server intends to send a final response after the request has been fully received and acted upon.",
	},
	101: {
		Subject:     "Switching Protocols",
		Description: "The 101 (Switching Protocols) status code indicates that the server understands and is willing to comply with the client's request, via the Upgrade header field, for a change in the application protocol being used on this connection. The server MUST generate an Upgrade header field in the response that indicates which protocol(s) will be in effect after this response.",
	},
	102: {
		Subject:     "Processing",
		Description: "The 102 (Processing) status code is an interim response used to inform the client that the server has accepted the complete request, but has not yet completed it.  This status code SHOULD only be sent when the server has a reasonable expectation that the request will take significant time to complete. As guidance, if a method is taking longer than 20 seconds (a reasonable, but arbitrary value) to process the server SHOULD return a 102 (Processing) response. The server MUST send a final response after the request has been completed.",
	},
	200: {
		Subject:     "OK",
		Description: "The 200 (OK) status code indicates that the request has succeeded. The content sent in a 200 response depends on the request method. For the methods defined by this specification, the intended meaning of the content can be summarized as:\nGET 	the target resource\nHEAD 	the target resource, like GET, but without transferring the representation data\nPOST 	the status of, or results obtained from, the action\nPUT, DELETE 	the status of the action\nOPTIONS 	communication options for the target resource\nTRACE 	the request message as received by the server returning the trace",
	},
	201: {
		Subject:     "Created",
		Description: "The 201 (Created) status code indicates that the request has been fulfilled and has resulted in one or more new resources being created. The primary resource created by the request is identified by either a Location header field in the response or, if no Location header field is received, by the target URI.",
	},
	202: {
		Subject:     "Accepted",
		Description: "The 202 (Accepted) status code indicates that the request has been accepted for processing, but the processing has not been completed. The request might or might not eventually be acted upon, as it might be disallowed when processing actually takes place. There is no facility in HTTP for re-sending a status code from an asynchronous operation.",
	},
	203: {
		Subject:     "Non-Authoritative Information",
		Description: "The 203 (Non-Authoritative Information) status code indicates that the request was successful but the enclosed content has been modified from that of the origin server's 200 (OK) response by a transforming proxy. This status code allows the proxy to notify recipients when a transformation has been applied, since that knowledge might impact later decisions regarding the content. For example, future cache validation requests for the content might only be applicable along the same request path (through the same proxies). A 203 response is heuristically cacheable; i.e., unless otherwise indicated by the method definition or explicit cache controls",
	},
	204: {
		Subject:     "No Content",
		Description: "The 204 (No Content) status code indicates that the server has successfully fulfilled the request and that there is no additional content to send in the response content. Metadata in the response header fields refer to the target resource and its selected representation after the requested action was applied.",
	},
	205: {
		Subject:     "Reset Content",
		Description: "The 205 (Reset Content) status code indicates that the server has fulfilled the request and desires that the user agent reset the \"document view\", which caused the request to be sent, to its original state as received from the origin server.",
	},
	206: {
		Subject:     "Partial Content",
		Description: "The 206 (Partial Content) status code indicates that the server is successfully fulfilling a range request for the target resource by transferring one or more parts of the selected representation.",
	},
	207: {
		Subject:     "Multi-Status",
		Description: "The 207 (Multi-Status) status code provides status for multiple independent operations",
	},
	208: {
		Subject:     "Already Reported",
		Description: "The 208 (Already Reported) status code can be used inside a DAV: propstat response element to avoid enumerating the internal members of multiple bindings to the same collection repeatedly.  For each binding to a collection inside the request's scope, only one will be reported with a 200 status, while subsequent DAV:response elements for all other bindings will use the 208 status, and no DAV:response elements for their descendants are included.",
	},
	226: {
		Subject:     "IM Used",
		Description: "The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance- manipulations applied to the current instance.",
	},
	300: {
		Subject:     "Multiple Choices",
		Description: "The 300 (Multiple Choices) status code indicates that the target resource has more than one representation, each with its own more specific identifier, and information about the alternatives is being provided so that the user (or user agent) can select a preferred representation by redirecting its request to one or more of those identifiers.",
	},
	301: {
		Subject:     "Moved Permanently",
		Description: "The 301 (Moved Permanently) status code indicates that the target resource has been assigned a new permanent URI and any future references to this resource ought to use one of the enclosed URIs. The server is suggesting that a user agent with link-editing capability can permanently replace references to the target URI with one of the new references sent by the server. However, this suggestion is usually ignored unless the user agent is actively editing references (e.g., engaged in authoring content), the connection is secured, and the origin server is a trusted authority for the content being edited.",
	},
	302: {
		Subject:     "Found",
		Description: "The 302 (Found) status code indicates that the target resource resides temporarily under a different URI. Since the redirection might be altered on occasion, the client ought to continue to use the target URI for future requests.",
	},
	303: {
		Subject:     "See Other",
		Description: "The 303 (See Other) status code indicates that the server is redirecting the user agent to a different resource, as indicated by a URI in the Location header field, which is intended to provide an indirect response to the original request. A user agent can perform a retrieval request targeting that URI (a GET or HEAD request if using HTTP), which might also be redirected, and present the eventual result as an answer to the original request. Note that the new URI in the Location header field is not considered equivalent to the target URI.",
	},
	304: {
		Subject:     "Not Modified",
		Description: "The 304 (Not Modified) status code indicates that a conditional GET or HEAD request has been received and would have resulted in a 200 (OK) response if it were not for the fact that the condition evaluated to false. In other words, there is no need for the server to transfer a representation of the target resource because the request indicates that the client, which made the request conditional, already has a valid representation; the server is therefore redirecting the client to make use of that stored representation as if it were the content of a 200 (OK) response.",
	},
	305: {
		Subject:     "Use Proxy",
		Description: "The 305 (Use Proxy) status code was defined in a previous version of this specification and is now deprecated",
	},
	306: {
		Subject:     "(Unused)",
		Description: "The 306 status code was defined in a previous version of this specification, is no longer used, and the code is reserved.",
	},
	307: {
		Subject:     "Temporary Redirect",
		Description: "The 307 (Temporary Redirect) status code indicates that the target resource resides temporarily under a different URI and the user agent MUST NOT change the request method if it performs an automatic redirection to that URI. Since the redirection can change over time, the client ought to continue using the original target URI for future requests.",
	},
	308: {
		Subject:     "Permanent Redirect",
		Description: "The 308 (Permanent Redirect) status code indicates that the target resource has been assigned a new permanent URI and any future references to this resource ought to use one of the enclosed URIs. The server is suggesting that a user agent with link-editing capability can permanently replace references to the target URI with one of the new references sent by the server. However, this suggestion is usually ignored unless the user agent is actively editing references (e.g., engaged in authoring content), the connection is secured, and the origin server is a trusted authority for the content being edited.",
	},
	400: {
		Subject:     "Bad Request",
		Description: "The 400 (Bad Request) status code indicates that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).",
	},
	401: {
		Subject:     "Unauthorized",
		Description: "The 401 (Unauthorized) status code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. The server generating a 401 response MUST send a WWW-Authenticate header field containing at least one challenge applicable to the target resource.",
	},
	402: {
		Subject:     "Payment Required",
		Description: "The 402 (Payment Required) status code is reserved for future use.",
	},
	403: {
		Subject:     "Forbidden",
		Description: "The 403 (Forbidden) status code indicates that the server understood the request but refuses to fulfill it. A server that wishes to make public why the request has been forbidden can describe that reason in the response content (if any).",
	},
	404: {
		Subject:     "Not Found",
		Description: "The 404 (Not Found) status code indicates that the origin server did not find a current representation for the target resource or is not willing to disclose that one exists. A 404 status code does not indicate whether this lack of representation is temporary or permanent; the 410 (Gone) status code is preferred over 404 if the origin server knows, presumably through some configurable means, that the condition is likely to be permanent.",
	},
	405: {
		Subject:     "Method Not Allowed",
		Description: "The 405 (Method Not Allowed) status code indicates that the method received in the request-line is known by the origin server but not supported by the target resource. The origin server MUST generate an Allow header field in a 405 response containing a list of the target resource's currently supported methods.",
	},
	406: {
		Subject:     "Not Acceptable",
		Description: "The 406 (Not Acceptable) status code indicates that the target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.",
	},
	407: {
		Subject:     "Proxy Authentication Required",
		Description: "The 407 (Proxy Authentication Required) status code is similar to 401 (Unauthorized), but it indicates that the client needs to authenticate itself in order to use a proxy for this request. The proxy MUST send a Proxy-Authenticate header field containing a challenge applicable to that proxy for the request. The client MAY repeat the request with a new or replaced Proxy-Authorization header field",
	},
	408: {
		Subject:     "Request Timeout",
		Description: "The 408 (Request Timeout) status code indicates that the server did not receive a complete request message within the time that it was prepared to wait. If the client has an outstanding request in transit, it MAY repeat that request. If the current connection is not usable (e.g., as it would be in HTTP/1.1 because request delimitation is lost), a new connection will be used.",
	},
	409: {
		Subject:     "Conflict",
		Description: "The 409 (Conflict) status code indicates that the request could not be completed due to a conflict with the current state of the target resource. This code is used in situations where the user might be able to resolve the conflict and resubmit the request. The server SHOULD generate content that includes enough information for a user to recognize the source of the conflict.",
	},
	410: {
		Subject:     "Gone",
		Description: "The 410 (Gone) status code indicates that access to the target resource is no longer available at the origin server and that this condition is likely to be permanent. If the origin server does not know, or has no facility to determine, whether or not the condition is permanent, the status code 404 (Not Found) ought to be used instead.",
	},
	411: {
		Subject:     "Length Required",
		Description: "The 411 (Length Required) status code indicates that the server refuses to accept the request without a defined Content-Length. The client MAY repeat the request if it adds a valid Content-Length header field containing the length of the request content.",
	},
	412: {
		Subject:     "Precondition Failed",
		Description: "The 412 (Precondition Failed) status code indicates that one or more conditions given in the request header fields evaluated to false when tested on the server. This response status code allows the client to place preconditions on the current resource state (its current representations and metadata) and, thus, prevent the request method from being applied if the target resource is in an unexpected state.",
	},
	413: {
		Subject:     "Content Too Large",
		Description: "The 413 (Content Too Large) status code indicates that the server is refusing to process a request because the request content is larger than the server is willing or able to process. The server MAY terminate the request, if the protocol version in use allows it; otherwise, the server MAY close the connection.",
	},
	414: {
		Subject:     "URI Too Long",
		Description: "The 414 (URI Too Long) status code indicates that the server is refusing to service the request because the target URI is longer than the server is willing to interpret. This rare condition is only likely to occur when a client has improperly converted a POST request to a GET request with long query information, when the client has descended into an infinite loop of redirection (e.g., a redirected URI prefix that points to a suffix of itself) or when the server is under attack by a client attempting to exploit potential security holes.",
	},
	415: {
		Subject:     "Unsupported Media Type",
		Description: "The 415 (Unsupported Media Type) status code indicates that the origin server is refusing to service the request because the content is in a format not supported by this method on the target resource.",
	},
	416: {
		Subject:     "Range Not Satisfiable",
		Description: "The 416 (Range Not Satisfiable) status code indicates that the set of ranges in the request's Range header field has been rejected either because none of the requested ranges are satisfiable or because the client has requested an excessive number of small or overlapping ranges (a potential denial of service attack).",
	},
	417: {
		Subject:     "Expectation Failed",
		Description: "The 417 (Expectation Failed) status code indicates that the expectation given in the request's Expect header field could not be met by at least one of the inbound servers.",
	},
	418: {
		Subject:     "(Unused)",
		Description: "[RFC2324] was an April 1 RFC that lampooned the various ways HTTP was abused; one such abuse was the definition of an application-specific 418 status code, which has been deployed as a joke often enough for the code to be unusable for any future use. Therefore, the 418 status code is reserved in the IANA HTTP Status Code Registry. This indicates that the status code cannot be assigned to other applications currently.",
	},
	421: {
		Subject:     "Misdirected Request",
		Description: "The 421 (Misdirected Request) status code indicates that the request was directed at a server that is unable or unwilling to produce an authoritative response for the target URI. An origin server (or gateway acting on behalf of the origin server) sends 421 to reject a target URI that does not match an origin for which the server has been configured or does not match the connection context over which the request was received",
	},
	422: {
		Subject:     "Unprocessable Content",
		Description: "The 422 (Unprocessable Content) status code indicates that the server understands the content type of the request content (hence a 415 (Unsupported Media Type) status code is inappropriate), and the syntax of the request content is correct, but it was unable to process the contained instructions. For example, this status code can be sent if an XML request content contains well-formed (i.e., syntactically correct), but semantically erroneous XML instructions.",
	},
	423: {
		Subject:     "Locked",
		Description: "The 423 (Locked) status code means the source or destination resource of a method is locked.  This response SHOULD contain an appropriate precondition or postcondition code, such as 'lock-token-submitted' or 'no-conflicting-lock'.",
	},
	424: {
		Subject:     "Failed Dependency",
		Description: "The 424 (Failed Dependency) status code means that the method could not be performed on the resource because the requested action depended on another action and that action failed.  For example, if a command in a PROPPATCH method fails, then, at minimum, the rest of the commands will also fail with 424 (Failed Dependency).",
	},
	425: {
		Subject:     "Too Early",
		Description: "A 425 (Too Early) status code indicates that the server is unwilling to risk processing a request that might be replayed. User agents that send a request in early data are expected to retry the request when receiving a 425 (Too Early) response status code.  A user agent SHOULD retry automatically, but any retries MUST NOT be sent in early data.",
	},
	426: {
		Subject:     "Upgrade Required",
		Description: "The 426 (Upgrade Required) status code indicates that the server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server MUST send an Upgrade header field in a 426 response to indicate the required protocol(s).",
	},
	428: {
		Subject:     "Precondition Required",
		Description: "The 428 status code indicates that the origin server requires the request to be conditional. Its typical use is to avoid the \"lost update\" problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.  By requiring requests to be conditional, the server can assure that clients are working with the correct copies.",
	},
	429: {
		Subject:     "Too Many Requests",
		Description: "The 429 status code indicates that the user has sent too many requests in a given amount of time (\"rate limiting\"). The response representations SHOULD include details explaining the condition, and MAY include a Retry-After header indicating how long to wait before making a new request.",
	},
	431: {
		Subject:     "Request Header Fields Too Large",
		Description: "The 431 status code indicates that the server is unwilling to process the request because its header fields are too large.  The request MAY be resubmitted after reducing the size of the request header fields.",
	},
	451: {
		Subject:     "Unavailable For Legal Reasons",
		Description: "This status code indicates that the server is denying access to the resource as a consequence of a legal demand. The server in question might not be an origin server.  This type of legal demand typically most directly affects the operations of ISPs and search engines.",
	},
	500: {
		Subject:     "Internal Server Error",
		Description: "The 500 (Internal Server Error) status code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.",
	},
	501: {
		Subject:     "Not Implemented",
		Description: "The 501 (Not Implemented) status code indicates that the server does not support the functionality required to fulfill the request. This is the appropriate response when the server does not recognize the request method and is not capable of supporting it for any resource.",
	},
	502: {
		Subject:     "Bad Gateway",
		Description: "The 502 (Bad Gateway) status code indicates that the server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while attempting to fulfill the request.",
	},
	503: {
		Subject:     "Service Unavailable",
		Description: "The 503 (Service Unavailable) status code indicates that the server is currently unable to handle the request due to a temporary overload or scheduled maintenance, which will likely be alleviated after some delay. The server MAY send a Retry-After header field to suggest an appropriate amount of time for the client to wait before retrying the request.",
	},
	504: {
		Subject:     "Gateway Timeout",
		Description: "The 504 (Gateway Timeout) status code indicates that the server, while acting as a gateway or proxy, did not receive a timely response from an upstream server it needed to access in order to complete the request.",
	},
	505: {
		Subject:     "HTTP Version Not Supported",
		Description: "The 505 (HTTP Version Not Supported) status code indicates that the server does not support, or refuses to support, the major version of HTTP that was used in the request message. The server is indicating that it is unable or unwilling to complete the request using the same major version as the client, other than with this error message. The server SHOULD generate a representation for the 505 response that describes why that version is not supported and what other protocols are supported by that server.",
	},
	506: {
		Subject:     "Variant Also Negotiates",
		Description: "The 506 status code indicates that the server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
	},
	507: {
		Subject:     "Insufficient Storage",
		Description: "The 507 (Insufficient Storage) status code means the method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request. This condition is considered to be temporary. If the request that received this status code was the result of a user action, the request MUST NOT be repeated until it is requested by a separate user action.",
	},
	508: {
		Subject:     "Loop Detected",
		Description: "The 508 (Loop Detected) status code indicates that the server terminated an operation because it encountered an infinite loop while processing a request with \"Depth: infinity\".  This status indicates that the entire operation failed.",
	},
	510: {
		Subject:     "Not Extended (OBSOLETED)",
		Description: "The policy for accessing the resource has not been met in the request.  The server should send back all the information necessary for the client to issue an extended request. It is outside the scope of this specification to specify how the extensions inform the client.",
	},
	511: {
		Subject:     "Network Authentication Required",
		Description: "The 511 status code is designed to mitigate problems caused by \"captive portals\" to software (especially non-browser agents) that is expecting a response from the server that a request was made to, not the intervening network infrastructure.  It is not intended to encourage deployment of captive portals -- only to limit the damage caused by them.",
	},
}

func GetRef() map[int]dictionary.Element {
	return Ref
}
