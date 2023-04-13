package main

import "sync"

type DB struct {
	m sync.Map
}

func NewDB() *DB {
	db := DB{}
	for s, p := range stocks {
		db.m.Store(s, p)
	}

	return &db
}

func (db *DB) Price(symbol string) (float64, bool) {
	p, ok := db.m.Load(symbol)
	if !ok {
		return 0, false
	}

	return p.(float64), true
}

// Static demo data (NSDAQ from 2023-04-13)
var stocks = map[string]float64{
	"AAPL":  160.10,
	"ABNB":  112.42,
	"ADBE":  369.89,
	"ADI":   187.29,
	"ADP":   215.61,
	"ADSK":  193.16,
	"AEP":   94.26,
	"ALGN":  332.25,
	"AMAT":  113.16,
	"AMD":   92.33,
	"AMGN":  249.50,
	"ANSS":  316.49,
	"ASML":  654.66,
	"ATVI":  84.95,
	"AVGO":  616.70,
	"AZN":   73.76,
	"BIIB":  285.79,
	"BKNG":  2547.25,
	"BKR":   29.25,
	"CDNS":  212.18,
	"CEG":   76.81,
	"CHTR":  343.28,
	"CMCSA": 37.64,
	"COST":  489.35,
	"CPRT":  76.14,
	"CRWD":  135.00,
	"CSCO":  50.11,
	"CSGP":  68.59,
	"CSX":   30.16,
	"CTAS":  453.94,
	"CTSH":  60.28,
	"DDOG":  66.84,
	"DLTR":  150.17,
	"DXCM":  114.82,
	"EA":    126.13,
	"EBAY":  42.78,
	"ENPH":  196.03,
	"EXC":   43.08,
	"FANG":  145.01,
	"FAST":  52.56,
	"FISV":  114.49,
	"FTNT":  67.19,
	"GFS":   66.06,
	"GILD":  82.16,
	"GOOG":  105.22,
	"GOOGL": 104.64,
	"HON":   193.23,
	"IDXX":  470.90,
	"ILMN":  226.88,
	"INTC":  32.02,
	"INTU":  435.09,
	"ISRG":  261.77,
	"JD":    36.95,
	"KDP":   35.21,
	"KHC":   39.28,
	"KLAC":  370.14,
	"LCID":  8.13,
	"LRCX":  497.07,
	"LULU":  363.54,
	"MAR":   161.47,
	"MCHP":  79.23,
	"MDLZ":  70.21,
	"MELI":  1256.04,
	"META":  214.00,
	"MNST":  52.47,
	"MRNA":  155.56,
	"MRVL":  39.53,
	"MSFT":  283.49,
	"MU":    61.96,
	"NFLX":  331.03,
	"NVDA":  264.95,
	"NXPI":  170.22,
	"ODFL":  340.06,
	"ORLY":  877.35,
	"PANW":  195.28,
	"PAYX":  108.41,
	"PCAR":  72.11,
	"PDD":   67.02,
	"PEP":   182.56,
	"PYPL":  73.50,
	"QCOM":  120.27,
	"REGN":  820.43,
	"RIVN":  14.19,
	"ROST":  102.89,
	"SBUX":  105.60,
	"SGEN":  205.20,
	"SIRI":  3.98,
	"SNPS":  376.52,
	"TEAM":  158.55,
	"TMUS":  149.51,
	"TSLA":  180.54,
	"TXN":   177.55,
	"VRSK":  192.17,
	"VRTX":  326.37,
	"WBA":   35.69,
	"WBD":   14.06,
	"WDAY":  191.74,
	"XEL":   71.04,
	"ZM":    67.49,
	"ZS":    105.87,
}
