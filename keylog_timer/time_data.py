def analyze_tex():
    file_op = open("timer_log.txt", "r")
    hour = 0
    minutes = 0
    sec = 0
    ms = 0
    time_count_ms = 0
    time_string = []
    time_integer = []
    press_key = []
    time_gap = [0]
    for ele in file_op:
        line = ele
        line = line.split()
        press_key.append(line[2])
        time_string.append(line[1])
    for item in range(0, len(time_string)):
        a = time_string[item]
        a = a.split(":")
        b = a[2].split(",")
        hour = int(a[0])
        minutes = int(a[1])
        sec = int(b[0])
        ms = int(b[1])
        time_count_ms = hour * 3600000 + minutes * 60000 + sec * 1000 + ms
        time_integer.append(time_count_ms)
        if item > 0:
            time_gap.append(time_integer[item] - time_integer[item - 1])
    file_op.close()
    for i in range(0, len(time_gap)):
        st = str(press_key[i]) + " " + str(time_string[i]) + " " + str(time_integer[i]) + " " + str(time_gap[i])
        with open("file.txt", "a") as file1:
            file1.writelines(st)
            file1.write("\n")


if __name__ == '__main__':
    analyze_tex()
