namespace go test

struct Req {
    1: required string Name (api.query="name")
}   

struct Resp {
    1: string RespBody
}

service TestService {
    Req TestMethod(1: Resp request) (api.get="/TestService/TestMethod")
}