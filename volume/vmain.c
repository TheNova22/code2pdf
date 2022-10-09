#include<stdio.h>
#include<stdlib.h>
#include<omp.h>
#include "vadd.c"

int main(){
    int n = 10;
    // Create three arrays
    double a[20], b[20], c[20];

    for (int i = 0; i < n; i++){
        // Assign random values
        a[i] = rand() % n;
        b[i] = rand() % n;
    }

    for (int i = 0; i < n; i++){
        // Print the summations
        printf("%f\n", vectAdd(c,a,b,n)[i]);
    }
    return 0;
}