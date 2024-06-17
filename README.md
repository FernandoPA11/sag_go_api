# sag_go_api
API REST para un sistema de gestión ganadera en el lenguaje Golang

# Estructura del proyecto

SAG_GO_API/
├── cmd/
│   └── server/
│       └── main.go                        # Punto de entrada principal de la aplicación
├── pkg/
│   ├── handlers/                          # Manejadores de rutas (endpoints)
│   │   ├── user_handlers.go
│   │   └── user_handlers_test.go          # Pruebas para los handlers
│   ├── models/                            # (O entidades, domain, etc.) Modelos de datos
│   │   ├── user.go
│   │   └── user_test.go                   # Pruebas para los modelos
│   ├── repository/                        # Lógica de acceso a datos
│   │   ├── user_repository.go
│   │   └── user_repository_test.go        # Pruebas para los repositorios
│   └── services/                          # Lógica de negocio
│       ├── user_service.go
│       └── user_service_test.go           # Pruebas para los servicios
├── core/
│   ├── config/                            # Manejo de configuración
│   │   ├── config.go
│   │   └── config_test.go                 # Pruebas para la configuración
│   └── db/                                # Inicialización y conexión a la base de datos
│       ├── db.go
│       └── db_test.go                     # Pruebas para la base de datos
├── .air.toml                              # Configuración de Air para hot reload
├── .gitignore                             # Archivos y carpetas que git debe ignorar
├── go.mod                                 # Dependencias del proyecto
└── go.sum
