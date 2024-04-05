def main():
    v = [1, 5, 10, 50, 100, 500]
    A = 0
    C = []
    solve(A, v, C)


def solve(A, v, C):
    ans = 0

    i = 5
    while True:
        if i < 0:
            break
        t = min(A / v[i], C[i])
        A -= t * v[i]
        ans += t

    print(ans)
