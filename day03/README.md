## Problem

The natural numbers are arranged in a 2D grid. 1 is in the middle element and the sequence increases, spiraling outward in the counter-clockwise direction. For example:

    17 16 15 14 13
    18  5  4  3 12
    19  6  1  2 11
    20  7  8  9 10
    21 22 23 24...

Given a natural number `i` we are asked to find the Manhattan distance in the grid from element `i` to element `1`. See the problem statement [here](http://adventofcode.com/2017/day/3).

## Analysis

We would like to find a mapping `f(i)` from the element offset `i` to the Manhattan distance of element `i` from element 1 in the grid.

Consider the sequence of squares having odd side length `a`. The element at offset `a^2` is always in the lower-right corner because we advance to the right and proceed counter-clockwise from the previous square of odd side length. For example, we draw the first 4 squares in the sequence:

The 1-square ending in element 1:

    1

The 3-square ending in element 9:

    5 4 3
    6 1 2
    7 8 9

The 5-square ending in element 25:

    17 16 15 14 13
    18  5  4  3 12
    19  6  1  2 11
    20  7  8  9 10
    21 22 23 24 25

The 7-square ending in element 49:

    37 36 35 34 33 32 31
    38 17 16 15 14 13 30
    39 18  5  4  3 12 29
    40 19  6  1  2 11 28
    41 20  7  8  9 10 27
    42 21 22 23 24 25 26
    43 44 45 46 47 48 49

Consider the boundaries of the squares. We first reduce the indexes `mod a-1`. Next we subtract `2` and add `a-1 (mod a-1)` to shift largest values to the corners while values remain positive. Finally, we add `1` to each value. Let the values reduced in this way be `j(i) = (i - 2 + a - 1) mod (a-1) + 1 = (i + a - 3) mod (a-1) + 1`. Ignoring the degenerate `a=1` case, we illustrate this procedure for the first few cases:

    5 4 3         1 0 1                1 0 1      2 1 2
    6   2  mod 2  0   0  (-2+2) mod 2  0   0  +1  1   1
    7 8 9         1 0 1                1 0 1      2 1 2

    17 16 15 14 13         1  0  3  2  1                3  2  1  0  3      4  3  2  1  4
    18          12         2           0                0           2      1           3
    19          11  mod 4  3           3  (-2+4) mod 4  1           1  +1  2           2
    20          10         0           2                2           0      3           1
    21 22 23 24 25         1  2  3  0  1                3  0  1  2  3      4  1  2  3  4

    37 36 35 34 33 32 31         1  0  5  4  3  2  1                5  4  3  2  1  0  5      6  5  4  3  2  1  6
    38                30         2                 0                0                 4      1                 5
    39                29         3                 5                1                 3      2                 4
    40                28  mod 6  4                 4  (-2+6) mod 6  2                 2  +1  3                 3
    41                27         5                 3                3                 1      4                 2
    42                26         0                 2                4                 0      5                 1
    43 44 45 46 47 48 49         1  2  3  4  5  0  1                5  0  1  2  3  4  5      6  1  2  3  4  5  6

Notice `j(i) = f(i)` for `j(i) >= floor(a/2)`. For `j(i) < floor(a/2)`, notice that `j(i)` is the number of steps from `i` to the nearest corner. Therefore for `j(i) < floor(a/2)`, `f(i)` is `f(a corner index)` minus `j(i)`: `f(i) = f(a^2) - j(i)`.

## Algorithm

Putting this together we obtain the following algorithm:

Given a natural number `i`.

1. Find the smallest odd `a` such that `a^2` >= `i`
1. Calculate `j(i) = (i + a - 3) mod (a-1) + 1`
    1. If `j(i) >= floor(a/2)` then `f(i) = j(i)`
    1. If `j(i) < floor(a/2)` then `f(i) = j(a^2) - j(i)`

## Solving the Problem

Our input is `361527`. The smallest odd `a` such that `a^2 >= 361527` is `603`. `j(361527) = (361527 + 603 - 3) mod (602) + 1 = 326`. `326 >= 301 = floor(603/2)`, so `f(361527) = 326`.

## Conclusion

Writing a program to simulate the problem would have been faster than working through the above analysis. On the other hand, the above algorithm is O(1) while simulating would have been O(i).
