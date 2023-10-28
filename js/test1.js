const fs = require('fs/promises')

const INPUT_CSV_FILE = './data/sales.csv'
const OUTPUT_CSV_FILE = './js/salesWithTaxes.csv'
const BENCHMARK_CSV_FILE = './output/benchmarks.csv'

const logBench = false
const TAX_PERCENT = 0.05

function benchmark(tag) {
  const startTime = performance.now()
  const memoryStart = process.memoryUsage().heapUsed
  return () => {
    const totalTime = (performance.now() - startTime).toFixed(2) + ' ms'
    const memoryUsage =
      ((process.memoryUsage().heapUsed - memoryStart) / 1024 / 1024).toFixed(
        2
      ) + ' mb'

    if (logBench) {
      console.log(`------ BENCHMARK ${tag.toUpperCase()} RESULTS ------`)
      console.log(`| Time: ${totalTime}, Memory: ${memoryUsage} |`)
      console.log('------ END BENCHMARK RESULTS ------\n')
    }

    // Log to file
    fs.appendFile(
      BENCHMARK_CSV_FILE,
      `javascript,${tag},${totalTime},${memoryUsage}\n`
    ).catch((error) => {
      throw error
    })
  }
}

function readFile() {
  const benchInput = benchmark('read file')
  const sales = []

  fs.readFile(INPUT_CSV_FILE, 'utf8')
    .then((data) => {
      data.split('\r\n').forEach((line) => {
        sales.push(line.split(','))
      })
      processData(sales)
      benchInput()
    })
    .catch((error) => {
      throw error
    })
}

function processData(sales) {
  const benchProcess = benchmark('process data')
  const salesWithTaxes = sales.map((data, i) => {
    if (i > 0 && data.length === 3) {
      const price = Number(data[2].slice(1))
      const priceWithTax = (price + price * TAX_PERCENT).toFixed(2)
      data[2] = `$${priceWithTax}`
    }
    return data
  })
  const writeData = salesWithTaxes.reduce(
    (prev, curr) => `${prev}${curr.join(',')}\n`,
    ''
  )
  writeFile(writeData)
  benchProcess()
}

function writeFile(writeData) {
  const benchOutput = benchmark('write file')

  fs.writeFile(OUTPUT_CSV_FILE, writeData, 'utf8').catch((error) => {
    throw error
  })
  benchOutput()
}

function main() {
  readFile()
}
main()
