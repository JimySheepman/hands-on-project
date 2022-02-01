#include <malloc.h>
#include "globals.h"

KEYSPEC keyboard_def[] =
    {
        {0x1, SDLK_4},
        {0x2, SDLK_5},
        {0x3, SDLK_6},
        {0x4, SDLK_r},
        {0x5, SDLK_t},
        {0x6, SDLK_y},
        {0x7, SDLK_f},
        {0x8, SDLK_g},
        {0x9, SDLK_h},
        {0xA, SDLK_v},
        {0xB, SDLK_b},
        {0xC, SDLK_7},
        {0xD, SDLK_u},
        {0xE, SDLK_j},
        {0xF, SDLK_n}};

SDLKey
keyboard_keycodetosymbol(int keycode)
{
    int i;

    for (i = 0; i < KEY_NUMBEROFKEYS; i++)
    {
        if (keyboard_def[i].keycode == keycode)
        {
            return keyboard_def[i].symbol;
        }
    }
    return SDLK_END;
}

int keyboard_symboltokeycode(SDLKey symbol)
{
    int i;

    for (i = 0; i < KEY_NUMBEROFKEYS; i++)
    {
        if (keyboard_def[i].symbol == symbol)
        {
            return keyboard_def[i].keycode;
        }
    }
    return KEY_NOKEY;
}

int keyboard_checkforkeypress(int keycode)
{
    Uint8 *keystates = SDL_GetKeyState(NULL);
    return keystates[keyboard_keycodetosymbol(keycode)];
}

int keyboard_waitgetkeypress(void)
{
    while (TRUE)
    {
        if (SDL_PollEvent(&event))
        {
            switch (event.type)
            {
            case SDL_KEYDOWN:
                return keyboard_symboltokeycode(event.key.keysym.sym);
                break;
                
            default:
                break;
            }
        }
        SDL_Delay(20);
    }
}