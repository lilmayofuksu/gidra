from gidra.plugins import (change_account, change_nickname,
                            commands, windseed_blocker, seed_exchange, checksum_bypass)


from gidra.proxy import GenshinProxy, PacketDirection
from gidra.proto import QueryCurrRegionHttpRsp
from gidra.mhycrypt import init_keys, decrypt, encrypt_and_sign
import ec2b

from bottle import route, run, request
import requests
import base64

PROXY_GATESERVER = ('127.0.0.1', 8888)

#Change these according to your account's region
TARGET_DISPATCH_URL = 'https://oseurodispatch.yuanshen.com/query_cur_region'
TARGET_GATESERVER = ('47.245.143.151', 22102)

@route('/query_cur_region')
def handle_query_cur():
    # Trick to bypass system proxy, this way we don't need to hardcode the ec2b key
    session = requests.Session()
    session.trust_env = False

    r = session.get(f'{TARGET_DISPATCH_URL}?{request.query_string}')

    if any(map(request.query.version.__contains__, ["2.7.5", "2.8", "2.8.5", "3.0"])):

        key_id = int(request.query.key_id)
        respdata = r.json()
        respdec = decrypt(base64.b64decode(respdata['content']), key_id)

        proto = QueryCurrRegionHttpRsp()
        proto.parse(respdec)

        if proto.retcode == 0:
            proto.region_info.gateserver_ip, proto.region_info.gateserver_port = PROXY_GATESERVER
            proxy.key = ec2b.derive(proto.client_secret_key)

        enc_data, sign = encrypt_and_sign(bytes(proto), key_id)
        return {'content': base64.b64encode(enc_data).decode(), 'sign': base64.b64encode(sign).decode()}

    else:
        proto = QueryCurrRegionHttpRsp()
        proto.parse(base64.b64decode(r.text))

        if proto.retcode == 0:
            proto.region_info.gateserver_ip, proto.region_info.gateserver_port = PROXY_GATESERVER
            proxy.key = ec2b.derive(proto.client_secret_key)

        return base64.b64encode(bytes(proto)).decode()

proxy = GenshinProxy(PROXY_GATESERVER, TARGET_GATESERVER)

def main():
    init_keys("./keys")
    #proxy.add(change_account.router)
    #proxy.add(change_nickname.router)
    proxy.add(windseed_blocker.router)
    proxy.add(seed_exchange.router)
    proxy.add(checksum_bypass.router)
    #proxy.add(commands.router)

    proxy.start()

    run(host='127.0.0.1', port=8081, debug=False)

if __name__ == '__main__':
    main()
