package handler

import (
	"fmt"
	"net/http"
)

var wakeResp = `
{
"response":
	{
	"api_version": 7,
	"backseat_token": null,
	"backseat_token_updated_at": null,
	"calendar_enabled": true,
	"color": null,
	"display_name": "Muinkpark model 3",
	"id": 1639654603719884,
	"id_s": "1639654603719884",
	"in_service": false,
	"option_codes": "AD15,MDL3,PBSB,RENA,BT37,ID3W,RF3G,S3PB,DRLH,DV2W,W39B,APF0,COUS,BC3B,CH07,PC30,FC3P,FG31,GLFR,HL31,HM31,IL31,LTPB,MR31,FM3B,RS3H,SA3P,STCP,SC04,SU3C,T3CA,TW00,TM00,UT3P,WR00,AU3P,APH3,AF00,ZCST,MI00,CDM0",
	"state": "online",
	"tokens": [
		"625cf7f19da1f780",
		"a0cbd7524a472de2"
	],
	"vehicle_id": 1552555543,
	"vin": "5YJ3E7EB2KF240544"
	}
}
`

func WakeHandler(w http.ResponseWriter, r *http.Request) {
	if !checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, wakeResp)
	w.Header().Set("Content-Type", "application/json")
}
