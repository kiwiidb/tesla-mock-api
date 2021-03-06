package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var data = `{
	"response": {
	  "id": 64951216767793660,
	  "user_id": 676781,
	  "vehicle_id": 1552555543,
	  "vin": "5YJ3E7EB2KF240544",
	  "display_name": "Muinkpark model 3",
	  "option_codes": "AD15,MDL3,PBSB,RENA,BT37,ID3W,RF3G,S3PB,DRLH,DV2W,W39B,APF0,COUS,BC3B,CH07,PC30,FC3P,FG31,GLFR,HL31,HM31,IL31,LTPB,MR31,FM3B,RS3H,SA3P,STCP,SC04,SU3C,T3CA,TW00,TM00,UT3P,WR00,AU3P,APH3,AF00,ZCST,MI00,CDM0",
	  "color": null,
	  "tokens": [
		"3964c6fb9141c541",
		"9083a16adcaa7b41"
	  ],
	  "state": "online",
	  "in_service": false,
	  "id_s": "64951216767793667",
	  "calendar_enabled": true,
	  "api_version": 7,
	  "backseat_token": null,
	  "backseat_token_updated_at": null,
	  "charge_state": {
		"battery_heater_on": false,
		"battery_level": 61,
		"battery_range": 176.59,
		"charge_current_request": 16,
		"charge_current_request_max": 16,
		"charge_enable_request": true,
		"charge_energy_added": 0,
		"charge_limit_soc": 79,
		"charge_limit_soc_max": 100,
		"charge_limit_soc_min": 50,
		"charge_limit_soc_std": 90,
		"charge_miles_added_ideal": 0,
		"charge_miles_added_rated": 0,
		"charge_port_cold_weather_mode": false,
		"charge_port_door_open": false,
		"charge_port_latch": "Engaged",
		"charge_rate": 0,
		"charge_to_max_range": false,
		"charger_actual_current": 0,
		"charger_phases": null,
		"charger_pilot_current": 16,
		"charger_power": 0,
		"charger_voltage": 2,
		"charging_state": "Disconnected",
		"conn_charge_cable": "<invalid>",
		"est_battery_range": 197.73,
		"fast_charger_brand": "<invalid>",
		"fast_charger_present": false,
		"fast_charger_type": "<invalid>",
		"ideal_battery_range": 176.59,
		"managed_charging_active": false,
		"managed_charging_start_time": null,
		"managed_charging_user_canceled": false,
		"max_range_charge_counter": 0,
		"minutes_to_full_charge": 0,
		"not_enough_power_to_heat": null,
		"scheduled_charging_pending": false,
		"scheduled_charging_start_time": null,
		"time_to_full_charge": 0,
		"timestamp": 1588501785238,
		"trip_charging": false,
		"usable_battery_level": 60,
		"user_charge_enable_request": null
	  },
	  "climate_state": {
		"battery_heater": false,
		"battery_heater_no_power": null,
		"climate_keeper_mode": "off",
		"defrost_mode": 0,
		"driver_temp_setting": 20,
		"fan_status": 0,
		"inside_temp": 21.5,
		"is_auto_conditioning_on": false,
		"is_climate_on": false,
		"is_front_defroster_on": false,
		"is_preconditioning": false,
		"is_rear_defroster_on": false,
		"left_temp_direction": 0,
		"max_avail_temp": 28,
		"min_avail_temp": 15,
		"outside_temp": 15.5,
		"passenger_temp_setting": 20,
		"remote_heater_control_enabled": false,
		"right_temp_direction": 0,
		"seat_heater_left": 0,
		"seat_heater_rear_center": 0,
		"seat_heater_rear_left": 0,
		"seat_heater_rear_right": 0,
		"seat_heater_right": 0,
		"side_mirror_heaters": false,
		"timestamp": 1588501785250,
		"wiper_blade_heater": false
	  },
	  "drive_state": {
		"gps_as_of": 1588497461,
		"heading": 120,
		"latitude": 51.058265,
		"longitude": 3.726171,
		"native_latitude": 51.058265,
		"native_location_supported": 1,
		"native_longitude": 3.726171,
		"native_type": "wgs",
		"power": 0,
		"shift_state": null,
		"speed": null,
		"timestamp": 1588501785257
	  },
	  "vehicle_config": {
		"can_accept_navigation_requests": true,
		"can_actuate_trunks": true,
		"car_special_type": "base",
		"car_type": "model3",
		"charge_port_type": "CCS",
		"ece_restrictions": true,
		"eu_vehicle": true,
		"exterior_color": "DeepBlue",
		"has_air_suspension": false,
		"has_ludicrous_mode": false,
		"key_version": 2,
		"motorized_charge_port": true,
		"plg": false,
		"rear_seat_heaters": 1,
		"rear_seat_type": null,
		"rhd": false,
		"roof_color": "Glass",
		"seat_type": null,
		"spoiler_type": "None",
		"sun_roof_installed": null,
		"third_row_seats": "<invalid>",
		"timestamp": 1588501785275,
		"use_range_badging": true,
		"wheel_type": "Stiletto19"
	  },
	  "gui_settings": {
		"gui_24_hour_time": true,
		"gui_charge_rate_units": "km/hr",
		"gui_distance_units": "km/hr",
		"gui_range_display": "Rated",
		"gui_temperature_units": "C",
		"show_range_units": false,
		"timestamp": 1588501785284
	  },
	  "vehicle_state": {
		"api_version": 7,
		"autopark_state_v2": "unavailable",
		"calendar_supported": true,
		"car_version": "2020.4.1 4a4ad401858f",
		"center_display_state": 0,
		"df": 0,
		"dr": 0,
		"fd_window": 0,
		"fp_window": 0,
		"ft": 0,
		"is_user_present": false,
		"locked": true,
		"media_state": {
		  "remote_control_enabled": true
		},
		"notifications_supported": true,
		"odometer": 33441.900896,
		"parsed_calendar_supported": true,
		"pf": 0,
		"pr": 0,
		"rd_window": 0,
		"remote_start": false,
		"remote_start_enabled": true,
		"remote_start_supported": true,
		"rp_window": 0,
		"rt": 0,
		"sentry_mode": false,
		"sentry_mode_available": true,
		"software_update": {
		  "download_perc": 0,
		  "expected_duration_sec": 2700,
		  "install_perc": 1,
		  "status": "downloading_wifi_wait",
		  "version": "2020.8.3"
		},
		"speed_limit_mode": {
		  "active": false,
		  "current_limit_mph": 79.407659,
		  "max_limit_mph": 90,
		  "min_limit_mph": 50,
		  "pin_code_set": false
		},
		"timestamp": 1588501785281,
		"valet_mode": false,
		"valet_pin_needed": true,
		"vehicle_name": "Muinkpark model 3"
	  }
	}
  }
`

func Handler(w http.ResponseWriter, r *http.Request) {
	if !checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, data)
	w.Header().Set("Content-Type", "application/json")
}

func checkAuth(r *http.Request) bool {
	header := fmt.Sprintf("Bearer %s", accessToken)
	logrus.Info(header)
	return r.Header.Get("Authorization") == header
}
