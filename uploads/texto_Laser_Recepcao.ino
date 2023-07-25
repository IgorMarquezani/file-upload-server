//Comunicação via laser
//Código para o receptor.

#include <SoftwareSerial.h>

SoftwareSerial mySerial(3, 2); // RX (OUT do Receptor), TX

void setup() {
  // Abre a serial para receber dados
  Serial.begin(57600);
  while (!Serial) {
    ; // Espera para Arduinos com ATmega32u4 (USB Nativa)
  }
  Serial.println("Aguarde pelos dados.");
  mySerial.begin(4800);
}

void loop() {
  if (mySerial.available()) {
    Serial.write(mySerial.read());
  }
}
