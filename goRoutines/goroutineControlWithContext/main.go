package main

import (
  "context"
  "fmt"
  "math/rand"
  "os"
  "time"
)

type Transaction struct {
  Name string
  ID   int64
}

func ProcessTransactions(ctx context.Context, Transaction *Transaction) {
  // sign: func (c *valueCtx) Value(key interface{}) interface{}
  //here Value will just inform us about the status of Transaction
  completed := ctx.Value("completed").(chan struct{}) // type assertion
  for {
    select {
    case <-completed:
      fmt.Printf("\nThe Transaction of [%d] has been completed. \n", Transaction.ID)
      return
    case <-ctx.Done():
      // Done() causes: either by a Cancellation or a Deadline being exceeded
      if ctx.Err() == context.Canceled {
        fmt.Printf("\nThe Transaction [%d] has been canceled.\n", Transaction.ID)
        return
      } else if ctx.Err() == context.DeadlineExceeded {
        fmt.Printf("\nThe DELAY for Transaction has expired. The transaction with ID [%d] is being terminated. \n", Transaction.ID)
        os.Exit(0)
      }
    default:
      time.Sleep(time.Duration(rand.Int63n(0.15 * 1e9)))
    }
  }
}
func main() {
  // let's define a signalling channel
  completed := make(chan struct{})
  var ctx context.Context
  var cancel context.CancelFunc
  // let's configure our context 
  ctx = context.WithValue(context.Background(), "completed", completed)
  ctx, cancel = context.WithTimeout(ctx, 8*time.Second)

}
