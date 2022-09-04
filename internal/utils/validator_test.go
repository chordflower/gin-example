package utils_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/chordflower/gin_example/internal/utils"
)

var _ = ginkgo.Describe("#Validator", func() {

	var validator *utils.Validate

	ginkgo.BeforeEach(func() {
		validator = utils.NewValidator()
	})

	ginkgo.It("it should validate a custom check", func() {
		validator.Check(true, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.Check(false, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a value is present", func() {
		validator.IsPresent("", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsPresent(nil, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a value is not present", func() {
		validator.IsNotPresent(nil, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsNotPresent("", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is alphanumeric", func() {
		validator.IsAlphaNumeric("phayt6uXpHdN6vszzXNEXkqZyydSYgMQ7JAV5psDew", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsAlphaNumeric("N^44P2qx6uP629dHXED6JuU6@g!@5PT#Y#w@2mEpca", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is in base64", func() {
		validator.IsBase64("VFZNSWNFMlExSjNlOUJLUVBwaklhZ3BaTlpRMWVlMkJUMEhLeFBUQURTZ2xPQlJDcjJwUk5VMFJJeTY0dkJMS3FNREhwc2xKVUYyN2xseVBiZkRPaXdIQjNWQ0NEWTVueURRcU96TzFwZGdCeHFzQzVONHZOeENQ", "")
		validator.IsBase64("UW5PdEdndkVlNndsZGs4bTNUN0d6Vk83RnV1aFY5V2xiRXVVT3V0WFdOeU93dWo1ZXdNUWJjODFUNEVaSVRpM1kza2szY09zUFZGalVkSGkyVHpoWXJCS0tPelFOTHEweDZjQlQ1c2V6YWh2cjlrMjJPaGJ0eHJD", "")
		validator.IsBase64("YzJmUUtOc29sVEU0dE1oSkx1enRSZjNTZkFmekxoYWJ5Sm5mc3BuWWkzN0ZCa3I3cXU2b3hoN0hwWURKVnlVNW0yeGhBa3FTUVl1d2w2T3dRekNFbVNJcHAyaTRFeG4wUGtXdk5hOERvMlRhZmVNWEFxNHpMMVhQ", "")
		validator.IsBase64("VE1zSmJsbHFrS1Bvd2RoYk1XQ1p1dnJXZE5KemE1Z2VIY09aWmtjSEk1NVFmbEVCWVl3TnRiZWpMY0EyOE5aSUVhWTVTcGRjaWp3bjMyV3pnR2VMbWhlTklJVjQ0ZU1tZk54S0sxTWZTR3U5dzdmMFpiS0xVN2p6", "")
		validator.IsBase64("cnBsZVdlZWFUWTVyaXpTVTVyYWN1Qm9kd3JZdTYyS0J1bTZxb3cxOWtsUnJUMHhqVXpjSkVwMW9MY09Wdm1LbWxPYWhDN044MllheHZYVVBOUlZkMVVVY0dVdHZvRXRMRURrV3BtdkZPYml6UG44T0pXTDZ6cUZT", "")
		validator.IsBase64("SW0zZmVMY0FKTGZLeTFyRWNXWGFVMmpZVllmRGxOUlRSS2hIQ253VGh1M2JpcmhOQkdBdjNRb043MlRkTXNjN0diVDF5bXN0MGN1ZmhOdjNBVFZoN09JQ0VuRWFFT2dZaFF2NjJFZ0hkU1cwR09aREl0NUtYQ0xF", "")
		validator.IsBase64("bDUxc0ZvTkVxVTU4d2xqcmlMUWRnbkZNaFVBcmVUSjBYeUY4SHZFdXJXR2dSWDUyazJBZTJYU3o1OEJSZmpOTmJtZWhwWDFpWUVwbUdGN25rR0REUlhPcm5OV1pLT1c3b0tFQkFZbzd5S3QyUmx3UEI2YVkwTTE0dw==", "")
		validator.IsBase64("YUdkMktyb2cyRVRsSlVsbHVRdGRYS0p2Y2tNVHRCT29zVXdveGVJbFY3bktaMHltTU5RNG5BNFo2NXFKR0RTVkM5WHVTc0l4S21OOTZaNVREN0JVYnRQQjVwQzFWeTVDWFJIcW43ZnhCMHRITjVaQ1JLRlgybzNZUg==", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsBase64("molasalertedoldrecensionscrowswanssewvasty", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is lowercase", func() {
		validator.IsLowercase("lowercase string", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsLowercase("Capitalcase String", "")
		validator.IsLowercase("UPPERCASE STRING", "")
		validator.IsLowercase("ߠߟߞߛߘ", "")                                      //N'Ko
		validator.IsLowercase("إذا لم يأتيك الجبل ، يجب أن تذهب إلى الجبل", "") //if the mountain won't come to you, you must go to the mountain
		gomega.Expect(validator.ErrorNumber()).To(gomega.Equal(uint(4)))
	})

	ginkgo.It("it should validate if a string is uppercase", func() {
		validator.IsUppercase("PARROTED COMPLETING GAGS COPPED BOYAR", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsUppercase("Parroted Completing Gags Copped Boyar", "")
		validator.IsUppercase("parroted completing gags copped boyar", "")
		validator.IsUppercase("ߠߟߞߛߘ", "")
		validator.IsUppercase("إذا لم يأتيك الجبل ، يجب أن تذهب إلى الجبل", "")
		gomega.Expect(validator.ErrorNumber()).To(gomega.Equal(uint(4)))
	})

	ginkgo.It("it should validate if a string is a credit card number", func() {
		validator.IsCreditCard("4194033457664927", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsCreditCard("4194055457347927", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid domain", func() {
		validator.IsDomain("example.com", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsDomain("example.$%#sd", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid email", func() {
		validator.IsEmail("kyle.deckow@gmail.com", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsEmail("kyle.deckow@", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid uuid", func() {
		validator.IsGUID("414f6a3f-b70c-4bca-b457-de56db223c7c", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsGUID("414f6a3f-b70c-4bca-b457-de56nb223c7z", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid hostname", func() {
		validator.IsHostname("example.com", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsHostname("example.$&%", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid ip", func() {
		validator.IsIP("54.120.224.153", "")
		validator.IsIP("daa6:c765:1421:7e01:2343:f9e7:bd47:3c82", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsIP("54.120.224.653", "")
		validator.IsIP("daa6:c765:1421:7j01:2343:f9e7:bd47:3c82", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid iso date", func() {
		validator.IsStdDate("2021-07-13T14:40:00Z", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsStdDate("2021-12-45T14:40:00Z", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid duration", func() {
		validator.IsDuration("2h3m", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsDuration("3z6m", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string has the given size", func() {
		validator.IsSize("XwU1K^fmLdSTsx97ryW3mAnccSqd6W", 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsSize("S3T9Enw%t6fvSfnG^sk^!3BdAGfRzd", 20, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is empty", func() {
		validator.IsEmpty("", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsEmpty("S3T9Enw%t6fvSfnG^sk^!3BdAGfRzd", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is not empty", func() {
		validator.IsNotEmpty("S3T9Enw%t6fvSfnG^sk^!3BdAGfRzd", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsNotEmpty("", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is between the given size", func() {
		validator.IsBetween("XwU1K^fmLdSTsx97ryW3mAnccSqd6W", 20, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsBetween("pF3WMdkpvdLvg2@VyCJ8%qcrU4hr6WauM", 20, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a string is a valid URL", func() {
		validator.IsURL("https://domain.example.com:453/api/users/articles?from=2343&size=20#example", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsURL("://domain.example.com:453/api/users/articles?from=2343&size=20#example", "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is greater than the given number", func() {
		validator.IsGreaterThan(32, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsGreaterThan(23, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is less than the given number", func() {
		validator.IsLessThan(23, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsLessThan(32, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is between the given numbers", func() {
		validator.IsBetweenNumbers(23, 20, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsBetweenNumbers(32, 20, 30, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is positive", func() {
		validator.IsPositive(23, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsPositive(-23, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is negative", func() {
		validator.IsNegative(-23, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsNegative(23, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

	ginkgo.It("it should validate if a number is a port", func() {
		validator.IsPort(65534, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeFalse())
		validator.IsPort(65537, "")
		gomega.Expect(validator.HasErrors()).To(gomega.BeTrue())
	})

})
