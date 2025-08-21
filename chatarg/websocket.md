# Documentación WebSocket Inspircd

Esta documentación describe los parámetros de conexión WebSocket y ejemplos de comandos (`cmd`) basados en el archivo HAR proporcionado para un servidor Inspircd.

## Parámetros de Conexión WebSocket

* **URI:** `wss://chatarg.com/webchat/`

* **Encabezados (Headers):**
    * `Host`: `chatarg.com`
    * `Connection`: `Upgrade`
    * `Pragma`: `no-cache`
    * `Cache-Control`: `no-cache`
    * `User-Agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0`
    * `Upgrade`: `websocket`
    * `Origin`: `https://chatarg.com`
    * `Sec-WebSocket-Version`: `13`
    * `Accept-Encoding`: `gzip, deflate, br, zstd`
    * `Accept-Language`: `es,es-ES;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6`
    * `Sec-WebSocket-Key`: `YOUR_GENERATED_KEY` (Este valor es dinámico y debe ser generado por el cliente)
    * `Sec-WebSocket-Extensions`: `permessage-deflate; client_max_window_bits`

## Comandos de Mensajes WebSocket Recibidos (`cmd`)

Aquí tienes un ejemplo de cada comando (`cmd`) único encontrado en los mensajes WebSocket recibidos del servidor:

* `status`: Mensaje de estado del servidor.
    ```json
    {"seq":1,"pendingDNS":1,"cmd":"status","pending":0,"hostname":"186.33.238.108","connections":0,"channel":"IRCClient","readyToConnect":true}
    ```

* `nicklist`: Lista de nicks en un canal.
    ```json
    {"seq":52,"cmd":"nicklist","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","nicks":[{"host":"pju9cr.9lo2.8070.g0ad65.IP","nick":"leonnel"},{"host":"g71.p33.jemp3k.IP","nick":"Guest29042C"},{"host":"3kr.7hl.9tl5rt.IP","nick":"eliot29"},{"host":"j5k.1dr.4fgvf5.IP","nick":"Daianasoledad34"},{"host":"ck5.u9m.unge1s.IP","nick":"joshepluis"},{"host":"ln2r28.cokb.kjks.g0ad65.IP","nick":"__Maxxii31"},{"host":"a15vae.mvmd.o779.g0ad65.IP","nick":"Lucas_39CBA"},{"host":"gogevj.jisr.8p0b.4eb7q5.IP","nick":"Santiago39"},{"host":"vob.1jb.00l87o.IP","nick":"Fed3"},{"host":"5hd.qup.sb5ljh.IP","nick":"Silvina44"},{"host":"k9e.tcd.9dc2nb.IP","nick":"al_rincon"},{"host":"2f9.2sa.67sdfm.IP","nick":"_PabloQQ"},{"host":"s76.bbh.n4b5fh.IP","nick":"Oscar6G"},{"host":"2gl.vpv.uaoa15.IP","nick":"daniell_"}],"channeltype":"chat"}
    ```

* `part`: Un usuario sale de un canal.
    ```json
    {"seq":64,"cmd":"part","nick":"MilagrosArg","user":"DaleChat","host":"4v6.ji6.7ksean.IP","quit":true,"message":"Salir: Desconectado ([https://chatarg.com](https://chatarg.com))","channel":"IRCClient:irc.dalechatea.me","ctime":"1539536029001"}
    ```

* `userinfo`: Información detallada de un usuario.
    ```json
    {"seq":65,"cmd":"userinfo","nick":"Metalhera","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ct":1747153164083,"info":"{\\\"userIcon\\\":\\\"[https://a.dalechatea.me/7894029744508564997462267.jpg](https://a.dalechatea.me/7894029744508564997462267.jpg)\\\"}","cc":"AR","tz":"-180"}
    ```

* `l`: (Probablemente 'list') Solicita información de canales (usualmente enviado por el cliente, pero este es un ejemplo de respuesta).
    ```json
    {"seq":1,"cmd":"l","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat"}
    ```

* `s`: (Probablemente 'send') Envía un mensaje a un canal (usualmente enviado por el cliente, este es un ejemplo de cómo se vería si se recibiera de vuelta, aunque la mayoría de las veces no se reenvía al propio cliente).
    ```json
    {"seq":2,"cmd":"s","channel":"#Argentina","text":"hola"}
    ```

* `usericon`: Actualiza el icono de usuario de un usuario.
    ```json
    {"seq":3,"cmd":"usericon","nick":"MiUsuario","channel":"IRCClient:irc.dalechatea.me:#Argentina","icon":"[https://dalechatea.me/images/JMMWhtk.png](https://dalechatea.me/images/JMMWhtk.png)","channeltype":"chat","localchannel":"#Argentina"}
    ```

* `usersetting`: Actualiza una configuración de usuario.
    ```json
    {"seq":4,"cmd":"usersetting","nick":"MiUsuario","channel":"IRCClient:irc.dalechatea.me:#Argentina","setting":"away","value":"Ausente","channeltype":"chat","localchannel":"#Argentina"}
    ```

* `topic`: Muestra o establece el tema de un canal.
    ```json
    {"seq":25,"cmd":"topic","channel":"IRCClient:irc.dalechatea.me:#Argentina","nick":"ChanServ","topic":"Bienvenidos al chat de Argentina »» RECORDAR LEER LAS NORMAS: [https://chatarg.com/normas/](https://chatarg.com/normas/) GRACIAS! »» SALAS: #Argentina #30argentina #40argentina #50argentina #LigoteoArgentina #TransArgentina #SexoArgentina #MasturbacionArgentina #CamArgentina","time":1747134282,"localchannel":"#Argentina","channeltype":"chat"}
    ```

* `servertext`: Mensaje de texto del servidor.
    ```json
    {"seq":26,"cmd":"servertext","text":"Now talking on #Argentina","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat"}
    ```

* `motd`: Mensaje del día del servidor.
    ```json
    {"seq":27,"cmd":"motd","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","motd":["     .-.           . . .-\",\"     | |.-..----. |-  | |-","     `-|-|  ||  .|| | | |-`","      `-' `|__||__|' ' ' '-"]}
    ```

* `settings`: Configuración del cliente/usuario.
    ```json
    {"seq":28,"cmd":"settings","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","settings":{"userListShowIcons":true,"userListSort":"nick","userListDirection":"asc","userListShowOffline":false,"f":"0","awayTime":0,"activity":0,"lastClick":0,"typingAwayAfter":300000,"typingActivityAfter":2000,"typingSend":"/me is typing...","typingStop":"/me is no longer typing.","userListGroupAway":false,"userListGroupOps":true,"userListGroupVoiced":true,"userListColorizeNicks":true,"noInvite":false,"nick":"MiUsuario","channel":"#Argentina","server":"irc.dalechatea.me","port":"+6697","ssl":true,"webircGateway":"[https://chatarg.com/webchat/](https://chatarg.com/webchat/)","chatId":"irc.dalechatea.me:#Argentina","userIcon":"dalechatea.me/images/JMMWhtk.png"}}
    ```

* `notice`: Mensaje de notificación del servidor.
    ```json
    {"seq":5,"cmd":"notice","nick":"","notice":"Current local users: 1429 Max: 1956","ctime":"1539536029001","channel":"IRCClient:irc.dalechatea.me"}
    ```

* `join`: Un usuario se une a un canal.
    ```json
    {"seq":9,"cmd":"join","user":"DaleChat","host":"b3v.np7.k217c4.IP","nick":"zero123x","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ctime":"1539536029001"}
    ```

* `connected`: Confirmación de conexión al servidor.
    ```json
    {"seq":2,"cmd":"connected","channel":"IRCClient","name":"irc.dalechatea.me","network":"irc.dalechatea.me"}
    ```

* `init`: Inicialización del cliente.
    ```json
    {"seq":4,"cmd":"init","nick":"zero123x","channel":"IRCClient:irc.dalechatea.me","ctime":"39536029001"}
    ```

* `mode`: Cambios en los modos de usuario o canal.
    ```json
    {"seq":7,"cmd":"mode","nick":"zero123x","msg":"zero123x estableció los modos +xiwcT","ctime":"1539536029001","channel":"IRCClient:irc.dalechatea.me"}
    ```

* `msg`: Un mensaje de chat enviado a un canal o usuario.
    ```json
    {"seq":171,"cmd":"msg","nick":"El-Mas-Dulce","localchannel":"#Argentina","channeltype":"chat","channel":"IRCClient:irc.dalechatea.me:#Argentina","msg":"\u0002\u00032Quiere un abrazo de oso capas. Jajajaja.","replyTo":" solangee: Osito"}
    ```

* `typing`: Indica si un usuario está escribiendo.
    ```json
    {"seq":689,"cmd":"typing","nick":"Metalero53","channel":"IRCClient:irc.dalechatea.me:#Argentina","typing":false}
    ```

* `emote`: Un mensaje de acción (emote).
    ```json
    {"seq":696,"cmd":"emote","ctime":"1539536029001","nick":"Azimut_","localchannel":"#Argentina","channeltype":"chat","channel":"IRCClient:irc.dalechatea.me:#Argentina","emote":":kiss Azura"}
    ```

* `avatar`: Información o cambio del avatar de un usuario.
    ```json
    {"seq":723,"cmd":"avatar","nick":"Georgina38","localchannel":"#argentina","channel":"IRCClient:irc.dalechatea.me:#argentina","channeltype":"chat","userIcon":"[https://a.dalechatea.me/5447788000559999991118822.jpg](https://a.dalechatea.me/5447788000559999991118822.jpg)"}
    ```

* `unban`: Un usuario o máscara es desbaneado de un canal.
    ```json
    {"seq":726,"cmd":"unban","by":"Bot","nick":"*!*@d5k.hvo.2dc8a2.IP","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ctime":"1539536029001"}
    ```

* `kick`: Un usuario es expulsado de un canal.
    ```json
    {"seq":829,"cmd":"kick","nick":"Gerardo7","kicker":"Bot","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ctime":"1539536029001","reason":"Requested (Matches *!*@54g.4l0.4serdp.IP) (_JoelDelOeste)"}
    ```

* `ban`: Un usuario o máscara es baneado de un canal.
    ```json
    {"seq":830,"cmd":"ban","by":"Bot","nick":"*!*@54g.4l0.4serdp.IP","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ctime":"1539536029001"}
    ```

* `removechat`: Indica la remoción de un chat (posiblemente un chat privado).
    ```json
    {"seq":840,"cmd":"removechat","chatid":"gerardo7","channel":"IRCClient:irc.dalechatea.me:#argentina"}
    ```

* `changenick`: Un usuario cambia su nick.
    ```json
    {"seq":901,"cmd":"changenick","nick":"Sergio","newnick":"Guest22918","channel":"IRCClient:irc.dalechatea.me","ctime":"1539536029001"}
    ```

* `away`: Un usuario cambia su estado a "ausente".
    ```json
    {"seq":1079,"cmd":"away","nick":"Briana_30","channel":"IRCClient:irc.dalechatea.me"}
    ```

* `usermode`: Cambios en los modos de un usuario específico.
    ```json
    {"seq":1983,"cmd":"usermode","op":true,"by":"Bot","nick":"Clare","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","ctime":"1539536029001"}
    ```
* `topicwho`: Información sobre quién estableció el tema de un canal y cuándo.
    ```json
    {"seq":12,"cmd":"topicwho","localchannel":"#Argentina","channel":"IRCClient:irc.dalechatea.me:#Argentina","channeltype":"chat","creator":"ElGranArdilla","date":"1657870309"}
    ```