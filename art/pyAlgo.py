from typing import List

def getPlusSignCount(N: int, L: List[int], D: str) -> int:
    pos = [0, 0]
    x_lines = []
    y_lines = []
    plus_sign_count = 0

    for stroke in range(0, len(L)):
        if D[stroke] == "U":
            y_lines.append([[pos[1], pos[1] + L[stroke]], pos[0]])
            pos[1] = pos[1] + L[stroke]
        elif D[stroke] == "D":
            y_lines.append([[pos[1], pos[1] - L[stroke]], pos[0]])
            pos[1] = pos[1] - L[stroke]
        elif D[stroke] == "L":
            x_lines.append([[pos[0], pos[0] - L[stroke]], pos[1]])
            pos[0] = pos[0] - L[stroke]
        elif D[stroke] == "R":
            x_lines.append([[pos[0], pos[0] + L[stroke]], pos[1]])
            pos[0] = pos[0] + L[stroke]

    # order line start / end values to be from least to greatest
    # in order to merge later
    for x in range(0, len(x_lines)):
        if x_lines[x][0][0] > x_lines[x][0][1]:
            x_lines[x] = [[x_lines[x][0][1], x_lines[x][0][0]], x_lines[x][1]]
    for y in range(0, len(y_lines)):
        if y_lines[y][0][0] > y_lines[y][0][1]:
            y_lines[y] = [[y_lines[y][0][1], y_lines[y][0][0]], y_lines[y][1]]


    # merge line that intersect
    def merge_lines(lines):
        merged_lines = []
        for line in lines:
            for existing_line in merged_lines:
                if line[1] == existing_line[1]:
                    if line[0][1] >= existing_line[0][0] and line[0][0] <= existing_line[0][1]:
                        existing_line[0] = [min(line[0][0], existing_line[0][0]), max(line[0][1], existing_line[0][1])]
                        break
            else:
                merged_lines.append(line)
        return merged_lines

    x_lines = merge_lines(x_lines)
    y_lines = merge_lines(y_lines)

    # get valid coordinates from both lines (ie don't include line start / end points
    vaLid_xline_coords = set()
    valid_yline_coords = set()

    for line in x_lines:
        for i in range(line[0][0] + 1, line[0][1]):
            vaLid_xline_coords.add((i, line[1]))

    for line in y_lines:
        for i in range(line[0][0] + 1, line[0][1]):
            valid_yline_coords.add((line[1], i))

    for coordinate in vaLid_xline_coords:
        if coordinate in valid_yline_coords:
            plus_sign_count += 1

    return plus_sign_count