from itertools import cycle
from typing import Tuple

from .mt64 import mt64

from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_v1_5
from Crypto.Signature import pkcs1_15
from Crypto.Hash import SHA256

import os

keys = {}

def init_keys(path: str):
    with open(os.path.join(path, 'MHYPrivCN.pem'), 'r') as f:
        keys[2] = RSA.import_key(f.read())

    with open(os.path.join(path, 'MHYPrivOS.pem'), 'r') as f:
        keys[3] = RSA.import_key(f.read())

    with open(os.path.join(path, 'SigningKey.pem'), 'r') as f:
        keys["SigningKey"] = RSA.import_key(f.read())

    with open(os.path.join(path, 'MHYSignRsaOS.pem'), 'r') as f:
        keys["MHYSignOS"] = RSA.import_key(f.read())

    with open(os.path.join(path, 'MHYSignRsaCN.pem'), 'r') as f:
        keys["MHYSignCN"] = RSA.import_key(f.read())


def new_key(seed: int) -> bytes:
    mt = mt64()
    mt.seed(seed)

    mt.seed(mt.int64())
    mt.int64()

    return bytes(byte for _ in range(512) for byte in mt.int64().to_bytes(8, "big"))

def xor(data: bytes, key: bytes) -> bytes:
    return bytes(v ^ k for (v, k) in zip(data, cycle(key)))

def decrypt(data: bytes, key_id) -> bytes:
    key = keys[key_id]
    dec = PKCS1_v1_5.new(key)

    chunk_size = 256
    out = b''

    for i in range(0, len(data), chunk_size):
        chunk = data[i:i + chunk_size]
        out += dec.decrypt(chunk, None)

    return out

def do_sign(message, key):
    signer = pkcs1_15.new(key)
    digest = SHA256.new(message)
    return signer.sign(digest)

def encrypt_and_sign(data: bytes, key_id) -> Tuple[bytes, bytes]:
    sign_key = keys["SigningKey"]

    key = keys[key_id]
    enc = PKCS1_v1_5.new(key)

    chunk_size = 256 - 11
    out = b''

    if len(data) > chunk_size:
        for i in range(0, len(data), chunk_size):
            chunk = data[i:i + chunk_size]
            out += enc.encrypt(chunk)
    else:
        out = enc.encrypt(data)

    signature = do_sign(data, sign_key)

    return out, signature