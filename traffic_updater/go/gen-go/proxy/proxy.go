// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package proxy

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - FromId
//  - ToId
//  - WayId
//  - Speed
//  - TrafficLevel
type Flow struct {
  FromId *int64 `thrift:"fromId,1" db:"fromId" json:"fromId,omitempty"`
  ToId *int64 `thrift:"toId,2" db:"toId" json:"toId,omitempty"`
  WayId int64 `thrift:"wayId,3,required" db:"wayId" json:"wayId"`
  Speed float64 `thrift:"speed,4,required" db:"speed" json:"speed"`
  TrafficLevel int32 `thrift:"trafficLevel,5,required" db:"trafficLevel" json:"trafficLevel"`
}

func NewFlow() *Flow {
  return &Flow{}
}

var Flow_FromId_DEFAULT int64
func (p *Flow) GetFromId() int64 {
  if !p.IsSetFromId() {
    return Flow_FromId_DEFAULT
  }
return *p.FromId
}
var Flow_ToId_DEFAULT int64
func (p *Flow) GetToId() int64 {
  if !p.IsSetToId() {
    return Flow_ToId_DEFAULT
  }
return *p.ToId
}

func (p *Flow) GetWayId() int64 {
  return p.WayId
}

func (p *Flow) GetSpeed() float64 {
  return p.Speed
}

func (p *Flow) GetTrafficLevel() int32 {
  return p.TrafficLevel
}
func (p *Flow) IsSetFromId() bool {
  return p.FromId != nil
}

func (p *Flow) IsSetToId() bool {
  return p.ToId != nil
}

func (p *Flow) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetWayId bool = false;
  var issetSpeed bool = false;
  var issetTrafficLevel bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
        issetWayId = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.DOUBLE {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
        issetSpeed = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
        issetTrafficLevel = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
  if !issetWayId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field WayId is not set"));
  }
  if !issetSpeed{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Speed is not set"));
  }
  if !issetTrafficLevel{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TrafficLevel is not set"));
  }
  return nil
}

func (p *Flow)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.FromId = &v
}
  return nil
}

func (p *Flow)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ToId = &v
}
  return nil
}

func (p *Flow)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.WayId = v
}
  return nil
}

func (p *Flow)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Speed = v
}
  return nil
}

func (p *Flow)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.TrafficLevel = v
}
  return nil
}

func (p *Flow) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Flow"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Flow) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetFromId() {
    if err := oprot.WriteFieldBegin("fromId", thrift.I64, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:fromId: ", p), err) }
    if err := oprot.WriteI64(int64(*p.FromId)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.fromId (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:fromId: ", p), err) }
  }
  return err
}

func (p *Flow) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetToId() {
    if err := oprot.WriteFieldBegin("toId", thrift.I64, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:toId: ", p), err) }
    if err := oprot.WriteI64(int64(*p.ToId)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.toId (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:toId: ", p), err) }
  }
  return err
}

func (p *Flow) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("wayId", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:wayId: ", p), err) }
  if err := oprot.WriteI64(int64(p.WayId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.wayId (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:wayId: ", p), err) }
  return err
}

func (p *Flow) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("speed", thrift.DOUBLE, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:speed: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Speed)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.speed (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:speed: ", p), err) }
  return err
}

func (p *Flow) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("trafficLevel", thrift.I32, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:trafficLevel: ", p), err) }
  if err := oprot.WriteI32(int32(p.TrafficLevel)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.trafficLevel (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:trafficLevel: ", p), err) }
  return err
}

func (p *Flow) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Flow(%+v)", *p)
}

type ProxyService interface {
  GetAllFlows(ctx context.Context) (r []*Flow, err error)
  // Parameters:
  //  - WayId
  GetFlowById(ctx context.Context, wayId int64) (r *Flow, err error)
  // Parameters:
  //  - WayIds
  GetFlowsByIds(ctx context.Context, wayIds []int64) (r []*Flow, err error)
}

type ProxyServiceClient struct {
  c thrift.TClient
}

func NewProxyServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ProxyServiceClient {
  return &ProxyServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewProxyServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ProxyServiceClient {
  return &ProxyServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewProxyServiceClient(c thrift.TClient) *ProxyServiceClient {
  return &ProxyServiceClient{
    c: c,
  }
}

func (p *ProxyServiceClient) Client_() thrift.TClient {
  return p.c
}
func (p *ProxyServiceClient) GetAllFlows(ctx context.Context) (r []*Flow, err error) {
  var _args0 ProxyServiceGetAllFlowsArgs
  var _result1 ProxyServiceGetAllFlowsResult
  if err = p.Client_().Call(ctx, "getAllFlows", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - WayId
func (p *ProxyServiceClient) GetFlowById(ctx context.Context, wayId int64) (r *Flow, err error) {
  var _args2 ProxyServiceGetFlowByIdArgs
  _args2.WayId = wayId
  var _result3 ProxyServiceGetFlowByIdResult
  if err = p.Client_().Call(ctx, "getFlowById", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

// Parameters:
//  - WayIds
func (p *ProxyServiceClient) GetFlowsByIds(ctx context.Context, wayIds []int64) (r []*Flow, err error) {
  var _args4 ProxyServiceGetFlowsByIdsArgs
  _args4.WayIds = wayIds
  var _result5 ProxyServiceGetFlowsByIdsResult
  if err = p.Client_().Call(ctx, "getFlowsByIds", &_args4, &_result5); err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

type ProxyServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler ProxyService
}

func (p *ProxyServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ProxyServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ProxyServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewProxyServiceProcessor(handler ProxyService) *ProxyServiceProcessor {

  self6 := &ProxyServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["getAllFlows"] = &proxyServiceProcessorGetAllFlows{handler:handler}
  self6.processorMap["getFlowById"] = &proxyServiceProcessorGetFlowById{handler:handler}
  self6.processorMap["getFlowsByIds"] = &proxyServiceProcessorGetFlowsByIds{handler:handler}
return self6
}

func (p *ProxyServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x7

}

type proxyServiceProcessorGetAllFlows struct {
  handler ProxyService
}

func (p *proxyServiceProcessorGetAllFlows) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ProxyServiceGetAllFlowsArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getAllFlows", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ProxyServiceGetAllFlowsResult{}
var retval []*Flow
  var err2 error
  if retval, err2 = p.handler.GetAllFlows(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getAllFlows: " + err2.Error())
    oprot.WriteMessageBegin("getAllFlows", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getAllFlows", thrift.REPLY, seqId); err2 != nil {
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

type proxyServiceProcessorGetFlowById struct {
  handler ProxyService
}

func (p *proxyServiceProcessorGetFlowById) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ProxyServiceGetFlowByIdArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getFlowById", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ProxyServiceGetFlowByIdResult{}
var retval *Flow
  var err2 error
  if retval, err2 = p.handler.GetFlowById(ctx, args.WayId); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getFlowById: " + err2.Error())
    oprot.WriteMessageBegin("getFlowById", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getFlowById", thrift.REPLY, seqId); err2 != nil {
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

type proxyServiceProcessorGetFlowsByIds struct {
  handler ProxyService
}

func (p *proxyServiceProcessorGetFlowsByIds) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ProxyServiceGetFlowsByIdsArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getFlowsByIds", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ProxyServiceGetFlowsByIdsResult{}
var retval []*Flow
  var err2 error
  if retval, err2 = p.handler.GetFlowsByIds(ctx, args.WayIds); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getFlowsByIds: " + err2.Error())
    oprot.WriteMessageBegin("getFlowsByIds", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getFlowsByIds", thrift.REPLY, seqId); err2 != nil {
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


// HELPER FUNCTIONS AND STRUCTURES

type ProxyServiceGetAllFlowsArgs struct {
}

func NewProxyServiceGetAllFlowsArgs() *ProxyServiceGetAllFlowsArgs {
  return &ProxyServiceGetAllFlowsArgs{}
}

func (p *ProxyServiceGetAllFlowsArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *ProxyServiceGetAllFlowsArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getAllFlows_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ProxyServiceGetAllFlowsArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetAllFlowsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ProxyServiceGetAllFlowsResult struct {
  Success []*Flow `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewProxyServiceGetAllFlowsResult() *ProxyServiceGetAllFlowsResult {
  return &ProxyServiceGetAllFlowsResult{}
}

var ProxyServiceGetAllFlowsResult_Success_DEFAULT []*Flow

func (p *ProxyServiceGetAllFlowsResult) GetSuccess() []*Flow {
  return p.Success
}
func (p *ProxyServiceGetAllFlowsResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ProxyServiceGetAllFlowsResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *ProxyServiceGetAllFlowsResult)  ReadField0(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*Flow, 0, size)
  p.Success =  tSlice
  for i := 0; i < size; i ++ {
    _elem8 := &Flow{}
    if err := _elem8.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem8), err)
    }
    p.Success = append(p.Success, _elem8)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ProxyServiceGetAllFlowsResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getAllFlows_result"); err != nil {
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

func (p *ProxyServiceGetAllFlowsResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Success {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ProxyServiceGetAllFlowsResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetAllFlowsResult(%+v)", *p)
}

// Attributes:
//  - WayId
type ProxyServiceGetFlowByIdArgs struct {
  WayId int64 `thrift:"wayId,1" db:"wayId" json:"wayId"`
}

func NewProxyServiceGetFlowByIdArgs() *ProxyServiceGetFlowByIdArgs {
  return &ProxyServiceGetFlowByIdArgs{}
}


func (p *ProxyServiceGetFlowByIdArgs) GetWayId() int64 {
  return p.WayId
}
func (p *ProxyServiceGetFlowByIdArgs) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *ProxyServiceGetFlowByIdArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.WayId = v
}
  return nil
}

func (p *ProxyServiceGetFlowByIdArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getFlowById_args"); err != nil {
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

func (p *ProxyServiceGetFlowByIdArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("wayId", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:wayId: ", p), err) }
  if err := oprot.WriteI64(int64(p.WayId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.wayId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:wayId: ", p), err) }
  return err
}

func (p *ProxyServiceGetFlowByIdArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetFlowByIdArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ProxyServiceGetFlowByIdResult struct {
  Success *Flow `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewProxyServiceGetFlowByIdResult() *ProxyServiceGetFlowByIdResult {
  return &ProxyServiceGetFlowByIdResult{}
}

var ProxyServiceGetFlowByIdResult_Success_DEFAULT *Flow
func (p *ProxyServiceGetFlowByIdResult) GetSuccess() *Flow {
  if !p.IsSetSuccess() {
    return ProxyServiceGetFlowByIdResult_Success_DEFAULT
  }
return p.Success
}
func (p *ProxyServiceGetFlowByIdResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ProxyServiceGetFlowByIdResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *ProxyServiceGetFlowByIdResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Flow{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ProxyServiceGetFlowByIdResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getFlowById_result"); err != nil {
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

func (p *ProxyServiceGetFlowByIdResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *ProxyServiceGetFlowByIdResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetFlowByIdResult(%+v)", *p)
}

// Attributes:
//  - WayIds
type ProxyServiceGetFlowsByIdsArgs struct {
  WayIds []int64 `thrift:"wayIds,1" db:"wayIds" json:"wayIds"`
}

func NewProxyServiceGetFlowsByIdsArgs() *ProxyServiceGetFlowsByIdsArgs {
  return &ProxyServiceGetFlowsByIdsArgs{}
}


func (p *ProxyServiceGetFlowsByIdsArgs) GetWayIds() []int64 {
  return p.WayIds
}
func (p *ProxyServiceGetFlowsByIdsArgs) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *ProxyServiceGetFlowsByIdsArgs)  ReadField1(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]int64, 0, size)
  p.WayIds =  tSlice
  for i := 0; i < size; i ++ {
var _elem9 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem9 = v
}
    p.WayIds = append(p.WayIds, _elem9)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ProxyServiceGetFlowsByIdsArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getFlowsByIds_args"); err != nil {
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

func (p *ProxyServiceGetFlowsByIdsArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("wayIds", thrift.LIST, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:wayIds: ", p), err) }
  if err := oprot.WriteListBegin(thrift.I64, len(p.WayIds)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.WayIds {
    if err := oprot.WriteI64(int64(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:wayIds: ", p), err) }
  return err
}

func (p *ProxyServiceGetFlowsByIdsArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetFlowsByIdsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ProxyServiceGetFlowsByIdsResult struct {
  Success []*Flow `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewProxyServiceGetFlowsByIdsResult() *ProxyServiceGetFlowsByIdsResult {
  return &ProxyServiceGetFlowsByIdsResult{}
}

var ProxyServiceGetFlowsByIdsResult_Success_DEFAULT []*Flow

func (p *ProxyServiceGetFlowsByIdsResult) GetSuccess() []*Flow {
  return p.Success
}
func (p *ProxyServiceGetFlowsByIdsResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ProxyServiceGetFlowsByIdsResult) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *ProxyServiceGetFlowsByIdsResult)  ReadField0(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*Flow, 0, size)
  p.Success =  tSlice
  for i := 0; i < size; i ++ {
    _elem10 := &Flow{}
    if err := _elem10.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem10), err)
    }
    p.Success = append(p.Success, _elem10)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ProxyServiceGetFlowsByIdsResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getFlowsByIds_result"); err != nil {
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

func (p *ProxyServiceGetFlowsByIdsResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Success {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ProxyServiceGetFlowsByIdsResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ProxyServiceGetFlowsByIdsResult(%+v)", *p)
}

