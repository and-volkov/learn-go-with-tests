package romannumerals

import "testing"

// Benchmarks for Conversion to Roman
func BenchmarkConvertToRoman(b *testing.B) {
	var results struct {
		RomanIterative float64
		RomanRecursive float64
	}
	b.Run("iterative", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range cases {
				ConvertToRoman(tt.Arabic)
			}
		}
		results.RomanIterative = float64(b.N)
	})
	b.Run("recursive", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range cases {
				ConvertToRomanRecursive(tt.Arabic)
			}
		}
		results.RomanRecursive = float64(b.N)
	})

	b.Logf("\nRoman Performance Summary:\n"+
		"  Iterative vs Recursive: %.2f%%\n",
		(results.RomanIterative/results.RomanRecursive-1)*100,
	)
}

// Benchmarks for Conversion to Arabic
func BenchmarkConertToArabic(b *testing.B) {
	var results struct {
		ArabicIterative float64
		ArabicRecursive float64
	}
	b.Run("iterative", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range cases {
				ConvertToArabic(tt.Roman)
			}
		}
		results.ArabicIterative = float64(b.N)
	})
	b.Run("recursive", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, tt := range cases {
				ConvertToArabicRecursive(tt.Roman)
			}
		}
		results.ArabicRecursive = float64(b.N)
	})
	b.Logf("\n Arabic Performance Summary:\n"+
		"  Iterative vs Recursive: %.2f%%\n",
		(results.ArabicIterative/results.ArabicRecursive-1)*100,
	)
}
