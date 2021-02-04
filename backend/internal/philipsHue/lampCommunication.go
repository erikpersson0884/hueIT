package philipsHue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/viddem/huego/internal/utilities"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SetLampCall(values *utilities.LampData, secrets *utilities.HueSecrets, lampNumber uint16) error {
	client := &http.Client{}
	jsonData, err := json.Marshal(values)

	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/lights/%d/state", secrets.BaseUrl, lampNumber)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

type LampsDataJson struct {
	Lights map[string]LampDataJson `json:"lights"`
	Groups map[string]interface{} `json:"groups"`
	Config map[string]interface{} `json:"config"`
	Schedules map[string]interface{} `json:"schedules"`
	Scenes map[string]interface{} `json:"scenes"`
	Rules map[string]interface{} `json:"rules"`
	Sensors map[string]interface{} `json:"sensors"`
	ResourceLinks map[string]interface{} `json:"resourceLinks"`
}

type LampDataJson struct {
	State utilities.LampData `json:"state"`
}

type LampWithCoordinates struct {
	Light utilities.Light `json:"light"`
	State utilities.LampDataRGB `json:"state"`
}


func GetLightsInfo(secrets *utilities.HueSecrets) ([]LampWithCoordinates, error) {
	res, err := http.Get(secrets.BaseUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var lampsData LampsDataJson
	err = json.Unmarshal(body, &lampsData)
	if err != nil {
		return nil, err
	}

	var lampDatas []LampWithCoordinates
	for key, val := range lampsData.Lights {
		id, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}

		light, err := secrets.GetLightFromMap(uint16(id))
		if err == nil {
			lampDatas = append(lampDatas, LampWithCoordinates{
				Light: light,
				State: val.State.ToRGB(),
			})
		}
	}

	return lampDatas, nil
}
