MAX_N = 4
used = [False] * MAX_N
perm = [0] * MAX_N


def permutation1(pos: int, n: int):
    if pos == n:
        return

    for i in range(n):
        if not used[i]:
            perm[pos] = i
            used[i] = True
            permutation1(pos + 1, n)
            used[i] = False
    print(perm)


if __name__ == "__main__":
    permutation1(0, 4)
