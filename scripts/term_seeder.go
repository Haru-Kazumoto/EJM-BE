package scripts

import (
	"EJM/pkg/models"

	"gorm.io/gorm"
)

type TermSeeder struct {
	Name string
	DB   *gorm.DB
}

func (termSeeder *TermSeeder) Execute() error {
	data := []models.Term{
		{
			TermID: "S1ADA276",
			ATMKey: "571592949117",
			TermType: "ATM",
			Merk: "DIEBOLD",
			ManagedBy: "Vendor-CAKDPK",
			Jarkom: "VSAT",
			Region: "REGION V / JAKARTA 3",
			Area: "127 / JAKARTA FATMAWATI",
			CardRetainedTes: "100",
		},
		{
			TermID: "S1AWKMMX",
			ATMKey: "7167842190509",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-PTUGBGR",
			Jarkom: "VSAT",
			Region: "REGION V / JAKARTA 3",
			Area: "133 / BOGOR",
			CardRetainedTes: "83",

		},
		{
			TermID: "S1AW10ZP",
			ATMKey: "10796142316459",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-PTUGMDN",
			Jarkom: "VSAT",
			Region: "REGION I / SUMATERA 1",
			Area: "106 / MEDAN BALAIKOTA",
			CardRetainedTes: "22",

		},
		{
			TermID: "S1AW1M89",
			ATMKey: "10980942194063",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-TAGSOLO",
			Jarkom: "VSAT",
			Region: "REGION VII / JAWA 2",
			Area: "138 / SOLO",
			CardRetainedTes: "150",

		},
		{
			TermID: "S1AW1536",
			ATMKey: "2340441376984070",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Branch-18600",
			Jarkom: "LEASED-LINE",
			Region: "REGION X / SULAWESI DAN MALUKU",
			Area: "186 / MALUKU",
			CardRetainedTes: "70",
		},
		{
			TermID: "S1AW16KV",
			ATMKey: "2581950835457800",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Branch-18108",
			Jarkom: "LEASED-LINE",
			Region: "REGION XI / BALI DAN NUSA TENGGARA",
			Area: "181 / KUPANG",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1RK17KS",
			ATMKey: "5251877953091",
			TermType: "CRM",
			Merk: "OKI",
			ManagedBy: "Vendor-TAGKDR",
			Jarkom: "LEASED-LINE",
			Region: "REGION VIII / JAWA 3",
			Area: "171 / KEDIRI",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1ADKD3I",
			ATMKey: "291729053961144",
			TermType: "ATM",
			Merk: "DIEBOLD",
			ManagedBy: "Vendor-CAKKDR",
			Jarkom: "VSAT",
			Region: "REGION VIII / JAWA 3",
			Area: "171 / KEDIRI",
			CardRetainedTes: "58",
		},
		{
			TermID: "S1AW14VG",
			ATMKey: "15556758677346",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-CAKMJK",
			Jarkom: "VSAT",
			Region: "REGION VIII / JAWA 3",
			Area: "142 / SURABAYA BASUKI RAHMAT",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1AWK43X",
			ATMKey: "78043351458506",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-CAKBWI",
			Jarkom: "VSAT",
			Region: "REGION VIII / JAWA 3",
			Area: "143 / JEMBER",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1AW14X8",
			ATMKey: "323333623236565",
			TermType: "ATM",
			Merk: "WINCOR",
			ManagedBy: "Vendor-PTUGMLG",
			Jarkom: "MPLS",
			Region: "REGION VIII / JAWA 3",
			Area: "144 / MALANG",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1ROA40X",
			ATMKey: "1589989624830809",
			TermType: "CRM",
			Merk: "HITACHI",
			ManagedBy: "Vendor-PTUGGSK",
			Jarkom: "LEASED-LINE",
			Region: "REGION VIII / JAWA 3",
			Area: "178 / GRESIK",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1RK1XAE",
			ATMKey: "1264209312939701",
			TermType: "CRM",
			Merk: "OKI",
			ManagedBy: "Vendor-CAKMJK",
			Jarkom: "LEASED-LINE",
			Region: "REGION VIII / JAWA 3",
			Area: "141 / SURABAYA GENTENGKALI",
			CardRetainedTes: "17",
		},
		{
			TermID: "S1ANA41U",
			ATMKey: "541378030143538",
			TermType: "ATM",
			Merk: "NCR",
			ManagedBy: "Vendor-PTUGSBY",
			Jarkom: "VSAT",
			Region: "REGION VIII / JAWA 3",
			Area: "141 / SURABAYA GENTENGKALI",
			CardRetainedTes: "17",
		},
		// {
		// 	TermID: "",
		// 	ATMKey: "",
		// 	TermType: "",
		// 	Merk: "",
		// 	ManagedBy: "",
		// 	Jarkom: "",
		// 	Region: "",
		// 	Area: "",
		// 	CardRetainedTes: "17",
		// },
		// {
		// 	TermID: "",
		// 	ATMKey: "",
		// 	TermType: "",
		// 	Merk: "",
		// 	ManagedBy: "",
		// 	Jarkom: "",
		// 	Region: "",
		// 	Area: "",
		// 	CardRetainedTes: "17",
		// },
	}

	if err := termSeeder.DB.Model(models.Term{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (termSeeder *TermSeeder) GetName() string {
	return termSeeder.Name
}
