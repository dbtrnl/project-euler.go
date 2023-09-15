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

func TestFindAllDivisorsOf(t *testing.T) {
	expected := []int{}
	result, err := FindAllDivisorsOf(0)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1}
	result, err = FindAllDivisorsOf(1)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2}
	result, err = FindAllDivisorsOf(2)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 3, 4, 6, 9, 12, 18, 36}
	result, err = FindAllDivisorsOf(36)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4, 5, 8, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500, 1000}
	result, err = FindAllDivisorsOf(1000)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 7, 11, 13, 77, 91, 143, 1001}
	result, err = FindAllDivisorsOf(1001)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 17, 59, 1003}
	result, err = FindAllDivisorsOf(1003)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	result, err = FindAllDivisorsOf(-1)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid value: -1")
}

func TestFindAllProperDivisorsOf(t *testing.T) {
	expected := []int{}
	result, err := FindAllProperDivisorsOf(0)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1}
	result, err = FindAllProperDivisorsOf(1)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1}
	result, err = FindAllProperDivisorsOf(2)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 3, 4, 6, 9, 12, 18, 36}
	result, err = FindAllDivisorsOf(36)
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

func TestReturnWrittenNumber(t *testing.T) {
	expected := "zero"
	result := ReturnNumberString(0)
	assert.Equal(t, expected, result)

	expected = "eleven"
	result = ReturnNumberString(11)
	assert.Equal(t, expected, result)

	expected = "one hundred"
	result = ReturnNumberString(100)
	assert.Equal(t, expected, result)

	expected = "one hundred and eleven"
	result = ReturnNumberString(111)
	assert.Equal(t, expected, result)

	expected = "one thousand"
	result = ReturnNumberString(1000)
	assert.Equal(t, expected, result)

	expected = "one thousand and one"
	result = ReturnNumberString(1001)
	assert.Equal(t, expected, result)

	expected = "one thousand and eleven"
	result = ReturnNumberString(1011)
	assert.Equal(t, expected, result)

	expected = "one thousand one hundred and eleven"
	result = ReturnNumberString(1111)
	assert.Equal(t, expected, result)

	/* TODO
	expected = "ten thousand"
	result = ReturnNumberString(10000)
	assert.Equal(t, expected, result)

	expected = "ten thousand and one"
	result = ReturnNumberString(10001)
	assert.Equal(t, expected, result)

	expected = "ten thousand eleven"
	result = ReturnNumberString(10011)
	assert.Equal(t, expected, result)

	expected = "ten thousand one hundred eleven"
	result = ReturnNumberString(10111)
	assert.Equal(t, expected, result)

	expected = "eleven thousand one hundred eleven"
	result = ReturnNumberString(11111)
	assert.Equal(t, expected, result)
	*/
}

func TestCountLettersInString(t *testing.T) {
	expected := 23
	result := CountLettersInString("three hundred and forty-two")
	assert.Equal(t, expected, result)
	result = CountLettersInString("2$three hundred and forty-two")
	assert.Equal(t, expected, result)

	expected = 20
	result = CountLettersInString("one hundred and fifteen")
	assert.Equal(t, expected, result)

	expected = 104
	result = CountLettersInString("AÁÅĆDΔEĔƐFĞHĦIÍJЙҖҘКĿMÑØΦPQŘŞŠTŦUÚVWXYÝŽαбcדΣεƒGнιjκlмиσρQяsтуυvωxчzאकbգԴΣϜհιյkլոpҩrՏτմվwхyzあ俄бхひสשָგم你여П")
	assert.Equal(t, expected, result)
}

func TestFindAndSumAllProperDivisorsOf(t *testing.T) {
	expected := 55
	result, err := FindAndSumAllProperDivisorsOf(36)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = 185
	result, err = FindAndSumAllProperDivisorsOf(2839)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = 14304
	result, err = FindAndSumAllProperDivisorsOf(5040)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestIsNumberAmicable(t *testing.T) {
	expected := &AmicableNumberObject{
		IsAmicable: false,
		Numbers:    [2]int{0, 0},
	}
	result, err := IsNumberAmicable(6)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = &AmicableNumberObject{
		IsAmicable: true,
		Numbers:    [2]int{220, 284},
	}
	result, err = IsNumberAmicable(220)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = &AmicableNumberObject{
		IsAmicable: false,
		Numbers:    [2]int{0, 0},
	}
	result, err = IsNumberAmicable(221)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestCalculateStringScore(t *testing.T) {
	expected := 27
	result := CalculateStringScore("az")
	assert.Equal(t, expected, result)
	result = CalculateStringScore(" 23 $;,. AZ ")
	assert.Equal(t, expected, result)
}

func TestGetNumberAbundancy(t *testing.T) {
	expected := "abundant"
	result, err := GetNumberAbundancy(12)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	result, err = GetNumberAbundancy(5775)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = "deficient"
	result, err = GetNumberAbundancy(10)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	result, err = GetNumberAbundancy(13)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = "perfect"
	result, err = GetNumberAbundancy(28)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	result, err = GetNumberAbundancy(33550336)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	result, err = GetNumberAbundancy(137438691328)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	result, err = GetNumberAbundancy(-1)
	assert.Equal(t, result, "")
	assert.EqualError(t, err, "number: -1 must be >= 0.")
}

func TestFindAbundantNumbersUntil(t *testing.T) {
	expected := []int{12, 18, 20, 24, 30, 36, 40, 42, 48, 54}
	result, err := FindAbundantNumbersUntil(54)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = []int{12, 18, 20, 24, 30, 36, 40, 42, 48, 54, 56, 60, 66, 70, 72, 78, 80, 84, 88, 90, 96, 100, 102, 104, 108, 112,
		114, 120, 126, 132, 138, 140, 144, 150, 156, 160, 162, 168, 174, 176, 180, 186, 192, 196, 198, 200, 204, 208, 210, 216, 220,
		222, 224, 228, 234, 240, 246, 252, 258, 260, 264, 270}
	result, err = FindAbundantNumbersUntil(270)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	result, err = FindAbundantNumbersUntil(-1)
	assert.Nil(t, result)
	assert.EqualError(t, err, "number: -1 must be >= 0.")
}

func TestFindUniqueCombinatorialSums(t *testing.T) {
	expected := map[int]bool(map[int]bool{2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true})
	result := FindUniqueCombinatorialSums([]int{1, 2, 3, 4}, 999)
	assert.Equal(t, expected, result)

	expected = map[int]bool(map[int]bool{2: true, 3: true, 4: true, 5: true, 6: true, 7: true})
	result = FindUniqueCombinatorialSums([]int{1, 2, 3, 4}, 7)
	assert.Equal(t, expected, result)

	expected = map[int]bool(map[int]bool{})
	result = FindUniqueCombinatorialSums([]int{1, 2, 3, 4}, 1)
	assert.Equal(t, expected, result)
}

func TestFibonacciSequence(t *testing.T) {
	expected := []int{0}
	result, err := FibonacciSequence(0)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = []int{0, 1}
	result, err = FibonacciSequence(1)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = []int{0, 1, 1}
	result, err = FibonacciSequence(2)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = []int{0, 1, 1, 2, 3, 5, 8}
	result, err = FibonacciSequence(6)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657,
		46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169,
		63245986, 102334155, 165580141, 267914296, 433494437, 701408733, 1134903170, 1836311903, 2971215073, 4807526976, 7778742049,
		12586269025, 20365011074, 32951280099, 53316291173, 86267571272, 139583862445, 225851433717, 365435296162, 591286729879, 956722026041,
		1548008755920, 2504730781961, 4052739537881, 6557470319842, 10610209857723, 17167680177565, 27777890035288, 44945570212853,
		72723460248141, 117669030460994, 190392490709135, 308061521170129, 498454011879264, 806515533049393, 1304969544928657,
		2111485077978050, 3416454622906707, 5527939700884757, 8944394323791464, 14472334024676221, 23416728348467685, 37889062373143906,
		61305790721611591, 99194853094755497, 160500643816367088, 259695496911122585, 420196140727489673, 679891637638612258,
		1100087778366101931, 1779979416004714189, 2880067194370816120, 4660046610375530309, 7540113804746346429,
	}
	result, err = FibonacciSequence(92)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestFibonacciSequenceBigInt(t *testing.T) {
	ninetyThirdFibNum := new(big.Int)
	ninetyThirdFibNum.SetString("12200160415121876738", 10)

	fibBigInt := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(8), big.NewInt(13),
		big.NewInt(21), big.NewInt(34), big.NewInt(55), big.NewInt(89), big.NewInt(144), big.NewInt(233), big.NewInt(377), big.NewInt(610), big.NewInt(987),
		big.NewInt(1597), big.NewInt(2584), big.NewInt(4181), big.NewInt(6765), big.NewInt(10946), big.NewInt(17711), big.NewInt(28657),
		big.NewInt(46368), big.NewInt(75025), big.NewInt(121393), big.NewInt(196418), big.NewInt(317811), big.NewInt(514229), big.NewInt(832040),
		big.NewInt(1346269), big.NewInt(2178309), big.NewInt(3524578), big.NewInt(5702887), big.NewInt(9227465), big.NewInt(14930352), big.NewInt(24157817),
		big.NewInt(39088169), big.NewInt(63245986), big.NewInt(102334155), big.NewInt(165580141), big.NewInt(267914296), big.NewInt(433494437),
		big.NewInt(701408733), big.NewInt(1134903170), big.NewInt(1836311903), big.NewInt(2971215073), big.NewInt(4807526976), big.NewInt(7778742049),
		big.NewInt(12586269025), big.NewInt(20365011074), big.NewInt(32951280099), big.NewInt(53316291173), big.NewInt(86267571272), big.NewInt(139583862445),
		big.NewInt(225851433717), big.NewInt(365435296162), big.NewInt(591286729879), big.NewInt(956722026041), big.NewInt(1548008755920), big.NewInt(2504730781961),
		big.NewInt(4052739537881), big.NewInt(6557470319842), big.NewInt(10610209857723), big.NewInt(17167680177565), big.NewInt(27777890035288),
		big.NewInt(44945570212853), big.NewInt(72723460248141), big.NewInt(117669030460994), big.NewInt(190392490709135), big.NewInt(308061521170129),
		big.NewInt(498454011879264), big.NewInt(806515533049393), big.NewInt(1304969544928657), big.NewInt(2111485077978050), big.NewInt(3416454622906707),
		big.NewInt(5527939700884757), big.NewInt(8944394323791464), big.NewInt(14472334024676221), big.NewInt(23416728348467685), big.NewInt(37889062373143906),
		big.NewInt(61305790721611591), big.NewInt(99194853094755497), big.NewInt(160500643816367088), big.NewInt(259695496911122585), big.NewInt(420196140727489673),
		big.NewInt(679891637638612258), big.NewInt(1100087778366101931), big.NewInt(1779979416004714189), big.NewInt(2880067194370816120),
		big.NewInt(4660046610375530309), big.NewInt(7540113804746346429), ninetyThirdFibNum,
	}
	expected := fibBigInt[:1]
	result, err := FibonacciSequenceBigInt(0)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = fibBigInt[:2]
	result, err = FibonacciSequenceBigInt(1)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = fibBigInt[:3]
	result, err = FibonacciSequenceBigInt(2)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = fibBigInt[:7]
	result, err = FibonacciSequenceBigInt(6)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	expected = fibBigInt
	result, err = FibonacciSequenceBigInt(93)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
