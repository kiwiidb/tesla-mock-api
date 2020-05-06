package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var vehicles = `{
    "count": 4,
    "response": [
        {
            "api_version": 7,
            "backseat_token": null,
            "backseat_token_updated_at": null,
            "calendar_enabled": true,
            "color": null,
            "display_name": "Ecomobiliteit Gent",
            "id": 5877462114549086,
            "id_s": "5877462114549086",
            "in_service": false,
            "option_codes": "AD15,MDL3,PBSB,RENA,BT37,ID3W,RF3G,S3PB,DRLH,DV2W,W39B,APF0,COUS,BC3B,CH07,PC30,FC3P,FG31,GLFR,HL31,HM31,IL31,LTPB,MR31,FM3B,RS3H,SA3P,STCP,SC04,SU3C,T3CA,TW00,TM00,UT3P,WR00,AU3P,APH3,AF00,ZCST,MI00,CDM0",
            "state": "online",
            "tokens": [
                "b77c6199315edeb9",
                "bcc1405ccd2144ae"
            ],
            "vehicle_id": 1198402930,
            "vin": "5YJXCCE21GF021890"
        },
        {
            "api_version": 7,
            "backseat_token": null,
            "backseat_token_updated_at": null,
            "calendar_enabled": true,
            "color": null,
            "display_name": "Erwin's model 3",
            "id": 69100559710657366,
            "id_s": "69100559710657366",
            "in_service": false,
            "option_codes": "AD15,MDL3,PBSB,RENA,BT37,ID3W,RF3G,S3PB,DRLH,DV2W,W39B,APF0,COUS,BC3B,CH07,PC30,FC3P,FG31,GLFR,HL31,HM31,IL31,LTPB,MR31,FM3B,RS3H,SA3P,STCP,SC04,SU3C,T3CA,TW00,TM00,UT3P,WR00,AU3P,APH3,AF00,ZCST,MI00,CDM0",
            "state": "online",
            "tokens": [
                "ca2f8bd05584ca2e",
                "79b2fe23d1edf28b"
            ],
            "vehicle_id": 1760563915,
            "vin": "5YJ3E7EC8LF652820"
        },
        {
            "api_version": 7,
            "backseat_token": null,
            "backseat_token_updated_at": null,
            "calendar_enabled": true,
            "color": null,
            "display_name": "Model S RED",
            "id": 2551490200514290,
            "id_s": "2551490200514290",
            "in_service": false,
            "option_codes": "AD15,MDL3,PBSB,RENA,BT37,ID3W,RF3G,S3PB,DRLH,DV2W,W39B,APF0,COUS,BC3B,CH07,PC30,FC3P,FG31,GLFR,HL31,HM31,IL31,LTPB,MR31,FM3B,RS3H,SA3P,STCP,SC04,SU3C,T3CA,TW00,TM00,UT3P,WR00,AU3P,APH3,AF00,ZCST,MI00,CDM0",
            "state": "online",
            "tokens": [
                "d40916e28d1fbbc4",
                "ecf3b4689a0a8d36"
            ],
            "vehicle_id": 1446007059,
            "vin": "5YJSA7E23KF302671"
        },
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
                "5564053bb6d505cc",
                "7bd224aee6ff37fc"
            ],
            "vehicle_id": 1552555543,
            "vin": "5YJ3E7EB2KF240544"
        }
    ]
}

`

func VehiclesHandler(w http.ResponseWriter, r *http.Request) {
	if !checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, vehicles)
	w.Header().Set("Content-Type", "application/json")
}

func checkAuth(r *http.Request) bool {
	header := fmt.Sprintf("Bearer %s", accessToken)
	logrus.Info(header)
	return r.Header.Get("Authorization") == header
}
