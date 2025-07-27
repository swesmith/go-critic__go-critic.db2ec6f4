package checker_test

func _() {
	{
		var xs []int
		/*! rewrite as for-range so compiler can recognize this pattern */
		for i := 0; i < len(xs); i++ {
			xs[i] = 0
		}
	}
	{
		var xs []byte
		/*! rewrite as for-range so compiler can recognize this pattern */
		for i := 0; i < len(xs); i++ {
			xs[i] = Zero
		}
	}
	const Zero = 0
}
