class Pair:
    def __init__(self) -> None:
        self.first = None
        self.second = None


def main():
    N = 10
    S = []
    T = []

    itv = []
    for _ in range(N):
        itv.append(Pair())
    solve(itv, N, S, T)


def solve(itv, N, S, T):
    for i in range(N):
        itv[i].first = T[i]
        itv[i].second = S[i]
    itv.sort()

    ans = 0
    t = 0
    for i in range(N):
        if t < itv[i].second:
            ans += 1
            t = itv[i].first

    print(ans)
