import re
p = re.compile(r'^(\d+)-(\d+)\s(\w):\s(\w+)$')

def count_valid_passwords(file_path):
    counter = 0

    with open(file_path) as f:
        for l in f.readlines():
            m = p.match(l.strip())
            first_index = int(m.group(1))
            second_index = int(m.group(2))
            char_check = m.group(3)
            password = m.group(4)

            match_first = password[first_index - 1] == char_check
            match_second = password[second_index - 1] == char_check

            if match_first and match_second:
                continue
            elif match_first or match_second:
                counter += 1

    return counter

if __name__ == "__main__":
    print(count_valid_passwords("2/input.txt"))