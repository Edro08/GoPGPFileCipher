## GoPGPFileCipher

**Descripción:** Aplicación de línea de comandos escrita en Go que permite encriptar y desencriptar archivos utilizando cifrado PGP.<br>
Soporta la gestión de múltiples perfiles, cada uno con sus propias claves públicas, privadas y passphrase.<br>
Es ideal para automatizar o facilitar el manejo seguro de archivos en entornos segmentados por cliente, grupo o entorno.<br>

***💡 NOTA: Se puede utilizar perfil default, de ser el caso skip paso 1.***

### 🔹<span style="color:orange">Paso 1: Crear paths/carpetas</span>
- Crear/agregar paths con el nombre del perfil, ejemplo de estructura:
  - keys/<span style="color:green">[Nombre del perfil]</span>/
  - files/<span style="color:green">[Nombre del perfil]</span>/request/
  - files/<span style="color:green">[Nombre del perfil]</span>/response/

### 🔹<span style="color:orange">Paso 2: Agregar/Cambiar keys</span>
- Agregar/cambiar las correspondientes keys del perfil, ejemplo:
    - keys/<span style="color:green">[Nombre del perfil]</span>/private.key
    - keys/<span style="color:green">[Nombre del perfil]</span>/public.key
    - keys/<span style="color:green">[Nombre del perfil]</span>/passphrase.txt

### 🔹<span style="color:orange">Paso 3: Ejecutar main.go</span>
- Ejecutar programa desde:
  - <span style="color:green">app/main.go</span>


- Pasos dentro del programa:
  - Ejecutar [Paso 4](#span-stylecolororangepaso-4-agregar-archivos-a-procesar-en-requestspan)
  - Ingresar el nombre del perfil cuando se solicite.
    - Si lo dejás vacío, se usará el perfil default. 
  - El sistema procesará los archivos detectados.
  - Ejecutar [Paso 5](#span-stylecolororangepaso-5-extraer-archivos-de-responsespan)
  - Podrás repetir el proceso o salir del programa.

### 🔹<span style="color:orange">Paso 4: Agregar archivos a procesar en 'request'</span>
- Agregar los archivos a procesar en request:
  - files/<span style="color:green">[Nombre del perfil]</span>/request/<span style="color:green">[archivos.txt]</span>
  - files/<span style="color:green">[Nombre del perfil]</span>/request/<span style="color:green">[archivos.pgp]</span>

- Archivos permitidos:
  - .txt → serán encriptados 
  - .pgp → serán desencriptados


### 🔹<span style="color:orange">Paso 5: Extraer archivos de 'response'</span>
- Una vez terminado el proceso, retirar sus archivos desde:
  - files/<span style="color:green">[Nombre del perfil]</span>/response/<span style="color:green">[archivos.txt]</span>
  - files/<span style="color:green">[Nombre del perfil]</span>/response/<span style="color:green">[archivos.pgp]</span>