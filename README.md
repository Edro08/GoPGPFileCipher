## 💻 go-pgp-file-cipher

**Descripción:** Aplicación de línea de comandos escrita en Go que permite encriptar y desencriptar archivos utilizando cifrado PGP.<br>
Soporta la gestión de múltiples perfiles, cada uno con sus propias claves públicas, privadas y passphrase.<br>
Es ideal para automatizar o facilitar el manejo seguro de archivos en entornos segmentados por cliente, grupo o entorno.<br>

***💡 NOTA: Se puede utilizar perfil default, de ser el caso skip paso 1.***

### ☕ Paso 1: Crear paths/carpetas
- Crear/agregar paths con el nombre del perfil, ejemplo de estructura:
  - keys/<span style="color:green">[Nombre del perfil]</span>/
  - files/<span style="color:green">[Nombre del perfil]</span>/request/
  - files/<span style="color:green">[Nombre del perfil]</span>/response/

### ☕ Paso 2: Agregar/Cambiar keys
- Agregar/cambiar las correspondientes keys del perfil, ejemplo:
    - keys/<span style="color:green">[Nombre del perfil]</span>/private.key
    - keys/<span style="color:green">[Nombre del perfil]</span>/public.key
    - keys/<span style="color:green">[Nombre del perfil]</span>/passphrase.txt

### ☕ Paso 3: Ejecutar main.go
- Ejecutar programa desde:
  - <span style="color:green">app/main.go</span>


- Pasos dentro del programa:
  - Ejecutar [Paso 4: Agregar archivos a procesar en 'request'](#paso-4-agregar-archivos-a-procesar-en-request)
  - Ingresar el nombre del perfil cuando se solicite.
    - Si lo dejás vacío, se usará el perfil default. 
  - El sistema procesará los archivos detectados.
  - Ejecutar [Paso 5: Extraer archivos de 'response'](#paso-5-extraer-archivos-de-response)
  - Podrás repetir el proceso o salir del programa.

### ☕ Paso 4: Agregar archivos a procesar en 'request'
- Agregar los archivos a procesar en request:
  - files/<span style="color:green">[Nombre del perfil]</span>/request/<span style="color:green">[archivos.txt]</span>
  - files/<span style="color:green">[Nombre del perfil]</span>/request/<span style="color:green">[archivos.pgp]</span>

- Archivos permitidos:
  - .txt → serán encriptados 
  - .pgp → serán desencriptados


### ☕ Paso 5: Extraer archivos de 'response'
- Una vez terminado el proceso, retirar sus archivos desde:
  - files/<span style="color:green">[Nombre del perfil]</span>/response/<span style="color:green">[archivos.txt]</span>
  - files/<span style="color:green">[Nombre del perfil]</span>/response/<span style="color:green">[archivos.pgp]</span>
