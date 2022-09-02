from gidra import proto
from gidra.proxy import GenshinProxy, HandlerRouter, PacketDirection
from gidra.proxy.cmdids import CmdID

router = HandlerRouter()

@router(CmdID.PlayerLoginReq, PacketDirection.Client)
def change_player_checksum(proxy: GenshinProxy, msg: proto.PlayerLoginReq):
    #msg.checksum = "ed9fb95b179f957394ef2d984a397f35e8b31b9850496833399c259b358c9ba723" #2.8
    msg.checksum = "c071e821a011fe7a5f6c791d4002dc4b2ed2e864481c6fe2e9db3b6379c18f6b25" #3.0
    proxy.send(msg, PacketDirection.Server)
