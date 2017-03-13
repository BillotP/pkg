// Autogenerated by Thrift Compiler (1.0.0-upfluence)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package transport

import (
	"bytes"
	"fmt"
	"github.com/upfluence/goutils/Godeps/_workspace/src/github.com/upfluence/base/service/thrift/transport/http"
	"github.com/upfluence/goutils/Godeps/_workspace/src/github.com/upfluence/base/service/thrift/transport/rabbitmq"
	"github.com/upfluence/goutils/Godeps/_workspace/src/github.com/upfluence/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal
var _ = time.Now()

var _ = http.GoUnusedProtection__
var _ = rabbitmq.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - HttpTransport
//  - RabbitmqTransport
type Transport struct {
	HttpTransport     *http.Transport     `thrift:"http_transport,1" json:"http_transport,omitempty"`
	RabbitmqTransport *rabbitmq.Transport `thrift:"rabbitmq_transport,2" json:"rabbitmq_transport,omitempty"`
}

func NewTransport() *Transport {
	return &Transport{}
}

var Transport_HttpTransport_DEFAULT *http.Transport

func (p *Transport) GetHttpTransport() *http.Transport {
	if !p.IsSetHttpTransport() {
		return Transport_HttpTransport_DEFAULT
	}
	return p.HttpTransport
}

var Transport_RabbitmqTransport_DEFAULT *rabbitmq.Transport

func (p *Transport) GetRabbitmqTransport() *rabbitmq.Transport {
	if !p.IsSetRabbitmqTransport() {
		return Transport_RabbitmqTransport_DEFAULT
	}
	return p.RabbitmqTransport
}
func (p *Transport) CountSetFieldsTransport() int {
	count := 0
	if p.IsSetHttpTransport() {
		count++
	}
	if p.IsSetRabbitmqTransport() {
		count++
	}
	return count

}

func (p *Transport) IsSetHttpTransport() bool {
	return p.HttpTransport != nil
}

func (p *Transport) IsSetRabbitmqTransport() bool {
	return p.RabbitmqTransport != nil
}

func (p *Transport) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Transport) ReadField1(iprot thrift.TProtocol) error {
	p.HttpTransport = http.NewTransport()
	if err := p.HttpTransport.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.HttpTransport), err)
	}
	return nil
}

func (p *Transport) ReadField2(iprot thrift.TProtocol) error {
	p.RabbitmqTransport = rabbitmq.NewTransport()
	if err := p.RabbitmqTransport.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RabbitmqTransport), err)
	}
	return nil
}

func (p *Transport) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsTransport(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("Transport"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Transport) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetHttpTransport() {
		if err := oprot.WriteFieldBegin("http_transport", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:http_transport: ", p), err)
		}
		if err := p.HttpTransport.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.HttpTransport), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:http_transport: ", p), err)
		}
	}
	return err
}

func (p *Transport) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetRabbitmqTransport() {
		if err := oprot.WriteFieldBegin("rabbitmq_transport", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:rabbitmq_transport: ", p), err)
		}
		if err := p.RabbitmqTransport.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RabbitmqTransport), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:rabbitmq_transport: ", p), err)
		}
	}
	return err
}

func (p *Transport) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Transport(%+v)", *p)
}
