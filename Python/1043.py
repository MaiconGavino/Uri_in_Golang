line = input().split(" ")
ladoA, ladoB, ladoC = line
ladoA = float(ladoA)
ladoB = float(ladoB)
ladoC = float(ladoC)
if ladoB - ladoC < ladoA and ladoB + ladoC > ladoA:
    if ladoA - ladoC < ladoB and ladoA + ladoC > ladoB:
        if ladoA - ladoB < ladoC and ladoA + ladoB > ladoC:
            perimetro = ladoA + ladoB + ladoC
            print("Perimetro = {:.1f}".format(perimetro))

else:
    areaTrapezio = ((ladoA + ladoB)*ladoC)/2
    print("Area = {:.1f}".format(areaTrapezio))