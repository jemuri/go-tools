package limit

var limit = make(chan struct{}, 3)

func ParallelLimit(taskq chan func()) {

	for task := range taskq {
		go func() {
			defer func() {
				<-limit
			}()

			limit <- struct{}{}
			//如果 task() 发生 panic，那“许可证”可能就还不回去了，因此需要使用 defer 来保证
			task()
		}()
	}
}
