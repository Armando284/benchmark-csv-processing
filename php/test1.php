<?php

$INPUT_CSV_FILE = "./data/sales.csv";
$OUTPUT_CSV_FILE = "./php/salesWithTaxes.csv";
$BENCHMARK_CSV_FILE = "./output/benchmarks.csv";
$GLOBALS['outputPath'] = $BENCHMARK_CSV_FILE;
$TAX_PERCENT = 0.05;

$GLOBALS['logBench'] = false;

function benchmark(string $tag, float $timeStart, float $memoryStart)
{
  $timeUsage = number_format((microtime(true) - $timeStart) * 1000, 2) . " ms";
  $memoryUsage = number_format((memory_get_usage() - $memoryStart) / 1024 / 1024, 2) . " mb";
  if ($GLOBALS['logBench']) {
    echo ("-------- BENCHMARK " . strtoupper($tag) . " RESULTS --------" . "\n");
    echo ("time: " . $timeUsage . " memory: " . $memoryUsage . "\n");
    echo ("-------- END BENCHMARK  RESULTS --------" . "\n" . "\n");
  }

  $fp = fopen($GLOBALS["outputPath"], 'a');
  fputcsv($fp, ["php", $tag, $timeUsage, $memoryUsage]);
  fclose($fp);
}

$timeStart = microtime(true);
$memoryStart = memory_get_usage();

$rows = [];
$handle = fopen($INPUT_CSV_FILE, "r"); // open in readonly mode
while (($row = fgetcsv($handle)) !== false) {
  $rows[] = $row;
}
fclose($handle);

benchmark("read file", $timeStart, $memoryStart);

$timeStart = microtime(true);
$memoryStart = memory_get_usage();

$headers = array_slice($rows, 0, 1);
$sales = array_slice($rows, 1);

$salesWithTaxes = [];

foreach ($sales as $row) {
  // parse string money format text into float for calculation
  $price = floatval(substr($row[2], 1));
  $addTax = number_format($price + $price * $TAX_PERCENT, 2);
  $salesWithTaxes[] = [$row[0], $row[1], '$' . $addTax];
}

$fullData = array_merge($headers, $salesWithTaxes);

benchmark("process data",  $timeStart, $memoryStart);

$timeStart = microtime(true);
$memoryStart = memory_get_usage();

$fp = fopen($OUTPUT_CSV_FILE, 'w');
foreach ($fullData as $row) {
  fputcsv($fp, $row);
}
fclose($fp);

benchmark("write file",  $timeStart, $memoryStart);
