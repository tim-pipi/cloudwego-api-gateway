namespace go api

struct EchoReq {
    1:required string message
}

struct EchoResp {
    1: string response
}

service EchoService {
    EchoResp echo(1: EchoReq request) (api.get="/EchoService/echo")
}
