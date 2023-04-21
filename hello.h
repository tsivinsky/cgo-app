#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *Hello(const char *name) {
  static char greeting[100];
  snprintf(greeting, sizeof(greeting), "Hello from C, %s!", name);
  return greeting;
}

void SayHello(char *name) {
  char *greeting = Hello(name);
  printf("%s\n", greeting);
}
