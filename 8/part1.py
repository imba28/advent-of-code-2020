NOP = "nop"
ACC = "acc"
JMP = "jmp"

def get_instructions(filePath):
    with open(filePath, "r") as fp:
        return [{"command": line[:3], "value": int(line[4:])} for line in fp.readlines()]

def execute(instructions):
    executed_lines = {}
    acc = 0
    line = 0
    instruction = instructions[line]

    while instruction:
        if line in executed_lines:
            return 1, acc
        executed_lines[line] = True

        if instruction["command"] == NOP or instruction["command"] == ACC:
            line += 1

            if instruction["command"] == ACC:
                acc += instruction["value"]
        elif instruction["command"] == JMP:
            line += instruction["value"]
        
        instruction = instructions[line]

    return 0, acc

if __name__ == "__main__":
    acc = 0
    instructions = get_instructions("8/input.txt")
    status, acc = execute(instructions)

    print(status, acc)