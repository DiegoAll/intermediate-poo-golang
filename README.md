# intermediate-poo-golang


## Visual Studio Code

https://github.com/golang/vscode-go

### 3. Repaso general: Variables y condicionales, slices y map

### 4. Repaso general: GoRoutines y apuntadores

### 5. ¿Es Go orientado a objetos?

### 6. Structs vs Clases

### 7. Métodos y funciones

### 8. Constructores 

### 9. Herencia

### 10. Interfaces

### 11. Aplicando interfaces con Abstract Factory 

### 12. Implementación final de Abstract Factory

### 13. Funciones anonimas

### 14. Funciones variadicas y retornos con nombre

### 15. Cómo utilizar los Go Modules

### 16. Creando nuestro módulo

### 17. Testing

### 18. Code Coverage 

### 19. Profiling

### 20. Testing usando Mocks


### 21. Implementando Mocks

Es muy importante guardar los valores originales de las funciones que estamos utilizando para realizar nuestros test en caso de que en test posteriores decidamos evaluar otras partes del código que también  dependen de esas funciones. 


### 22. Unbuffered Channels y buffered channels 

Go utiliza dos tipos de canales para las comunicaciones, los unbuffered channels son aquellos que tienen una cola de longitud 0, es decir, se utilizan para comunicación síncrona. 

//Creating de Buffered channel with size = 3 
c := make(chan int, 3) 

//Sending the message 
c <- 1 

// Is waiting the message. 
fmt.Println(<-c) 

fatal error: all goroutines are asleep - deadlock! 

Hay alternativas a los canales. No necesariamente puedes manejar tus Go Routines utilizando canales, hay otro mecanismo que sera muy simple muy intuitivo de entender que nos permitira lograr el mismo efecto que con los canales.  


### 23. Waitgroup 

import ( "sync" ) 

var wg sync.WaitGroup 
wg.Add() 
wg.Done() 

Cada vez que lanzamos una subrutina nueva, le agregamos 1 al contador y cada vez que termina la restamos 1. 
go doSomething(i, &wg)  

Esto se va a pasar como copia si yo no lo envio como referencia.   

### Summary 

Cuando se trabaja con GoRoutines no siempre se planea que estas envíen datos a través de channels entre unas y otras, en estos casos se puede utilizar un Wait Group 
Para alcanzar ese bloqueo de rutinas que es necesario. 
 

### 24. Buffered Channels como semaforos 

 

 



