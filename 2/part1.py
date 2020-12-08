import re
p = re.compile(r'^(\d+)-(\d+)\s(\w):\s(\w+)$')

def count_valid_passwords(file_path):
    counter = 0

    with open(file_path) as f:
        for l in f.readlines():
            m = p.match(l.strip())
            lower_bound = int(m.group(1))
            upper_bound = int(m.group(2))
            char_check = m.group(3)
            password = m.group(4)

            occurrences = password.count(char_check)
            if lower_bound <= occurrences <= upper_bound:  
                counter += 1

    return counter

if __name__ == "__main__":
    print(count_valid_passwords("2/input.txt"))