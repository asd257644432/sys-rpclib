namespace go lql.sys.test

struct Request {
        1: string message
}

struct Response {
        1: string message
}

service Hello {
    Response echo(1: Request req)
}