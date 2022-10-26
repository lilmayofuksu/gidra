from . import core
import sys, threading

__all__ = ["KcpObj"]

i = 0

class KcpObj:
    def __init__(self, conv, token, callback):
        self.conv = conv
        self.token = token
        self.mutex = threading.Lock()

        global i
        self.cobj = core.lkcp_create(conv, token, i, callback)
        i += 1

    def wndsize(self, sndwnd, rcvwnd):
        core.lkcp_wndsize(self.cobj, sndwnd, rcvwnd)

    def nodelay(self, nodelay, interval, resend, nc):
        return core.lkcp_nodelay(self.cobj, nodelay, interval, resend, nc)

    def check(self, current):
        self.mutex.acquire()
        ret = core.lkcp_check(self.cobj, current)
        self.mutex.release()
        return ret

    def update(self, current):
        self.mutex.acquire()
        core.lkcp_update(self.cobj, current)
        self.mutex.release()

    def send(self, data):
        self.mutex.acquire()
        if sys.version_info.major == 3 and isinstance(data, str):
            data = data.encode("UTF-8")
        ret = core.lkcp_send(self.cobj, data)
        self.mutex.release()
        return ret

    def input(self, data):
        self.mutex.acquire()
        ret = core.lkcp_input(self.cobj, data)
        self.mutex.release()
        return ret

    def recv(self):
        self.mutex.acquire()
        ret = core.lkcp_recv(self.cobj)
        self.mutex.release()
        return ret

    def flush(self):
        self.mutex.acquire()
        core.lkcp_flush(self.cobj)
        self.mutex.release()

    
    def setmtu(self, mtu):
        core.lkcp_setmtu(self.cobj, mtu)
