#include <Accelerate/Accelerate.h>
#include <stdint.h>
#include <stdlib.h>


void variance_uint32_darwin(uint32_t buf[], int len, double *result) {
    double* varianceVector = malloc(len * sizeof(double));
    if (varianceVector == NULL) {
        *result = -1.0;
        return;
    }

    //Converts an array of unsigned 32-bit integers to single-precision floating-point values.
    vDSP_Stride noStride = 1;
    vDSP_vfltu32D(buf, noStride, varianceVector, noStride, len);

    double mean = 0.0;
    double variance = 0.0;
    // variance: v = 1 / (n-1) * Sum (x - my)**2
    vDSP_meanvD(varianceVector, noStride, &mean, len); 
    double sqrtLengthMinusOne = sqrt((double)(len - 1));
    double oneOverSqrtLengthMinusOne = 1.0 / sqrtLengthMinusOne;
    double minusMeanOverSqrtLengthMinusOne = -mean / sqrtLengthMinusOne;

    // Multiplies the product of a double-precision vector and a double-precision scalar value by a double-precision scalar value.
    // void vDSP_vsmsaD(const double *__A, vDSP_Stride __IA, const double *__B, const double *__C, double *__D, vDSP_Stride __ID, vDSP_Length __N);
    // for (n = 0; n < N; ++n)
    //   D[n] = A[n]*B + C;
    vDSP_vsmsaD(varianceVector, noStride, &oneOverSqrtLengthMinusOne, &minusMeanOverSqrtLengthMinusOne, varianceVector, noStride, len);
    for (int i = 0; i < len; i++) {
        varianceVector[i] = (varianceVector[i] * varianceVector[i]);
    }

    // Calculates the sum of values in a double-precision vector.
    // void vDSP_sveD(const double *__A, vDSP_Stride __I, double *__C, vDSP_Length __N);
    vDSP_sveD(varianceVector, noStride, &variance, len);
    
    *result = variance;

    free(varianceVector);
}