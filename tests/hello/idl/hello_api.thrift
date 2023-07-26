namespace go api

struct HelloReq {
    1: required string Name (api.query="name")
}   

struct HelloResp {
    1: string RespBody
}

struct EchoReq {
    1: required string message (vt.min_size = "1")
}

struct EchoResp {
    1: string response
}

service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/HelloService/HelloMethod")
    EchoResp echo(1: EchoReq request) (api.get="/HelloService/echo")
}
