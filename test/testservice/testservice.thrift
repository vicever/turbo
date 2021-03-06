namespace go gen
include "shared.thrift"

struct SayHelloResponse {
  1: string message,
}

service TestService {
    SayHelloResponse sayHello (1:shared.CommonValues values, 2:string yourName, 3:i64 int64Value, 4:bool boolValue,
     5:double float64Value, 6:i64 uint64Value, 7:i32 int32Value, 8:i16 int16Value)
}