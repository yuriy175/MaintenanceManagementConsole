package BL

import (
	"fmt"
	"net/http"

	"../DAL"
)

func HttpServer() {

	/*http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
	    fmt.Fprint(w, "Contact Page")
	})*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		devices := DAL.DalGetDeviceConnections()
		for _, device := range devices {
			fmt.Fprintln(w, "Number : %d name %s type %s", device.DeviceId, device.Name, device.Type)
		}
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
