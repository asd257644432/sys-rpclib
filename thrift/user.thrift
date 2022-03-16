namespace go user

include "./common.thrift"
include "./base.thrift"

struct GetUserRequest {
        1: i64 UserId
        255: base.Base Base
}

struct GetUserResponse {
        1: common.User User
        255: base.BaseResp BaseResp
}

service User {
    GetUserResponse GetUser(1: GetUserRequest req)
}