#include<stdio.h>
#include<stdlib.h>
#include<omp.h>

// Function to add elements to arrrays into another
double *vectAdd(double *c, double *a, double *b, int n){
    #pragma omp parallel for
    for(int i = 0; i < n; i++){
        c[i] = a[i] + b[i];
    }
    // Return the updated final array
    return c;
}