package logger

import "log"

// LogProcessedRows prints on every 10 lines imported
func LogProcessedRows(verbose bool, logMsg string){
	if !verbose {
		defer func() {
			log.Printf("100%s",logMsg)
		}()
	}

	for i := 1; i <= 100; i++ {
		if verbose && i%10 == 0 {
			log.Printf("%d%s", i,logMsg)
		}
	}

}


