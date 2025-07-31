package routes

import (
	"qr-code-generator/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterQRCodeGeneratorFirebasestoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/firebase/qrcode/generate", controllers.QrcodeGenerateFirebase).Methods("POST")
	router.HandleFunc("/firebase/qrcode/{token}", controllers.ValidateQRCodeFirebase).Methods("GET")

}
