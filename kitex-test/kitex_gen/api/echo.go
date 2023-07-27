// Code generated by thriftgo (0.2.11). DO NOT EDIT.

package api

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type EchoReq struct {
	Message string `thrift:"message,1,required" frugal:"1,required,string" json:"message"`
}

func NewEchoReq() *EchoReq {
	return &EchoReq{}
}

func (p *EchoReq) InitDefault() {
	*p = EchoReq{}
}

func (p *EchoReq) GetMessage() (v string) {
	return p.Message
}
func (p *EchoReq) SetMessage(val string) {
	p.Message = val
}

var fieldIDToName_EchoReq = map[int16]string{
	1: "message",
}

func (p *EchoReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetMessage bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetMessage = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetMessage {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EchoReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_EchoReq[fieldId]))
}

func (p *EchoReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *EchoReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EchoReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *EchoReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *EchoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoReq(%+v)", *p)
}

func (p *EchoReq) DeepEqual(ano *EchoReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Message) {
		return false
	}
	return true
}

func (p *EchoReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Message, src) != 0 {
		return false
	}
	return true
}

type EchoResp struct {
	Response string `thrift:"response,1" frugal:"1,default,string" json:"response"`
}

func NewEchoResp() *EchoResp {
	return &EchoResp{}
}

func (p *EchoResp) InitDefault() {
	*p = EchoResp{}
}

func (p *EchoResp) GetResponse() (v string) {
	return p.Response
}
func (p *EchoResp) SetResponse(val string) {
	p.Response = val
}

var fieldIDToName_EchoResp = map[int16]string{
	1: "response",
}

func (p *EchoResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EchoResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EchoResp) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Response = v
	}
	return nil
}

func (p *EchoResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EchoResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *EchoResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("response", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Response); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *EchoResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResp(%+v)", *p)
}

func (p *EchoResp) DeepEqual(ano *EchoResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Response) {
		return false
	}
	return true
}

func (p *EchoResp) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Response, src) != 0 {
		return false
	}
	return true
}

type EchoService interface {
	Echo(ctx context.Context, request *EchoReq) (r *EchoResp, err error)
}

type EchoServiceClient struct {
	c thrift.TClient
}

func NewEchoServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *EchoServiceClient {
	return &EchoServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewEchoServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *EchoServiceClient {
	return &EchoServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewEchoServiceClient(c thrift.TClient) *EchoServiceClient {
	return &EchoServiceClient{
		c: c,
	}
}

func (p *EchoServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *EchoServiceClient) Echo(ctx context.Context, request *EchoReq) (r *EchoResp, err error) {
	var _args EchoServiceEchoArgs
	_args.Request = request
	var _result EchoServiceEchoResult
	if err = p.Client_().Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type EchoServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      EchoService
}

func (p *EchoServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *EchoServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *EchoServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewEchoServiceProcessor(handler EchoService) *EchoServiceProcessor {
	self := &EchoServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("echo", &echoServiceProcessorEcho{handler: handler})
	return self
}
func (p *EchoServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type echoServiceProcessorEcho struct {
	handler EchoService
}

func (p *echoServiceProcessorEcho) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EchoServiceEchoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := EchoServiceEchoResult{}
	var retval *EchoResp
	if retval, err2 = p.handler.Echo(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing echo: "+err2.Error())
		oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("echo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type EchoServiceEchoArgs struct {
	Request *EchoReq `thrift:"request,1" frugal:"1,default,EchoReq" json:"request"`
}

func NewEchoServiceEchoArgs() *EchoServiceEchoArgs {
	return &EchoServiceEchoArgs{}
}

func (p *EchoServiceEchoArgs) InitDefault() {
	*p = EchoServiceEchoArgs{}
}

var EchoServiceEchoArgs_Request_DEFAULT *EchoReq

func (p *EchoServiceEchoArgs) GetRequest() (v *EchoReq) {
	if !p.IsSetRequest() {
		return EchoServiceEchoArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *EchoServiceEchoArgs) SetRequest(val *EchoReq) {
	p.Request = val
}

var fieldIDToName_EchoServiceEchoArgs = map[int16]string{
	1: "request",
}

func (p *EchoServiceEchoArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *EchoServiceEchoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EchoServiceEchoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EchoServiceEchoArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewEchoReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *EchoServiceEchoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("echo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *EchoServiceEchoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *EchoServiceEchoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoArgs(%+v)", *p)
}

func (p *EchoServiceEchoArgs) DeepEqual(ano *EchoServiceEchoArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *EchoServiceEchoArgs) Field1DeepEqual(src *EchoReq) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type EchoServiceEchoResult struct {
	Success *EchoResp `thrift:"success,0,optional" frugal:"0,optional,EchoResp" json:"success,omitempty"`
}

func NewEchoServiceEchoResult() *EchoServiceEchoResult {
	return &EchoServiceEchoResult{}
}

func (p *EchoServiceEchoResult) InitDefault() {
	*p = EchoServiceEchoResult{}
}

var EchoServiceEchoResult_Success_DEFAULT *EchoResp

func (p *EchoServiceEchoResult) GetSuccess() (v *EchoResp) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResp)
}

var fieldIDToName_EchoServiceEchoResult = map[int16]string{
	0: "success",
}

func (p *EchoServiceEchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EchoServiceEchoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EchoServiceEchoResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewEchoResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *EchoServiceEchoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("echo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *EchoServiceEchoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *EchoServiceEchoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoResult(%+v)", *p)
}

func (p *EchoServiceEchoResult) DeepEqual(ano *EchoServiceEchoResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *EchoServiceEchoResult) Field0DeepEqual(src *EchoResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}