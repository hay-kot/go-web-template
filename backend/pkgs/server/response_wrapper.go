package server

// Wrapper is a helper type to wrap a response struct in a JSON response.
type Wrapper map[string]interface{}

// Wrap creates a Wrapper instance and adds the initial namespace and data to be returned.
func Wrap(namespace string, data interface{}) Wrapper {
	return Wrapper{
		namespace: data,
	}
}

// Add a key value pair to the top level of the response wrapper.
func (w Wrapper) Add(key string, value interface{}) Wrapper {
	w[key] = value
	return w
}

// Message adds a "message" field to the json response.
func (w Wrapper) Message(msg string) Wrapper {
	return w.Add("message", msg)
}

// Details adds a "details" field to the json response
func (w Wrapper) Details(details interface{}) Wrapper {
	return w.Add("details", details)
}
