# Iridium-gidra

A proxy between the client and the server of a certain anime game

## Usage (Go version)
[Click Here](https://github.com/MoonlightPS/Iridium-gidra/tree/master/gover) to view usage for go version

Highly recommended to use this version as there is a huge performance increase thanks to [Night12138](https://github.com/Night12138)

## Warning
WindSeed is not blocked on either of the proxies! for your [safety](https://github.com/MoonlightPS/Iridium-gidra/commit/a58ac1263e85539e7475100ea50ba86a1ea7ed7b#commitcomment-87842613)

## Usage (Python version)

1. Clone the Github Repository

```powershell
git clone https://github.com/MoonlightPS/Iridium-gidra.git
cd Iridium-gidra
```

2. Install [Poetry](https://python-poetry.org/docs/#installing-with-the-official-installer) and use the following command to install the required dependencies

```powershell
poetry install
```

3. Install python-kcp
```powershell
cd python-kcp
python setup.py install
```

4. Run the following commands to start proxy

```powershell
py -m gidra
```

5. Use the following [fiddler script](https://github.com/MoonlightPS/Iridium-gidra#fiddler-script) to redirect dispatch

6. **Use patched `UserAssembly.dll` or the proxy won't work!! and be sure to change it back when you are not using the proxy!!**

7. Start the game and have fun!

## Fiddler Script
```cs
/* Gidra proxy fiddler script */
import System;
import System.Windows.Forms;
import Fiddler;
import System.Text.RegularExpressions;

class Handlers
{
    static function OnBeforeRequest(oS: Session) {

        if(oS.host.EndsWith("dispatch.yuanshen.com")) {
            oS.oRequest.headers.UriScheme = "http";
            oS.oRequest.headers.Add('url',oS.host);
            oS.host = "localhost";
            oS.port = 8081;
        }

        if(oS.host.Contains("overseauspider.yuanshen.com")){
            oS.oRequest.FailSession(404, "Blocked", "your mom");
        }
    }
};
```

## Note

- Packets captured are saved after you exit out of gidra in console and can be found in `./gidra/packet_dump`
- proxy auto detects dispatch url and gateserver address when using the above fiddler script, you do not have to hardcode any of these!

### Format of packet capture:

```jsonc
[
  {
    "index": int,
    "packetId": int,
    "protoName": string,
    "source": string,
    "time": float,
    "object": protobuf object
  }
]
```
