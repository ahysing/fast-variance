#ifndef __STATS_BLOCKING__
#define __STATS_BLOCKING__

#include <stdint.h>

void variance_uint32_loopunrolled(uint32_t buf[], int len, double *result);

#endif
