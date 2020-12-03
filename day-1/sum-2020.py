# This is crazy brute force, probably a more elegant way to do this

input = []
with open('input') as input_file:
    input = [int(s) for s in input_file.read().splitlines()]
input.sort(reverse=True)

for i, n in enumerate(input, start=1):
    for x in input[i::]:
        if n + x == 2020:
            print(f"{n}+{x}=2020: {n}*{x}={n*x}")

for i, n in enumerate(input, start=1):
    for x in input[i::]:
        for y in input[i+1::]:
            if n + x + y == 2020:
                print(f"{n}+{x}+{y}=2020: {n}*{x}*{y}={n*x*y}")
