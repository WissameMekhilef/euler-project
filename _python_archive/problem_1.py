#!/usr/bin/env python3


def get_multiples_of_under(a1, a2, b):
    result = 0
    for i in range(b):
        if i % a1 == 0 or i % a2 == 0:
            result += i
    return result


def main():
    x = get_multiples_of_under(3, 5, 100)

    print(x)


main()
