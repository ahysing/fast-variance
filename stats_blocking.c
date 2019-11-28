#include <stdint.h>

#define SIZE 16

void variance_uint32_loopunrolled(uint32_t buf[], int len, double *result) {
    double intermediate[SIZE];

    uint32_t sum = 0;
    for (int i = 0; i < len; i++) {
        uint32_t value = buf[i];
        sum += value;
    }

    double mean = sum / ((double)len);


    double accumulator = 0.0;
    for (int j = 0; j < len / SIZE; j++) {
        for (int i = 0; i < SIZE; i++) {
            intermediate[i] = (double)buf[(j * SIZE) + i];
        }
        for (int i = 0; i < SIZE; i++) {
            double value = intermediate[i];
            double diff = value - mean;
            intermediate[i] = diff;
        }
        for (int i = 0; i < SIZE; i++) {
            double diff = intermediate[i];
            accumulator += diff * diff;
        }
    }

    int start = (len - (len % SIZE));
    for (int i = start; i < len; i++) {
        intermediate[i - start] = (double)buf[i];
    }

    int end = (len % SIZE);
    for (int i = 0; i < end; i++) {
        double value = intermediate[i];
        double diff = value - mean;
        intermediate[i] = diff;
    }

    for (int i = 0; i < end; i++) {
        double diff = intermediate[i];
        accumulator += diff * diff;
    }

    double variance = accumulator / (double)(len - 1);
    *result = variance;
}
