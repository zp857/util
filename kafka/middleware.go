package kafka

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

func MiddlewareChain(next HandlerFunc, m ...MiddlewareFunc) HandlerFunc {
	if len(m) == 0 {
		return next
	}
	return m[0](MiddlewareChain(next, m[1:]...))
}
