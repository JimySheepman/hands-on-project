#include <stdlib.h>
#include "globals.h"

int loadrom(char *romfilename, int offset)
{
    FILE *fp;
    int result = TRUE;

    fp = fopen(romfilename, "r");
    if (fp == NULL)
    {
        printf("Error: could not open ROM image: %s\n", romfilename);
        result = FALSE;
    }
    else
    {
        while (!feof(fp))
        {
            fread(&memory[offset], 1, 1, fp);
            offset++;
        }
        fclose(fp);
    }
    return result;
}

void print_help(void)
{
    printf("usage: yac8e [-h] [-s SCALE] [-d OP_DELAY] [-t] rom\n\n");
    printf("Starts a simple Chip 8 emulator. See README.md for more ");
    printf("information, and\n LICENSE for terms of use.\n\n");
    printf("positional arguments:\n");
    printf("  rom          the ROM file to load on startup\n\n");
    printf("optional arguments:\n");
    printf("  -h           show this help message and exit\n");
    printf("  -s SCALE     the scale factor to apply to the display ");
    printf("(default is 5)\n");
    printf("  -d OP_DELAY  sets the CPU operation to take at least the ");
    printf("specified number of milliseconds to execute (default is 1)\n");
    printf("  -t           starts the CPU up in trace mode\n");
}

char *
parse_options(int argc, char **argv)
{
    int arg;
    char *filename = NULL;

    for (arg = 1; arg < argc; arg++)
    {
        if ((argv[arg][0] == '-') && (strlen(argv[arg]) != 2))
        {
            printf("Unrecognized option: %s\n", argv[arg]);
            print_help();
            exit(1);
        }
        else if ((argv[arg][0] == '-') && (strlen(argv[arg]) == 2))
        {
            switch (argv[arg][1])
            {
            case ('h'):
                print_help();
                exit(0);
                break;

            case ('s'):
                arg++;
                if (arg < argc)
                {
                    scale_factor = atoi(argv[arg]);
                }
                break;

            case ('d'):
                arg++;
                if (arg < argc)
                {
                    op_delay = atoi(argv[arg]);
                }
                break;

            case ('t'):
                cpu.state = CPU_DEBUG;
                break;

            default:
                printf("Unrecognized option: %s\n", argv[arg]);
                print_help();
                exit(1);
                break;
            }
        }
        else
        {
            if (filename == NULL)
            {
                filename = argv[arg];
            }
            else
            {
                printf("Unrecognized parameter: %s\n", argv[arg]);
                print_help();
                exit(1);
            }
        }
    }

    if (filename == NULL)
    {
        printf("ROM file not specified\n");
        print_help();
        exit(1);
    }

    return filename;
}

int main(int argc, char **argv)
{
    char *filename;

    scale_factor = SCALE_FACTOR;
    cpu_reset();
    cpu.state = CPU_RUNNING;

    filename = parse_options(argc, argv);

    if (SDL_Init(SDL_INIT_VIDEO) < 0)
    {
        printf("Fatal: Unable to initialize SDL\n%s\n", SDL_GetError());
        exit(1);
    }

    if (!memory_init(MEM_4K))
    {
        printf("Fatal: Unable to allocate emulator memory\n");
        SDL_Quit();
        exit(1);
    }

    if (!loadrom("FONTS.chip8", 0))
    {
        printf("Fatal: Could not load FONTS.chip8\n");
        memory_destroy();
        SDL_Quit();
        exit(1);
    }

    if (!loadrom(filename, ROM_DEFAULT))
    {
        printf("Fatal: Emulator shutdown due to errors\n");
        memory_destroy();
        SDL_Quit();
        exit(1);
    }

    if (!screen_init())
    {
        printf("Fatal: Emulator shutdown due to errors\n");
        memory_destroy();
        SDL_Quit();
        exit(0);
    }

    if (!cpu_timerinit())
    {
        printf("Fatal: emulator shutdown due to errors\n");
        memory_destroy();
        SDL_Quit();
        exit(1);
    }

    cpu_execute();

    memory_destroy();
    screen_destroy();
    SDL_Quit();
    return 0;
}
