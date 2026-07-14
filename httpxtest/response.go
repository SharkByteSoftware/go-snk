package httpxtest

import "net/http"

// ResponseFunc is a function that can be used to respond to an HTTP request.
type ResponseFunc func(w http.ResponseWriter, r *http.Request)


// SequencedResponse is a single entry in a sequence of canned responses.
type SequencedResponse struct {
	statusCode int
	body       any
	options    []Option
}

// Response constructs a SequencedResponse for use with OnSequence or OnRouteSequence.
func Response(statusCode int, body any, options ...Option) SequencedResponse {
	return SequencedResponse{statusCode: statusCode, body: body, options: options}
}
