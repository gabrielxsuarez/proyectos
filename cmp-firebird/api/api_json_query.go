package api

import (
	"cmp-firebird/config"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

/* ========== REQUEST ========== */
type ApiJsonQueryRequest struct {
	Base       string `json:"base"`
	Query      string `json:"query"`
	Parametros []any  `json:"parametros"`
	Timeout    int    `json:"timeout"`
}

/* ========== RESPONSE ========== */
type ApiJsonQueryResponse struct {
	Estado string           `json:"estado"`
	Datos  []map[string]any `json:"datos,omitempty"`
	Error  string           `json:"error,omitempty"`
}

/* ========== ENDPOINT ========== */
func ApiJsonQueryEndpoint(w http.ResponseWriter, r *http.Request) {

	// 1. Parsear request
	var request ApiJsonQueryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		jsonQueryError(w, 400, "ERROR", "Error al parsear request: %v", err)
		return
	}
	defer r.Body.Close()

	// 2. Validar request
	if request.Base == "" || request.Query == "" {
		jsonQueryError(w, 400, "ERROR", "Parametros incorrectos", nil)
		return
	}

	// 3. Obtener pool de conexiónes
	db, ok := config.Pool(request.Base)
	if !ok {
		jsonQueryError(w, 400, "ERROR", "Base de datos '%s' no existe", nil)
		return
	}

	// 4. Configurar timeout
	timeout := request.Timeout
	if timeout <= 0 {
		timeout = config.Timeout()
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeout)*time.Second)
	defer cancel()

	// 5. Preparar query
	stmt, err := db.PrepareContext(ctx, request.Query)
	if err != nil {
		jsonQueryError(w, 500, "ERROR", "Error al preparar query: %v", err)
		return
	}
	defer stmt.Close()

	// 6. Ejecutar query
	rows, err := stmt.QueryContext(ctx, request.Parametros...)
	if err != nil {
		jsonQueryError(w, 500, "ERROR", "Error al ejecutar query: %v", err)
		return
	}
	defer rows.Close()

	// 7. Obtener columnas
	columnas, err := rows.Columns()
	if err != nil {
		jsonQueryError(w, 500, "ERROR", "Error al obtener columnas: %v", err)
		return
	}
	columnasMinusculas := make([]string, len(columnas))
	for i, columna := range columnas {
		columnasMinusculas[i] = strings.ToLower(columna)
	}
	tiposColumnas, err := rows.ColumnTypes()
	if err != nil {
		jsonQueryError(w, 500, "ERROR", "Error al obtener tipos en columnas: %v", err)
		return
	}

	// 8. Valores
	valores := make([]any, len(columnas))
	for i, tipoColumna := range tiposColumnas {
		switch tipoColumna.DatabaseTypeName() {
		case "INTEGER", "BIGINT", "SMALLINT":
			valores[i] = new(sql.NullInt64)
		case "DECIMAL", "NUMERIC", "FLOAT", "DOUBLE":
			valores[i] = new(sql.NullFloat64)
		case "VARCHAR", "CHAR", "TEXT":
			valores[i] = new(sql.NullString)
		case "TIMESTAMP", "DATE", "TIME":
			valores[i] = new(sql.NullTime)
		case "BOOLEAN":
			valores[i] = new(sql.NullBool)
		default:
			valores[i] = new(any)
		}
	}

	// 9. Punteros
	punteros := make([]any, len(columnas))
	for i := range columnas {
		punteros[i] = &valores[i]
	}

	// 10. Headers
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	// 11. Stream
	primeraFila := true
	stream := json.NewStream(json.ConfigCompatibleWithStandardLibrary, w, config.BufferSize())
	stream.WriteObjectStart()
	stream.WriteObjectField("datos")
	stream.WriteArrayStart()
	for rows.Next() {

		// 12. Leer registro
		if err := rows.Scan(punteros...); err != nil {
			jsonQueryFinalizarStreamErr(stream, "ERROR", "Error al leer fila: %v", err)
			return
		}

		// 13. Primera fila
		if !primeraFila {
			stream.WriteMore()
		} else {
			primeraFila = false
		}

		// 14. Escribir registro
		stream.WriteObjectStart()
		for i, columna := range columnasMinusculas {
			if i > 0 {
				stream.WriteMore()
			}
			stream.WriteObjectField(columna)

			// 15. Escribir valor
			valor := valores[i]
			switch v := valor.(type) {
			case nil:
				stream.WriteNil()
			case string:
				stream.WriteString(v)
			case int64:
				if v > 9007199254740992 || v < -9007199254740992 {
					stream.WriteString(strconv.FormatInt(v, 10))
				} else {
					stream.WriteInt64(v)
				}
			case int32:
				stream.WriteInt32(v)
			case int16:
				stream.WriteInt16(v)
			case decimal.Decimal:
				stream.WriteRaw(v.String())
			case float64:
				stream.WriteFloat64(v)
			case float32:
				stream.WriteFloat32(v)
			case bool:
				stream.WriteBool(v)
			case time.Time:
				switch tiposColumnas[i].DatabaseTypeName() {
				case "DATE":
					stream.WriteString(v.Format(time.DateOnly))
				case "TIME":
					stream.WriteString(v.Format(time.TimeOnly))
				default:
					stream.WriteString(v.Format(time.DateTime))
				}
			default:
				stream.WriteVal(v)
			}
		}
		stream.WriteObjectEnd()
	}

	// 16. Fin Response
	jsonQueryFinalizarStream(stream, "OK", "")
}

/* ========== FUNCIONES AUXILIARES ========== */
func jsonQueryError(w http.ResponseWriter, codigoHttp int, estado string, error string, err error) {
	w.WriteHeader(codigoHttp)
	w.Header().Set("Content-Type", "application/json")
	resp := ApiJsonQueryResponse{
		Estado: estado,
		Error:  fmt.Sprintf(error, err),
	}
	json.NewEncoder(w).Encode(resp)
}

func jsonQueryFinalizarStream(stream *json.Stream, estado string, error string) {
	jsonQueryFinalizarStreamErr(stream, estado, error, nil)
}

func jsonQueryFinalizarStreamErr(stream *json.Stream, estado string, error string, err error) {
	stream.WriteArrayEnd()
	stream.WriteMore()
	stream.WriteObjectField("estado")
	stream.WriteString(estado)
	if error != "" {
		stream.WriteMore()
		stream.WriteObjectField("error")
		if err != nil {
			stream.WriteString(fmt.Sprintf("%s: %v", error, err))
		} else {
			stream.WriteString(error)
		}
	}
	stream.WriteObjectEnd()
	stream.Flush()
}
