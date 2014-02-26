def check(arr,i,j):
    return arr[i][j] == 2

def look(arr, x1,y1,y2,i,j):
    if x1 < 0:
        return True
    return (i >= x1 and j >= y1 and i <= x1+y2-y1 and j <= y2) == check(arr, i, j)

def detect(inp):
    First = True
    startx=-1
    starty=-1
    endx=-1
    for i in range(len(inp)):
        for j in range(len(inp[i])):
            #got the first one
            if First and check(inp,i,j):
                startx=i
                starty=j
                while j<len(inp[i]) and check(inp, i, j):
                    j=j+1
                if j == len(inp[i]):
                    j=j-1
                endx=j
                First = False
                print startx,starty,endx
            if look(inp, startx, starty, endx, i, j) == False:
                print i,j
                return "NO"
    if First:
        return "NO"
    return "YES"

f=[[1,2,2,2,1],[1,2,2,2,1],[1,2,2,2,1]]
print detect(f)
