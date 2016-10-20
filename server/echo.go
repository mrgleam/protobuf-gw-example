package main

import (
	"fmt"

	"github.com/golang/glog"
	examples "github.com/mrgleam/protobuf-gw-example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implements of EchoServiceServer

type echoServer struct{}

// func readHttpBody(response *http.Response) string {
//
// 	fmt.Println("Reading body")
//
// 	bodyBuffer := make([]byte, 5000)
// 	var str string
//
// 	count, err := response.Body.Read(bodyBuffer)
//
// 	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
//
// 		if err != nil {
//
// 		}
//
// 		str += string(bodyBuffer[:count])
// 	}
//
// 	return str
//
// }
//
// //Converts a code to an Auth_Token
// func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
// 	fmt.Println("GetAccessToken")
// 	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
// 	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
// 		client_id + "&redirect_uri=" + callbackUri +
// 		"&client_secret=" + secret + "&code=" + code)
//
// 	if err == nil {
//
// 		auth := readHttpBody(response)
//
// 		var token AccessToken
//
// 		tokenArr := strings.Split(auth, "&")
//
// 		token.Token = strings.Split(tokenArr[0], "=")[1]
// 		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])
//
// 		if err == nil {
// 			token.Expiry = int64(expireInt)
// 		}
//
// 		return token
// 	}
//
// 	var token AccessToken
//
// 	return token
// }

func newEchoServer() examples.YourServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *examples.StringMessage) (*examples.StringMessage, error) {
	fmt.Println(msg)
	// // generate loginURL
	// fbConfig := &oauth2.Config{
	// 	// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
	// 	// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"
	//
	// 	ClientID:     "", // change this to yours
	// 	ClientSecret: "",
	// 	RedirectURL:  "http://localhost:8080/v1/example/echo", // change this to your webserver adddress
	// 	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	// 	Endpoint: oauth2.Endpoint{
	// 		AuthURL:  "https://www.facebook.com/dialog/oauth",
	// 		TokenURL: "https://graph.facebook.com/oauth/access_token",
	// 	},
	// }
	//
	// url := fbConfig.AuthCodeURL("")
	// response, _ := http.Get(url)
	// fmt.Println(url)
	// fmt.Println(response)
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *examples.StringMessage) (*examples.StringMessage, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}
