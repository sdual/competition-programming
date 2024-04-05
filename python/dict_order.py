from typing import assert_never, get_origin


def main():
    N = 10
    S = []

    solve(N, S)


def solve(N, S):
    a = 0
    b = N - 1
    while a <= b:
        left = False

        i = 0
        while a + i <= b:
            if S[a + i] < S[b - i]:
                left = True
                break
            elif S[a + i] > S[b - i]:
                left = False
                break
        if left:
            print(S[a])
            a += 1
        else:
            print(S[b])
            b -= 1
