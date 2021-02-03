while True:
    try:
        line = input().split(" ")
        veloc, temp = line
        veloc = int(veloc)
        temp = int(temp)
        result = veloc * (temp*2)
        print("{}".format(result))
    except EOFError:
        break