STATUS_OK = 0
STATUS_ERR = 1

NOP = "nop"
ACC = "acc"
JMP = "jmp"

def get_instructions(filePath):
    with open(filePath, "r") as fp:
        return [{"command": line[:3], "value": int(line[4:])} for line in fp.readlines()]

def execute(instructions):
    executed_lines = set()
    acc = 0
    line = 0
    instruction = instructions[line]

    while instruction and line not in executed_lines and line < len(instructions):
        instruction = instructions[line]
        executed_lines.add(line)

        if instruction["command"] == NOP or instruction["command"] == ACC:
            line += 1

            if instruction["command"] == ACC:
                acc += instruction["value"]
        elif instruction["command"] == JMP:
            line += instruction["value"]
    
    return STATUS_OK if line >= len(instructions) else STATUS_ERR, acc

def fix_instructions(instructions):
    for i, instruction in enumerate(instructions):
        if instruction["command"] not in (NOP, JMP):
            continue

        prev_command = instructions[i]["command"]
        instruction["command"] = JMP if prev_command == NOP else NOP

        status, acc = execute(instructions)
        if status == STATUS_OK:
            return STATUS_OK, acc

        instruction["command"] = prev_command

    return STATUS_ERR, -1

if __name__ == "__main__":
    instructions = get_instructions("8/input.txt")
    status, acc = fix_instructions(instructions)

    print(status, acc)