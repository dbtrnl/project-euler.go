package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLargestPrimeFactor(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 5},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, test := range tests {
		actual := FindLargestPrimeFactor(test.input)
		assert.Equal(t, test.expected, actual, "For input %d", test.input)
	}
}

func TestFindSumOfNumberIntervalSquares(t *testing.T) {
	assert.Equal(t, 385, FindSumOfNumberIntervalSquares(1, 10))
}

func TestFindSquareOfNumberIntervalSum(t *testing.T) {
	assert.Equal(t, 3025, FindSquareOfNumberIntervalSum(1, 10))
}

func TestFindNthPrime(t *testing.T) {
	assert.Equal(t, 2, FindNthPrime(1))
	assert.Equal(t, 11, FindNthPrime(5))
	assert.Equal(t, 229, FindNthPrime(50))
	assert.Equal(t, 86719, FindNthPrime(8428))
	assert.Equal(t, 104743, FindNthPrime(10001))
}

func TestFindProductOfDigitsInNumberSeries(t *testing.T) {
	assert.Equal(t, 6, FindProductOfDigitsInNumberSeries("123"))
	assert.Equal(t, 120, FindProductOfDigitsInNumberSeries("12345"))
	assert.Equal(t, 4354560, FindProductOfDigitsInNumberSeries("8652393278"))
	assert.Equal(t, 0, FindProductOfDigitsInNumberSeries("86523932780"))
}

func TestFindAllPrimesSmallerThan(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	}
	assert.Equal(t, expected, FindAllPrimesSmallerThan(174))
	assert.Equal(t, expected, FindAllPrimesSmallerThan(179))

	expected = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179,
	}
	assert.Equal(t, expected, FindAllPrimesSmallerThan(180))
}

func TestFindAllProperDivisorsOf(t *testing.T) {
	expected := []int{1}
	result, err := FindAllProperDivisorsOf(1)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1}
	result, err = FindAllProperDivisorsOf(2)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4}
	result, err = FindAllProperDivisorsOf(8)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4, 5, 10, 20, 25, 50}
	result, err = FindAllProperDivisorsOf(100)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4, 5, 8, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500}
	result, err = FindAllProperDivisorsOf(1000)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 7, 11, 13, 77, 91, 143}
	result, err = FindAllProperDivisorsOf(1001)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 17, 59}
	result, err = FindAllProperDivisorsOf(1003)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	result, err = FindAllProperDivisorsOf(-1)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid value: -1")
}

func TestFindCollatzSequenceOf(t *testing.T) {
	result := FindCollatzSequenceOf(5)
	expected := []int{5, 16, 8, 4, 2, 1}
	assert.Equal(t, result, expected)

	result = FindCollatzSequenceOf(1236)
	expected = []int{1236, 618, 309, 928, 464, 232, 116, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	assert.Equal(t, result, expected)

	result = FindCollatzSequenceOf(784923)
	expected = []int{784923, 2354770, 1177385, 3532156, 1766078, 883039, 2649118, 1324559, 3973678, 1986839, 5960518, 2980259, 8940778, 4470389, 13411168, 6705584, 3352792, 1676396, 838198, 419099, 1257298, 628649, 1885948, 942974, 471487, 1414462, 707231, 2121694, 1060847, 3182542, 1591271, 4773814, 2386907, 7160722, 3580361, 10741084, 5370542, 2685271, 8055814, 4027907, 12083722, 6041861, 18125584, 9062792, 4531396, 2265698, 1132849, 3398548, 1699274, 849637, 2548912, 1274456, 637228, 318614, 159307, 477922, 238961, 716884, 358442, 179221, 537664, 268832, 134416, 67208, 33604, 16802, 8401, 25204, 12602, 6301, 18904, 9452, 4726, 2363, 7090, 3545, 10636, 5318, 2659, 7978, 3989, 11968, 5984, 2992, 1496, 748, 374, 187, 562, 281, 844, 422, 211, 634, 317, 952, 476, 238, 119, 358, 179, 538, 269, 808, 404, 202, 101, 304, 152, 76, 38, 19, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	assert.Equal(t, result, expected)
}

func TestFindLongestCollatzSequenceUnder(t *testing.T) {
	result := FindLongestCollatzSequenceUnder(5)
	expected := []int{3, 10, 5, 16, 8, 4, 2, 1}
	assert.Equal(t, result, expected)

	result = FindLongestCollatzSequenceUnder(1000)
	expected = []int{871, 2614, 1307, 3922, 1961, 5884, 2942, 1471, 4414, 2207, 6622, 3311, 9934, 4967, 14902, 7451, 22354, 11177, 33532, 16766, 8383, 25150, 12575, 37726, 18863, 56590, 28295, 84886, 42443, 127330, 63665, 190996, 95498, 47749, 143248, 71624, 35812, 17906, 8953, 26860, 13430, 6715, 20146, 10073, 30220, 15110, 7555, 22666, 11333, 34000, 17000, 8500, 4250, 2125, 6376, 3188, 1594, 797, 2392, 1196, 598, 299, 898, 449, 1348, 674, 337, 1012, 506, 253, 760, 380, 190, 95, 286, 143, 430, 215, 646, 323, 970, 485, 1456, 728, 364, 182, 91, 274, 137, 412, 206, 103, 310, 155, 466, 233, 700, 350, 175, 526, 263, 790, 395, 1186, 593, 1780, 890, 445, 1336, 668, 334, 167, 502, 251, 754, 377, 1132, 566, 283, 850, 425, 1276, 638, 319, 958, 479, 1438, 719, 2158, 1079, 3238, 1619, 4858, 2429, 7288, 3644, 1822, 911, 2734, 1367, 4102, 2051, 6154, 3077, 9232, 4616, 2308, 1154, 577, 1732, 866, 433, 1300, 650, 325, 976, 488, 244, 122, 61, 184, 92, 46, 23, 70, 35, 106, 53, 160, 80, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	assert.Equal(t, result, expected)
}

func TestFactorial(t *testing.T) {
	expected := 120
	result, err := Factorial(5)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = 3628800
	result, err = Factorial(10)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestFactorialBig(t *testing.T) {
	expected := big.NewInt(0)

	expected.SetString("1", 10)
	result, err := FactorialBig(1)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected.SetString("2", 10)
	result, err = FactorialBig(2)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected.SetString("120", 10)
	result, err = FactorialBig(5)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected.SetString("51090942171709440000", 10)
	result, err = FactorialBig(21)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected.SetString("265252859812191058636308480000000", 0)
	result, err = FactorialBig(30)
	assert.Nil(t, err)
	assert.Equal(t, result, expected)
}
