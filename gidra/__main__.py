import datetime
import json
import os

from loguru import logger
import gidra
from gidra.plugins import (change_account, change_nickname,
                            commands, windseed_blocker, seed_exchange, checksum_bypass)


from gidra.proxy import GenshinProxy, PacketDirection
from gidra.proto import QueryCurrRegionHttpRsp
from gidra.mhycrypt import init_keys, decrypt, encrypt_and_sign
import ec2b

from bottle import route, run, request
import requests
import base64

PROXY_GATESERVER = ('127.0.0.1', 8888) # Game Proxy
PROXY_DISPATCH = ('127.0.0.1',8081) # Http Proxy

TARGET_DISPATCH_URL = None
TARGET_GATESERVER = None
proxy: GenshinProxy = None

@route('/query_cur_region')
def handle_query_cur():

    global TARGET_DISPATCH_URL

    region = request.get_header('url')
    logger.info(f'Found Region: {str(region).split("dispatch")[0]}')

    TARGET_DISPATCH_URL =  "https://" + str(region) + "/query_cur_region"

    # Trick to bypass system proxy, this way we don't need to hardcode the ec2b key
    session = requests.Session()
    session.trust_env = False

    r = session.get(f'{TARGET_DISPATCH_URL}?{request.query_string}')

    if any(map(request.query.version.__contains__, ["2.7.5", "2.8.","3."])):

        key_id = int(request.query.key_id)
        respdata = r.json()
        respdec = decrypt(base64.b64decode(respdata['content']), key_id)

        proto = QueryCurrRegionHttpRsp()
        proto.parse(respdec)

        global TARGET_GATESERVER
        TARGET_GATESERVER = (proto.region_info.gateserver_ip,proto.region_info.gateserver_port)
        logger.info(f'Connecting to GateServer: {TARGET_GATESERVER}')

        global proxy
        proxy = GenshinProxy(PROXY_GATESERVER,TARGET_GATESERVER,key_id)
        initProxy()

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


def initProxy():
    #proxy.add(change_account.router)
    #proxy.add(change_nickname.router)
    # proxy.add(windseed_blocker.router) # Already blocked in proxy as kcp crashes
    proxy.add(seed_exchange.router)
    proxy.add(checksum_bypass.router)
    #proxy.add(commands.router)
    proxy.start()

def main():
    logger.info('Starting gidra proxy...')
    logger.debug(f'Proxy running on {PROXY_DISPATCH[0]}:{PROXY_DISPATCH[1]}, you can start your game now!')
    init_keys("./keys")
    run(host=PROXY_DISPATCH[0], port=PROXY_DISPATCH[1], debug=False, quiet = True)

if __name__ == '__main__':
    main()

def handleExit():
    logger.info('Captured a total of {} packets!'.format(len(proxy.packet_data)))
    try:
        logger.debug('Saving the dump file to IG_' +  datetime.datetime.now().strftime("%Y-%m-%d-%H.%M.%S")+'.json')
        if(len(proxy.packet_data)!=0):
            json.dump(proxy.packet_data,open("gidra/packet_dump/IG_"+datetime.datetime.now().strftime("%Y-%m-%d-%H.%M.%S")+'.json','w'),indent=2)
            logger.debug('Done saving..')
    except Exception as e:
        logger.error(f'Error while saving dump ({e})')


import atexit
atexit.register(handleExit)

