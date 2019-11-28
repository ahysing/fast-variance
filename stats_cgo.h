// +build !darwin
#ifdef CGO

#ifndef _STATS_CGO_H_
#define _STATS_CGO_H_

#include <stdint.h>
void variance_uint32(uint32_t buf[], int len, double *result);

#endif /* _STATS_CGO_H_ */

#endif /* CGO */
