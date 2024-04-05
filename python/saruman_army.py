def main():
    pass


def solve(N, R, X: list):
    X.sort()

    i = 0
    ans = 0

    while i < N:
        s = X[i]
        i += 1
        while i < N and X[i] <= s + R:
            i += 1
        p = X[i - 1]
        while i < N and X[i] <= p + R:
            i += 1
        ans += 1
    print(ans)
