with open("in_large", "w") as f:
    f.write("1\n")
    f.write("2000000\n")
    for i in range(2000000):
        f.write("1 ")