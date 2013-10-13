#include<iostream>
using namespace std;

int getC(int n) {
    return n & (n-1);
}

int getN(int n) {
    int c = getC(n);
    return n & ~c;
}

int getLastNumber(int n) {
    int k = n, val = 0;
    do {
        val = k;         
    } while ((k = getC(k))!=0);
    return val;
}

void insert(int *arr, int len, int k) {
    int y = getN(k);
    int l = k;
    while (k < len) {
        arr[k] += 1;
        k += k & -k;
    }
}

int getValX(int *arr, int len, int n) {
    int val = 0;
    while (n) {
        val += arr[n];
        n = getC(n);
    }
    return val;
}

int find(int *arr, int len, int m, int n) {
    return getValX(arr, len, n) - getValX(arr, len, m);                        
}

void printArr(int *arr, int len) {
    for (int i=0;i<len;i++) {
        if (arr[i] != 0) {
            cout << "index " << i << "value " << arr[i];
        }
    }
}

int main() {
    int arr[32] = {0};
    insert(arr, 32, 3);
    insert(arr, 32, 4);
    insert(arr, 32, 5);
    insert(arr, 32, 6);
    insert(arr, 32, 7);
    insert(arr, 32, 8);
    //printArr(arr, 32);
    cout << find(arr, 32, 4, 8);
}
