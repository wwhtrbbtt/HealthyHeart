import serial
import time

arduino = serial.Serial(port="/dev/cu.usbmodem1301", baudrate=9600, timeout=0.1)


def parse(raw):
    dec = raw.decode("utf-8")
    try:
        s = str(dec).strip()
        return int(s)
    except Exception as e:
        print(e)
        return 0


while True:
    raw = arduino.readline()
    value = parse(raw)
    print(value)
    time.sleep(0.05)
