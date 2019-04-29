"""
	| Solving The Gardian's math puzzle for vietnamese 3rd graders (efficiently)
	| Python implementation
	|
	| _ + 13 * _ / _ + _ + 12 * _ - _ - 11 + _ * _ / _ - 10 = 66
	|
	| Author: nperraud <np@bitbox.io>
"""

from datetime import datetime
from itertools import permutations
from multiprocessing import Pool


def unpuzzle(a):
    solutions = list()
    dataset = frozenset(range(1, 10)) - {a}
    for x in permutations(dataset):
        if (
            round(
                a + (13 * x[0] / x[1]) + x[2] + (12 * x[3]) - x[4] - 11 + (x[5] * x[6] / x[7]) - 10,
                10,
            )
            == 66
        ):
            solutions.append((a,) + x)
    return solutions


if __name__ == "__main__":
    # /!\ require Python 3
    if 5 / 2 == 2:
        print("Requires Python3.")
        exit(1)
    print(__doc__)

    # use multiple cores
    starttime = datetime.now()
    with Pool(3) as p:
        solutions = p.map(unpuzzle, range(1, 10))
    endtime = datetime.now()
    n = 0
    for s in solutions:
        for r in s:
            n += 1
            print(r)
    print(
        "\n-- %d solutions found in %s seconds. --\n" % (n, (endtime - starttime).total_seconds())
    )
    exit(0)
