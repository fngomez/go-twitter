# go-twitter

Es necesario correr los siguientes comandos para que el proyecto compile:

go get github.com/abiosoft/ishell

--

go get -u github.com/kardianos/govendor

govendor init

govendor add +external

govendor fetch github.com/gin-gonic/gin@v1.3

--

Para traerse todas las dependencias hay que pararse en la carpeta vendor y ejecutar: govendor sync
