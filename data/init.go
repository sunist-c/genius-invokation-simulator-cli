package data

var (
	context *Context
)

func GetContext() (ctx *Context) {
	if context == nil {
		context = &Context{
			Initialized: true,
		}
	}
	return context
}
