# RESOLUCIÓN DEL DESAFÍO FINAL 

Según el desafío a resolver decidimos crear metodos par obtener listados completos de pacientes, turnos y odontologos. 
Por lógica de negocio decidimos que los metodos POST de paciente y odontólogo no pueden crearse con DNI o matrícula utilizadas previamente. Por ello, según la consigna que establecía reservar el método PATCH para actualizar la estructura por un único campo, establecimos reservarlo para modificación de estos datos sensibles. Es por esto que el método PUT (para modificar la estructura completa) solo será exitoso cuando DNI o matrícula coincidan con los existentes en la base de datos para su respectivo id. En el caso de turnos, el metodo PATCH permite actualizar únicamente fecha y hora como alternativa para reprogramar un turno.

En el caso del método DELETE no se permite eliminar pacientes y odontólogos que tengan turnos asignados.

Para el caso del crear un turno con DNI de paciente y matrícula se creó una estructura auxiliar que recuperará los datos persistidos para realizar el POST de correspondiente.

En cuanto a la seguridad mediante middleware se creó un handler de autenticación que extrae y valida un token del header de la request.

Creamos otro middleware para loggeo de la aplicacion (se encuentra comentado en el main), pero a fines prácticos decidimos no utilizarlo.

Por último, en cuanto a la documentación de nuestra API-REST se implementó swagger en todos los handlers correspondientes a odontólogo, paciente y turno.

<u> Para ejecutar este proyecto: </u>
1. Ejecute el script de base de datos.
2. Establecer variables de entorno según el ejemplo otorgado en ".example.env".
3. Ejecutar en la consola `go run ./cmd/server/main.go`  desde la raiz del proyecto.
4. Una vez este corriendo, se podrá consultar swagger en esta direccion "http://localhost:8080/docs/index.html#/".   
5. Podrá consultar la colección de postman adjunta para mayor claridad de los métodos de la API.

Gracias Lore  y Ceci :)!!
