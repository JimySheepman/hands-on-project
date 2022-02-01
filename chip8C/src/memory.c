#include <malloc.h>
#include "globals.h"

int memory_init(int memorysize)
{
    memory = (byte *)malloc(sizeof(byte) * memorysize);
    return memory != NULL;
}

void memory_destroy(void)
{
    free(memory);
}