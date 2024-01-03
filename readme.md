## First Time
go get      github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

## run
CompileDaemon -command="./crs"


Consideraciones:

Al iniciar la aplicación:

► Verifica si existe el archivo "lockfile.lock", allí se almacena el PID de la aplicación que está corriendo.

► Verifica si existe el json de configuración de funcionalidades llamado "crs-init-config.json"