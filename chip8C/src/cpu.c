#include <time.h>
#include "globals.h"

Uint32
cpu_timerinterrupt(Uint32 interval, void *parameters)
{
    decrement_timers = TRUE;
    return interval;
}

int cpu_timerinit(void)
{
    int result = TRUE;
    SDL_InitSubSystem(SDL_INIT_TIMER);
    cpu_timer = SDL_AddTimer(17, cpu_timerinterrupt, NULL);

    if (cpu_timer == NULL)
    {
        printf("Error: could not create timer: %s\n", SDL_GetError());
        result = FALSE;
    }

    return result;
}

void cpu_reset(void)
{
    cpu.v[0x0] = 0;
    cpu.v[0x1] = 0;
    cpu.v[0x2] = 0;
    cpu.v[0x3] = 0;
    cpu.v[0x4] = 0;
    cpu.v[0x5] = 0;
    cpu.v[0x6] = 0;
    cpu.v[0x7] = 0;
    cpu.v[0x8] = 0;
    cpu.v[0x9] = 0;
    cpu.v[0xA] = 0;
    cpu.v[0xB] = 0;
    cpu.v[0xC] = 0;
    cpu.v[0xD] = 0;
    cpu.v[0xE] = 0;
    cpu.v[0xF] = 0;

    cpu.rpl[0x0] = 0;
    cpu.rpl[0x1] = 0;
    cpu.rpl[0x2] = 0;
    cpu.rpl[0x3] = 0;
    cpu.rpl[0x4] = 0;
    cpu.rpl[0x5] = 0;
    cpu.rpl[0x6] = 0;
    cpu.rpl[0x7] = 0;
    cpu.rpl[0x8] = 0;
    cpu.rpl[0x9] = 0;
    cpu.rpl[0xA] = 0;
    cpu.rpl[0xB] = 0;
    cpu.rpl[0xC] = 0;
    cpu.rpl[0xD] = 0;
    cpu.rpl[0xE] = 0;
    cpu.rpl[0xF] = 0;

    cpu.i.WORD = 0;
    cpu.sp.WORD = SP_START;
    cpu.dt = 0;
    cpu.st = 0;
    cpu.pc.WORD = CPU_PC_START;
    cpu.oldpc.WORD = CPU_PC_START;
    cpu.operand.WORD = 0;

    srand(time(0));
    cpu.state = CPU_PAUSED;

    cpu.opdesc = (char *)malloc(MAXSTRSIZE);
}

void cpu_process_sdl_events(void)
{
    if (SDL_PollEvent(&event))
    {
        switch (event.type)
        {
        case SDL_QUIT:
            cpu.state = CPU_STOP;
            break;

        case SDL_KEYDOWN:
            if (event.key.keysym.sym == QUIT_KEY)
            {
                cpu.state = CPU_STOP;
            }
            else if (event.key.keysym.sym == DEBUG_KEY)
            {
                cpu.state = CPU_DEBUG;
            }
            else if (event.key.keysym.sym == TRACE_KEY)
            {
                cpu.state = CPU_TRACE;
            }
            else if (event.key.keysym.sym == NORMAL_KEY)
            {
                cpu.state = CPU_RUNNING;
            }
            else if (event.key.keysym.sym == STEP_KEY)
            {
                cpu.state = CPU_STEP;
            }
            break;

        default:
            break;
        }
    }
}

void cpu_execute_single(void)
{
    byte x;
    byte y;
    byte src;
    byte tgt;
    byte tbyte;
    int temp;
    int i;
    int j;
    int k;
    int color;
    int currentcolor;
    word tword;
    int xcor;
    int ycor;

    cpu.oldpc = cpu.pc;

    cpu.operand.BYTE.high = memory_read(cpu.pc.WORD);
    cpu.pc.WORD++;
    cpu.operand.BYTE.low = memory_read(cpu.pc.WORD);
    cpu.pc.WORD++;

    switch (cpu.operand.BYTE.high & 0xF0)
    {
    case 0x00:
        switch (cpu.operand.BYTE.low)
        {
        case 0xC0:
        case 0xC1:
        case 0xC2:
        case 0xC3:
        case 0xC4:
        case 0xC5:
        case 0xC6:
        case 0xC7:
        case 0xC8:
        case 0xC9:
        case 0xCA:
        case 0xCB:
        case 0xCC:
        case 0xCD:
        case 0xCE:
        case 0xCF:
            temp = cpu.operand.BYTE.low & 0x0F;
            screen_scroll_down(temp);
            sprintf(cpu.opdesc, "SCRD %d", temp);
            break;

        case 0xE0:
            screen_blank();
            sprintf(cpu.opdesc, "CLS");
            break;

        case 0xEE:
            cpu.sp.WORD--;
            cpu.pc.BYTE.high = memory_read(cpu.sp.WORD);
            cpu.sp.WORD--;
            cpu.pc.BYTE.low = memory_read(cpu.sp.WORD);
            sprintf(cpu.opdesc, "RTS");
            break;

        case 0xFB:
            screen_scroll_right();
            sprintf(cpu.opdesc, "SCRR");
            break;

        case 0xFC:
            screen_scroll_left();
            sprintf(cpu.opdesc, "SCRL");
            break;

        case 0xFD:
            cpu.state = CPU_STOP;
            sprintf(cpu.opdesc, "EXIT");
            break;

        case 0xFE:
            screen_disable_extended();
            sprintf(cpu.opdesc, "EXTD");
            break;

        case 0xFF:
            screen_set_extended();
            sprintf(cpu.opdesc, "EXTE");
            break;

        default:
            break;
        }
        break;

    case 0x10:
        cpu.pc.WORD = (cpu.operand.WORD & 0x0FFF);
        sprintf(cpu.opdesc, "JUMP %03X", cpu.pc.WORD);
        break;

    case 0x20:
        memory_write(cpu.sp, cpu.pc.BYTE.low);
        cpu.sp.WORD++;
        memory_write(cpu.sp, cpu.pc.BYTE.high);
        cpu.sp.WORD++;
        cpu.pc.WORD = (cpu.operand.WORD & 0x0FFF);
        sprintf(cpu.opdesc, "CALL %03X", cpu.pc.WORD);
        break;

    case 0x30:
        src = cpu.operand.BYTE.high & 0xF;
        if (cpu.v[src] == cpu.operand.BYTE.low)
        {
            cpu.pc.WORD += 2;
        }
        sprintf(cpu.opdesc, "SKE V%X, %02X", src,
                cpu.operand.BYTE.low);
        break;

    case 0x40:
        src = cpu.operand.BYTE.high & 0xF;
        if (cpu.v[src] != cpu.operand.BYTE.low)
        {
            cpu.pc.WORD += 2;
        }
        sprintf(cpu.opdesc, "SKNE V%X, %02X", src,
                cpu.operand.BYTE.low);
        break;

    case 0x50:
        src = cpu.operand.BYTE.high & 0xF;
        tgt = (cpu.operand.BYTE.low & 0xF0) >> 4;
        if (cpu.v[src] == cpu.v[tgt])
        {
            cpu.pc.WORD += 2;
        }
        sprintf(cpu.opdesc, "SKE V%X, V%X", src, tgt);
        break;

    case 0x60:
        src = cpu.operand.BYTE.high & 0xF;
        cpu.v[src] = cpu.operand.BYTE.low;
        sprintf(cpu.opdesc, "LOAD V%X, %02X", src,
                cpu.operand.BYTE.low);
        break;

    case 0x70:
        src = cpu.operand.BYTE.high & 0xF;
        temp = cpu.v[src] + cpu.operand.BYTE.low;
        cpu.v[src] = (temp > 255) ? temp - 256 : temp;
        sprintf(cpu.opdesc, "ADD V%X, %02X", src,
                cpu.operand.BYTE.low);
        break;

    case 0x80:
        switch (cpu.operand.BYTE.low & 0x0F)
        {
        case 0x0:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            cpu.v[tgt] = cpu.v[src];
            sprintf(cpu.opdesc, "LOAD V%X, V%X", tgt, src);
            break;

        case 0x1:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            cpu.v[tgt] = cpu.v[tgt] | cpu.v[src];
            sprintf(cpu.opdesc, "OR V%X, V%X", tgt, src);
            break;

        case 0x2:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            cpu.v[tgt] = cpu.v[tgt] & cpu.v[src];
            sprintf(cpu.opdesc, "AND V%X, V%X", tgt, src);
            break;

        case 0x3:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            cpu.v[tgt] = cpu.v[tgt] ^ cpu.v[src];
            sprintf(cpu.opdesc, "XOR V%X, V%X", tgt, src);
            break;

        case 0x4:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            temp = cpu.v[src] + cpu.v[tgt];
            cpu.v[0xF] = (temp > 255) ? 1 : 0;
            cpu.v[tgt] = (temp > 255) ? temp - 256 : temp;
            sprintf(cpu.opdesc, "ADD V%X, V%X", tgt, src);
            break;

        case 0x5:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            if (cpu.v[tgt] > cpu.v[src])
            {
                cpu.v[0xF] = 1;
                cpu.v[tgt] -= cpu.v[src];
            }
            else
            {
                cpu.v[0xF] = 0;
                cpu.v[tgt] = 256 + cpu.v[tgt] - cpu.v[src];
            }
            sprintf(cpu.opdesc, "SUB V%X, V%X", tgt, src);
            break;

        case 0x6:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.v[0xF] = (cpu.v[src] & 1) ? 1 : 0;
            cpu.v[src] = cpu.v[src] >> 1;
            sprintf(cpu.opdesc, "SHR V%X", src);
            break;

        case 0x7:
            tgt = cpu.operand.BYTE.high & 0xF;
            src = (cpu.operand.BYTE.low & 0xF0) >> 4;
            if (cpu.v[src] < cpu.v[tgt])
            {
                cpu.v[0xF] = 1;
                cpu.v[tgt] = cpu.v[src] - cpu.v[tgt];
            }
            else
            {
                cpu.v[0xF] = 0;
                cpu.v[tgt] = 256 + cpu.v[src] - cpu.v[tgt];
            }
            sprintf(cpu.opdesc, "SUBN V%X, V%X", tgt, src);
            break;

        case 0xE:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.v[0xF] = (cpu.v[src] & 0x80) ? 1 : 0;
            cpu.v[src] = cpu.v[src] << 1;
            sprintf(cpu.opdesc, "SHL V%X", src);
            break;

        default:
            break;
        }
        break;

    case 0x90:
        src = cpu.operand.BYTE.high & 0xF;
        tgt = (cpu.operand.BYTE.low & 0xF0) >> 4;
        if (cpu.v[src] != cpu.v[tgt])
        {
            cpu.pc.WORD += 2;
        }
        sprintf(cpu.opdesc, "SKNE V%X, V%X", src, tgt);
        break;

    case 0xA0:
        cpu.i.WORD = (cpu.operand.WORD & 0x0FFF);
        sprintf(cpu.opdesc, "LOAD I, %03X", (cpu.operand.WORD & 0x0FFF));
        break;

    case 0xB0:
        cpu.pc.WORD = (cpu.operand.WORD & 0x0FFF) + cpu.i.WORD;
        sprintf(cpu.opdesc, "JUMP I + %03X", (cpu.operand.WORD & 0x0FFF));
        break;

    case 0xC0:
        tgt = cpu.operand.BYTE.high & 0xF;
        cpu.v[tgt] = (rand() % 255) & cpu.operand.BYTE.low;
        sprintf(cpu.opdesc, "RAND V%X, %02X", tgt,
                (cpu.operand.BYTE.low));
        break;

    case 0xD0:
        x = cpu.operand.BYTE.high & 0xF;
        y = (cpu.operand.BYTE.low & 0xF0) >> 4;
        tword.WORD = cpu.i.WORD;
        tbyte = cpu.operand.BYTE.low & 0xF;
        cpu.v[0xF] = 0;

        if (screen_extended_mode && tbyte == 0)
        {
            for (i = 0; i < 16; i++)
            {
                for (k = 0; k < 2; k++)
                {
                    tbyte = memory_read(cpu.i.WORD + (i * 2) + k);
                    ycor = cpu.v[y] + i;
                    ycor = ycor % screen_get_height();

                    for (j = 0; j < 8; j++)
                    {
                        xcor = cpu.v[x] + j + (k * 8);
                        xcor = xcor % screen_get_width();

                        color = (tbyte & 0x80) ? 1 : 0;
                        currentcolor = screen_getpixel(xcor, ycor);

                        cpu.v[0xF] = (currentcolor && color) ? 1 : cpu.v[0xF];
                        color = color ^ currentcolor;

                        screen_draw(xcor, ycor, color);
                        tbyte = tbyte << 1;
                    }
                }
            }
            sprintf(cpu.opdesc, "DRAWEX V%X, V%X, %X", x, y,
                    (cpu.operand.BYTE.low & 0xF));
        }
        else
        {
            for (i = 0; i < (cpu.operand.BYTE.low & 0xF); i++)
            {
                tbyte = memory_read(cpu.i.WORD + i);
                ycor = cpu.v[y] + i;
                ycor = ycor % screen_get_height();

                for (j = 0; j < 8; j++)
                {
                    xcor = cpu.v[x] + j;
                    xcor = xcor % screen_get_width();

                    color = (tbyte & 0x80) ? 1 : 0;
                    currentcolor = screen_getpixel(xcor, ycor);

                    cpu.v[0xF] = (currentcolor && color) ? 1 : cpu.v[0xF];
                    color = color ^ currentcolor;

                    screen_draw(xcor, ycor, color);
                    tbyte = tbyte << 1;
                }
            }
            sprintf(cpu.opdesc, "DRAW V%X, V%X, %X", x, y,
                    (cpu.operand.BYTE.low & 0xF));
        }

        if ((cpu.state != CPU_DEBUG) && (cpu.state != CPU_TRACE))
        {
            screen_refresh(FALSE);
        }
        break;

    case 0xE0:
        switch (cpu.operand.BYTE.low)
        {
        case 0x9E:
            src = cpu.operand.BYTE.high & 0xF;
            if (keyboard_checkforkeypress(cpu.v[src]))
            {
                cpu.pc.WORD += 2;
            }
            sprintf(cpu.opdesc, "SKPR V%X", src);
            break;

        case 0xA1:
            src = cpu.operand.BYTE.high & 0xF;
            if (!keyboard_checkforkeypress(cpu.v[src]))
            {
                cpu.pc.WORD += 2;
            }
            sprintf(cpu.opdesc, "SKUP V%X", src);
            break;

        default:
            break;
        }
        break;

    case 0xF0:
        switch (cpu.operand.BYTE.low)
        {
        case 0x07:
            tgt = cpu.operand.BYTE.high & 0xF;
            cpu.v[tgt] = cpu.dt;
            sprintf(cpu.opdesc, "LOAD V%X, DELAY", tgt);
            break;

        case 0x0A:
            tgt = cpu.operand.BYTE.high & 0xF;
            cpu.v[tgt] = keyboard_waitgetkeypress();
            sprintf(cpu.opdesc, "KEYD V%X", tgt);
            break;

        case 0x15:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.dt = cpu.v[src];
            sprintf(cpu.opdesc, "LOAD DELAY, V%X", src);
            break;

        case 0x18:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.st = cpu.v[src];
            sprintf(cpu.opdesc, "LOAD SOUND, V%X", src);
            break;

        case 0x1E:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.i.WORD += cpu.v[src];
            sprintf(cpu.opdesc, "ADD I, V%X", src);
            break;

        case 0x29:
            src = cpu.operand.BYTE.high & 0xF;
            cpu.i.WORD = cpu.v[src] * 5;
            sprintf(cpu.opdesc, "LOAD I, V%X", src);
            break;

        case 0x33:
            src = cpu.operand.BYTE.high & 0xF;

            tword.WORD = cpu.i.WORD;
            i = cpu.v[src] / 100;
            memory_write(tword, i);

            tword.WORD++;
            i = (cpu.v[src] % 100) / 10;
            memory_write(tword, i);

            tword.WORD++;
            i = (cpu.v[src] % 100) % 10;
            memory_write(tword, i);
            sprintf(cpu.opdesc, "BCD V%X (%03d)", src, cpu.v[src]);
            break;

        case 0x55:
            tword.WORD = cpu.i.WORD;
            for (i = 0; i <= (cpu.operand.BYTE.high & 0xF); i++)
            {
                memory_write(tword, cpu.v[i]);
                tword.WORD++;
            }
            sprintf(cpu.opdesc, "STOR %X", (cpu.operand.BYTE.high & 0xF));
            break;

        case 0x65:
            temp = cpu.i.WORD;
            tbyte = cpu.operand.BYTE.high & 0xF;
            for (i = 0; i <= tbyte; i++)
            {
                cpu.v[i] = memory_read(temp);
                temp++;
            }
            sprintf(cpu.opdesc, "LOAD %X", tbyte);
            break;

        case 0x75:
            tbyte = cpu.operand.BYTE.high & 0xF;
            for (i = 0; i <= tbyte; i++)
            {
                cpu.rpl[i] = cpu.v[i];
            }
            sprintf(cpu.opdesc, "SRPL %X", tbyte);
            break;

        case 0x85:
            tbyte = cpu.operand.BYTE.high & 0xF;
            for (i = 0; i <= tbyte; i++)
            {
                cpu.v[i] = cpu.rpl[i];
            }
            sprintf(cpu.opdesc, "LRPL %X", tbyte);
            break;

        default:
            break;
        }
        break;

    default:
        break;
    }
}

void cpu_execute(void)
{
    while (cpu.state != CPU_STOP)
    {
        cpu_execute_single();

        if ((cpu.state == CPU_DEBUG) || (cpu.state == CPU_TRACE))
        {
            screen_refresh(TRUE);
            while (cpu.state == CPU_DEBUG)
            {
                cpu_process_sdl_events();
                SDL_Delay(20);
            }
            cpu.state = (cpu.state == CPU_STEP) ? CPU_DEBUG : cpu.state;
        }
        else
        {
            SDL_Delay(op_delay);
        }

        if (decrement_timers)
        {
            cpu.dt -= (cpu.dt > 0) ? 1 : 0;
            cpu.st -= (cpu.st > 0) ? 1 : 0;
            decrement_timers = FALSE;
        }

        cpu_process_sdl_events();
    }
}
