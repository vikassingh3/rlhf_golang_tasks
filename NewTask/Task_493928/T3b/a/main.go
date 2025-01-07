package main

const batchSize = 100

func calculateYieldConcurrentlyBatch(fields []Field, resultCh chan []Result) {
    var batch []Result
    for _, field := range fields {
        result := Result{Field: field.Name, Yield: field.Yield * 10}
        batch = append(batch, result)
        if len(batch) == batchSize {
            resultCh <- batch
            batch = make([]Result, 0, batchSize)
        }
    }
    // Send the remaining batch if any
    if len(batch) > 0 {
        resultCh <- batch
    }
    close(resultCh)
}
// rest of the code remains the same