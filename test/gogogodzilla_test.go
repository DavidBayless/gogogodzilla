package test_test

import (
	// . "gogogodzilla/"
	// . "gogogodzilla/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Gogogodzilla", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should show an awesome header", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Eventually(page.FindByID("godzirra")).Should(BeFound())
		})
	})

	It("should show a 'dank' meme as the kidz in the cool cidz club say", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Eventually(page.FindByID("dankMeme")).Should(BeFound()) // Also be dank
		})
	})
	It("should allow user to enter new Godzilla in input field", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("newGodzilla").Fill("Hello World")).To(Succeed())
		})
	})
	It("Should allow user to enter ew Godzilla height in input field", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("newGodzillaHeight").Fill("25m^3")).To(Succeed())
		})
	})
	It("Should allow a user to go to /godzillas", func() {
		By("clicking on 'New Godzillas'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("submit").Click()).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:9001/godzirras"))
		})
	})
	It("Should ")

})
