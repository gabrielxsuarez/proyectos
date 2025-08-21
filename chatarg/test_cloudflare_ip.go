package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// IPs conocidas de Cloudflare
var cloudflareIPs = []string{
	"104.16.132.229", // cloudflare.com
	"104.16.133.229", // cloudflare.com
	"1.1.1.1",        // DNS Cloudflare
	"1.0.0.1",        // DNS Cloudflare
	"104.26.14.97",   // IP vista en los logs del chat
	"172.64.0.1",     // Rango típico de Cloudflare
	"198.41.128.1",   // Rango típico de Cloudflare
}

// Servidores a probar
var testServers = []string{
	"wss://ws03.dalechatea.me/",     // Servidor 1 (funciona)
	"wss://wss.dalechatea.me:1239/", // Servidor 2 (a probar)
	"wss://51.161.13.145:1239/",     // Servidor 2 por IP directa
}

func testWebSocketConnection(wsURL, cloudflareIP string) error {
	fmt.Printf("\n🔌 Probando: %s\n", wsURL)
	fmt.Printf("📡 Simulando desde IP Cloudflare: %s\n", cloudflareIP)

	// Configurar headers simulando Cloudflare
	headers := http.Header{}
	headers.Set("User-Agent", "Mozilla/5.0 (compatible; CloudflareBot/1.0)")
	headers.Set("Origin", "https://chatarg.com")

	// Headers críticos de Cloudflare - simular que venimos DESDE Cloudflare
	headers.Set("X-Forwarded-For", "200.223.222.222, "+cloudflareIP) // Usuario + IP de CF
	headers.Set("X-Real-IP", cloudflareIP)                           // IP real = Cloudflare
	headers.Set("CF-Connecting-IP", "200.223.222.222")               // IP del usuario final
	headers.Set("CF-Ray", "8a1b2c3d4e5f6789-EZE")
	headers.Set("CF-IPCountry", "AR")
	headers.Set("CF-Visitor", `{"scheme":"https"}`)
	headers.Set("CDN-Loop", "cloudflare")
	headers.Set("X-Forwarded-Proto", "https")

	// Headers adicionales que Cloudflare suele agregar
	headers.Set("X-Forwarded-Host", "chatarg.com")
	headers.Set("X-Forwarded-Server", "ws03.dalechatea.me")
	headers.Set("Via", "1.1 cloudflare")

	// Si es conexión por IP, configurar TLS
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	if strings.Contains(wsURL, "51.161.13.145") {
		dialer.TLSClientConfig = &tls.Config{
			ServerName: "wss.dalechatea.me",
		}
		headers.Set("Host", "wss.dalechatea.me")
	}

	// Intentar conexión
	conn, resp, err := dialer.Dial(wsURL, headers)
	if err != nil {
		if resp != nil {
			fmt.Printf("❌ Error: %v (Status: %d)\n", err, resp.StatusCode)
			// Mostrar algunos headers de respuesta
			for key, values := range resp.Header {
				if strings.HasPrefix(strings.ToLower(key), "cf-") ||
					strings.HasPrefix(strings.ToLower(key), "x-") ||
					key == "Server" {
					for _, value := range values {
						fmt.Printf("   %s: %s\n", key, value)
					}
				}
			}
		} else {
			fmt.Printf("❌ Error: %v\n", err)
		}
		return err
	}
	defer conn.Close()

	fmt.Println("✅ ¡ÉXITO! Conexión establecida")

	// Leer mensajes iniciales
	messageCount := 0
	timeout := time.After(5 * time.Second)

readLoop:
	for {
		select {
		case <-timeout:
			break readLoop
		default:
			conn.SetReadDeadline(time.Now().Add(time.Second))
			_, message, err := conn.ReadMessage()
			if err != nil {
				if !strings.Contains(err.Error(), "timeout") {
					fmt.Printf("Error leyendo: %v\n", err)
				}
				break readLoop
			}
			messageCount++
			fmt.Printf("📨 Mensaje %d: %s\n", messageCount, string(message))
			if messageCount >= 2 {
				break readLoop
			}
		}
	}

	// Cerrar limpiamente
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	return nil
}

func main3() {
	fmt.Println("🚀 Prueba intensiva con IPs de Cloudflare")
	fmt.Println("=========================================")
	fmt.Println("Objetivo: Conectar a wss.dalechatea.me:1239 simulando tráfico desde Cloudflare")

	successCount := 0
	totalTests := 0

	// Probar cada servidor con cada IP de Cloudflare
	for _, wsURL := range testServers {
		fmt.Printf("\n📡 === Probando servidor: %s ===\n", wsURL)

		for i, cfIP := range cloudflareIPs {
			totalTests++
			fmt.Printf("\n--- Intento %d/%d ---", i+1, len(cloudflareIPs))

			err := testWebSocketConnection(wsURL, cfIP)
			if err == nil {
				successCount++
				fmt.Printf("🎉 ¡CONEXIÓN EXITOSA! Servidor: %s, IP CF: %s\n", wsURL, cfIP)

				// Si logramos conectar al servidor 2, es un gran éxito
				if strings.Contains(wsURL, ":1239") {
					fmt.Println("🎊 ¡¡¡BREAKTHROUGH!!! Logramos conectar al servidor backend!!!")
				}
			}

			// Pausa entre intentos para evitar rate limiting
			if i < len(cloudflareIPs)-1 {
				time.Sleep(2 * time.Second)
			}
		}
	}

	fmt.Printf("\n\n=== 📊 RESUMEN FINAL ===\n")
	fmt.Printf("Conexiones exitosas: %d/%d\n", successCount, totalTests)

	if successCount == 0 {
		fmt.Println("\n💡 Conclusión: El servidor wss.dalechatea.me:1239 parece tener")
		fmt.Println("   verificación adicional más allá de solo la IP de Cloudflare.")
		fmt.Println("   Podría estar verificando:")
		fmt.Println("   - Certificados específicos de Cloudflare")
		fmt.Println("   - Headers de autenticación internos")
		fmt.Println("   - Conexiones específicas desde infraestructura de CF")
	} else {
		fmt.Printf("\n🎯 Se encontraron %d configuraciones que funcionan!\n", successCount)
	}
}
