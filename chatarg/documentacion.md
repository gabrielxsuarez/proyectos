# Documentación Cliente WebSocket IRC - Chat Argentina

Esta documentación describe el protocolo completo para implementar un cliente WebSocket que se conecte al servidor IRC de Chat Argentina.

## Información de Conexión

### Endpoints WebSocket
- **Servidor Activo**: `ws://wss.dalechatea.me:1245/` (sin SSL) - **ÚNICO FUNCIONAL**
- ~~**Producción**: `wss://ws03.dalechatea.me/` (puerto 443) - INACTIVO~~
- ~~**Alternativo 1**: `wss://wss.dalechatea.me:1242/` (con SSL) - INACTIVO~~
- ~~**Backend**: `wss://wss.dalechatea.me:1239/` (con restricciones de seguridad) - INACTIVO~~

### Headers Requeridos para Conexión

```
Host: chatarg.com (o el host correspondiente)
Connection: Upgrade
Pragma: no-cache
Cache-Control: no-cache
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Upgrade: websocket
Origin: https://chatarg.com
Sec-WebSocket-Version: 13
Accept-Encoding: gzip, deflate, br, zstd
Accept-Language: es,es-ES;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6
Sec-WebSocket-Key: [GENERADO_AUTOMATICAMENTE]
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
X-Forwarded-For: [IP_DEL_CLIENTE] (opcional, para enmascaramiento)
```

## Flujo de Conexión

### 1. Establecer Conexión WebSocket
Conectar al endpoint WebSocket con los headers especificados.

### 2. Mensaje Inicial del Servidor
El servidor envía un mensaje con sessionid:
```json
{
  "seq": 0,
  "sessionid": "client00.NwPT7QRr94"
}
```

### 3. Secuencia de Inicialización del Cliente
Después de recibir el mensaje inicial (seq=0), enviar en orden:

#### A. Cliente Info
```json
{
  "seq": 1,
  "cmd": "clientinfo",
  "localtime": 1747153164083,
  "tzoffset": -180,
  "clientinfo": "{}"
}
```

#### B. Embed Info
```json
{
  "seq": 2,
  "cmd": "embed",
  "channel": "IRCClient",
  "referrer": "https://chatarg.com/webchat/"
}
```

#### C. Conectar al Servidor IRC
```json
{
  "seq": 3,
  "cmd": "connect",
  "channel": "IRCClient",
  "data": "irc.dalechatea.me:6697",
  "nick": "tu_nickname",
  "pass": "",
  "authMethod": "nickserv",
  "joinchannels": "#Argentina",
  "charset": "utf-8"
}
```

## Protocolo de Mensajes

### Control de Secuencia
- Cada mensaje tiene un campo `seq` (secuencia)
- El servidor inicia con `seq: 0` y envía `sessionid`
- El cliente debe iniciar con `seq: 1` e incrementar para cada mensaje enviado
- El servidor también usa `seq` para mensajes recibidos

### Keep-Alive
Enviar ping cada 30 segundos:
```json
{
  "seq": [SIGUIENTE_SEQ],
  "cmd": "nping",
  "m": "checking"
}
```

## Mensajes del Servidor (Recibidos)

### Estados y Conexión

#### `sessionid` - Mensaje inicial con ID de sesión
```json
{
  "seq": 0,
  "sessionid": "client00.NwPT7QRr94"
}
```

#### `status` - Estado del servidor
```json
{
  "seq": 1,
  "pendingDNS": 1,
  "cmd": "status",
  "pending": 0,
  "hostname": "186.33.238.108",
  "connections": 0,
  "channel": "IRCClient",
  "readyToConnect": true
}
```

#### `connected` - Confirmación de conexión
```json
{
  "seq": 2,
  "cmd": "connected",
  "channel": "IRCClient",
  "name": "irc.dalechatea.me",
  "network": "irc.dalechatea.me"
}
```

#### `init` - Inicialización del cliente
```json
{
  "seq": 4,
  "cmd": "init",
  "nick": "tu_nickname",
  "channel": "IRCClient:irc.dalechatea.me",
  "ctime": "39536029001"
}
```

### Información del Canal

#### `nicklist` - Lista de usuarios en canal
```json
{
  "seq": 52,
  "cmd": "nicklist",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "nicks": [
    {
      "host": "pju9cr.9lo2.8070.g0ad65.IP",
      "nick": "leonnel"
    }
  ],
  "channeltype": "chat"
}
```

#### `topic` - Tema del canal
```json
{
  "seq": 25,
  "cmd": "topic",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "nick": "ChanServ",
  "topic": "Bienvenidos al chat de Argentina",
  "time": 1747134282,
  "localchannel": "#Argentina",
  "channeltype": "chat"
}
```

#### `topicwho` - Información del tema
```json
{
  "seq": 12,
  "cmd": "topicwho",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "creator": "ElGranArdilla",
  "date": "1657870309"
}
```

### Mensajes de Chat

#### `msg` - Mensaje de chat
```json
{
  "seq": 171,
  "cmd": "msg",
  "nick": "El-Mas-Dulce",
  "localchannel": "#Argentina",
  "channeltype": "chat",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "msg": "\\u0002\\u00032Hola mundo",
  "replyTo": "usuario: mensaje original"
}
```

#### `emote` - Mensaje de acción
```json
{
  "seq": 696,
  "cmd": "emote",
  "ctime": "1539536029001",
  "nick": "Azimut_",
  "localchannel": "#Argentina",
  "channeltype": "chat",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "emote": ":kiss Azura"
}
```

#### `typing` - Usuario escribiendo
```json
{
  "seq": 689,
  "cmd": "typing",
  "nick": "Metalero53",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "typing": false
}
```

### Eventos de Usuarios

#### `join` - Usuario se une al canal
```json
{
  "seq": 9,
  "cmd": "join",
  "user": "DaleChat",
  "host": "b3v.np7.k217c4.IP",
  "nick": "zero123x",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "ctime": "1539536029001"
}
```

#### `part` - Usuario sale del canal
```json
{
  "seq": 64,
  "cmd": "part",
  "nick": "MilagrosArg",
  "user": "DaleChat",
  "host": "4v6.ji6.7ksean.IP",
  "quit": true,
  "message": "Salir: Desconectado",
  "channel": "IRCClient:irc.dalechatea.me",
  "ctime": "1539536029001"
}
```

#### `changenick` - Cambio de nickname
```json
{
  "seq": 901,
  "cmd": "changenick",
  "nick": "Sergio",
  "newnick": "Guest22918",
  "channel": "IRCClient:irc.dalechatea.me",
  "ctime": "1539536029001"
}
```

#### `away` - Usuario ausente
```json
{
  "seq": 1079,
  "cmd": "away",
  "nick": "Briana_30",
  "channel": "IRCClient:irc.dalechatea.me"
}
```

### Información de Usuarios

#### `userinfo` - Información detallada del usuario
```json
{
  "seq": 65,
  "cmd": "userinfo",
  "nick": "Metalhera",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "ct": 1747153164083,
  "info": "{\\\"userIcon\\\":\\\"https://a.dalechatea.me/image.jpg\\\"}",
  "cc": "AR",
  "tz": "-180"
}
```

#### `usericon` - Icono de usuario
```json
{
  "seq": 3,
  "cmd": "usericon",
  "nick": "MiUsuario",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "icon": "https://dalechatea.me/images/image.png",
  "channeltype": "chat",
  "localchannel": "#Argentina"
}
```

#### `avatar` - Avatar de usuario
```json
{
  "seq": 723,
  "cmd": "avatar",
  "nick": "Georgina38",
  "localchannel": "#argentina",
  "channel": "IRCClient:irc.dalechatea.me:#argentina",
  "channeltype": "chat",
  "userIcon": "https://a.dalechatea.me/avatar.jpg"
}
```

### Moderación

#### `kick` - Usuario expulsado
```json
{
  "seq": 829,
  "cmd": "kick",
  "nick": "Gerardo7",
  "kicker": "Bot",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "ctime": "1539536029001",
  "reason": "Razón de la expulsión"
}
```

#### `ban` - Usuario baneado
```json
{
  "seq": 830,
  "cmd": "ban",
  "by": "Bot",
  "nick": "*!*@54g.4l0.4serdp.IP",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "ctime": "1539536029001"
}
```

#### `unban` - Usuario desbaneado
```json
{
  "seq": 726,
  "cmd": "unban",
  "by": "Bot",
  "nick": "*!*@d5k.hvo.2dc8a2.IP",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "ctime": "1539536029001"
}
```

### Configuración y Sistema

#### `settings` - Configuración del cliente
```json
{
  "seq": 28,
  "cmd": "settings",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "settings": {
    "userListShowIcons": true,
    "userListSort": "nick",
    "nick": "MiUsuario",
    "channel": "#Argentina",
    "server": "irc.dalechatea.me",
    "port": "+6697",
    "ssl": true,
    "webircGateway": "https://chatarg.com/webchat/"
  }
}
```

#### `motd` - Mensaje del día
```json
{
  "seq": 27,
  "cmd": "motd",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat",
  "motd": ["Línea 1", "Línea 2", "Línea 3"]
}
```

#### `notice` - Notificación del servidor
```json
{
  "seq": 5,
  "cmd": "notice",
  "nick": "",
  "notice": "Current local users: 1429 Max: 1956",
  "ctime": "1539536029001",
  "channel": "IRCClient:irc.dalechatea.me"
}
```

#### `servertext` - Texto del servidor
```json
{
  "seq": 26,
  "cmd": "servertext",
  "text": "Now talking on #Argentina",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat"
}
```

## Mensajes del Cliente (Enviados)

### Enviar Mensaje de Texto
```json
{
  "seq": [SIGUIENTE_SEQ],
  "cmd": "text",
  "chan": "#Argentina",
  "data": "\\u00031Hola mundo",
  "channel": "IRCClient:irc.dalechatea.me"
}
```

### Enviar Respuesta a Mensaje
```json
{
  "seq": [SIGUIENTE_SEQ],
  "cmd": "text",
  "chan": "#Argentina",
  "data": "\\u00031Mi respuesta",
  "reply": " usuario: mensaje original",
  "channel": "IRCClient:irc.dalechatea.me"
}
```

### Listar Información del Canal
```json
{
  "seq": [SIGUIENTE_SEQ],
  "cmd": "l",
  "localchannel": "#Argentina",
  "channel": "IRCClient:irc.dalechatea.me:#Argentina",
  "channeltype": "chat"
}
```

### Enviar Mensaje Simple (alternativo)
```json
{
  "seq": [SIGUIENTE_SEQ],
  "cmd": "s",
  "channel": "#Argentina",
  "text": "hola"
}
```

## Canales Disponibles
- `#Argentina` - Canal principal
- `#30argentina` - Usuario de 30+ años
- `#40argentina` - Usuario de 40+ años  
- `#50argentina` - Usuario de 50+ años
- `#LigoteoArgentina` - Canal de citas
- `#TransArgentina` - Canal trans
- `#SexoArgentina` - Canal adulto
- `#MasturbacionArgentina` - Canal adulto
- `#CamArgentina` - Canal de webcam

## Consideraciones de Implementación

### Manejo de Errores
- Implementar reconnection automática
- Validar formato JSON de mensajes
- Manejar desconexiones inesperadas

### Formato de Texto
- Los mensajes pueden contener códigos de formato IRC (`\\u0002`, `\\u00032`, etc.)
- El prefijo `\\u00031` es común en mensajes de texto

### Timestamps
- `ctime`: timestamp en milisegundos
- Algunos campos de fecha están en formato Unix timestamp

### Seguridad
- Nunca incluir credenciales en el código
- Validar todos los inputs del usuario
- Sanitizar mensajes antes de mostrar

### Performance
- Implementar buffer para mensajes recibidos
- Controlar rate limiting para envío de mensajes
- Usar goroutines para operaciones no bloqueantes

## Ejemplo de Implementación Básica en Go

```go
// Estructura básica para manejar mensajes
type Message struct {
    Seq     int    `json:"seq"`
    Cmd     string `json:"cmd"`
    Channel string `json:"channel,omitempty"`
    Nick    string `json:"nick,omitempty"`
    Data    string `json:"data,omitempty"`
    Text    string `json:"text,omitempty"`
}

// Headers de conexión
headers := http.Header{
    "Origin":     []string{"https://chatarg.com"},
    "User-Agent": []string{"Mozilla/5.0..."},
}

// Conectar
conn, _, err := websocket.DefaultDialer.Dial("ws://wss.dalechatea.me:1245/", headers)
```

Esta documentación proporciona toda la información necesaria para implementar un cliente WebSocket completo para el sistema de chat IRC de Chat Argentina.