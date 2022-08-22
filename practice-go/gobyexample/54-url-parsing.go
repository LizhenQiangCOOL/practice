package main

import (
	"fmt"
	"net"
	"net/url"
)

func main()  {

	s := "postgres://usr:pass@host.com:5432/path?k=v&k1=v1#f"

	u, err := url.Parse(s)
	if err != nil{
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	//#号之后的内容，锚
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}