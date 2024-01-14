#include "csort.h"

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

void Sort(int *arr, int len)
{
    // Record start time
    clock_t start = clock();
    // Sort arr using qsort
    qsort(arr, len, sizeof(int), cmpfunc);
    // Record end time
    clock_t end = clock();
    // Calculate time elapsed
    double time_elapsed = (double)(end - start) / CLOCKS_PER_SEC * 1000;
    // Print time elapsed
    printf("C sort Time elapsed: %f\n", time_elapsed);
}