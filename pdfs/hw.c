#include<omp.h>
#include<stdio.h>
#include<stdlib.h>

int main(){
    
    // Init a parallel section
	#pragma omp parallel	
	{
		// Finding thread number
		int t = omp_get_thread_num();

		// output
		printf("Hello world from thread number %d \n",t);
	}
	
	return 0;
}

// gcc hw.c -fopenmp
// ./a.out