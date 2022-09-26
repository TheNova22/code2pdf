from collections import defaultdict


def solution(map):

    fromA = defaultdict(lambda : float('inf'))
    fromB = defaultdict(lambda : float('inf'))

    r, c = len(map) , len(map[0])
    fromA[(0,0)] = 1

    q = [(0,0,1)]

    while q:
        x, y, score = q.pop(0)

        

        for (i,j) in [(0,1), (1,0),(-1,0), (0, -1)]:
            if 0 <= x + i < r and 0 <= y + j < c and fromA[(x + i,y + j)] == float('inf'):
                
                fromA[(x + i,y + j)] = score + 1

                if map[x + i][y + j] == 0 : q.append((x + i, y + j, score + 1))


    q = [(r - 1, c - 1, 1)]
    fromB[(r-1,c-1)] = 1

    while q:
        x, y, score = q.pop(0)


        for (i,j) in [(0,1), (1,0),(-1,0), (0, -1)]:
            if 0 <= x + i < r and 0 <= y + j < c and fromB[(x + i,y + j)] == float('inf'):
                
                fromB[(x + i,y + j)] = score + 1

                if map[x + i][y + j] == 0 : q.append((x + i, y + j, score + 1))


    minVal = float('inf')

    

    for i in range(r):
        for j in range(c):
            # print(fromA[(i,j)], fromB[(i,j)], fromA[(i,j)] + fromB[(i,j)])
            minVal = min(minVal, fromA[(i,j)] + fromB[(i,j)])


    return minVal - 1





print(solution([[0, 1, 1, 0], [0, 0, 0, 1], [1, 1, 0, 0], [1, 1, 1, 0]]))