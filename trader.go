package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/urfave/cli"
)

var russell2k = []string{
	"FLWS", "FCCY", "SRCE", "XXII", "DDD", "EGHT", "AVHI", "ATEN", "AAC", "AAON", "AIR", "AAN", "ABAX", "ABEO", "ANF", "ABM", "AXAS", "ACIA", "ACTG", "ACAD", "AKR", "AXDX", "XLRN", "ANCX", "ACCO", "ARAY", "AKAO", "ACHN", "ACIW", "ACRS", "ACMR", "ACNB", "ACOR", "ATU", "GOLF", "ACXM", "ADMS", "AE", "ADUS", "IOTS", "ADMA", "ATGE", "ADTN", "ADRO", "ADSW", "WMS", "ADES", "AEIS", "ASIX", "ADVM", "AEGN", "AGLE", "AERI", "HIVE", "AJRD", "AVAV", "MITT", "AGEN", "AGYS", "ADC", "AGFS", "AIMT", "ATSG", "AYR", "AKS", "AKCA", "AKBA", "AKRX", "ALG", "ALRM", "AIN", "ALBO", "ALDR", "ALDX", "ALEX", "ALX", "ALCO", "ATI", "ABTX", "ALGT", "ALNA", "ALE", "AOI", "AMOT", "MDRX", "AOSL", "AMR", "ALTR", "AYX", "ASPS", "AIMC", "AMAG", "AMBC", "AMBA", "AMBR", "AMC", "AMED", "APEI", "AMRC", "AAT", "AXL", "AEO", "AEL", "AMNB", "AOBC", "ARII", "ARA", "ARL", "AMSWA", "AWR", "AVD", "AMWD", "CRMT", "COLD", "ABCB", "AMSF", "ATLO", "FOLD", "AMKR", "AMN", "AMRX", "AMPH", "AMPE", "AFSI", "AMRS", "ANAB", "ANDE", "ANGO", "ANIP", "ANIK", "AXE", "ATRS", "ANH", "APLS", "APOG", "ARI", "AMEH", "APPF", "AIT", "AAOI", "AREX", "PETX", "AVID", "CAR", "AVA", "AVX", "ACLS", "AXGN", "AAXN", "AXTI", "AZZ", "BGS", "RILY", "BW", "BMI", "BCPC", "BWINB", "BANC", "BANF", "BLX", "TBBK", "BXS", "BOCH", "BMRC", "NTB", "BPRN", "BFIN", "BWFG", "BANR", "BHB", "BNED", "BKS", "B", "BBSI", "BAS", "BSET", "BCML", "BBX", "BCBP", "BECN", "BBGI", "BZH", "BBBY", "BELFB", "BDC", "BLCM", "BEL", "BHE", "BNCL", "CDMO", "BGSF", "BGFV", "BIG", "A", "BH", "BCRX", "BHVN", "BIOS", "BSTC", "BEAT", "BTX", "BJRI", "BKH", "BLKB", "BL", "BXMT", "BLMN", "BCOR", "BLBD", "BHBK", "BXG", "BXC", "BPMC", "BRG", "BMCH", "BOFI", "WIFI", "BCC", "BOJA", "BCEI", "BOOT", "SAM", "BOMN", "BPFH", "EPAY", "BOX", "BYD", "BRC", "BHR", "BDGE", "BWB", "BGG", "BCOV", "BSIG", "EAT", "BCO", "BRS", "BNFT", "BHLB", "BRKS", "BRT", "BMTC", "BLMT", "BKE", "BLDR", "BFST", "BY", "CFFI", "CJ", "CCMP", "CACI", "WHD", "CADE", "CDZI", "CSTE", "CAI", "CALM", "CAMP", "CVGW", "CAL", "CRC", "CWT", "CALA", "CALX", "ELY", "CPE", "CLXT", "ABCD", "CBM", "CATC", "CAC", "CWH", "CNNE", "CPLA", "CCBG", "CSU", "CFFN", "CSTR", "CMO", "CARA", "CRR", "CBLK", "CARB", "CSII", "CDLX", "BKD", "BRKL", "CECO", "CTRE", "CARG", "CARO", "CRS", "CSV", "CRZO", "TAST", "CARS", "CVNA", "CASA", "CWST", "CASI", "CASS", "ROX", "CSLT", "CBIO", "CPRX", "CTT", "CATY", "CATO", "CVCO", "CBFV", "CBZ", "CBL", "CBTX", "CECE", "CDR", "CELC", "CBMG", "CELH", "CSFL", "CETV", "CENT", "CENTA", "CPF", "CVCY", "CENX", "CNBKA", "CNTY", "CCS", "CERS", "CEVA", "ECOM", "GTLS", "CHFN", "CATM", "CRCM", "CDNA", "CHEF", "CHGG", "CHFC", "CCXI", "CHMG", "CHMI", "CHSP", "CPK", "CHS", "PLCE", "CMRX", "CDXC", "CHDN", "CHUY", "CIEN", "CMPR", "CBB", "CIR", "CRUS", "CISN", "CTRN", "CZNC", "CIA", "CHCO", "CIO", "CIVB", "CIVI", "CLAR", "CLNE", "CCO", "CLFD", "CLSD", "CLW", "CLF", "CLPR", "CLD", "CLDR", "CLVS", "CCNE", "CNO", "COBZ", "COKE", "CDXS", "CVLY", "CDE", "CCF", "CLDT", "CAKE", "COHU", "COLL", "CLNC", "COLB", "CLBK", "CMCO", "FIX", "CMC", "CVGI", "CBU", "ESXB", "TCFC", "CYH", "CHCT", "CTBI", "CVLT", "CMP", "CPSI", "CIX", "CMTL", "CNCE", "CNMD", "CTWS", "CNOB", "CONN", "CEIX", "CNSL", "CTO", "CWCO", "CBPX", "CTRL", "CVON", "CVG", "CPS", "CTB", "CRBP", "CORT", "CORE", "CXW", "CORR", "CPLG", "CORI", "CSOD", "CRVL", "CRVS", "CCOI", "CWBR", "CNS", "CHRS", "CUZ", "CVA", "CVTI", "CVIA", "COWN", "CRAI", "CBRL", "BREW", "B", "CRAY", "CREE", "CROX", "CCRN", "CRY", "CYRX", "CSGS", "CSWI", "CTIC", "CTS", "CUB", "CUE", "CULP", "CURO", "CUBI", "CUTR", "CVBF", "CVI", "CBAY", "CYS", "CYTK", "CTMX", "CTSO", "DJCO", "DAKT", "DAN", "DAR", "DZSI", "DSKE", "PLAY", "DWSN", "DF", "DCPH", "DECK", "DFRG", "CMRE", "COTV", "ICBK", "COUP", "DENN", "DEPO", "DERM", "DHT", "DHIL", "DO", "DRH", "DRNA", "DBD", "DGII", "DMRC", "DDS", "DCOM", "DIN", "DIOD", "DPLO", "BOOM", "DGICA", "DFIN", "RRD", "LPG", "DORM", "PLOW", "DOVA", "DRQ", "DS", "DSW", "DCO", "DLTH", "DRRX", "DXPE", "DY", "DVAX", "DX", "EGBN", "EGLE", "EGRX", "ESTE", "DEA", "EML", "EGP", "KODK", "EBIX", "ECHO", "TACO", "DK", "DLX", "DNLI", "DNR", "EHTH", "EE", "LOCO", "ERI", "ESIO", "EFII", "ELVT", "ELF", "ELLI", "PERY", "ELOX", "EMCI", "EME", "EEX", "EBS", "NYNY", "EIG", "ENTA", "ECPG", "WIRE", "ENDP", "ECYT", "ELGX", "EIGI", "WATT", "UUUU", "ERII", "EGC", "ENS", "EGL", "EBF", "ENVA", "ENPH", "NPO", "ENSG", "ESGR", "ENFC", "ENTG", "ETM", "EBTC", "EFSC", "EVC", "ENV", "ECR", "EPC", "EDIT", "EDR", "EGAN", "ERA", "EROS", "ESCA", "ESE", "ESPR", "ESQ", "ESSA", "ESND", "ESNT", "ESL", "ETH", "ETSY", "EVBN", "EVLO", "EVBG", "EVRI", "EVTC", "EVH", "EOLS", "EPM", "AQUA", "XAN", "XELA", "EXLS", "EXPO", "EXPR", "EXTN", "EXTR", "EZPW", "FN", "FARM", "FMAO", "FFKT", "FMNB", "FPI", "FARO", "FATE", "FBK", "FFG", "FCB", "AGM", "FSS", "FII", "EVI", "ENZ", "EPE", "EPZM", "PLUS", "EQBK", "LION", "FRGI", "FNGN", "FISI", "FNSR", "FNLC", "FBNC", "FBP", "FBMS", "FRBA", "BUSE", "FBIZ", "FCBP", "FCBC", "FCCO", "FCF", "FBNK", "FDEF", "FFBC", "THFF", "FFIN", "FFNW", "FFWM", "FGBI", "FR", "INBK", "FIBK", "FLIC", "FRME", "FMBH", "FMBI", "FNWB", "FSFG", "FUNC", "FCFS", "FIT", "FIVE", "FPRX", "FIVN", "FBC", "FLXN", "FLXS", "FNHC", "FENC", "FOE", "FG", "FGEN", "FDBC", "FOR", "FORM", "FORR", "FRTA", "FBIO", "FET", "FWRD", "FOSL", "FSTR", "FBM", "FMI", "FCPT", "FOXF", "FRAN", "FC", "FELE", "FSB", "FSP", "FI", "RAIL", "FDP", "FRPT", "RESI", "FTR", "FRO", "FRPH", "FSBW", "FCN", "FTSI", "FCEL", "FULT", "FNKO", "FSNN", "FF", "GTHX", "GAIA", "GCAP", "GBL", "GME", "GCI", "GLOG", "GATX", "FTK", "FLNT", "FLDM", "FFIC", "FNBG", "FONR", "FSCT", "GEN", "GNMK", "GHDX", "THRM", "GNW", "GEO", "GABC", "GERN", "GTY", "ROCK", "GBCI", "GOOD", "LAND", "GLT", "GKOS", "GBT", "BRSS", "GBLI", "GMRE", "GNL", "GWRS", "GMED", "GLUU", "GLYC", "GMS", "GNC", "GOGO", "GLNG", "GORO", "GDEN", "GDP", "GSHD", "GPRO", "GRC", "GOV", "GPX", "GHM", "GPT", "GVA", "GPMT", "GTN", "GCP", "GNK", "GENC", "GNRC", "GFN", "GCO", "GIII", "GBX", "GCBC", "GHL", "GLRE", "GEF", "B", "GRIF", "GFF", "GPI", "GRPN", "GTT", "GTXI", "GBNK", "GNTY", "GES", "GLF", "GPOR", "HEES", "HABT", "HCKT", "HAE", "HK", "HNRG", "HALL", "HALO", "HYH", "HBB", "HLNE", "HWC", "HAFC", "HASI", "HONE", "HLIT", "HSC", "HBIO", "HVT", "HA", "HCOM", "HWKN", "HAYN", "FUL", "AJX", "GLDD", "GSBC", "GWB", "GNBC", "GRBK", "GDOT", "GPRE", "HTLF", "HL", "HSII", "HELE", "HSDT", "HLX", "HMTV", "HRI", "HTBK", "HCCI", "HFWA", "HRTG", "MLHR", "HRTX", "HT", "HTZ", "HSKA", "HF", "HIBB", "HPR", "HIL", "HI", "HTH", "HIFS", "HMSY", "HNI", "HBCP", "HOMB", "HMST", "HTBI", "FIXX", "HOFT", "HOPE", "HMN", "HBNC", "HZNP", "HDP", "TWNK", "HMHC", "HLI", "HCHC", "HCI", "HIIQ", "HR", "HCSG", "HQY", "HSTM", "HTLD", "IBKC", "ICFI", "ICHR", "IDA", "IDRA", "IESC", "IIVI", "ILG", "IMAX", "IMMR", "IMDZ", "IMGN", "IMMU", "IMH", "IMPV", "PI", "ICD", "IHC", "IRT", "IBCP", "INDB", "IBTX", "ILPT", "INFN", "IPCC", "III", "HIFR", "IEA", "NGVT", "IMKTA", "INWK", "IPHS", "IOSP", "IIPR", "INVA", "INGN", "INOV", "INO", "IPHI", "NSIT", "HOV", "HBMD", "HRG", "HUBG", "HUBS", "HUD", "HURC", "HURN", "HY", "NTLA", "I", "IPAR", "ICPT", "IDCC", "TILE", "INAP", "IBOC", "INSW", "ISCA", "XENT", "INTL", "ITCI", "IPI", "XON", "IIN", "IVC", "IVR", "ISTR", "ITG", "ISBC", "IRET", "ITIC", "NVTA", "IO", "IOVA", "IRMD", "IRTC", "IRDM", "IRBT", "IRWD", "ISRL", "STAR", "ITI", "ITRI", "JJSF", "JAX", "JCOM", "JACK", "INSM", "NSP", "INSP", "IBP", "IIIN", "INST", "INSY", "ITGR", "IDTI", "KAI", "KDMN", "KALU", "KALA", "KAMN", "KS", "KPTI", "KBH", "KBR", "FRAC", "KRNY", "KELYA", "KEM", "KMPR", "KMT", "KW", "KERX", "KEG", "KEYW", "KFRC", "KE", "KBAL", "KIN", "KND", "KINS", "KNSL", "KIRK", "KRG", "KREF", "KLDX", "KLXI", "KMG", "KNL", "KN", "KOPN", "KOP", "KFY", "KRA", "KTOS", "JAG", "JRVR", "JELD", "JCAP", "JILL", "JBT", "JOUT", "JNCE", "LRN", "KTWO", "LNDC", "LE", "LCI", "LNTH", "LPI", "LHO", "LSCC", "LAUR", "LAWS", "LCII", "LCNB", "LFGR", "LTXB", "LMAT", "LC", "TREE", "LEVL", "LXRX", "LXP", "LGIH", "LHCG", "BATRA", "BATRK", "LEXEA", "LILA", "LILAK", "LBRT", "LTRPA", "LPNT", "LCUT", "LGND", "LLEX", "LLNW", "LMNR", "LIND", "LNN", "LQDT", "LAD", "KRO", "KURA", "KVHI", "LJPC", "LZB", "LADR", "LTS", "LBAI", "LKFN", "LANC", "LITE", "LMNX", "LBC", "LDL", "MHO", "MCBC", "CLI", "MTSI", "MGNX", "SHOO", "MDGL", "MGLN", "MHLD", "MJCO", "MBUU", "MNK", "MLVF", "TUSK", "MNTX", "MTW", "MNKD", "MANT", "MMI", "MCS", "MPX", "HZO", "MRNS", "MRLN", "VAC", "MBII", "MRTN", "DOOR", "MTZ", "MTDR", "MTRN", "MTRX", "MATX", "MATW", "LIVN", "LOB", "LPSN", "LIVX", "LORL", "LPX", "LOXO", "LXU", "LKSD", "LTC", "LL", "MDCA", "MDC", "MRT", "MDCO", "MNOV", "MDSO", "MED", "MEDP", "MLNT", "MNLO", "MBWM", "MBIN", "MRCY", "MDP", "EBSB", "VIVO", "MMSI", "MTH", "MTOR", "MRSN", "MLAB", "CASH", "MEI", "MCB", "MGEE", "MTG", "MGPI", "MSTR", "MPB", "MBCN", "MSEX", "MSBI", "MSL", "MPO", "MOFG", "MCRN", "MLR", "MLP", "MAXR", "MMS", "MXL", "MXWL", "MBFI", "MBI", "MBTF", "MCFT", "MDR", "MGRC", "MC", "MTEM", "MNTA", "MCRI", "MGI", "MNR", "TYPE", "MNRO", "INNT", "A", "MPAA", "MOV", "MRC", "MSA", "MSGN", "MTGE", "MTSC", "MLI", "MWA", "LABL", "MUSA", "MBIO", "MFSF", "MVBF", "MYE", "MYOK", "MYRG", "MYGN", "NC", "NANO", "NSTG", "NH", "NK", "NSSC", "NTRA", "NATH", "NBHC", "MDXG", "MB", "MTX", "NERV", "MGEN", "MRTX", "MG", "MITK", "MINI", "MOBL", "MODN", "MOD", "NSM", "NGS", "NGVC", "NHTC", "NATR", "BABY", "NLS", "NCI", "NAVG", "NAV", "NBTB", "NCS", "NCSM", "NP", "NNI", "NEOG", "NEO", "NPTN", "NEOS", "NTGR", "NTCT", "NVRO", "NWHM", "NJR", "NEWM", "NEWR", "SNR", "NWY", "NYMT", "NYT", "NLNK", "NMRK", "NR", "NXEO", "NXRT", "NXST", "NKSH", "FIZZ", "NCMI", "NCOM", "NGHC", "NHI", "NHC", "NPK", "NRC", "NSA", "EYE", "NWLI", "NBN", "NOG", "NFBK", "NRIM", "NRE", "NWBI", "NWN", "NWPX", "NWE", "NWFL", "NOVT", "NVAX", "NVCR", "DNOW", "A", "NYLD", "NTRI", "NUVA", "NVTR", "NES", "NVEE", "NVEC", "NXTM", "NYMX", "OVLY", "OAS", "ORIG", "OII", "OCFC", "OCLR", "OFED", "OCUL", "OCN", "ODT", "ODP", "OFG", "NEXT", "NODK", "EGOV", "NCBS", "NIHD", "NINE", "NL", "LASR", "NMIH", "NNBR", "NE", "NDLS", "NAT", "OGS", "OLP", "OSPN", "OOMA", "OPBK", "OPK", "OPY", "OPTN", "OPB", "OSUR", "ORBC", "ORC", "ONVO", "OBNK", "ORN", "ORIT", "ORA", "ORRF", "OFIX", "KIDS", "OSIS", "OTTR", "OSG", "OSTK", "OVID", "OMI", "OXFD", "OXM", "PTSI", "CNXN", "PACB", "PMBC", "PPBI", "PCRX", "PTN", "OVBC", "ODC", "OIS", "OLBK", "ONB", "OSBC", "OLLI", "ZEUS", "OFLX", "OMER", "OMCL", "OMN", "ONDK", "PDCO", "PCTY", "PCSB", "PDCE", "PDFS", "PDLI", "PDLB", "PDVW", "BTU", "PGC", "PEB", "PENN", "PVAC", "JCP", "PWOD", "PEI", "PFSI", "PMT", "PEBO", "PEBK", "PFIS", "PUB", "PRFT", "PFGC", "PRSP", "PETQ", "PETS", "PFNX", "PFSW", "PGTI", "PHH", "PHIIK", "PAHC", "PLAB", "DOC", "P", "PHX", "PZZA", "PARR", "PAR", "PRTK", "PCYG", "PKE", "PRK", "PKOH", "PKBK", "PRTY", "PATK", "PEGI", "PNM", "COOL", "POL", "POR", "PTLA", "PBPB", "PCH", "POWL", "POWI", "PQG", "PRAA", "APTS", "PFBC", "PLPC", "PFBI", "PSDO", "PBH", "PRGX", "PSMT", "PRI", "PRMW", "PRIM", "PRA", "PFIE", "PGNX", "PRGS", "PUMP", "PRO", "PTI", "PRTA", "PRLB", "PRSC", "PVBC", "PFS", "PICO", "PDM", "PIR", "PIRS", "PNK", "PES", "PJC", "PBI", "PJT", "PLNT", "PLT", "AGS", "PLXS", "PLUG", "QSII", "QLYS", "NX", "QTNA", "QTRX", "QDEL", "QNST", "QES", "QHC", "QUOT", "RCM", "RARX", "RDN", "RLGT", "RDUS", "RDNT", "METC", "RMBS", "RPT", "RPD", "RAVN", "RYAM", "RBB", "ROLL", "RICK", "RMAX", "RDI", "RETA", "REPH", "RLH", "RRGB", "RRR", "RDFN", "RWT", "PBIP", "PSB", "PTCT", "PLSE", "PBYI", "PCYO", "PRPL", "PZN", "QTWO", "QADA", "QCRH", "QTS", "QUAD", "KWR", "QCP", "ROIC", "RTRX", "REVG", "RVNC", "REV", "REX", "REXR", "RXN", "RGCO", "RH", "RYTM", "RBBN", "RIGL", "RNET", "RMNI", "REI", "RAD", "RVSB", "RLI", "RLJ", "RMR", "RCKT", "RMTI", "RCKY", "ROG", "ROKU", "ROSE", "RST", "RDC", "RTIX", "RTEC", "RUSHA", "RUSHB", "RGNX", "RM", "RGS", "REIS", "RBNC", "MARK", "RNST", "REGI", "RCII", "RGEN", "RBCAA", "FRBK", "REN", "RECN", "TORC", "SASR", "JBSS", "SGMO", "SANM", "BFS", "SVRA", "SBBX", "SCSC", "SCHN", "SCHL", "SHLM", "SWM", "SAIC", "SGMS", "SALT", "STNG", "SCPH", "SSP", "SBCF", "CKH", "SMHI", "SHLD", "SPNE", "SEAS", "SCWX", "SLCT", "WTTR", "SIR", "SEM", "SELB", "SIGI", "SEMG", "SMTC", "RUTH", "RYI", "RHP", "STBA", "SBRA", "SB", "SFE", "SAFT", "SAFE", "SGA", "SAIA", "SAIL", "SBH", "SN", "SAFM", "SD", "SSTK", "SIFI", "SIEB", "SNNA", "SIEN", "BSRR", "SIGA", "SIGM", "SIG", "SLAB", "SBOW", "SAMG", "SFNC", "SMPL", "SSD", "SLP", "SBGI", "SITE", "SJW", "SKY", "SKYW", "SNBR", "SFS", "SGH", "SND", "SMBK", "SOI", "SLDB", "SAH", "SONC", "SRNE", "BID", "SEND", "SENEA", "SENS", "SXT", "MCRB", "SRG", "SREV", "SFBS", "SHAK", "SHEN", "SHLO", "SFL", "SCVL", "SHBI", "SSTI", "SFLY", "SR", "SAVE", "SMTA", "STXB", "SPOK", "SPWH", "SBPH", "SPSC", "SPXC", "FLOW", "SRCI", "JOE", "STAA", "STAG", "STMP", "SMP", "SXI", "STFC", "STBZ", "SCS", "STML", "SCL", "SBT", "STRL", "STC", "SF", "SYBT", "SRI", "SSYS", "STRS", "STRA", "RGR", "SJI", "SSB", "SFST", "SMBC", "SONA", "SBSI", "SWX", "SWN", "SP", "SPKE", "ONCE", "SPAR", "SPTN", "SPA", "SPPI", "TRK", "SPRO", "SLD", "SYKE", "SYNL", "SYNA", "SNDX", "SYNH", "SGYP", "SYBX", "SNX", "SYNT", "SYRS", "SYX", "TTOO", "TRHC", "TCMD", "TAHO", "TLRD", "TALO", "TNDM", "SKT", "TMHC", "TISI", "TECD", "TTGT", "TK", "TNK", "TGNA", "TRC", "TDOC", "TLRA", "TNAV", "SMMF", "INN", "SUM", "SNHY", "SXC", "SPWR", "RUN", "SHO", "SMCI", "SPN", "SGC", "SUP", "SUPN", "SVU", "SURF", "SGRY", "SRDX", "TBPH", "THR", "TPRE", "TDW", "TIER", "TTS", "TLYS", "TSBK", "TMST", "TIPT", "TWI", "TITN", "TVTY", "TIVO", "TOCA", "TMP", "TR", "BLD", "TOWR", "CLUB", "TOWN", "TRTX", "TPIC", "TCI", "TRXC", "TVPT", "TZOO", "TREC", "TG", "TREX", "TPH", "TLGT", "TELL", "THC", "TNC", "TEN", "TERP", "TRNO", "TBNK", "TTEK", "TTI", "TTPH", "TXRH", "TGH", "TGTX", "TCS", "MEET", "TTD", "TXMD", "TTMI", "TCX", "TUP", "TPB", "HEAR", "TPC", "TWIN", "TYME", "USCR", "USPH", "SLCA", "UFPT", "UCTT", "UPL", "RARE", "UMBF", "UMH", "UFI", "UNF", "UBSH", "UNB", "UIS", "UNT", "UBSI", "UCFC", "UCBI", "UBNK", "UFCS", "UIHC", "UNFI", "TCBK", "TRS", "TNET", "TPHS", "TSE", "GTS", "TSC", "TRTN", "TBK", "TGI", "TRNC", "TROX", "TBI", "TRUE", "TRUP", "TRST", "TRMK", "TTEC", "USAT", "USAK", "USNA", "UTMD", "VHI", "VLY", "VALU", "VNDA", "VREX", "VRNS", "VGR", "VEC", "VECO", "VRA", "VCYT", "VSTM", "VCEL", "PAY", "VRNT", "VBTX", "VRTV", "VERI", "VRS", "VVI", "VSAT", "VIAV", "VICR", "VRAY", "VKTX", "VLGEA", "UBFO", "USLM", "UTL", "UNTY", "UBX", "USAP", "UVV", "UEIC", "UFPI", "UHT", "UVE", "ULH", "UVSP", "UMRX", "UPLD", "UEC", "UE", "UBA", "ECOL", "WAFD", "WPG", "WRE", "WASH", "WSBF", "WTS", "WVE", "WDFC", "WEB", "WTW", "WMK", "WERN", "WSBC", "WAIR", "WTBA", "WABC", "WMC", "WNEB", "WHG", "WEYS", "WGL", "WSR", "WOW", "WRD", "WLDN", "WLH", "WLFC", "WSC", "WIN", "VHC", "VRTS", "VRTU", "VSH", "VPG", "VSTO", "VTL", "VSLR", "VCRA", "VG", "VYGR", "VSEC", "VUZI", "WTI", "WNC", "WDR", "WAGE", "WD", "HCC", "YELP", "YEXT", "YORW", "YRCW", "ZFGN", "ZAGG", "ZN", "ZIOP", "ZIXI", "ZOES", "ZGNX", "ZOM", "ZS", "ZUMZ", "WING", "WINA", "WGO", "WETF", "WMIH", "WWW", "WWD", "WK", "WRLD", "INT", "WWE", "WOR", "WMGI", "WSFS", "XCRA", "XNCR", "XHR", "XOXO", "XOMA", "XPER",
}

var sp500 = []string{
	"AAPL", "MSFT", "AMZN", "FB", "JPM", "BRK.B", "GOOG", "GOOGL", "JNJ", "XOM", "BAC", "WFC", "V", "UNH", "PFE", "CVX", "T", "INTC", "HD", "VZ", "PG", "CSCO", "BA", "MA", "C", "KO", "MRK", "DIS", "PEP", "CMCSA", "DWDP", "NVDA", "NFLX", "ABBV", "ORCL", "PM", "AMGN", "WMT", "ADBE", "IBM", "MCD", "MDT", "MMM", "HON", "UNP", "ABT", "MO", "GE", "TXN", "ACN", "NKE", "GILD", "CRM", "BKNG", "UTX", "COST", "QCOM", "LLY", "BMY", "PYPL", "TMO", "SLB", "AVGO", "COP", "CAT", "GS", "USB", "UPS", "NEE", "LOW", "LMT", "AXP", "SBUX", "BIIB", "EOG", "PNC", "MS", "AMT", "BDX", "CVS", "ANTM", "CB", "MDLZ", "CSX", "CELG", "OXY", "DHR", "AGN", "TJX", "AET", "MU", "SCHW", "FDX", "ADP", "BLK", "ISRG", "CL", "WBA", "DUK", "RTN", "CHTR", "CME", "SPG", "ATVI", "BK", "GD", "SYK", "NOC", "PSX", "INTU", "SPGI", "AMAT", "SO", "VLO", "NSC", "ILMN", "FOXA", "AIG", "GM", "COF", "D", "MET", "DE", "CI", "CCI", "CTSH", "BSX", "PX", "EMR", "ZTS", "VRTX", "HUM", "TGT", "ESRX", "ITW", "MMC", "ICE", "PRU", "EXC", "KMB", "BBT", "EA", "HPQ", "F", "MAR", "ECL", "KHC", "MPC", "HAL", "SHW", "LYB", "ADI", "AFL", "BAX", "EQIX", "WM", "HCA", "PGR", "STZ", "ETN", "PLD", "TRV", "APD", "DAL", "APC", "AON", "AEP", "JCI", "FIS", "ALL", "KMI", "ROST", "STI", "SYY", "TEL", "PSA", "STT", "PXD", "FISV", "EBAY", "LRCX", "LUV", "VFC", "ROP", "SRE", "EW", "EL", "REGN", "ADSK", "TROW", "MCO", "APH", "ADM", "OKE", "GIS", "CNC", "PPG", "ALXN", "GLW", "ALGN", "YUM", "APTV", "PEG", "WMB", "ORLY", "DFS", "WY", "CXO", "MCK", "ZBH", "MTB", "RHT", "DLR", "DG", "FTV", "AVB", "DXC", "EQR", "ED", "HPE", "IR", "MNST", "KR", "XEL", "CCL", "WELL", "NTRS", "PH", "PCG", "PCAR", "DVN", "PAYX", "KEY", "MCHP", "ROK", "COL", "SWK", "CMI", "EIX", "NTAP", "HLT", "IP", "TWTR", "DLTR", "RF", "A", "CERN", "IDXX", "SYF", "WEC", "FCX", "VTR", "ANDV", "WDC", "NUE", "AMP", "PPL", "FITB", "DTE", "BXP", "WLTW", "FOX", "IQV", "CFG", "FLT", "ES", "MYL", "AZO", "MSI", "NEM", "INFO", "UAL", "HRS", "RCL", "GPN", "HIG", "BBY", "XLNX", "CLX", "CTL", "LH", "KLAC", "SBAC", "CBS", "TDG", "CTAS", "K", "GWW", "NOV", "TSN", "VRSK", "AME", "HES", "MRO", "APA", "SWKS", "HBAN", "SIVB", "TXT", "CMA", "RSG", "LLL", "O", "FAST", "EXPE", "FE", "HST", "ESS", "ETFC", "AAL", "AMD", "ABMD", "AWK", "STX", "CAH", "NBL", "WAT", "MSCI", "TSS", "EFX", "OMC", "AEE", "EVRG", "RMD", "VMC", "MTD", "ETR", "DHI", "PFG", "CBRE", "EMN", "ANSS", "CAG", "DGX", "BHGE", "MKC", "MGM", "VRSN", "LNC", "XL", "WRK", "GPC", "BLL", "CTXS", "LEN", "CHD", "BF.B", "HSY", "TTWO", "TIF", "EXPD", "XYL", "SNPS", "GGP", "DRI", "CA", "BR", "ULTA", "CMS", "ABC", "KMX", "CHRW", "FTI", "L", "ARE", "TPR", "AJG", "SJM", "MLM", "AKAM", "WYNN", "HSIC", "TAP", "CDNS", "DOV", "EQT", "IT", "COO", "MAS", "URI", "VNO", "HCP", "KSU", "KSS", "CNP", "SYMC", "HFC", "CPRT", "M", "MHK", "CMG", "FMC", "RJF", "EXR", "PVH", "HOLX", "MAA", "CF", "CINF", "HAS", "XRAY", "NWL", "UHS", "INCY", "ADS", "AAP", "FFIV", "QRVO", "ZION", "COG", "BEN", "NDAQ", "IFF", "MOS", "VAR", "JBHT", "HII", "UDR", "DRE", "PKG", "CBOE", "IVZ", "VIAB", "ALB", "LKQ", "NCLH", "DVA", "PRGO", "NRG", "IRM", "RHI", "HRL", "LNT", "AVY", "KORS", "TSCO", "SNA", "TMK", "SLG", "PKI", "REG", "FRT", "JNPR", "AES", "PNW", "XEC", "BWA", "NI", "ARNC", "RE", "NKTR", "WU", "IPG", "JEC", "AMG", "FBHS", "WHR", "AOS", "DISCK", "DISH", "UNM", "CPB", "FLIR", "FLR", "LB", "ALK", "ALLE", "PHM", "HOG", "NLSN", "RL", "JEF", "GRMN", "PNR", "SEE", "KIM", "AIV", "HBI", "GPS", "HP", "IPGP", "PBCT", "MAC", "COTY", "GT", "TRIP", "JWN", "FLS", "SCG", "AIZ", "LEG", "FL", "NWSA", "NFX", "MAT", "XRX", "EVHC", "SRCL", "BHF", "HRB", "PWR", "DISCA", "UAA", "UA", "NWS",
}

func main() {
	app := cli.NewApp()
	app.Name = "trader"
	app.Usage = "lets get rich"

	app.Commands = []cli.Command{
		{
			Name:    "simulate",
			Aliases: []string{"s"},
			Usage:   "simulate the strategy",
			Action:  simulate,
		},
		{
			Name:    "earnings",
			Aliases: []string{"e"},
			Usage:   "check earnings releases",
			Action:  earnings,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type Strategy struct {
	Name         string
	StartDate    time.Time
	Index        []string
	ThresholdPct float64
	StartCash    float64
	Increment    float64
}

var now = time.Now()

var start = time.Date(now.Year()-5, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

var strategies = []Strategy{
	{
		Name:         "5yr, russell2k, 4% thresh, 2.5k increment, 20k start",
		StartDate:    start,
		Index:        russell2k,
		ThresholdPct: 0.04,
		StartCash:    20000,
		Increment:    2500,
	},
	{
		Name:         "5yr, russell2k, 5% thresh, 2k increment, 20k start",
		StartDate:    start,
		Index:        russell2k,
		ThresholdPct: 0.05,
		StartCash:    20000,
		Increment:    2000,
	},
	{
		Name:         "5yr, russell2k, 5% thresh, 2.5k increment, 20k start",
		StartDate:    start,
		Index:        russell2k,
		ThresholdPct: 0.05,
		StartCash:    20000,
		Increment:    2500,
	},
	{
		Name:         "5yr, russell2k, 5% thresh, 3k increment, 20k start",
		StartDate:    start,
		Index:        russell2k,
		ThresholdPct: 0.05,
		StartCash:    20000,
		Increment:    3000,
	},
	{
		Name:         "5yr, russell2k, 6% thresh, 2.5k increment, 20k start",
		StartDate:    start,
		Index:        russell2k,
		ThresholdPct: 0.06,
		StartCash:    20000,
		Increment:    2500,
	},
}

func simulate(c *cli.Context) error {
	fmt.Println("simulating strategies:")
	for _, s := range strategies {
		total := simulateStrat(s.Index, s.ThresholdPct, s.Increment, s.StartCash, s.StartDate)
		fmt.Printf("%s --> %f\n", s.Name, total)
	}
	return nil
}

func simulateStrat(index []string, percent, increment, startAmount float64, startDate time.Time) float64 {
	day := time.Duration(time.Hour * 24)
	amountHave := startAmount
	portfolio := make(map[string]int)
	for i := startDate; time.Now().Sub(i) > 0; i = i.Add(day) {
		stocks := fetchEarnings(i, index)
		for _, stock := range stocks.Stocks {
			closePrice, change := quoteForDate(stock, i)
			if change < (-1 * percent) { //buy low
				if amountHave > increment {
					amountToBuy := int(increment / closePrice)
					amountHave -= (float64(amountToBuy) * closePrice)
					portfolio[stock] += amountToBuy
				}
			}
			if change > percent { //sell high
				if portfolio[stock] > 0 {
					amountHave += float64(portfolio[stock]) * closePrice
					portfolio[stock] = 0
				}
			}
		}
	}
	return calculateTotal(amountHave, portfolio)
}

func calculateTotal(cash float64, portfolio map[string]int) float64 {
	day := time.Duration(time.Hour * 24)
	yesterday := time.Now().Add(-1 * day)
	for ticker, amount := range portfolio {
		closePrice, _ := quoteForDate(ticker, yesterday)
		cash += (float64(amount) * closePrice)
	}
	return cash
}

type Quotes struct {
	Ticker      string
	ClosePrices map[time.Time]float64
	Changes     map[time.Time]float64
}

func quoteForDate(ticker string, date time.Time) (closePrice float64, change float64) {
	file := fmt.Sprintf("quotes/%s", ticker)
	var result = new(Quotes)
	if _, err := os.Stat(file); err == nil {
		// file exists, check it
		err = Load(file, result)
		Check(err)
		if closePrice, ok := result.ClosePrices[date]; ok {
			if change, ok := result.Changes[date]; ok {
				return closePrice, change
			}
		}
	}
	if result.Changes == nil {
		result.Changes = make(map[time.Time]float64)
	}
	if result.ClosePrices == nil {
		result.ClosePrices = make(map[time.Time]float64)
	}

	day := time.Duration(time.Hour * 24)
	enddate := date.Add(day)
	start := datetime.New(&date)
	end := datetime.New(&enddate)
	params := &chart.Params{
		Symbol:   ticker,
		Interval: datetime.OneDay,
		Start:    start,
		End:      end,
	}
	iter := chart.Get(params)
	for iter.Next() {
		b := iter.Bar()
		diff := b.Close.Sub(b.Open)
		if b.Open.Sign() == 0 {
			return 0.0, 0.0
		}
		chg := diff.Div(b.Open)
		change, _ := chg.Float64()
		closePrice, _ := b.Close.Float64()
		result.Ticker = ticker
		result.Changes[date] = change
		result.ClosePrices[date] = closePrice
		err := Save(file, result)
		Check(err)
		return closePrice, change
	}
	return 0.0, 0.0
}

func earnings(c *cli.Context) error {
	stocks := fetchEarnings(time.Now(), sp500)
	fmt.Print(stocks.Stocks)
	return nil
}

type EarningDate struct {
	Date   time.Time
	Stocks []string
}

func fetchEarnings(date time.Time, index []string) EarningDate {
	url := fmt.Sprintf("https://www.bloomberg.com/markets/api/calendar/earnings/US?locale=en&date=%s", date.Format("2006-01-02"))
	file := fmt.Sprintf("earningdate/%s", date.Format("2006-01-02"))

	var result = new(EarningDate)
	if _, err := os.Stat(file); err == nil {
		// file exists
		err = Load(file, result)
		Check(err)
		return *result
	}
	fmt.Printf("did not find earningdate %s, fetching...\n", date.Format("2006-01-02"))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	stocks := filterEarnings(string(body), index)
	result.Date = date
	result.Stocks = stocks
	err = Save(file, result)
	Check(err)
	return *result
}

func filterEarnings(body string, index []string) []string {
	re := regexp.MustCompile(`/companies/security/(\w{1,4}):US`)
	matches := re.FindAllStringSubmatch(body, -1)
	winners := []string{}
	for _, match := range matches {
		ticker := match[1]
		if isInArray(ticker, index) {
			winners = append(winners, ticker)
		}
	}
	return winners
}

func isInArray(s string, list []string) bool {
	for _, str := range list {
		if s == str {
			return true
		}
	}
	return false
}

// Encode via Gob to file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
