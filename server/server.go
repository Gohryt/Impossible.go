package server

import (
	"crypto/tls"
	"net/http"
	"time"

	gjson "github.com/Gohryt/Impossible.go/json"
)

type (
	//Configuration is struct wich contains server configuration
	Configuration struct {
		Domain    string `json:"Domain"`
		Port      int    `json:"Port"`
		PortTLS   int    `json:"PortTLS"`
		FolderTLS string `json:"FolderTLS"`
	}
)

var (
	//ConfigTLS is predefined http tls config
	ConfigTLS = tls.Config{MinVersion: tls.VersionTLS12, CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}, PreferServerCipherSuites: true, CipherSuites: []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA, tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, tls.TLS_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_RSA_WITH_RC4_128_SHA, tls.TLS_RSA_WITH_AES_128_CBC_SHA, tls.TLS_RSA_WITH_AES_256_CBC_SHA, tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA}}
	//Default is predefined http server
	Default = http.Server{Handler: nil, IdleTimeout: 120 * time.Second, MaxHeaderBytes: 16 * 1024, ReadHeaderTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
	//DefaultTLS is predefined http tls server
	DefaultTLS = http.Server{TLSConfig: &ConfigTLS, Handler: nil, IdleTimeout: 120 * time.Second, MaxHeaderBytes: 16 * 1024, ReadHeaderTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}
)

//Listen function starts server listeners
func Listen(server, serverTLS *http.Server, TLSFolderAddress string, errorHandler func(*error)) {
	go listenTLS(serverTLS, TLSFolderAddress, errorHandler)
	go listenForTLSRedirect(server, errorHandler)
	return
}

func listenTLS(serverTLS *http.Server, TLSFolderAddress string, errorHandler func(*error)) {
	var (
		err error
	)
	err = serverTLS.ListenAndServeTLS(TLSFolderAddress+"certificate.pem", TLSFolderAddress+"key.pem")
	errorHandler(&err)
	return
}

func listenForTLSRedirect(server *http.Server, errorHandler func(*error)) {
	var (
		err error
	)
	err = server.ListenAndServe()
	errorHandler(&err)
	return
}

//FromMemory function reads json from byte array
func (c *Configuration) FromMemory(data []byte, errorHandler func(*error)) {
	gjson.FromMemory(data, c, errorHandler)
	return
}

//FromFile function reads json from file
func (c *Configuration) FromFile(filePath string, errorHandler func(*error)) {
	gjson.FromFile(filePath, c, errorHandler)
	return
}

//ToMemory function returns json as byte array
func (c *Configuration) ToMemory(errorHandler func(*error)) (data []byte) {
	data = gjson.ToMemory(*c, errorHandler)
	return
}

//ToFile function saves json to file
func (c *Configuration) ToFile(filePath string, errorHandler func(*error)) {
	gjson.ToFile(*c, filePath, errorHandler)
	return
}
