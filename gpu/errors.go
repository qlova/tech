package gpu

//Error is a gpu error.
type Error string

func (err Error) Error() string { return string(err) }

//Error values.
const (
	ErrNotOpen      Error = "gpu.Open has not been called"
	ErrUnimplmented Error = "this gpu function is unimplemented"
	ErrNoShader     Error = "gpu.ErrNoShader"
)
