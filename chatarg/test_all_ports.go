package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Puertos abiertos según el escaneo
var openPorts = []int{
	21, 25, 53, 80, 110, 111, 143, 443, 465, 587, 993, 995,
	1239, 1240, 1241, 1242, 1245,
	2077, 2078, 2082, 2083, 2086, 2087, 2091, 2095, 2096,
	7001, 7003, 7004, 7005,
	55642,
}

func testWebSocketPort(host string, port int, useWSS bool) {
	protocol := "ws"
	if useWSS {
		protocol = "wss"
	}

	wsURL := fmt.Sprintf("%s://%s:%d/", protocol, host, port)
	fmt.Printf("\n🔌 Probando: %s\n", wsURL)

	// Headers mínimos necesarios
	headers := map[string][]string{
		"Origin":     {"https://chatarg.com"},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
	}

	// Configurar dialer con timeout
	dialer := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Para pruebas, acepta cualquier certificado
		},
	}

	// Intentar conexión
	conn, resp, err := dialer.Dial(wsURL, headers)
	if err != nil {
		if resp != nil {
			fmt.Printf("❌ Error: HTTP %d - %v\n", resp.StatusCode, err)
		} else if strings.Contains(err.Error(), "timeout") {
			fmt.Printf("⏱️ Timeout en puerto %d\n", port)
		} else if strings.Contains(err.Error(), "connection refused") {
			fmt.Printf("🚫 Conexión rechazada en puerto %d\n", port)
		} else {
			fmt.Printf("❌ Error: %v\n", err)
		}
		return
	}
	defer conn.Close()

	fmt.Printf("✅ ¡ÉXITO! Conexión WebSocket establecida en puerto %d\n", port)

	// Intentar leer un mensaje inicial
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

	// Cerrar conexión limpiamente
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	fmt.Printf("🎯 Puerto %d: SERVIDOR WEBSOCKET ACTIVO\n", port)
}

func main4() {
	host := "wss.dalechatea.me"

	fmt.Println("🚀 Escáner de WebSocket para Chat Argentina")
	fmt.Println("==========================================")
	fmt.Printf("🎯 Servidor objetivo: %s\n", host)
	fmt.Printf("📊 Total de puertos a probar: %d\n", len(openPorts))
	fmt.Println("\n⚡ Iniciando pruebas de conexión WebSocket...")
	fmt.Println("Probando tanto WS como WSS en cada puerto...")

	successfulPorts := []string{}

	for i, port := range openPorts {
		fmt.Printf("\n=== [%d/%d] Puerto %d ===", i+1, len(openPorts), port)

		// Probar con WSS (seguro)
		testWebSocketPort(host, port, true)

		// Si no es un puerto típicamente seguro, probar también con WS
		if port != 443 && port != 1239 && port != 2083 && port != 2087 {
			time.Sleep(500 * time.Millisecond) // Pequeña pausa entre pruebas
			testWebSocketPort(host, port, false)
		}

		// Pausa entre puertos para evitar sobrecarga
		if i < len(openPorts)-1 {
			time.Sleep(time.Second)
		}
	}

	// Resumen final
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
		fmt.Println("💡 Nota: El puerto 1239 está abierto pero tiene restricciones de seguridad")
	}

	fmt.Println("\n💡 Puertos conocidos:")
	fmt.Println("   - Puerto 443: Servidor público funcional (ws03.dalechatea.me)")
	fmt.Println("   - Puerto 1239: Backend con restricciones de seguridad")
}
