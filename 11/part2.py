import copy

STATE_FLOOR = "."
STATE_OCCUPIED = "#"
STATE_EMPTY = "L"

def get_board(filePath):
    with open(filePath, "r") as fp:
        return [list(line.rstrip()) for line in fp.readlines()]

def get_adjacent_occupied_seats(board, y, x):
    adjanced_occupied_seats = 0

    for i in (-1, 0, 1):
        for j in (-1, 0, 1):
            if j == 0 and i == 0:
                continue

            next_y = y + i
            next_x = x + j
            while 0 <= next_x < len(board[y]) and 0 <= next_y < len(board):
                if board[next_y][next_x] == STATE_OCCUPIED:
                    adjanced_occupied_seats += 1
                    break
                if board[next_y][next_x] == STATE_EMPTY:
                    break
                
                next_x += j
                next_y += i

    return adjanced_occupied_seats

def get_seat_state(board, y, x):
    if board[y][x] == STATE_FLOOR:
        return STATE_FLOOR

    adjanced_occupied_seats = get_adjacent_occupied_seats(board, y, x)    

    if board[y][x] == STATE_EMPTY and adjanced_occupied_seats == 0:
        return STATE_OCCUPIED

    if board[y][x] == STATE_OCCUPIED and adjanced_occupied_seats >= 5:
        return STATE_EMPTY

    return board[y][x]

def run_round(board):
    new_board = copy.deepcopy(board)
    changed_states = 0

    for y in range(len(board)):
        for x in range(len(board[y])):
            new_state = get_seat_state(board, y, x)
            if new_state != board[y][x]:
                changed_states += 1
            new_board[y][x] = new_state

    return new_board, changed_states > 0

def count_state_type(board, state_type):
    counter = 0
    for y in range(len(board)):
        for x in range(len(board[y])):
            if board[y][x] == state_type:
                counter += 1
    return counter

if __name__ == "__main__":
    board = get_board("11/input.txt")
    has_changed = True

    while has_changed:
        board, has_changed = run_round(board)

    print(count_state_type(board, STATE_OCCUPIED))
