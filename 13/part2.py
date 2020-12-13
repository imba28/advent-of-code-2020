from math import floor, ceil
from functools import reduce

def get_data(filePath):
    with open(filePath, "r") as f:
        data = f.readlines()
        return [v.strip() for v in data[1].split(",")]

# code from https://en.wikibooks.org/wiki/Algorithm_Implementation/Mathematics/Extended_Euclidean_algorithm
def extended_euclid(a, b):
    """return (g, x, y) such that a*x + b*y = g = gcd(a, b)"""
    x0, x1, y0, y1 = 0, 1, 1, 0
    while a != 0:
        (q, a), b = divmod(b, a), a
        y0, y1 = y1, y0 - q * y1
        x0, x1 = x1, x0 - q * x1
    return b, x0, y0

# theory from https://www.youtube.com/watch?v=rUoKzTZewT8
def crt(m, a):
    # (i), calculate step size
    N = reduce(lambda x, y: x * y, m)

    # (ii)
    n = [int(N / v) for v in m]

    # (iii)
    x = []
    for i in range(len(n)):
        _, x0, _ = extended_euclid(n[i], m[i])
        x.append(x0)

    # (iv)
    g = sum([va * vx * vn for va, vx, vn in zip(a, x, n)])

    # (v)
    # find smallest positive result
    if g > 0:
        g = g - floor(g / N) * N
    else:
        g = g + ceil(abs(g) / N) * N
    return g

if __name__ == "__main__":
    data = get_data("13/input.txt")

    offsets = [k for k, v in enumerate(data) if v != "x"]
    ids = [int(id) for id in data if id != "x"]

    # somewhat hacky .... 5 * t => make sure dividend is always greater than the divisor => mods must be positive!
    mods = [0 if t == ids[0] else 5*t - offset for t, offset in zip(ids, offsets)]
    x = crt(ids, mods)

    print(x)