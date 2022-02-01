#include <SDL.h>
#include <SDL_ttf.h>
#include <stdlib.h>
#include "globals.h"

TTF_Font *message_font;

void init_colors(SDL_Surface *surface)
{
    COLOR_BLACK = SDL_MapRGB(surface->format, 0, 0, 0);
    COLOR_WHITE = SDL_MapRGB(surface->format, 250, 250, 250);
    COLOR_DGREEN = SDL_MapRGB(surface->format, 0, 70, 0);
    COLOR_LGREEN = SDL_MapRGB(surface->format, 0, 200, 0);
    COLOR_TEXT.r = 255;
    COLOR_TEXT.g = 255;
    COLOR_TEXT.b = 255;
}

int screen_getpixel(int x, int y)
{
    Uint8 r, g, b;
    Uint32 color = 0;
    int pixelcolor = 0;
    x = x * scale_factor;
    y = y * scale_factor;
    Uint32 *pixels = (Uint32 *)virtscreen->pixels;
    Uint32 pixel = pixels[(virtscreen->w * y) + x];
    SDL_GetRGB(pixel, virtscreen->format, &r, &g, &b);
    color = SDL_MapRGB(virtscreen->format, r, g, b);
    if (color == COLOR_WHITE)
    {
        pixelcolor = 1;
    }
    return pixelcolor;
}

void screen_blit_surface(SDL_Surface *src, SDL_Surface *dest, int x, int y)
{
    SDL_Rect location;
    location.x = x;
    location.y = y;
    SDL_BlitSurface(src, NULL, dest, &location);
}

SDL_Surface *
screen_render_message(char *msg, SDL_Color text_color)
{
    SDL_Surface *message_surface;

    message_surface = TTF_RenderText_Solid(message_font, msg, text_color);
    return message_surface;
}

void screen_trace_message(void)
{
    SDL_Surface *msg_surface;
    SDL_Rect box;

    box.x = 5;
    box.y = screen_height - 58;
    box.w = 342;
    box.h = 53;
    SDL_FillRect(overlay, &box, COLOR_LGREEN);

    box.x = 6;
    box.y = screen_height - 57;
    box.w = 340;
    box.h = 51;
    SDL_FillRect(overlay, &box, COLOR_DGREEN);

    char *buffer = (char *)malloc(MAXSTRSIZE);

    sprintf(buffer, "I:%04X DT:%02X ST:%02X PC:%04X %04X %s",
            cpu.i.WORD, cpu.dt, cpu.st, cpu.oldpc.WORD, cpu.operand.WORD,
            cpu.opdesc);
    msg_surface = screen_render_message(buffer, COLOR_TEXT);
    screen_blit_surface(msg_surface, overlay, 10, screen_height - 53);
    SDL_FreeSurface(msg_surface);

    sprintf(buffer,
            "V0:%02X V1:%02X V2:%02X V3:%02X V4:%02X V5:%02X V6:%02X V7:%02X",
            cpu.v[0], cpu.v[1], cpu.v[2], cpu.v[3], cpu.v[4], cpu.v[5],
            cpu.v[6], cpu.v[7]);
    msg_surface = screen_render_message(buffer, COLOR_TEXT);
    screen_blit_surface(msg_surface, overlay, 10, screen_height - 38);
    SDL_FreeSurface(msg_surface);

    sprintf(buffer,
            "V8:%02X V9:%02X VA:%02X VB:%02X VC:%02X VD:%02X VE:%02X VF:%02X",
            cpu.v[8], cpu.v[9], cpu.v[10], cpu.v[11], cpu.v[12], cpu.v[13],
            cpu.v[14], cpu.v[15]);
    msg_surface = screen_render_message(buffer, COLOR_TEXT);
    screen_blit_surface(msg_surface, overlay, 10, screen_height - 23);
    SDL_FreeSurface(msg_surface);

    free(buffer);
}

void screen_blank(void)
{
    screen_clear(virtscreen, COLOR_BLACK);
}

void screen_clear(SDL_Surface *surface, Uint32 color)
{
    SDL_Rect rect;
    rect.x = 0;
    rect.y = 0;
    rect.w = screen_width;
    rect.h = screen_height;
    SDL_FillRect(surface, &rect, color);
}

void screen_refresh(int overlay_on)
{
    screen_blit_surface(virtscreen, screen, 0, 0);
    if (overlay_on)
    {
        screen_trace_message();
        screen_blit_surface(overlay, screen, 0, 0);
    }
    screen_clear(overlay, COLOR_BLACK);
    SDL_UpdateRect(screen, 0, 0, 0, 0);
}

void screen_draw(int x, int y, int color)
{
    SDL_Rect pixel;
    Uint32 pixelcolor = COLOR_BLACK;

    pixel.x = x * scale_factor;
    pixel.y = y * scale_factor;
    pixel.w = scale_factor;
    pixel.h = scale_factor;
    if (color)
    {
        pixelcolor = COLOR_WHITE;
    }
    SDL_FillRect(virtscreen, &pixel, pixelcolor);
}

SDL_Surface *
screen_create_surface(int width, int height, int alpha, Uint32 color_key)
{
    Uint32 rmask, gmask, bmask, amask;
    SDL_Surface *tempsurface, *newsurface;

#if SDL_BYTEORDER == SDL_BIG_ENDIAN
    rmask = 0xff000000;
    gmask = 0x00ff0000;
    bmask = 0x0000ff00;
    amask = 0x000000ff;
#else
    rmask = 0x000000ff;
    gmask = 0x0000ff00;
    bmask = 0x00ff0000;
    amask = 0xff000000;
#endif

    tempsurface = SDL_CreateRGBSurface(SDL_SWSURFACE | SDL_SRCALPHA,
                                       screen_width, screen_height, SCREEN_DEPTH,
                                       rmask, gmask, bmask, amask);
    newsurface = SDL_DisplayFormat(tempsurface);
    SDL_FreeSurface(tempsurface);

    if (newsurface == NULL)
    {
        printf("Error: unable to create new surface\n");
    }
    else
    {
        SDL_SetAlpha(newsurface, SDL_SRCALPHA, alpha);
        if (color_key != -1)
        {
            SDL_SetColorKey(newsurface, SDL_SRCCOLORKEY, color_key);
        }
        screen_clear(newsurface, COLOR_BLACK);
    }

    return newsurface;
}

int screen_init(void)
{
    int result = FALSE;

    TTF_Init();
    message_font = TTF_OpenFont("VeraMono.ttf", 11);

    screen_width = (screen_extended_mode ? SCREEN_EXT_WIDTH : SCREEN_WIDTH) * scale_factor;
    screen_height = (screen_extended_mode ? SCREEN_EXT_HEIGHT : SCREEN_HEIGHT) * scale_factor;
    screen = SDL_SetVideoMode(screen_width,
                              screen_height,
                              SCREEN_DEPTH,
                              SDL_SWSURFACE);

    if (screen == NULL)
    {
        printf("Error: Unable to set video mode: %s\n", SDL_GetError());
    }
    else
    {
        SDL_SetAlpha(screen, SDL_SRCALPHA, 255);
        SDL_WM_SetCaption("YAC8 Emulator", NULL);
        init_colors(screen);
        virtscreen = screen_create_surface(screen_width, screen_height,
                                           255, -1);
        overlay = screen_create_surface(screen_width, screen_height,
                                        200, COLOR_BLACK);
        result = TRUE;
    }

    return ((virtscreen != NULL) && (overlay != NULL) && result);
}

void screen_destroy(void)
{
    SDL_FreeSurface(virtscreen);
    SDL_FreeSurface(overlay);
    SDL_FreeSurface(screen);
}

void screen_set_extended(void)
{
    screen_destroy();
    screen_extended_mode = TRUE;
    screen_init();
}

void screen_disable_extended(void)
{
    screen_destroy();
    screen_extended_mode = FALSE;
    screen_init();
}

void screen_scroll_left(void)
{
    SDL_Rect source_rect, dest_rect;

    int width = screen_get_width() * scale_factor;
    int height = screen_get_height() * scale_factor;

    SDL_Surface *tempsurface = screen_create_surface(width, height, 255, -1);

    source_rect.x = 0;
    source_rect.y = 0;
    source_rect.w = width;
    source_rect.h = height;

    dest_rect.x = (-4 * scale_factor);
    dest_rect.y = 0;
    dest_rect.w = 0;
    dest_rect.h = 0;

    SDL_BlitSurface(virtscreen, &source_rect, tempsurface, &dest_rect);
    SDL_FreeSurface(virtscreen);
    virtscreen = tempsurface;
}

void screen_scroll_right(void)
{
    SDL_Rect source_rect, dest_rect;

    int width = screen_get_width() * scale_factor;
    int height = screen_get_height() * scale_factor;

    SDL_Surface *tempsurface = screen_create_surface(width, height, 255, -1);

    source_rect.x = 0;
    source_rect.y = 0;
    source_rect.w = width - (4 * scale_factor);
    source_rect.h = height;

    dest_rect.x = (4 * scale_factor);
    dest_rect.y = 0;
    dest_rect.w = 0;
    dest_rect.h = 0;

    SDL_BlitSurface(virtscreen, &source_rect, tempsurface, &dest_rect);
    SDL_FreeSurface(virtscreen);
    virtscreen = tempsurface;
}

void screen_scroll_down(int num_pixels)
{
    SDL_Rect source_rect, dest_rect;

    int width = screen_get_width() * scale_factor;
    int height = screen_get_height() * scale_factor;

    SDL_Surface *tempsurface = screen_create_surface(width, height, 255, -1);

    source_rect.x = 0;
    source_rect.y = 0;
    source_rect.w = width;
    source_rect.h = height;

    dest_rect.x = 0;
    dest_rect.y = (num_pixels * scale_factor);
    dest_rect.w = 0;
    dest_rect.h = 0;

    SDL_BlitSurface(virtscreen, &source_rect, tempsurface, &dest_rect);
    SDL_FreeSurface(virtscreen);
    virtscreen = tempsurface;
}

int screen_get_height(void)
{
    return screen_extended_mode ? SCREEN_EXT_HEIGHT : SCREEN_HEIGHT;
}

int screen_get_width(void)
{
    return screen_extended_mode ? SCREEN_EXT_WIDTH : SCREEN_WIDTH;
}
