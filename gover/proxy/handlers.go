package proxy

import (
	"encoding/base64"

	"github.com/MoonlightPS/Iridium-gidra/gover/gen"
	"github.com/MoonlightPS/Iridium-gidra/gover/utils"
)

type Handler = func(*KCPConn, []byte) ([]byte, error)

var handlersMap = map[int]Handler{}

var b64 = base64.StdEncoding

func HandleGetPlayerTokenReq(conn *KCPConn, data []byte) ([]byte, error) {
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.GetPlayerTokenReq)

	seedEncrypted, err := b64.DecodeString(body.GetClientSeed())
	if err != nil {
		return nil, err
	}

	seedBytes, err := utils.Decrypt(seedEncrypted, utils.SIGN_KEY)
	if err != nil {
		return nil, err
	}
	conn.seed = be.Uint64(seedBytes)

	keyID := int(body.GetKeyId())
	if keyID == utils.CN_KEY {
		keyID = utils.CN_SIGN_KEY
	} else {
		keyID = utils.OS_SIGN_KEY
	}
	seedEncrypted, err = utils.Encrypt(seedBytes, keyID)
	if err != nil {
		return nil, err
	}

	body.ClientSeed = b64.EncodeToString(seedEncrypted)
	return conn.parser.Compose(msg)
}

func HandleGetPlayerTokenRsp(conn *KCPConn, data []byte) ([]byte, error) {
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.GetPlayerTokenRsp)

	seedEncrypted, err := b64.DecodeString(body.GetEncryptedSeed())
	if err != nil {
		return nil, err
	}

	keyID := int(body.GetKeyId())
	seedBytes, err := utils.Decrypt(seedEncrypted, keyID)
	if err != nil {
		return nil, err
	}
	conn.seed = be.Uint64(seedBytes) ^ conn.seed

	signature, err := utils.Sign(seedBytes, utils.SIGN_KEY)
	if err != nil {
		return nil, err
	}

	body.SeedSignature = b64.EncodeToString(signature)
	ret, err := conn.parser.Compose(msg)
	conn.key.GenKey(conn.seed)
	return ret, err
}

func HandlePlayerLoginReq(conn *KCPConn, data []byte) ([]byte, error) {
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.PlayerLoginReq)

	if conn.keyID == utils.CN_KEY {
		body.Checksum = "64309cf5f6d6b7c427d3e15622636372c14bc8ce7252be4bd27e9a1866b688c226"
	} else {
		body.Checksum = "eb8aeaf9f40c5bc5af2ac93ad1da07fa05acf5206fe08c10290357a414aecb7c24"
	}

	return conn.parser.Compose(msg)
}

func init() {
	handlersMap[utils.GetPlayerTokenReq] = HandleGetPlayerTokenReq
	handlersMap[utils.GetPlayerTokenRsp] = HandleGetPlayerTokenRsp
	handlersMap[utils.PlayerLoginReq] = HandlePlayerLoginReq
}
