package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id  	int		 `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

// p Plant 为接受器，给Plant结构体添加了一个 String 方法
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main()  {
	coffee := &Plant{Id:27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))


	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil{
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "Californian"}

	type Nesting struct {
		XMLName xml.Name 	`xml."nesting"`
		Plants  []*Plant	`xml:"parent>child>plant"`
	}

	netsting := &Nesting{}
	netsting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(netsting, " ", "  ")
	fmt.Println(string(out))
}
