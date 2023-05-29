namespace go hello

struct Request {
	1: string Message
}

struct Response {
	1: string message
}

service Echo {
    Response echo(1: Request req)
}

