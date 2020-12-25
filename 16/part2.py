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

def is_valid_ticket(ranges, ticket):
    for number in ticket:
        valid = False
        for r in ranges:
            if r[0] <= number <= r[1]:
                valid = True
                break
        if not valid:
            return False
    return True

flatten = lambda t: [v for l in t for v in l]


def get_valid_tickets(rules, tickets):
    all_ranges = flatten([ranges for k, ranges in rules.items()])

    valid_tickets = list(filter(lambda ticket: is_valid_ticket(all_ranges, ticket), tickets))

    return valid_tickets

def satisfies_rule(number, ranges):
    valid = False
    for r in ranges:
        if r[0] <= number <= r[1]:
            valid = True
            break
    return valid

def identify_columns(rules, tickets):
    column_to_rules = [[] for i in range(len(tickets[0]))]

    for rule, ranges in rules.items():
        for column in range(len(tickets[0])):
            satisfies_all_tickets = True
            for ticket in tickets:
                if not satisfies_rule(ticket[column], ranges):
                    satisfies_all_tickets = False
                    break
            if satisfies_all_tickets:
                column_to_rules[column] += [rule]

    sorted_rules_to_column = sorted(column_to_rules, key=lambda v: len(v))
    for column, _ in enumerate(sorted_rules_to_column):
        if len(sorted_rules_to_column[column]) == 1:
            rule = sorted_rules_to_column[column][0]

            for i in range(column+1, len(sorted_rules_to_column)):
                if rule in sorted_rules_to_column[i]:
                    sorted_rules_to_column[i].remove(rule)
    return flatten(column_to_rules)

            
if __name__ == "__main__":
    data = get_data("16/input")
    valid_tickets = get_valid_tickets(data["rules"], data["other_tickets"])

    columns = identify_columns(data["rules"], valid_tickets)

    departure_columns = []
    for i, column in enumerate(columns):
        if 'departure' in column:
            departure_columns.append(i)

    result = 1
    for i in departure_columns:
        result *= data["my_ticket"][i]

    print(result)