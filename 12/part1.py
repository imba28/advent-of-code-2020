def get_instructions(filePath):
    with open(filePath, "r") as fp:
        return [{"direction": line[0], "units": int(line[1:])} for line in fp.readlines()]

DIRECTIONS = ('N', 'E', 'S', 'W')
COORDINATES = ([0,-1], [1,0], [0, 1], [-1, 0])

def get_new_position(currentPosition, currentDirection, direction, units):
    if direction in DIRECTIONS:
        coordinates = COORDINATES[DIRECTIONS.index(direction)]
        currentPosition = [currentPosition[0] + coordinates[0] * units, currentPosition[1] + coordinates[1] * units]
    elif direction == 'L':
        currentDirection = DIRECTIONS[(DIRECTIONS.index(currentDirection) - int(units / 90)) % len(COORDINATES)]
    elif direction == 'R':
        currentDirection = DIRECTIONS[(DIRECTIONS.index(currentDirection) + int(units / 90)) % len(COORDINATES)]
    elif direction == 'F':
        coordinates = COORDINATES[DIRECTIONS.index(currentDirection)]
        currentPosition = [currentPosition[0] + coordinates[0] * units, currentPosition[1] + coordinates[1] * units]
    
    return currentPosition, currentDirection


if __name__ == "__main__":
    instructions = get_instructions("12/input.txt")

    currentPosition = [0,0]
    currentDirection = 'E'

    for instruction in instructions:
        currentPosition, currentDirection = get_new_position(currentPosition, currentDirection, instruction["direction"], instruction["units"])

    print(abs(currentPosition[0]) + abs(currentPosition[1]))