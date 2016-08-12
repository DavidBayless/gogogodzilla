package test_test

import (
	// . "gogogodzilla/"
	// . "gogogodzilla/test"

	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var DB *sql.DB

// var _ = BeforeSuite(func() {
// 	connstring := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "localadmin", "godzirras")
// 	var err error
// 	DB, err = sql.Open("postgres", connstring)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = DB.Ping()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// })

var _ = Describe("Gogogodzilla", func() {
	var page *agouti.Page

	BeforeEach(func() {
		connstring := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "localadmin", "godzirras")
		var err error
		DB, err = sql.Open("postgres", connstring)
		if err != nil {
			log.Fatal(err)
		}
		err = DB.Ping()
		if err != nil {
			fmt.Println(err)
		}

		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// DB.Query("DELETE FROM godzillas")
		DB.Close()
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
	It("should prevent user from entering new Godzilla containing illegal characters", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("newGodzilla").Fill("Godzill@")).To(Succeed())
			Expect(page.FindByID("submit").Click()).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:9001/"))
		})
	})
	It("Should allow user to enter ew Godzilla height in input field", func() {
		By("visiting '/'", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("newGodzillaHeight").Fill("25m^3")).To(Succeed())
		})
	})
	It("Should populate the database", func() {
		By("filling the form, and clicking the submit button", func() {
			Expect(page.Navigate("http://localhost:9001")).To(Succeed())
			Expect(page.FindByID("newGodzilla").Fill("Hydra")).To(Succeed())
			Expect(page.FindByID("newGodzillaHeight").Fill("3ft")).To(Succeed())
			Expect(page.FindByID("submit").Click()).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:9001/godzirras"))
			rows, err := DB.Query("Select * from godzillas where name='Hydra'")
			if err != nil {
				log.Fatal(err)
			}
			var nameExpect string
			var heightExpect string
			var idInsert int

			defer rows.Close()
			for rows.Next() {
				var id int
				var name string
				var height string
				if err := rows.Scan(&id, &name, &height); err != nil {
					fmt.Println(err)
				}
				nameExpect = name
				heightExpect = height
				idInsert = id
			}
			Expect(nameExpect).To(Equal("Hydra"))
			Expect(heightExpect).To(Equal("3ft"))

			rowz, errz := DB.Query("DELETE FROM godzillas WHERE id=" + strconv.Itoa(idInsert))
			fmt.Println(errz)
			fmt.Println(rowz)
		})
	})

})
