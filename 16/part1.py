import re

def _parse_int_list(string):
    return [int(v) for v in string.split(",")]

def other_tickets_state(line, data):
    if line == "\n":
        return None

    if "nearby tickets" not in line:
        if "other_tickets" not in data:
            data["other_tickets"] = []

        data["other_tickets"].append(_parse_int_list(line))
    
    return other_tickets_state

def my_tickets_state(line, data):
    if line == "\n":
        return other_tickets_state

    if "your ticket" not in line:
        assert "my_ticke" not in data

        data["my_ticket"] =_parse_int_list(line)

    return my_tickets_state

def rules_state(line, data):
    if line == "\n":
        return my_tickets_state
    
    if 'rules' not in data:
        data['rules'] = {} 

    match = re.match(r"^([\w\s]+):\s([\d\-\d]+)\D+([\d\-\d]+)$", line)
    if match:
        groups = match.groups()
        rule_name = groups[0]
        assert rule_name not in data['rules']

        ranges = [[int(v) for v in group.split("-")] for group in groups[1:]]
        data['rules'][rule_name] = ranges
    
    return rules_state

def get_data(filePath):
    with open(filePath, "r") as f:
        state = rules_state
        data = {}

        for line in f.readlines():
            state = state(line, data)
            if state == None:
                return data
    return data

def get_invalid_numbers(rules, tickets):
    invalid_numbers = []

    flatten = lambda t: [v for l in t for v in l]
    all_ranges = flatten([ranges for k, ranges in rules.items()])

    for ticket in tickets:
        for number in ticket:
            valid = False
            for ranges in all_ranges:
                if ranges[0] <= number <= ranges[1]:
                    valid = True
                    break
            if not valid:
                invalid_numbers.append(number)

    return invalid_numbers

if __name__ == "__main__":
    data = get_data("16/input")
    invalid_numbers = get_invalid_numbers(data["rules"], data["other_tickets"])

    print(sum(invalid_numbers))