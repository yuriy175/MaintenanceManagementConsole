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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetDeviceConnections()
		for _, device := range devices {
			fmt.Fprintf(w, "Number : %d name %s type %s time %s\n",
				device.DeviceId, device.Name, device.Type, device.DateTime.Format(time.RFC3339))
		}
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
