package responder

type Response struct {
	Res  interface{}
	Code int
}

func (r Response) Metadata() map[string]interface{} {
	return map[string]interface{}{
		"author":      "nitwhiz",
		"license":     "wtfpl",
		"license-url": "http://www.wtfpl.net/about/",
	}
}

func (r Response) Result() interface{} {
	return r.Res
}

func (r Response) StatusCode() int {
	return r.Code
}
