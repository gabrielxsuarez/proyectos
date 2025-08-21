package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)


type ChatClient struct {
	conn           *websocket.Conn
	sendSeq        int
	receiveSeq     int
	nick           string
	channel        string
	sessionID      string
	connected      bool
	userIP         string
	seenCommands   map[string]bool
	wsURL          string
	debugMode      bool
}

type NickInput struct {
	nick        string
	channel     string
	testAllPort bool
	testServer  string
	debugMode   bool
}

func main() {
	fmt.Println("🚀 Cliente WebSocket para Chat Argentina")
	fmt.Println("==========================================")

	for {
		client := &ChatClient{}

		nickInput := client.requestNick()
		if nickInput == nil {
			fmt.Println("❌ Nick requerido para continuar")
			continue
		}

		// Si es test all port, ejecutar testing y salir
		if nickInput.testAllPort {
			server := nickInput.testServer
			if server == "" {
				// Cargar servidor desde config
				if !client.loadConfig() {
					fmt.Println("❌ Error cargando configuración para obtener servidor")
					continue
				}
				server = client.extractServerFromURL()
			}
			testAllPorts(server)
			continue
		}

		// Configurar nick, canal y modo debug
		client.nick = nickInput.nick
		client.debugMode = nickInput.debugMode
		if nickInput.channel != "" {
			client.channel = nickInput.channel
		}

		if !client.loadConfig() {
			fmt.Println("❌ Error cargando configuración")
			continue
		}

		// Si se especificó canal en el nick, sobrescribir el del config
		if nickInput.channel != "" {
			client.channel = nickInput.channel
		}

		// Mostrar indicador de modo debug si está activo
		if client.debugMode {
			fmt.Println("🐛 MODO DEBUG ACTIVADO - Se mostrarán todos los mensajes del servidor")
			fmt.Println("==========================================")
		}

		client.selectRandomIP()

		if client.connect() {
			client.handleChat()
		}

		client.disconnect()
		fmt.Println("\n🔄 Reiniciando...")
	}
}

func (c *ChatClient) requestNick() *NickInput {
	fmt.Print("📝 Ingrese su nick: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("❌ Error leyendo nick: %v\n", err)
		return nil
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("❌ El nick no puede estar vacío")
		return nil
	}

	return parseNickInput(input)
}

func (c *ChatClient) loadConfig() bool {
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Printf("❌ Error leyendo config.txt: %v\n", err)
		return false
	}

	content := strings.TrimSpace(string(data))
	lines := strings.Split(content, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		
		switch key {
		case "sala":
			if c.channel == "" { // Solo usar si no se especificó canal en nick
				c.channel = "#" + value
			}
		case "servidor":
			// Si el valor ya tiene protocolo, usarlo tal cual
			if strings.HasPrefix(value, "ws://") || strings.HasPrefix(value, "wss://") {
				c.wsURL = value
				// Agregar slash final si no lo tiene
				if !strings.HasSuffix(c.wsURL, "/") {
					c.wsURL += "/"
				}
			} else {
				// Si no tiene protocolo, usar ws:// por defecto
				c.wsURL = "ws://" + value + "/"
			}
		}
	}
	
	if c.channel == "" || c.wsURL == "" {
		fmt.Println("❌ config.txt debe contener 'sala' y 'servidor'")
		return false
	}
	
	fmt.Printf("📂 Canal: %s\n", c.channel)
	fmt.Printf("🌐 Servidor: %s\n", c.wsURL)
	return true
}

func (c *ChatClient) selectRandomIP() {
	data, err := ioutil.ReadFile("ip.txt")
	if err != nil {
		fmt.Printf("⚠️ Error leyendo ip.txt: %v. Usando IP por defecto.\n", err)
		c.userIP = "200.223.222.222"
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) == 0 {
		c.userIP = "200.223.222.222"
		return
	}

	rand.Seed(time.Now().UnixNano())
	selectedIP := strings.TrimSpace(lines[rand.Intn(len(lines))])
	if selectedIP == "" {
		c.userIP = "200.223.222.222"
	} else {
		c.userIP = selectedIP
	}

	fmt.Printf("🌐 IP seleccionada: %s\n", c.userIP)
}

func (c *ChatClient) connect() bool {
	headers := map[string][]string{
		"Origin":          {"https://chatarg.com"},
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
		"X-Forwarded-For": {c.userIP},
	}

	fmt.Printf("📡 Conectando a: %s\n", c.wsURL)

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.Dial(c.wsURL, headers)
	if err != nil {
		fmt.Printf("❌ Error conectando: %v\n", err)
		return false
	}

	c.conn = conn
	c.sendSeq = 0
	c.receiveSeq = 0
	c.connected = false
	c.seenCommands = make(map[string]bool)

	fmt.Println("✅ Conexión WebSocket establecida")

	go c.readMessages()

	if !c.waitForSessionID() {
		return false
	}

	if !c.initializeConnection() {
		return false
	}

	return true
}

func (c *ChatClient) waitForSessionID() bool {
	timeout := time.NewTimer(10 * time.Second)
	defer timeout.Stop()

	for {
		select {
		case <-timeout.C:
			fmt.Println("❌ Timeout esperando sessionid")
			return false
		default:
			if c.sessionID != "" {
				fmt.Printf("🔑 SessionID recibido: %s\n", c.sessionID)
				return true
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (c *ChatClient) initializeConnection() bool {
	fmt.Println("🔄 Inicializando conexión...")

	c.sendSeq = 0
	if !c.sendClientInfo() {
		return false
	}

	c.sendSeq++
	if !c.sendEmbedInfo() {
		return false
	}

	c.sendSeq++
	if !c.sendConnectIRC() {
		return false
	}

	fmt.Println("⏳ Esperando confirmación de conexión al canal...")

	timeout := time.NewTimer(15 * time.Second)
	defer timeout.Stop()

	for {
		select {
		case <-timeout.C:
			fmt.Println("❌ Timeout esperando conexión al canal")
			return false
		default:
			if c.connected {
				fmt.Printf("✅ Conectado al canal %s como %s\n", c.channel, c.nick)
				return true
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (c *ChatClient) sendClientInfo() bool {
	msg := ClientInfo{
		Seq:        c.sendSeq,
		Cmd:        "clientinfo",
		LocalTime:  time.Now().UnixNano() / int64(time.Millisecond),
		TzOffset:   -180,
		ClientInfo: `{"userIcon":"https://dalechatea.me/images/aflugbx.png"}`,
	}

	return c.sendMessage(msg)
}

func (c *ChatClient) sendEmbedInfo() bool {
	msg := EmbedInfo{
		Seq:      c.sendSeq,
		Channel:  "IRCClient",
		Cmd:      "embed",
		Referrer: "https://chatarg.com/webchat/",
	}

	return c.sendMessage(msg)
}

func (c *ChatClient) sendConnectIRC() bool {
	msg := ConnectIRC{
		Seq:          c.sendSeq,
		Cmd:          "connect",
		Channel:      "IRCClient",
		Data:         "irc.dalechatea.me:+6697",
		Nick:         c.nick,
		Pass:         "",
		AuthMethod:   "nickserv",
		JoinChannels: c.channel,
		Charset:      "utf-8",
	}

	return c.sendMessage(msg)
}

func (c *ChatClient) sendMessage(msg interface{}) bool {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("❌ Error serializando mensaje: %v\n", err)
		return false
	}

	// En modo debug, imprimir todos los mensajes enviados
	if c.debugMode {
		fmt.Printf("[DEBUG] ⬆️ ENVIANDO: %s\n", string(data))
	}

	err = c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		fmt.Printf("❌ Error enviando mensaje: %v\n", err)
		return false
	}

	return true
}

func (c *ChatClient) readMessages() {
	for {
		if c.conn == nil {
			return
		}

		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("❌ Error leyendo mensaje: %v\n", err)
			}
			return
		}

		if messageType == websocket.TextMessage {
			c.handleMessage(message)
		}
	}
}

func (c *ChatClient) handleMessage(data []byte) {
	// En modo debug, imprimir todos los mensajes recibidos
	if c.debugMode {
		fmt.Printf("[DEBUG] ⬇️ RECIBIDO: %s\n", string(data))
	}

	var generic GenericMessage
	if err := json.Unmarshal(data, &generic); err != nil {
		return
	}

	c.saveMessageExample(generic.Cmd, data)

	switch generic.Cmd {
	case "":
		var sessionMsg SessionID
		if err := json.Unmarshal(data, &sessionMsg); err == nil {
			c.sessionID = sessionMsg.SessionID
		}

	case "connected":
		// Conectado silenciosamente

	case "init":
		c.connected = true

	case "msg":
		// Mensaje recibido silenciosamente

	case "join":
		// Join silencioso

	case "part":
		// Part silencioso

	case "topic":
		// Topic silencioso

	case "notice":
		// Notice silencioso

	default:
		// Otros mensajes silenciosos
	}
}

func (c *ChatClient) saveMessageExample(cmd string, data []byte) {
	if cmd == "" {
		cmd = "sessionid"
	}
	
	if !c.seenCommands[cmd] {
		c.seenCommands[cmd] = true
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		example := fmt.Sprintf("[%s] CMD_%s: %s\n", timestamp, cmd, string(data))
		
		file, err := os.OpenFile("ejemplos.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			file.WriteString(example)
			file.Close()
		}
	}
}

func (c *ChatClient) cleanIRCFormat(text string) string {
	replacements := map[string]string{
		"\\u0002":  "",
		"\\u00032": "",
		"\\u00031": "",
		"\\u000f":  "",
	}

	for old, new := range replacements {
		text = strings.ReplaceAll(text, old, new)
	}

	return text
}

func (c *ChatClient) SendTexto(texto string) {
	c.sendSeq++
	msg := SendText{
		Seq:     c.sendSeq,
		Cmd:     "text",
		Chan:    c.channel,
		Data:    "\u00031" + texto,
		Channel: "IRCClient:irc.dalechatea.me",
	}

	c.sendMessage(msg)
}

func (c *ChatClient) SendTextoWithReply(texto, replica string) {
	c.sendSeq++
	msg := SendText{
		Seq:     c.sendSeq,
		Cmd:     "text",
		Chan:    c.channel,
		Data:    "\u00031" + texto,
		Channel: "IRCClient:irc.dalechatea.me",
		Reply:   " " + replica,
	}

	c.sendMessage(msg)
}

func (c *ChatClient) handleChat() {
	fmt.Println("\n💬 ¡Chat iniciado! Escriba sus mensajes (o 'q' para desconectar)")
	fmt.Println("💡 Formato con respuesta: 'texto de respuesta | mensaje'")
	fmt.Println("========================================")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	reader := bufio.NewReader(os.Stdin)

	go c.Mping()

	for {
		select {
		case <-interrupt:
			fmt.Println("\n🛑 Desconectando...")
			return
		default:
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("❌ Error leyendo entrada: %v\n", err)
				continue
			}

			input = strings.TrimSpace(input)

			if input == "q" {
				fmt.Println("🛑 Desconectando por solicitud del usuario...")
				return
			}

			if input == "" {
				continue
			}

			if strings.Contains(input, "|") {
				parts := strings.Split(input, "|")
				if len(parts) >= 2 {
					lastPipeIndex := strings.LastIndex(input, "|")
					reply := strings.TrimSpace(input[:lastPipeIndex])
					message := strings.TrimSpace(input[lastPipeIndex+1:])

					if message != "" {
						c.SendTextoWithReply(message, reply)
						// Enviado silenciosamente
					}
				}
			} else {
				c.SendTexto(input)
				// Enviado silenciosamente
			}
		}
	}
}

func (c *ChatClient) Mping() {
	for c.conn != nil {
		time.Sleep(30 * time.Second)
		if c.conn != nil {
			c.sendSeq++
			keepAlive := KeepAlive{
				Seq: c.sendSeq,
				Cmd: "nping",
				M:   "checking",
			}
			c.sendMessage(keepAlive)
		}
	}
}

func (c *ChatClient) disconnect() {
	if c.conn != nil {
		err := c.conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		if err != nil {
			// Error silencioso
		}

		c.conn.Close()
		c.conn = nil
	}
}

// parseNickInput analiza la entrada del nick y determina el modo de operación
func parseNickInput(input string) *NickInput {
	// Detectar "test all port" con servidor opcional
	if strings.HasPrefix(input, "test all port") {
		parts := strings.Fields(input)
		result := &NickInput{
			testAllPort: true,
		}
		if len(parts) > 3 {
			result.testServer = parts[3]
		}
		return result
	}
	
	// Detectar formato "#debug Nick" para modo debug
	if strings.HasPrefix(input, "#debug ") {
		nickPart := strings.TrimPrefix(input, "#debug ")
		return &NickInput{
			nick:      strings.TrimSpace(nickPart),
			debugMode: true,
		}
	}
	
	// Detectar formato "#Canal Nick"
	if strings.HasPrefix(input, "#") {
		parts := strings.Fields(input)
		if len(parts) >= 2 {
			return &NickInput{
				channel: parts[0],
				nick:    strings.Join(parts[1:], " "),
			}
		}
	}
	
	// Nick simple
	return &NickInput{
		nick: input,
	}
}

// extractServerFromURL extrae el servidor del wsURL configurado
func (c *ChatClient) extractServerFromURL() string {
	if c.wsURL == "" {
		return "wss.dalechatea.me"
	}
	
	// Remover protocolo y trailing slash
	server := strings.TrimPrefix(c.wsURL, "ws://")
	server = strings.TrimPrefix(server, "wss://")
	server = strings.TrimSuffix(server, "/")
	
	// Remover puerto si existe para obtener solo el host
	if strings.Contains(server, ":") {
		parts := strings.Split(server, ":")
		return parts[0]
	}
	
	return server
}

// readPortsFromFile lee la lista de puertos desde puertos.txt
func readPortsFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ports []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		port, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("⚠️ Línea ignorada (no es un puerto válido): %s\n", line)
			continue
		}
		ports = append(ports, port)
	}
	return ports, scanner.Err()
}

// testWebSocketPort prueba la conectividad WebSocket en un puerto específico
func testWebSocketPort(host string, port int, useWSS bool) bool {
	protocol := "ws"
	if useWSS {
		protocol = "wss"
	}

	wsURL := fmt.Sprintf("%s://%s:%d/", protocol, host, port)
	fmt.Printf("\n🔌 Probando: %s\n", wsURL)

	headers := map[string][]string{
		"Origin":     {"https://chatarg.com"},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	conn, resp, err := dialer.Dial(wsURL, headers)
	if err != nil {
		if resp != nil {
			fmt.Printf("❌ Error: HTTP %d - %v\n", resp.StatusCode, err)
		} else if strings.Contains(err.Error(), "timeout") {
			fmt.Printf("⏱️ Timeout en puerto %d\n", port)
		} else if strings.Contains(err.Error(), "connection refused") {
			fmt.Printf("🙫 Conexión rechazada en puerto %d\n", port)
		} else {
			fmt.Printf("❌ Error: %v\n", err)
		}
		return false
	}
	defer conn.Close()

	fmt.Printf("✅ ¡ÉXITO! Conexión WebSocket establecida en puerto %d\n", port)

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	messageType, message, err := conn.ReadMessage()
	if err == nil {
		if messageType == websocket.TextMessage {
			fmt.Printf("📨 Mensaje recibido: %s\n", string(message))
		} else {
			fmt.Printf("📦 Mensaje binario recibido (tipo %d)\n", messageType)
		}
	} else if !strings.Contains(err.Error(), "timeout") {
		fmt.Printf("⚠️ Error leyendo mensaje: %v\n", err)
	}

	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	fmt.Printf("🎯 Puerto %d: SERVIDOR WEBSOCKET ACTIVO\n", port)
	return true
}

// testAllPorts ejecuta el testing de puertos contra un servidor
func testAllPorts(host string) {
	fmt.Println("🚀 Escáner de WebSocket para Chat Argentina")
	fmt.Println("==========================================")
	fmt.Printf("🎯 Servidor objetivo: %s\n", host)

	ports, err := readPortsFromFile("puertos.txt")
	if err != nil {
		fmt.Printf("❌ Error leyendo archivo puertos.txt: %v\n", err)
		return
	}

	fmt.Printf("📊 Total de puertos a probar: %d\n", len(ports))
	fmt.Println("\n⚡ Iniciando pruebas de conexión WebSocket...")

	successfulPorts := []string{}

	for i, port := range ports {
		fmt.Printf("\n=== [%d/%d] Puerto %d ===", i+1, len(ports), port)

		if testWebSocketPort(host, port, true) {
			successfulPorts = append(successfulPorts, fmt.Sprintf("wss://%s:%d/", host, port))
		}

		if port != 443 && port != 1239 && port != 2083 && port != 2087 {
			time.Sleep(500 * time.Millisecond)
			if testWebSocketPort(host, port, false) {
				successfulPorts = append(successfulPorts, fmt.Sprintf("ws://%s:%d/", host, port))
			}
		}

		if i < len(ports)-1 {
			time.Sleep(time.Second)
		}
	}

	fmt.Println("\n\n========================================")
	fmt.Println("📋 RESUMEN DE RESULTADOS")
	fmt.Println("========================================")

	if len(successfulPorts) > 0 {
		fmt.Printf("✅ Puertos con WebSocket activo:\n")
		for _, p := range successfulPorts {
			fmt.Printf("   - %s\n", p)
		}
	} else {
		fmt.Println("⚠️ No se encontraron servidores WebSocket activos")
	}

	fmt.Println("\n💡 Testing completado.")
}