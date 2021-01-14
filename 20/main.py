def parse_ingredients(line):
    parts = line.strip().split(" (")

    ingredients = set([v.strip() for v in parts[0].split(" ")])
    allergeens = set([v.strip() for v in parts[1][8:-1].split(",")])
    
    return {
        "ingredients": ingredients,
        "allergeens": allergeens
    }

def get_data():
    with open("20/input.txt") as f:
        return [parse_ingredients(line) for line in f.readlines()]

def get_possible_allergeens(data):
    possible_allergeens = {}

    for row in data:
        for allergeen in row["allergeens"]:
            possible_allergeens.setdefault(allergeen, set(row["ingredients"]))
            possible_allergeens[allergeen] &= row["ingredients"]
    
    return possible_allergeens

def identify_allergeens(possible_allergeens):
    allergeens = {}
    while len(possible_allergeens) > 0:
        for allergeen, ingredients in tuple(possible_allergeens.items()):
            if len(ingredients) > 1:
                continue

            ingredient = list(ingredients)[0]
            allergeens[allergeen] = ingredient
            del possible_allergeens[allergeen]
            for v in possible_allergeens.values():
                v.discard(ingredient)
            break
    return allergeens

if __name__ == "__main__":
    data = get_data()
    possible_allergeens = get_possible_allergeens(data)
    allergeens = identify_allergeens(possible_allergeens)

    counter = 0
    allergenic_ingredients = set(allergeens.values())
    for row in data:
        counter += len(row["ingredients"] - allergenic_ingredients)


    # part 1
    print(f"Number non allergenic ingredients: {counter}")
    print(f"Save ingredients: {allergeens}")

    # part 2
    print(",".join([ingredient for _, ingredient in sorted(allergeens.items())]))