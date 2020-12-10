def get_numbers(filePath):
    with open(filePath, "r") as fp:
        return [int(line.rstrip()) for line in fp.readlines()]

def get_device_joltage(numbers):
    numbers = sorted(numbers)
    numbers += (numbers[-1] + 3,)

    outlet = 0
    differences_in_jolts = {jolt: 0 for jolt in (1,2,3,)}

    for jolt in numbers:
        difference = jolt - outlet
        if difference < 1 or difference > 3:
            print("end reached!", difference, outlet)
            break

        differences_in_jolts[difference] += 1
        outlet += difference
    
    return outlet, differences_in_jolts

def get_value(dictionary, index):
    return dictionary.get(index) or 0

def get_sum_of_last_three(dictionary, index):
    return get_value(dictionary, index-1) + get_value(dictionary, index-2) + get_value(dictionary, index-3)

def get_solutions(numbers):
    numbers = sorted(numbers)
    solutions = {0: 1}

    for jolt in numbers:
        solutions[jolt] = get_sum_of_last_three(solutions, jolt)

    return solutions[numbers[-1]]

if __name__ == "__main__":
    numbers = get_numbers("10/input.txt")
    _, jolts_counter = get_device_joltage(numbers)
    print("device jolts:", jolts_counter[1] * jolts_counter[3])
    print("number of solutions:", get_solutions(numbers))
    
    
        



