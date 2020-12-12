def get_instructions(filePath):
    with open(filePath, "r") as fp:
        return [{"direction": line[0], "units": int(line[1:])} for line in fp.readlines()]

DIRECTIONS = ('N', 'E', 'S', 'W')
COORDINATES = ([0,1], [1,0], [0, -1], [-1, 0])

def rotate_left(vector, deg):
    if deg == 90:
        return [-vector[1], vector[0]]
    if deg == 180:
        return [-vector[0], -vector[1]]
    if deg == 270:
        return [vector[1], -vector[0]]

    return vector

def get_new_position(currentPosition, waypoint, direction, units):
    if direction in DIRECTIONS:
        coordinates = COORDINATES[DIRECTIONS.index(direction)]
        waypoint = [waypoint[0] + coordinates[0] * units, waypoint[1] + coordinates[1] * units]
    elif direction == 'L':
        waypoint = rotate_left(waypoint, units)
    elif direction == 'R':
        waypoint = rotate_left(waypoint, (360 - units) % 360)
    elif direction == 'F':
        currentPosition = [currentPosition[0] + waypoint[0] * units, currentPosition[1] + waypoint[1] * units]
    
    return currentPosition, waypoint


if __name__ == "__main__":
    instructions = get_instructions("12/input.txt")

    currentPosition = [0,0]
    waypoint = [10, 1]

    for instruction in instructions:
        currentPosition, waypoint = get_new_position(currentPosition, waypoint, instruction["direction"], instruction["units"])

    print(currentPosition)
    print(abs(currentPosition[0]) + abs(currentPosition[1]))