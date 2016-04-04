package message

import (
	"bytes"
	"encoding/json"
)

const ClientTopic string = "client" // Subscription topic for clients
const MasterTopic string = "master" // Subscription topic for masters
const WorkerTopic string = "worker" // Subscription topic for workers

// A directed message.
type Message struct ***REMOVED***
	Topic   string `json:"-"`
	Type    string `json:"type"`
	Fields  Fields `json:"fields,omitempty"`
	Payload []byte `json:"payload,omitempty"`
***REMOVED***

// A set of fields in a message.
type Fields map[string]interface***REMOVED******REMOVED***

// Creates a message directed to the master server.
func NewToMaster(t string, f Fields) Message ***REMOVED***
	return Message***REMOVED***
		Topic:  MasterTopic,
		Type:   t,
		Fields: f,
	***REMOVED***
***REMOVED***

// Creates a message directed to clients.
func NewToClient(t string, f Fields) Message ***REMOVED***
	return Message***REMOVED***
		Topic:  ClientTopic,
		Type:   t,
		Fields: f,
	***REMOVED***
***REMOVED***

// Creates a message directed to workers.
func NewToWorker(t string, f Fields) Message ***REMOVED***
	return Message***REMOVED***
		Topic:  WorkerTopic,
		Type:   t,
		Fields: f,
	***REMOVED***
***REMOVED***

func (msg Message) WithPayload(src interface***REMOVED******REMOVED***) (Message, error) ***REMOVED***
	payload, err := json.Marshal(src)
	if err != nil ***REMOVED***
		return msg, err
	***REMOVED***
	msg.Payload = payload
	return msg, nil
***REMOVED***

func (msg Message) With(src interface***REMOVED******REMOVED***) Message ***REMOVED***
	msg, err := msg.WithPayload(src)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return msg
***REMOVED***

func (msg Message) Take(dst interface***REMOVED******REMOVED***) error ***REMOVED***
	return json.Unmarshal(msg.Payload, dst)
***REMOVED***

func To(topic, t string) Message ***REMOVED***
	return Message***REMOVED***Topic: topic, Type: t***REMOVED***
***REMOVED***

func ToClient(t string) Message ***REMOVED***
	return To(ClientTopic, t)
***REMOVED***

func ToMaster(t string) Message ***REMOVED***
	return To(MasterTopic, t)
***REMOVED***

func ToWorker(t string) Message ***REMOVED***
	return To(WorkerTopic, t)
***REMOVED***

// Decodes a message from wire format.
func Decode(data []byte, msg *Message) (err error) ***REMOVED***
	nullIndex := bytes.IndexByte(data, 0)
	msg.Topic = string(data[:nullIndex])
	return json.Unmarshal(data[nullIndex+1:], msg)
***REMOVED***

// Encodes a message to wire format.
func (msg *Message) Encode() ([]byte, error) ***REMOVED***
	jdata, err := json.Marshal(msg)
	if err != nil ***REMOVED***
		return jdata, err
	***REMOVED***

	buf := bytes.NewBufferString(msg.Topic)
	buf.WriteByte(0)
	buf.Write(jdata)
	return buf.Bytes(), nil
***REMOVED***
