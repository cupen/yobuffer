package idl

const testdata_full = `package pkgname;

// 结构体定义
@meta(id=10001, desc="玩家信息")
struct PlayerInfo {
    userId: int64 = 1;
    name: string = 2;
    avatar: string = 3;
    shortUID: int64 = 4;
    level: int16 = 5;
    isAmind: bool = 6;
    friends: list<int64> = 7 ;
    fightRecords: map<int64, int32> = 8;
}

// // 外部结构体定义(用于引用外部 IDL 如 protobuf 的定义)
// @meta(id=10002)
// import protobuf PlayerInfo


@meta(id=10003, desc="作为RPC参数")
struct PlayerID {
    string: userId  = 1;
}

// enum Color : byte {
//     Red = 1;
//     Green = 2;
//     Blue =3;
// }

// gloabl options
// option(rate1, 50/s)
// option(rate2, 100/m)
// option(rate3, 1000/h)

service Backend {
    // 获取玩家信息
    @meta(id=20001, rate=60/m)
    rpc GetPlayerInfo(pid: int64 = 1, index: int32 = 2) -> PlayerInfo

    // 发送聊天信息
    @meta(id=20003, rate=10/m)
    rpc SendChatMessage(id:PlayerID) -> void

    // 记录日志
    // fire-and-forget 
    @meta(id=20004, rate=100/s)
    faf SendLog(pid:uint64 = 1, event:string = 2, args:string = 3) -> void
}

// }
// rpc CallAPI {
//    args {
//        string UserID  = 1;
//    }
//    retrun PlayerInfo
// }
`
