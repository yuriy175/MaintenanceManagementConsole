package BL

import (
	"fmt"
	"net/http"
	"time"

	"../DAL"
)

func HttpServer() {

	/*http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
	    fmt.Fprint(w, "Contact Page")
	})*/
	http.HandleFunc("/devices/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetDeviceConnections()
		for _, device := range devices {
			fmt.Fprintf(w, "time %s device : %d name : %s type : %s connection : %d\n",
				device.DateTime.Format(time.RFC3339), device.DeviceId, device.DeviceName, device.DeviceType, device.DeviceConnection)
		}

		fmt.Fprint(w, "Index Page")
	})

	http.HandleFunc("/studies/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetStudiesInWork()
		for _, study := range devices {
			fmt.Fprintf(w, "time %s study : %d dicom : %s name : %s\n",
				study.DateTime.Format(time.RFC3339), study.StudyId, study.StudyDicomUid, study.StudyName)
		}

		fmt.Fprint(w, "Index Page")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
