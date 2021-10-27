package gpu

//Error is a gpu error.
type gpuError string

func (err gpuError) Error() string { return string(err) }

//Error values.
const (
	ErrNotOpen      gpuError = "gpu.Open has not been called"
	ErrUnimplmented gpuError = "this gpu function is unimplemented"
	ErrNoShader     gpuError = "gpu.ErrNoShader"
)
