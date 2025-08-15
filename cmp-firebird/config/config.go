package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/nakagami/firebirdsql"
	"gopkg.in/yaml.v3"
)

/* ========== VARIABLES GLOBALES ========== */
var config configuracion
var pools map[string]*sql.DB

/* ========== ESTRUCTURAS ========== */
type configuracion struct {
	Servidor string            `yaml:"servidor"`
	Bases    map[string]string `yaml:"bases"`
	Timeout  int               `yaml:"timeout"`
}

/* ========== INIT ========== */
func init() {
	configYaml, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error al leer config.yaml: %v", err)
	}

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		log.Fatalf("Error al parsear config.yaml: %v", err)
	}

	pools = make(map[string]*sql.DB)
	for nombre, dsn := range config.Bases {
		db, err := sql.Open("firebirdsql", dsn)
		if err != nil {
			log.Fatalf("Error al abrir base de datos %s: %v", nombre, err)
		}
		pools[nombre] = db
	}
}

/* ========== METODOS PUBLICOS ========== */
func Servidor() string {
	return config.Servidor
}

func Pool(nombreBase string) (*sql.DB, bool) {
	db, ok := pools[nombreBase]
	return db, ok
}

func Timeout() int {
	return config.Timeout
}

func BufferSize() int {
	return 65536 // 64KB
}
