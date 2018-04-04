package main

import "encoding/json"

type SubData struct {
	Name    string `json:"name,omitempty" enable:"true" hidden:"false" usage:"name of data"`
	Value   int    `json:"value,omitempty" enable:"true" hidden:"false" usage:"value of data"`
	Present bool   `json:"present,omitempty" enable:"true" hidden:"false" usage:"present in data"`
}

func (s *SubData) Init() *SubData {
	return s
}

type Data struct {
	MyName  string   `json:"my-name,omitempty" enable:"true" hidden:"false" usage:"my name" description:"my name"`
	MyValue int      `json:"my-value,omitempty" enable:"true" hidden:"false" usage:"my value" description:"my value"`
	A       *SubData `json:"a,omitempty" enable:"true" hidden:"false" usage:"command a" description:"command a"`
	B       *SubData `json:"b,omitempty" enable:"true" hidden:"false" usage:"command b" description:"command b"`
}

func (s *Data) Init() *Data {
	s.A = new(SubData).Init()
	s.B = new(SubData).Init()
	return s
}

func (s *Data) Marshal() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (s *Data) Unmarshal(b []byte) error {
	return json.Unmarshal(b, s)
}
