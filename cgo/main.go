package main

/*
#include <stdio.h>   // Это НЕ просто комментарий!
#include <stdlib.h>  // Это код на C, который будет обработан CGO

void hello() {       // Объявление C функции
    printf("Hello from C!\n");
}
*/
import "C" // Этот import говорит компилятору: "эй, прочитай комментарий выше как C код!"

func main() {
	C.hello() // Вызываем C функцию
}
