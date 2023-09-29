package SpotlightControl

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	LightHost string `md:"LightHost,required"`
	X string `md:"X,required"`
	Y string `md:"Y,required"`
	StartX string `md:"StartX,required"`
	StartY string `md:"StartY,required"`
	Color string `md:"Color,required"`
	ResetTime string `md:"ResetTime,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["LightHost"])
	r.LightHost = strVal

	strVal, _ = coerce.ToString(values["X"])
	r.X = strVal

	strVal, _ = coerce.ToString(values["Y"])
	r.Y = strVal

	strVal, _ = coerce.ToString(values["StartX"])
	r.StartX = strVal

	strVal, _ = coerce.ToString(values["StartY"])
	r.StartY = strVal

	strVal, _ = coerce.ToString(values["Color"])
	r.Color = strVal

	strVal, _ = coerce.ToString(values["ResetTime"])
	r.ResetTime = strVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"LightHost": r.LightHost,
		"X": r.X,
		"Y": r.Y,
		"StartX": r.StartX,
		"StartY": r.StartY,
		"Color": r.Color,
		"ResetTime": r.ResetTime,
	}
}

type Output struct {
	Status bool `md:"Status"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToBool(values["Status"])
	o.Status = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Status": o.Status,
	}
}

type Response struct {
	Message string `json:"message"`
}

type SpotlightCommand struct {
    X float64 
    Y float64 
    Z float64 
	Color int 
}