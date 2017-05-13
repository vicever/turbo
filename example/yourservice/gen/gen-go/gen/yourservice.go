// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package gen

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Message
type SayHelloResponse struct {
  Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewSayHelloResponse() *SayHelloResponse {
  return &SayHelloResponse{}
}


func (p *SayHelloResponse) GetMessage() string {
  return p.Message
}
func (p *SayHelloResponse) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
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

func (p *SayHelloResponse)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Message = v
}
  return nil
}

func (p *SayHelloResponse) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SayHelloResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SayHelloResponse) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err) }
  if err := oprot.WriteString(string(p.Message)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err) }
  return err
}

func (p *SayHelloResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SayHelloResponse(%+v)", *p)
}

// Attributes:
//  - Message
type EatAppleResponse struct {
  Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewEatAppleResponse() *EatAppleResponse {
  return &EatAppleResponse{}
}


func (p *EatAppleResponse) GetMessage() string {
  return p.Message
}
func (p *EatAppleResponse) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
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

func (p *EatAppleResponse)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Message = v
}
  return nil
}

func (p *EatAppleResponse) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("EatAppleResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EatAppleResponse) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err) }
  if err := oprot.WriteString(string(p.Message)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err) }
  return err
}

func (p *EatAppleResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EatAppleResponse(%+v)", *p)
}

type YourService interface {
  // Parameters:
  //  - YourName
  //  - Values
  //  - HelloValues
  SayHello(yourName string, values *CommonValues, helloValues *HelloValues) (r *SayHelloResponse, err error)
  // Parameters:
  //  - Num
  //  - StringValue
  //  - BoolValue
  EatApple(num int32, stringValue string, boolValue bool) (r *EatAppleResponse, err error)
}

type YourServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewYourServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *YourServiceClient {
  return &YourServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewYourServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *YourServiceClient {
  return &YourServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - YourName
//  - Values
//  - HelloValues
func (p *YourServiceClient) SayHello(yourName string, values *CommonValues, helloValues *HelloValues) (r *SayHelloResponse, err error) {
  if err = p.sendSayHello(yourName, values, helloValues); err != nil { return }
  return p.recvSayHello()
}

func (p *YourServiceClient) sendSayHello(yourName string, values *CommonValues, helloValues *HelloValues)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("sayHello", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := YourServiceSayHelloArgs{
  YourName : yourName,
  Values : values,
  HelloValues : helloValues,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *YourServiceClient) recvSayHello() (value *SayHelloResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "sayHello" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "sayHello failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "sayHello failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "sayHello failed: invalid message type")
    return
  }
  result := YourServiceSayHelloResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Num
//  - StringValue
//  - BoolValue
func (p *YourServiceClient) EatApple(num int32, stringValue string, boolValue bool) (r *EatAppleResponse, err error) {
  if err = p.sendEatApple(num, stringValue, boolValue); err != nil { return }
  return p.recvEatApple()
}

func (p *YourServiceClient) sendEatApple(num int32, stringValue string, boolValue bool)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("eatApple", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := YourServiceEatAppleArgs{
  Num : num,
  StringValue : stringValue,
  BoolValue : boolValue,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *YourServiceClient) recvEatApple() (value *EatAppleResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "eatApple" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "eatApple failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "eatApple failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "eatApple failed: invalid message type")
    return
  }
  result := YourServiceEatAppleResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type YourServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler YourService
}

func (p *YourServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *YourServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *YourServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewYourServiceProcessor(handler YourService) *YourServiceProcessor {

  self4 := &YourServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["sayHello"] = &yourServiceProcessorSayHello{handler:handler}
  self4.processorMap["eatApple"] = &yourServiceProcessorEatApple{handler:handler}
return self4
}

func (p *YourServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x5

}

type yourServiceProcessorSayHello struct {
  handler YourService
}

func (p *yourServiceProcessorSayHello) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := YourServiceSayHelloArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("sayHello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := YourServiceSayHelloResult{}
var retval *SayHelloResponse
  var err2 error
  if retval, err2 = p.handler.SayHello(args.YourName, args.Values, args.HelloValues); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing sayHello: " + err2.Error())
    oprot.WriteMessageBegin("sayHello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("sayHello", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type yourServiceProcessorEatApple struct {
  handler YourService
}

func (p *yourServiceProcessorEatApple) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := YourServiceEatAppleArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("eatApple", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := YourServiceEatAppleResult{}
var retval *EatAppleResponse
  var err2 error
  if retval, err2 = p.handler.EatApple(args.Num, args.StringValue, args.BoolValue); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing eatApple: " + err2.Error())
    oprot.WriteMessageBegin("eatApple", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("eatApple", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - YourName
//  - Values
//  - HelloValues
type YourServiceSayHelloArgs struct {
  YourName string `thrift:"yourName,1" db:"yourName" json:"yourName"`
  Values *CommonValues `thrift:"values,2" db:"values" json:"values"`
  HelloValues *HelloValues `thrift:"helloValues,3" db:"helloValues" json:"helloValues"`
}

func NewYourServiceSayHelloArgs() *YourServiceSayHelloArgs {
  return &YourServiceSayHelloArgs{}
}


func (p *YourServiceSayHelloArgs) GetYourName() string {
  return p.YourName
}
var YourServiceSayHelloArgs_Values_DEFAULT *CommonValues
func (p *YourServiceSayHelloArgs) GetValues() *CommonValues {
  if !p.IsSetValues() {
    return YourServiceSayHelloArgs_Values_DEFAULT
  }
return p.Values
}
var YourServiceSayHelloArgs_HelloValues_DEFAULT *HelloValues
func (p *YourServiceSayHelloArgs) GetHelloValues() *HelloValues {
  if !p.IsSetHelloValues() {
    return YourServiceSayHelloArgs_HelloValues_DEFAULT
  }
return p.HelloValues
}
func (p *YourServiceSayHelloArgs) IsSetValues() bool {
  return p.Values != nil
}

func (p *YourServiceSayHelloArgs) IsSetHelloValues() bool {
  return p.HelloValues != nil
}

func (p *YourServiceSayHelloArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
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

func (p *YourServiceSayHelloArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.YourName = v
}
  return nil
}

func (p *YourServiceSayHelloArgs)  ReadField2(iprot thrift.TProtocol) error {
  p.Values = &CommonValues{}
  if err := p.Values.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Values), err)
  }
  return nil
}

func (p *YourServiceSayHelloArgs)  ReadField3(iprot thrift.TProtocol) error {
  p.HelloValues = &HelloValues{}
  if err := p.HelloValues.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.HelloValues), err)
  }
  return nil
}

func (p *YourServiceSayHelloArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sayHello_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *YourServiceSayHelloArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("yourName", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:yourName: ", p), err) }
  if err := oprot.WriteString(string(p.YourName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.yourName (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:yourName: ", p), err) }
  return err
}

func (p *YourServiceSayHelloArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("values", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:values: ", p), err) }
  if err := p.Values.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Values), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:values: ", p), err) }
  return err
}

func (p *YourServiceSayHelloArgs) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("helloValues", thrift.STRUCT, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:helloValues: ", p), err) }
  if err := p.HelloValues.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.HelloValues), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:helloValues: ", p), err) }
  return err
}

func (p *YourServiceSayHelloArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("YourServiceSayHelloArgs(%+v)", *p)
}

// Attributes:
//  - Success
type YourServiceSayHelloResult struct {
  Success *SayHelloResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewYourServiceSayHelloResult() *YourServiceSayHelloResult {
  return &YourServiceSayHelloResult{}
}

var YourServiceSayHelloResult_Success_DEFAULT *SayHelloResponse
func (p *YourServiceSayHelloResult) GetSuccess() *SayHelloResponse {
  if !p.IsSetSuccess() {
    return YourServiceSayHelloResult_Success_DEFAULT
  }
return p.Success
}
func (p *YourServiceSayHelloResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *YourServiceSayHelloResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
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

func (p *YourServiceSayHelloResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &SayHelloResponse{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *YourServiceSayHelloResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("sayHello_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *YourServiceSayHelloResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *YourServiceSayHelloResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("YourServiceSayHelloResult(%+v)", *p)
}

// Attributes:
//  - Num
//  - StringValue
//  - BoolValue
type YourServiceEatAppleArgs struct {
  Num int32 `thrift:"num,1" db:"num" json:"num"`
  StringValue string `thrift:"stringValue,2" db:"stringValue" json:"stringValue"`
  BoolValue bool `thrift:"boolValue,3" db:"boolValue" json:"boolValue"`
}

func NewYourServiceEatAppleArgs() *YourServiceEatAppleArgs {
  return &YourServiceEatAppleArgs{}
}


func (p *YourServiceEatAppleArgs) GetNum() int32 {
  return p.Num
}

func (p *YourServiceEatAppleArgs) GetStringValue() string {
  return p.StringValue
}

func (p *YourServiceEatAppleArgs) GetBoolValue() bool {
  return p.BoolValue
}
func (p *YourServiceEatAppleArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
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

func (p *YourServiceEatAppleArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Num = v
}
  return nil
}

func (p *YourServiceEatAppleArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.StringValue = v
}
  return nil
}

func (p *YourServiceEatAppleArgs)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.BoolValue = v
}
  return nil
}

func (p *YourServiceEatAppleArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("eatApple_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *YourServiceEatAppleArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("num", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num: ", p), err) }
  if err := oprot.WriteI32(int32(p.Num)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.num (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num: ", p), err) }
  return err
}

func (p *YourServiceEatAppleArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("stringValue", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stringValue: ", p), err) }
  if err := oprot.WriteString(string(p.StringValue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.stringValue (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stringValue: ", p), err) }
  return err
}

func (p *YourServiceEatAppleArgs) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("boolValue", thrift.BOOL, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:boolValue: ", p), err) }
  if err := oprot.WriteBool(bool(p.BoolValue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.boolValue (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:boolValue: ", p), err) }
  return err
}

func (p *YourServiceEatAppleArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("YourServiceEatAppleArgs(%+v)", *p)
}

// Attributes:
//  - Success
type YourServiceEatAppleResult struct {
  Success *EatAppleResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewYourServiceEatAppleResult() *YourServiceEatAppleResult {
  return &YourServiceEatAppleResult{}
}

var YourServiceEatAppleResult_Success_DEFAULT *EatAppleResponse
func (p *YourServiceEatAppleResult) GetSuccess() *EatAppleResponse {
  if !p.IsSetSuccess() {
    return YourServiceEatAppleResult_Success_DEFAULT
  }
return p.Success
}
func (p *YourServiceEatAppleResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *YourServiceEatAppleResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
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

func (p *YourServiceEatAppleResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &EatAppleResponse{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *YourServiceEatAppleResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("eatApple_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *YourServiceEatAppleResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *YourServiceEatAppleResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("YourServiceEatAppleResult(%+v)", *p)
}


