
import heapq
def main():
    a = [0 for _ in range(9)]
    an = True
    for i  in range(123,333):
        k1 = 2*i
        k2 = 3*i
        a[0]=i%10
        a[1]=i//10%10
        a[2]=i//100
        a[3]=k1%10
        a[4]=k1//10%1
        a[5]=k1//100
        a[6]=k2%10
        a[7]=k2//10%10
        a[8]=k2//100
        a.sort()
        for j in range(9):
            if a[j]!=j+1:
                an = False
                break
        if an == True:
            print(i,2*i,3*i)
        an = True

if __name__ == '__main__':
    main()
    pass

