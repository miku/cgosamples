import random

for i in range(1_000_000_000):
    # TODO: bufio
    print(0 if random.random() > 0.5 else 1)

