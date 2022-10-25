# Iridium-gidra

A proxy between the client and the server of a certain anime game (Written in GoLang by [Night12138](https://github.com/Night12138))

## Usage (Go version)
Highly recommended to use this version as there is a huge performance increase
1. Clone the Github Repository
```powershell
git clone https://github.com/MoonlightPS/Iridium-gidra.git
cd Iridium-gidra/gover
```

2. Install [GoLang](https://go.dev/doc/install)

3. Run the following commands to start the proxy 

```powershell
go install
go run main.go
```

4. Use the following [fiddler script](https://github.com/MoonlightPS/Iridium-gidra#fiddler-script) to redirect dispatch

5. **Use patched `UserAssembly.dll` or the proxy won't work!! and be sure to change it back when you are not using the proxy!!**

6. Start the game and have fun!

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

- Packets captured by gover are stored in `./gover` and are saved as soon as you quit the game or exit the console
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
