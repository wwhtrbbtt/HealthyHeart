int pin = A0; 
int val;    

void setup() {
  Serial.begin(9600); 
}

void loop() {
  Sensorwert = analogRead(pin); 
  Serial.println(val);     
  delay(5); 
}