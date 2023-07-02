
Segun el desafio a resolver decidimos crear metodos par obtener listados completos de pacientes, turnos y odontologos. 
Por logica de negocio decidimos que los metodos POST de paciente y odontologo no pueden crearse con DNI o matricula utilizadas previamente. Por ello, segun la consigna que establecia reservar el metodo PATCH para actualizar la estructura por un unico campo, establecimos reservarlo para modificacion de estos datos sensibles. Es por esto que el metodo PUT(para modificar la estructura completa) solo sera exitoso cuando DNI o matricula coincidan con los existentes en la base de datos para su respectivo id. En el caso de turnos el metodo PATCH permite actualizar unicamente fecha y hora como alternativa para reprogramar un turno.

En el caso del metodo DELETE no se permite eliminar pacientes y odontologos que tengan turnos asignados.

Para el caso del crear un turno con DNI de paciente y matricula de dentista se creo una estructura auxiliar que recuperara los datos persistidos para realizar el POST de turno.

En cuanto a la seguridad mediante middleware se creo un handler de autenticacion que extrae y valida un token del header de la request.

Creamos otro middleware para loggeo de la aplicacion (se encuentra comentado en el main), pero a fines practicos decidimos no utilizarlo.

Por ultimo en cuanto a la docuemntacion de nuestra API-REST se implemento swagger en todos los hadlers correspondientes a odontologo, paciente y turno.

Para ejecutar este proyecto:
1. Ejecute el script de base de datos
2. Establecer variables de entorno segun el ejemplo otorgado en ".example.env"
3. Ejecutar en la consola "go run ./cmd/server/main.go" desde la raiz del proyecto
4. Una vez este corriento se podra consultar swagger en esta direccion "http://localhost:8080/docs/index.html#/"   
5. Podra consultar la coleccion de postman adjunta para mayor claridad de los metodos de la API

Gracias Lore  y Ceci :)!!
