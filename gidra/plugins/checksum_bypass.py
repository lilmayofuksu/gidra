from gidra import proto
from gidra.proxy import GenshinProxy, HandlerRouter, PacketDirection
from gidra.proxy.cmdids import CmdID

router = HandlerRouter()

@router(CmdID.PlayerLoginReq, PacketDirection.Client)
def change_player_checksum(proxy: GenshinProxy, msg: proto.PlayerLoginReq):
    msg.checksum = "ed9fb95b179f957394ef2d984a397f35e8b31b9850496833399c259b358c9ba723" #2.8
    proxy.send(msg, PacketDirection.Server)