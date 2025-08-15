package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"fyne.io/systray"
	"github.com/go-vgo/robotgo"
)

/* ========== VARIABLES GLOBALES ========== */
var (
	pantallaActiva bool
	pantallaItem   *systray.MenuItem
	intervalo      time.Duration
)

/* ========== MAIN ========== */
func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	err := cargarIcono()
	if err != nil {
		fmt.Println(err)
		systray.Quit()
		return
	}

	err = cargarMenu()
	if err != nil {
		fmt.Println(err)
		systray.Quit()
		return
	}
}

func onExit() {
}

/* ========== ESTRUCTURAS ========== */
type MenuItem struct {
	Texto   string
	Comando string
	SubMenu []MenuItem
}

/* ========== CARGAR ICONO ========== */
func cargarIcono() error {
	var ruta string

	if runtime.GOOS == "windows" {
		ruta = "config/icono.ico"
	} else {
		ruta = "config/icono.png"
	}

	bytes, err := os.ReadFile(ruta)
	if err != nil {
		return fmt.Errorf("error al cargar el icono: %w", err)
	}

	systray.SetIcon(bytes)
	return nil
}

/* ========== CARGAR MENU ========== */
func cargarMenu() error {
	err := cargarMenuUsuario()
	if err != nil {
		return err
	}

	cargarMenuSalir()
	return nil
}

func cargarMenuUsuario() error {
	archivos, err := archivosMenu()
	if err != nil {
		return fmt.Errorf("error al leer los archivos del menú: %w", err)
	}

	menu, err := menuItems(archivos)
	if err != nil {
		return fmt.Errorf("error al procesar los archivos del menú: %w", err)
	}

	dibujarMenu(menu, nil)
	return nil
}

func cargarMenuSalir() {
	salirItem := systray.AddMenuItem("Salir", "")
	go func() {
		<-salirItem.ClickedCh
		systray.Quit()
	}()
}

/* ========== PROCESAR MENU ========== */
func archivosMenu() ([]string, error) {
	var archivos []string

	err := filepath.WalkDir("config", func(ruta string, archivo fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		esArchivo := !archivo.IsDir()
		esTxt := strings.HasSuffix(archivo.Name(), ".txt")
		contieneMenu := strings.Contains(archivo.Name(), "menu")

		if esArchivo && esTxt && contieneMenu {
			archivos = append(archivos, ruta)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error al recorrer la carpeta config: %w", err)
	}

	sort.Strings(archivos)
	return archivos, nil
}

func menuItems(archivos []string) ([]MenuItem, error) {
	var menu []MenuItem
	var pila []*[]MenuItem

	pila = append(pila, &menu)

	for _, archivo := range archivos {
		contenido, err := os.ReadFile(archivo)
		if err != nil {
			return nil, fmt.Errorf("error al leer el archivo %s: %w", archivo, err)
		}

		lineas := strings.SplitSeq(string(contenido), "\n")

		for linea := range lineas {
			if strings.TrimSpace(linea) == "" {
				continue
			}

			sangria := 0
			for strings.HasPrefix(linea, "\t") {
				sangria++
				linea = strings.TrimPrefix(linea, "\t")
			}

			var nuevoElemento MenuItem

			if strings.TrimSpace(linea) == "-" {
				nuevoElemento.Texto = "-"
			} else {
				partes := strings.SplitN(linea, ":", 2)
				nuevoElemento.Texto = strings.TrimSpace(partes[0])
				if len(partes) > 1 {
					nuevoElemento.Comando = strings.TrimSpace(partes[1])
				}
			}

			for len(pila) > sangria+1 {
				pila = pila[:len(pila)-1]
			}

			menuPadre := pila[len(pila)-1]
			*menuPadre = append(*menuPadre, nuevoElemento)

			if nuevoElemento.Comando == "" {
				pila = append(pila, &(*menuPadre)[len(*menuPadre)-1].SubMenu)
			}
		}
	}

	return menu, nil
}

/* ========== DIBUJAR MENU ========== */
func dibujarMenu(items []MenuItem, parent *systray.MenuItem) {
	for _, item := range items {
		if item.Texto == "-" {
			agregarSeparador(parent)
			continue
		}

		menuItem := agregarMenuItem(item, parent)

		if len(item.SubMenu) > 0 {
			dibujarMenu(item.SubMenu, menuItem)
		} else if item.Comando != "" {
			if strings.HasPrefix(item.Comando, "pantalla_activa") {
				asignarPantallaActiva(menuItem, item.Comando)
			} else {
				asignarComando(menuItem, item.Comando)
			}
		}
	}
}

func agregarSeparador(parent *systray.MenuItem) {
	if parent == nil {
		systray.AddSeparator()
	} else {
		parent.AddSeparator()
	}
}

func agregarMenuItem(item MenuItem, parent *systray.MenuItem) *systray.MenuItem {
	if parent == nil {
		return systray.AddMenuItem(item.Texto, item.Comando)
	} else {
		return parent.AddSubMenuItem(item.Texto, item.Comando)
	}
}

func asignarComando(menuItem *systray.MenuItem, comando string) {
	go func() {
		for range menuItem.ClickedCh {
			partes := parsearComando(comando)
			ejecutarComando(partes)
		}
	}()
}

func parsearComando(comando string) []string {
	var partes []string
	dentroComillas := false
	parteActual := ""

	for _, caracter := range comando {
		if caracter == '"' {
			dentroComillas = !dentroComillas
			continue
		}

		if caracter == ' ' && !dentroComillas {
			if parteActual != "" {
				partes = append(partes, parteActual)
				parteActual = ""
			}
		} else {
			parteActual += string(caracter)
		}
	}

	if parteActual != "" {
		partes = append(partes, parteActual)
	}

	return partes
}

func ejecutarComando(partes []string) {
	if len(partes) == 0 {
		return
	}

	if runtime.GOOS == "windows" && strings.HasPrefix(partes[0], "http") {
		nuevasPartes := append([]string{"explorer"}, partes...)
		partes = nuevasPartes
	}

	fmt.Printf("Ejecutando comando: %s con argumentos: %v\n", partes[0], partes[1:])
	cmd := exec.Command(partes[0], partes[1:]...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
	}
}

/* ========== PANTALLA ACTIVA ========== */
func asignarPantallaActiva(menuItem *systray.MenuItem, comando string) {
	pantallaItem = menuItem

	partes := strings.Fields(comando)
	if len(partes) < 3 {
		return
	}

	activo := partes[1] == "true"
	segundos := strings.TrimSuffix(partes[2], "s")
	tiempo, err := strconv.Atoi(segundos)
	if err != nil {
		tiempo = 30
	}
	intervalo = time.Duration(tiempo) * time.Second

	go func() {
		for range menuItem.ClickedCh {
			alternarPantallaActiva()
		}
	}()

	if activo {
		iniciarPantallaActiva()
	}
}

func alternarPantallaActiva() {
	if pantallaActiva {
		detenerPantallaActiva()
	} else {
		iniciarPantallaActiva()
	}
}

func iniciarPantallaActiva() {
	pantallaActiva = true
	pantallaItem.Check()
	go mantenerPantallaActiva()
}

func detenerPantallaActiva() {
	pantallaActiva = false
	pantallaItem.Uncheck()
}

func mantenerPantallaActiva() {
	for pantallaActiva {
		robotgo.KeyTap("f15")

		for i := 0; i < int(intervalo.Seconds()*10) && pantallaActiva; i++ {
			time.Sleep(100 * time.Millisecond)
		}
	}
}
