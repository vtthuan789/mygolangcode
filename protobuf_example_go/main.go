package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	complexpb "protobuf_example_go/src/complex"
	enumpb "protobuf_example_go/src/enum_example"
	simplepb "protobuf_example_go/src/simple"
)

func main() {
	sm := doSimple()
	readWriteToFileDemo(sm)
	convertFromToJSONDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First dummy message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second dummy message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third dummy message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           20,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}
	fmt.Println(em)
	em.DayOfTheWeek = enumpb.DayOfTheWeek_SATURDAY
	fmt.Println(em)
}

func readWriteToFileDemo(pb proto.Message) {
	err := writeToFile("simple.bin", pb)
	if err != nil {
		log.Fatal(err)
	}
	sm := &simplepb.SimpleMessage{}
	err = readFromFile("simple.bin", sm)
	fmt.Println(sm)
}

func convertFromToJSONDemo(pb proto.Message) {
	smAsString := toJSON(pb)
	fmt.Println(smAsString)

	sm := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm)
	fmt.Println(sm)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatal("Can't convert to JSON!", err)
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatal("Can't unmarshal JSON to the pb struct!", err)
	}
}

func writeToFile(fname string, pb proto.Message) error {
	data, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialize to bytes! Error:", err)
		return err
	}

	err = ioutil.WriteFile(fname, data, 0666)

	if err != nil {
		log.Fatalln("Can't write to file! Error:", err)
	}

	fmt.Println("Data has been written!")

	return err
}

func readFromFile(fname string, pb proto.Message) error {
	data, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Read failed! Error:", err)
		return err
	}

	err = proto.Unmarshal(data, pb)

	if err != nil {
		log.Fatalln("Can't deserialize to pb! Error:", err)
		return err
	}

	return err
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "Not simple anymore"
	fmt.Println(sm)

	fmt.Println("The id is:", sm.GetId())

	return &sm
}
