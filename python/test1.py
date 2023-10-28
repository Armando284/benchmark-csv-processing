import time
import psutil

INPUT_CSV_FILE = "./data/sales.csv"
OUTPUT_CSV_FILE = "./python/salesWithTaxes.csv"
BENCHMARK_CSV_FILE = "./output/benchmarks.csv"
TAX_PERCENT = 0.05
log_bench = False


def benchmark(tag):
    time_start = time.perf_counter()
    memory_start = psutil.virtual_memory().used

    def finish_benchmark():
        # must be taken to miliseconds
        time_usage = f"{round((time.perf_counter() - time_start) * 1000, 2)} ms"
        memory_usage = f"{round((psutil.virtual_memory().used - memory_start)/1024/1024, 2)} mb"
        if log_bench:
            print("-------- BENCHMARK", tag, "RESULTS --------")
            print("time:", time_usage, "memory:", memory_usage)
            print("-------- END BENCHMARK RESULTS --------\n")
        with open(BENCHMARK_CSV_FILE, 'a') as f:
            f.write(f"python,{tag},{time_usage},{memory_usage}\n")
    return finish_benchmark


def read_file():
    read_bench = benchmark("read file")
    with open(INPUT_CSV_FILE) as file:
        content = file.readlines()
    header = content[:1]
    sales = []
    for line in content[1:]:
        sales.append(line.split(","))
    read_bench()
    process_data(sales, header)


def process_data(sales, header):
    process_bench = benchmark("process data")
    sales_with_taxes = []
    for sale in sales:
        price = float(sale[2][1:])
        priceWithTax = round(price + (price * TAX_PERCENT), 2)
        sales_with_taxes.append([sale[0], sale[1], f"${priceWithTax}"])
    lines = [",".join(header)]
    for sale in sales_with_taxes:
        lines.append(",".join(sale)+"\n")
    write_data = "".join(lines)
    write_file(write_data)
    process_bench()


def write_file(write_data):
    write_bench = benchmark("write file")
    with open(OUTPUT_CSV_FILE, "w") as output:
        output.write(write_data)
    write_bench()


def main():
    read_file()


if __name__ == "__main__":
    main()
