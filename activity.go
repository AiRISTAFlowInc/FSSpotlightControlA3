package SpotlightControl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"time"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})


// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	sent := RestCallMoveLightAndChangeColor(input.LightHost, input.Color, input.X, input.Y) // change color
	sleepTime, _:= strconv.Atoi(input.ResetTime)
	time.Sleep(time.Minute * time.Duration(sleepTime)) // Wait for reset
	RestCallMoveLightAndChangeColor(input.LightHost, "0", input.StartX, input.StartY) // reset

	output := &Output{Status: sent}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func RestCallMoveLightAndChangeColor(host string, color string, x string, y string) bool {
	// Create an HTTP client
	client := &http.Client{}

	intColor, _ := strconv.Atoi(color)

	SpotLightId := "0" // HARDCODED
	xRectified, _ := strconv.ParseFloat(x, 64)
	yRectified, _ := strconv.ParseFloat(y, 64)
	spotlightCommand := SpotlightCommand{
		X: xRectified,
		Y: yRectified,
		Z: 0,
		Color: intColor, // 10 is red, 0 is white
	}
	spotlightCommandJSON, err := json.Marshal(spotlightCommand) // serialize to json
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	// Create the request
	url := "http://"+host+"/api/move/"+SpotLightId
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(spotlightCommandJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Add("Content-Type", "application/json") // expect JSON body

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the config JSON into an object	
	var response Response
	errUnmarshal := json.Unmarshal(body, &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
	}

	if response.Message == "Ok"{
		return true
	}

	return false
}
