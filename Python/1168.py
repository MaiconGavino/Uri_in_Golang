casos = int(input())
for i in range(0,casos):
    leds = 0
    entrada = input()
    for i in range(0, len(entrada)):
        if entrada[i] == '1':
            leds += 2
        if entrada[i] == '2':
            leds += 5
        if entrada[i] == '3':
            leds += 5
        if entrada[i] == '4':
            leds += 4
        if entrada[i] == '5':
            leds += 5
        if entrada[i] == '6':
            leds += 6
        if entrada[i] == '7':
            leds += 3
        if entrada[i] == '8':
            leds += 7
        if entrada[i] == '9':
            leds += 6
        if entrada[i] == '0':
            leds += 6
    print("{} leds".format(int(leds)))