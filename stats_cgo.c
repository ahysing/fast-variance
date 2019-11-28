#ifdef CGO

#include "stats_cgo.h"
#include <stdlib.h>

void variance_uint32(uint32_t buf[], int len, double *result) {
    uint32_t sum = 0;
    for(int i = 0; i < len; i++) {
        sum += buf[i];
    }
    
    double mean = (double)sum / (double)len;

    double accumulator = 0.0;
    for (int i = 0; i < len; i++) {
        double value = (double)buf[i];
		double diff = value - mean;
		accumulator += diff * diff;
	}
    
    accumulator /= (double)(len - 1);
    *result = accumulator;
}

#endif
