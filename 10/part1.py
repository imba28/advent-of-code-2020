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

if __name__ == "__main__":
    numbers = get_numbers("10/input.txt")
    _, jolts_counter = get_device_joltage(numbers)
    print(jolts_counter[1] * jolts_counter[3])
    
        



