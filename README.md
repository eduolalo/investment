# Ejercicio de Go

## Requerimientos

* Go 1.15 o superior
* Variables de entorno seteadas:
    * RDS_STRING - cadena de conexión de formato redis (redis://{usuario}:{pass}@{url}:{puerto}/0?ssl=true)
    * RDS_POOL - tamaño del pool de conexiones (int)
    * RDS_NAME - Nombre del cliente de Redis (string)
    * PORT - Número del puesto que escuchará (por defecto es 3000)

## Instrucciones

### Si lo quieres correr en modo debug en VSCode:

* [Configura tu Delve](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)
* Abre el proyecto en VSCode
* configura tu _launch.json_
* Pícale en _Run_, presiona _F5_ o _SHIFT + CMD + D_

### Modo terminal:

* Ándate al path donde clonaste el proyecto y
```
$ go run main.go
```

## Jugar con el proyecto

[Aquí te dejo una colección de Postman](https://www.getpostman.com/collections/6401765c93da452dbddd) para que pruebes el proyecto, sólo le cambias el baseURL al tuyo :P

