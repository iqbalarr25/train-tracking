package telematic

import (
	"TrainTracking/internal/features/model"
	"encoding/json"
	"gopkg.in/guregu/null.v3"
	"strconv"
)

type Icon struct {
	Id       int         `json:"id"`
	UserId   interface{} `json:"user_id"`
	Type     string      `json:"type"`
	Order    int         `json:"order"`
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Path     string      `json:"path"`
	ByStatus int         `json:"by_status"`
}

type IconColors struct {
	Moving  string `json:"moving"`
	Stopped string `json:"stopped"`
	Offline string `json:"offline"`
	Engine  string `json:"engine"`
}

type DriverData struct {
	Id          interface{} `json:"id"`
	UserId      interface{} `json:"user_id"`
	DeviceId    interface{} `json:"device_id"`
	Name        interface{} `json:"name"`
	Rfid        interface{} `json:"rfid"`
	Phone       interface{} `json:"phone"`
	Email       interface{} `json:"email"`
	Description interface{} `json:"description"`
	CreatedAt   interface{} `json:"created_at"`
	UpdatedAt   interface{} `json:"updated_at"`
}

type Sensor struct {
	Id          int         `json:"id"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	ShowInPopup int         `json:"show_in_popup"`
	Value       string      `json:"value"`
	Val         interface{} `json:"val"`
	ScaleValue  *int        `json:"scale_value"`
}

type Location struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Traccar struct {
	Id                 int         `json:"id"`
	Name               string      `json:"name"`
	UniqueId           string      `json:"uniqueId"`
	LatestPositionId   int         `json:"latestPosition_id"`
	LastValidLatitude  float64     `json:"lastValidLatitude"`
	LastValidLongitude float64     `json:"lastValidLongitude"`
	Other              string      `json:"other"`
	Speed              string      `json:"speed"`
	Time               string      `json:"time"`
	DeviceTime         string      `json:"device_time"`
	ServerTime         string      `json:"server_time"`
	AckTime            string      `json:"ack_time"`
	Altitude           float64     `json:"altitude"`
	Course             float64     `json:"course"`
	Power              interface{} `json:"power"`
	Address            interface{} `json:"address"`
	Protocol           string      `json:"protocol"`
	LatestPositions    string      `json:"latest_positions"`
	MovedAt            string      `json:"moved_at"`
	StopedAt           string      `json:"stoped_at"`
	EngineOnAt         string      `json:"engine_on_at"`
	StopBeginAt        string      `json:"stop_begin_at"`
	EngineOffAt        string      `json:"engine_off_at"`
	EngineChangedAt    string      `json:"engine_changed_at"`
	DatabaseId         interface{} `json:"database_id"`
}

type Pivot struct {
	UserId          int         `json:"user_id"`
	DeviceId        int         `json:"device_id"`
	GroupId         int         `json:"group_id"`
	CurrentDriverId interface{} `json:"current_driver_id"`
	Active          int         `json:"active"`
	TimezoneId      interface{} `json:"timezone_id"`
}

type DeviceData struct {
	Id                     int           `json:"id"`
	UserId                 int           `json:"user_id"`
	CurrentDriverId        interface{}   `json:"current_driver_id"`
	TimezoneId             interface{}   `json:"timezone_id"`
	TraccarDeviceId        int           `json:"traccar_device_id"`
	IconId                 int           `json:"icon_id"`
	IconColors             IconColors    `json:"icon_colors"`
	Active                 int           `json:"active"`
	Kind                   int           `json:"kind"`
	Deleted                int           `json:"deleted"`
	Name                   string        `json:"name"`
	Imei                   interface{}   `json:"imei"`
	FuelMeasurementId      int           `json:"fuel_measurement_id"`
	FuelQuantity           string        `json:"fuel_quantity"`
	FuelPrice              string        `json:"fuel_price"`
	FuelPerKm              string        `json:"fuel_per_km"`
	FuelPerH               string        `json:"fuel_per_h"`
	SimNumber              null.String   `json:"sim_number"`
	Msisdn                 interface{}   `json:"msisdn"`
	DeviceModel            string        `json:"device_model"`
	PlateNumber            string        `json:"plate_number"`
	Vin                    string        `json:"vin"`
	RegistrationNumber     string        `json:"registration_number"`
	ObjectOwner            string        `json:"object_owner"`
	AdditionalNotes        string        `json:"additional_notes"`
	Comment                string        `json:"comment"`
	ExpirationDate         interface{}   `json:"expiration_date"`
	SimExpirationDate      interface{}   `json:"sim_expiration_date"`
	SimActivationDate      interface{}   `json:"sim_activation_date"`
	InstallationDate       interface{}   `json:"installation_date"`
	TailColor              string        `json:"tail_color"`
	TailLength             int           `json:"tail_length"`
	EngineHours            string        `json:"engine_hours"`
	DetectEngine           string        `json:"detect_engine"`
	DetectSpeed            string        `json:"detect_speed"`
	DetectDistance         interface{}   `json:"detect_distance"`
	MinMovingSpeed         int           `json:"min_moving_speed"`
	MinFuelFillings        int           `json:"min_fuel_fillings"`
	MinFuelThefts          int           `json:"min_fuel_thefts"`
	SnapToRoad             int           `json:"snap_to_road"`
	GprsTemplatesOnly      bool          `json:"gprs_templates_only"`
	ValidByAvgSpeed        int           `json:"valid_by_avg_speed"`
	Parameters             string        `json:"parameters"`
	Currents               interface{}   `json:"currents"`
	CreatedAt              string        `json:"created_at"`
	UpdatedAt              string        `json:"updated_at"`
	Forward                interface{}   `json:"forward"`
	DeviceTypeId           interface{}   `json:"device_type_id"`
	AppTrackerLogin        int           `json:"app_tracker_login"`
	Users                  []model.User  `json:"users"`
	Pivot                  Pivot         `json:"pivot"`
	Icon                   Icon          `json:"icon"`
	Traccar                Traccar       `json:"traccar"`
	Sensors                []Sensor      `json:"sensors"`
	Services               []interface{} `json:"services"`
	Driver                 interface{}   `json:"driver"`
	LastValidLatitude      float64       `json:"lastValidLatitude"`
	LastValidLongitude     float64       `json:"lastValidLongitude"`
	LatestPositions        string        `json:"latest_positions"`
	IconType               string        `json:"icon_type"`
	GroupId                int           `json:"group_id"`
	UserTimezoneId         interface{}   `json:"user_timezone_id"`
	Time                   string        `json:"time"`
	Course                 float64       `json:"course"`
	Speed                  int           `json:"speed"`
	FuelDetectSecAfterStop null.Int      `json:"fuel_detect_sec_after_stop,omitempty"`
}

type Item struct {
	Id                null.Int      `json:"id"`
	Alarm             int           `json:"alarm"`
	Name              string        `json:"name"`
	Online            string        `json:"online"`
	Time              string        `json:"time"`
	Timestamp         int64         `json:"timestamp"`
	Acktimestamp      int           `json:"acktimestamp"`
	Lat               float64       `json:"lat"`
	Lng               float64       `json:"lng"`
	Course            float64       `json:"course"`
	Speed             float64       `json:"speed"`
	Altitude          float64       `json:"altitude"`
	IconType          string        `json:"icon_type"`
	IconColor         string        `json:"icon_color"`
	IconColors        IconColors    `json:"icon_colors"`
	Icon              Icon          `json:"icon"`
	Power             string        `json:"power"`
	Address           string        `json:"address"`
	Protocol          string        `json:"protocol"`
	Driver            string        `json:"driver"`
	DriverData        DriverData    `json:"driver_data"`
	Sensors           []Sensor      `json:"sensors"`
	Services          []interface{} `json:"services"`
	Tail              []Location    `json:"tail"`
	DistanceUnitHour  string        `json:"distance_unit_hour"`
	UnitOfDistance    string        `json:"unit_of_distance"`
	UnitOfAltitude    string        `json:"unit_of_altitude"`
	UnitOfCapacity    string        `json:"unit_of_capacity"`
	StopDuration      string        `json:"stop_duration"`
	StopDurationSec   int           `json:"stop_duration_sec"`
	MovedTimestamp    int           `json:"moved_timestamp"`
	EngineStatus      null.Bool     `json:"engine_status"`
	DetectEngine      string        `json:"detect_engine"`
	EngineHours       string        `json:"engine_hours"`
	TotalDistance     float64       `json:"total_distance"`
	Inaccuracy        interface{}   `json:"inaccuracy"`
	SimExpirationDate interface{}   `json:"sim_expiration_date"`
	DeviceData        DeviceData    `json:"device_data"`
}

type GroupTelematic struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type DevicesResponse []GroupTelematic

type Device []struct {
	Address      string      `json:"address"`
	DeviceModel  string      `json:"deviceModel"`
	DeviceStatus string      `json:"deviceStatus"`
	ID           int         `json:"id"`
	Imei         interface{} `json:"imei"`
	Lat          float64     `json:"lat"`
	Lng          float64     `json:"lng"`
	Name         string      `json:"name"`
	PlatNumber   string      `json:"platNumber"`
	Protocol     string      `json:"protocol"`
	SimNumber    interface{} `json:"simNumber"`
	Timestamp    int         `json:"timestamp"`
	UserID       int         `json:"userId"`
	Vin          string      `json:"vin"`
}

func (d *DeviceData) UnmarshalJSON(data []byte) error {
	type Alias DeviceData
	aux := &struct {
		*Alias
		GprsTemplatesOnly interface{} `json:"gprs_templates_only"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch aux.GprsTemplatesOnly.(type) {
	case int:
		d.GprsTemplatesOnly = aux.GprsTemplatesOnly == int(1)
	case float64:
		d.GprsTemplatesOnly = aux.GprsTemplatesOnly == float64(1)
	}

	return nil
}

func (d *Sensor) UnmarshalJSON(data []byte) error {
	type Alias Sensor
	contract := &struct {
		*Alias
		Value interface{} `json:"value"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &contract); err != nil {
		return err
	}
	switch contract.Value.(type) {
	case int:
		d.Value = strconv.Itoa(contract.Value.(int))
	}

	return nil
}
