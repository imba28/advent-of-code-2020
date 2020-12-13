from math import floor

def get_data(filePath):
    with open(filePath, "r") as f:
        data = f.readlines()
        return int(data[0]), [int(v) for v in data[1].split(",") if v != "x"]

if __name__ == "__main__":
    timestamp, ids = get_data("13/input.txt")

    next_departures = [floor(timestamp/id)*id+id for id in ids]
    earliest_departure = min(next_departures)
    earliest_departing_bus = ids[next_departures.index(earliest_departure)]

    print(f"bus {earliest_departing_bus} departs in {earliest_departure - timestamp} minutes ({earliest_departure})")
    print(earliest_departing_bus * (earliest_departure - timestamp))