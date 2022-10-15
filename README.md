# Iridium-gidra

A proxy between the client and the server of a certain anime game

## Usage

1. Clone the Github Repository

```powershell
git clone https://github.com/MoonlightPS/Iridium-gidra.git
cd Iridium-gidra
```

2. Install [Poetry](https://python-poetry.org/docs/#installing-with-the-official-installer) and use the following command to install the required dependencies

```powershell
poetry install
```

3. Run the following commands to start proxy

```powershell
py -m gidra
```

4. Use the following fiddler script to redirect dispatch

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

5. **Use patched `UserAssembly.dll` or the proxy won't work!! and be sure to change it back when you are not using the proxy!!**

6. Start the game and have fun!

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
