package main

type ChemSpecies struct {
Category    string
Class       string
Name        string
Formula     string
Toxicity    float64
Nutrient    float64
Oxidant     float64
Radiation   float64
pH          float64
Description string
}

var ChemSpeciesTable = []ChemSpecies{
// -------- Water & ROS --------
{"Water & ROS", "neutral", "Water",             "H2O",   0.0, 1.0, 0.0, 0.0, 7.0, "Neutral; mild nutrient"},
{"Water & ROS", "neutral", "Deuterium oxide",   "D2O",   0.3, 0.5, 0.0, 0.0, 7.0, "Heavy water"},
{"Water & ROS", "oxidizer","Hydrogen peroxide", "H2O2",  0.3, 0.0, 0.8, 0.0, 5.5, "Oxidant"},
{"Water & ROS", "oxidizer","Hydroxyl radical",  "OH*",   0.9, 0.0, 1.0, 0.0, 7.0, "Highly reactive ROS"},
{"Water & ROS", "oxidizer","Superoxide",        "O2-",   0.4, 0.0, 0.9, 0.0, 8.0, "ROS"},
{"Water & ROS", "oxidizer","Peroxynitrite",     "ONOO-", 0.8, 0.0, 1.0, 0.0, 7.0, "RNS/ROS hybrid"},
{"Water & ROS", "oxidizer","Singlet oxygen",    "1O2",   0.7, 0.0, 0.9, 0.0, 7.0, "Excited-state O2"},

// -------- Gases --------
{"Gases", "neutral",  "Hydrogen gas",    "H2",   0.0, 0.3, 0.0, 0.0, 7.0, "Electron donor"},
{"Gases", "oxidizer", "Oxygen",          "O2",   0.0, 0.8, 0.0, 0.0, 7.0, "Terminal acceptor"},
{"Gases", "oxidizer", "Ozone",           "O3",   0.6, 0.0, 0.9, 0.0, 7.0, "Strong oxidant"},
{"Gases", "oxidizer", "Chlorine gas",    "Cl2",  0.9, 0.0, 0.7, 0.0, 3.0, "WWI gas"},
{"Gases", "neutral",  "Nitric oxide",    "NO",   0.3, 0.0, 0.5, 0.1, 7.0, "Signaling"},
{"Gases", "acid",     "Sulfur dioxide",  "SO2",  0.4, 0.0, 0.4, 0.0, 3.5, "Irritant"},
{"Gases", "acid",     "Hydrogen sulfide","H2S",  0.5, 0.0, 0.3, 0.0, 4.0, "Signaling at low conc"},
{"Gases", "base",     "Ammonia",         "NH3",  0.4, 0.5, 0.0, 0.0, 11.0, "Nitrogen source"},
{"Gases", "acid",     "Carbon dioxide",  "CO2",  0.0, 0.2, 0.0, 0.0, 4.0, "Respiratory"},
{"Gases", "neutral",  "Carbon monoxide", "CO",   0.9, 0.0, 0.0, 0.0, 7.0, "Hemoglobin binder"},
{"Gases", "acid",     "Hydrogen cyanide","HCN",  0.95, 0.0, 0.0, 0.0, 5.0, "Cytochrome oxidase inhibitor"},
{"Gases", "acid",     "Nitrogen dioxide", "NO2", 0.7, 0.0, 0.4, 0.0, 3.0, "Irritant"},
{"Gases", "neutral",  "Methane",         "CH4",  0.1, 0.3, 0.0, 0.0, 7.0, "Anaerobic product"},

// -------- Acids --------
{"Acids", "acid", "Proton",           "H+",   0.1, 0.0, 0.2, 0.0, 1.0, "pH effect"},
{"Acids", "acid", "Hydrochloric acid","HCl",  0.6, 0.0, 0.1, 0.0, 1.0, "Strong acid"},
{"Acids", "acid", "Sulfuric acid",    "H2SO4",0.7, 0.0, 0.1, 0.0, 1.0, "Strong acid"},
{"Acids", "acid", "Nitric acid",      "HNO3", 0.7, 0.0, 0.3, 0.0, 1.0, "Oxidizing acid"},
{"Acids", "acid", "Phosphoric acid",  "H3PO4",0.5, 0.0, 0.1, 0.0, 2.0, "Weak acid"},
{"Acids", "acid", "Acetic acid",      "AcOH", 0.3, 0.2, 0.1, 0.0, 3.0, "Weak acid"},
{"Acids", "acid", "Lactic acid",      "LacH", 0.2, 0.3, 0.1, 0.0, 3.5, "Fermentation"},
{"Acids", "acid", "Citric acid",      "CitH", 0.2, 0.3, 0.1, 0.0, 3.5, "TCA cycle"},
{"Acids", "acid", "Formic acid",      "HCOOH",0.5, 0.0, 0.1, 0.0, 2.0, "C1 acid"},
{"Acids", "acid", "Oxalic acid",      "H2C2O4",0.6, 0.0, 0.1, 0.0, 1.5, "Dicarboxylic"},
{"Acids", "acid", "Trichloroacetic acid", "TCA", 0.6, 0.0, 0.1, 0.0, 0.5, "Protein precipitant"},

// -------- Bases --------
{"Bases", "base", "Sodium hydroxide",   "NaOH",   0.7, 0.0, 0.0, 0.0, 13.0, "Strong base"},
{"Bases", "base", "Potassium hydroxide","KOH",    0.7, 0.0, 0.0, 0.0, 13.0, "Strong base"},
{"Bases", "base", "Calcium hydroxide",  "Ca(OH)2",0.5, 0.0, 0.0, 0.0, 12.0, "Strong base"},
{"Bases", "base", "Ammonium hydroxide", "NH4OH",  0.4, 0.0, 0.0, 0.0, 11.0, "Weak base"},
{"Bases", "base", "Lithium hydroxide",  "LiOH",   0.7, 0.0, 0.0, 0.0, 13.0, "Strong base"},
{"Bases", "base", "Barium hydroxide",   "Ba(OH)2",0.6, 0.0, 0.0, 0.0, 12.0, "Strong base"},

// -------- Carbonates --------
{"Carbonates", "carbonate", "Sodium carbonate",  "Na2CO3", 0.2, 0.0, 0.0, 0.0, 11.0, "Basic salt"},
{"Carbonates", "carbonate", "Sodium bicarbonate","NaHCO3", 0.1, 0.0, 0.0, 0.0, 8.3,  "Buffer"},
{"Carbonates", "carbonate", "Calcium carbonate", "CaCO3",  0.0, 0.0, 0.0, 0.0, 9.0,  "Shells; chalk"},
{"Carbonates", "carbonate", "Potassium bicarbonate","KHCO3",0.1, 0.0, 0.0, 0.0, 8.3,  "Buffer"},

// -------- Chlorine & Halogens --------
{"Chlorine & Halogens", "oxidizer", "Hypochlorous acid",   "HOCl",  0.5, 0.0, 0.9, 0.0, 4.0,  "Bleach"},
{"Chlorine & Halogens", "oxidizer", "Sodium hypochlorite", "NaClO", 0.4, 0.0, 0.8, 0.0, 11.0, "Household bleach"},
{"Chlorine & Halogens", "oxidizer", "Chlorine dioxide",    "ClO2",  0.6, 0.0, 1.0, 0.0, 3.0,  "Industrial oxidant"},
{"Chlorine & Halogens", "neutral",  "Chloride ion",        "Cl-",   0.0, 0.4, 0.0, 0.0, 7.0,  "Essential anion"},
{"Chlorine & Halogens", "neutral",  "Fluoride ion",        "F-",    0.4, 0.0, 0.0, 0.0, 8.0,  "Enolase inhibitor"},
{"Chlorine & Halogens", "neutral",  "Bromide ion",         "Br-",   0.1, 0.1, 0.0, 0.0, 7.0,  "Trace nutrient"},
{"Chlorine & Halogens", "neutral",  "Iodide ion",          "I-",    0.0, 0.2, 0.0, 0.0, 7.0,  "Thyroid precursor"},

// -------- Sulfur --------
{"Sulfur", "neutral", "Elemental sulfur", "S8",   0.0, 0.0, 0.0, 0.0, 7.0,  "Inert solid"},
{"Sulfur", "base",    "Sodium sulfide",   "Na2S", 0.6, 0.0, 0.2, 0.0, 12.0, "H2S source"},
{"Sulfur", "neutral", "DMSO",             "DMSO", 0.2, 0.0, 0.0, 0.0, 7.0,  "Cryoprotectant"},

// -------- Nitrogen --------
{"Nitrogen", "nutrient", "Nitrate", "NO3-", 0.0, 0.6, 0.0, 0.0, 7.0, "Nitrogen source"},
{"Nitrogen", "nutrient", "Nitrite", "NO2-", 0.3, 0.4, 0.2, 0.0, 7.5, "Toxic >1mM"},
{"Nitrogen", "neutral",  "Urea",    "Urea", 0.6, 0.0, 0.0, 0.0, 7.2, "Chaotrope"},

// -------- Essential metals --------
{"Essential metals", "metal",  "Magnesium ion", "Mg2+", 0.0, 0.3, 0.0, 0.0, 7.0, "Cofactor"},
{"Essential metals", "metal",  "Calcium ion",   "Ca2+", 0.1, 0.2, 0.0, 0.0, 7.0, "Signaling"},
{"Essential metals", "metal",  "Potassium ion", "K+",   0.0, 0.3, 0.0, 0.0, 7.0, "Osmolyte"},
{"Essential metals", "metal",  "Iron(II)",      "Fe2+", 0.1, 0.5, 0.1, 0.0, 6.0, "Fenton"},
{"Essential metals", "metal",  "Iron(III)",     "Fe3+", 0.1, 0.4, 0.0, 0.0, 4.0, "Carrier"},
{"Essential metals", "metal",  "Zinc ion",      "Zn2+", 0.1, 0.4, 0.0, 0.0, 6.0, "Cofactor"},
{"Essential metals", "metal",  "Copper(II)",    "Cu2+", 0.2, 0.3, 0.2, 0.0, 5.5, "Redox-active"},
{"Essential metals", "metal",  "Cobalt(II)",    "Co2+", 0.2, 0.3, 0.0, 0.0, 6.0, "B12 precursor"},
{"Essential metals", "metal",  "Manganese(II)", "Mn2+", 0.1, 0.3, 0.0, 0.0, 6.0, "Cofactor"},
{"Essential metals", "metal",  "Barium ion",    "Ba2+", 0.6, 0.0, 0.0, 0.0, 7.0, "Toxic metal"},
{"Essential metals", "acid",   "Aluminum ion",  "Al3+", 0.4, 0.0, 0.0, 0.0, 3.0, "Hydrolyzes acidic"},
{"Essential metals", "metal",  "Silver ion",    "Ag+",  0.7, 0.0, 0.2, 0.0, 6.0, "Antimicrobial"},

// -------- Heavy metals --------
{"Heavy metals", "metal",    "Lead ion",     "Pb2+", 0.85, 0.0, 0.0, 0.05, 6.0, "Neurotoxin"},
{"Heavy metals", "metal",    "Mercury ion",  "Hg2+", 0.95, 0.0, 0.0, 0.05, 5.0, "Thiol-reactive"},
{"Heavy metals", "metal",    "Cadmium ion",  "Cd2+", 0.85, 0.0, 0.0, 0.03, 6.0, "Nephrotoxin"},
{"Heavy metals", "metal",    "Arsenate",     "AsO4", 0.70, 0.0, 0.0, 0.10, 7.0, "DNA damage"},
{"Heavy metals", "oxidizer", "Chromate",     "Cr6+", 0.90, 0.0, 0.3, 0.00, 4.0, "Carcinogen"},
{"Heavy metals", "metal",    "Nickel ion",   "Ni2+", 0.60, 0.0, 0.0, 0.02, 6.0, "Allergen"},

// -------- Salts & Buffers --------
{"Salts & Buffers", "salt", "Sodium chloride",    "NaCl",  0.0, 0.3, 0.0, 0.0, 7.0, "Osmolyte"},
{"Salts & Buffers", "salt", "Potassium chloride", "KCl",   0.0, 0.3, 0.0, 0.0, 7.0, "K source"},
{"Salts & Buffers", "salt", "Magnesium chloride", "MgCl2", 0.0, 0.3, 0.0, 0.0, 7.0, "Cofactor"},
{"Salts & Buffers", "salt", "Calcium chloride",   "CaCl2", 0.1, 0.2, 0.0, 0.0, 7.0, "Signaling"},
{"Salts & Buffers", "salt", "Tris buffer",        "Tris",  0.0, 0.0, 0.0, 0.0, 7.5, "pH buffering"},
{"Salts & Buffers", "salt", "Sodium sulfate",     "Na2SO4",0.1, 0.0, 0.0, 0.0, 7.0, "Salt"},
{"Salts & Buffers", "salt", "Potassium nitrate",  "KNO3",  0.1, 0.2, 0.0, 0.0, 7.0, "Oxidizer salt"},
{"Salts & Buffers", "salt", "Sodium nitrate",     "NaNO3", 0.1, 0.2, 0.0, 0.0, 7.0, "Preservative"},
{"Salts & Buffers", "salt", "Ammonium chloride",  "NH4Cl", 0.2, 0.1, 0.0, 0.0, 5.5, "Slightly acidic"},
{"Salts & Buffers", "salt", "Copper sulfate",     "CuSO4", 0.4, 0.0, 0.1, 0.0, 4.5, "Algicide"},
{"Salts & Buffers", "salt", "Iron sulfate",       "FeSO4", 0.3, 0.1, 0.0, 0.0, 4.0, "Iron supplement"},
{"Salts & Buffers", "salt", "Zinc sulfate",       "ZnSO4", 0.2, 0.1, 0.0, 0.0, 5.5, "Zinc supplement"},
{"Salts & Buffers", "salt", "Manganese sulfate",  "MnSO4", 0.2, 0.1, 0.0, 0.0, 5.5, "Mn supplement"},

// -------- Antibiotics --------
{"Antibiotics", "neutral", "Ampicillin",      "Amp",  0.6, 0.0, 0.0, 0.0, 7.0, "beta-lactam"},
{"Antibiotics", "neutral", "Kanamycin",       "Kan",  0.6, 0.0, 0.0, 0.0, 7.0, "Aminoglycoside"},
{"Antibiotics", "neutral", "Chloramphenicol", "Cm",   0.5, 0.0, 0.0, 0.0, 7.0, "Ribosome"},
{"Antibiotics", "neutral", "Tetracycline",    "Tet",  0.5, 0.0, 0.0, 0.0, 7.0, "Bacteriostatic"},
{"Antibiotics", "neutral", "Puromycin",       "Puro", 0.7, 0.0, 0.0, 0.0, 7.0, "Premature stop"},
{"Antibiotics", "neutral", "Rifampicin",      "Rif",  0.6, 0.0, 0.0, 0.0, 7.0, "RNA pol"},
{"Antibiotics", "neutral", "Ciprofloxacin",   "Cip",  0.6, 0.0, 0.0, 0.0, 7.0, "Gyrase"},
{"Antibiotics", "neutral", "Streptomycin",    "Str",  0.7, 0.0, 0.0, 0.0, 7.0, "Aminoglycoside"},
{"Antibiotics", "neutral", "Erythromycin",    "Ery",  0.6, 0.0, 0.0, 0.0, 7.0, "Macrolide"},
{"Antibiotics", "neutral", "Vancomycin",      "Van",  0.8, 0.0, 0.0, 0.0, 7.0, "Cell wall"},
{"Antibiotics", "neutral", "Azithromycin",    "Azi",  0.6, 0.0, 0.0, 0.0, 7.0, "Macrolide"},

// -------- Solvents & Detergents --------
{"Solvents & Detergents", "organic",  "Ethanol",        "EtOH",  0.5, 0.0, 0.1, 0.0, 7.0, "Membrane disrupt"},
{"Solvents & Detergents", "organic",  "Methanol",       "MeOH",  0.7, 0.0, 0.1, 0.0, 7.0, "Toxic solvent"},
{"Solvents & Detergents", "detergent","SDS",            "SDS",   0.9, 0.0, 0.0, 0.0, 7.0, "Membrane disrupt"},
{"Solvents & Detergents", "detergent","Triton X-100",   "TX100", 0.7, 0.0, 0.0, 0.0, 7.0, "Nonionic detergent"},
{"Solvents & Detergents", "organic",  "Chloroform",     "CHCl3", 0.6, 0.0, 0.0, 0.0, 7.0, "Hepatotoxic"},
{"Solvents & Detergents", "organic",  "Formaldehyde",   "FA",    0.7, 0.0, 0.0, 0.0, 7.0, "Crosslinker"},
{"Solvents & Detergents", "organic",  "Acetonitrile",   "ACN",   0.5, 0.0, 0.0, 0.0, 7.0, "Aprotic"},
{"Solvents & Detergents", "organic",  "Benzene",        "Bz",    0.8, 0.0, 0.0, 0.0, 7.0, "Carcinogen"},
{"Solvents & Detergents", "organic",  "Toluene",        "Tol",   0.6, 0.0, 0.0, 0.0, 7.0, "Solvent"},
{"Solvents & Detergents", "acid",     "Phenol",         "PhOH",  0.7, 0.0, 0.0, 0.0, 6.0, "Toxic"},
{"Solvents & Detergents", "organic",  "Acetone",        "Ac2O",  0.3, 0.0, 0.0, 0.0, 7.0, "Polar aprotic"},
{"Solvents & Detergents", "organic",  "Isopropanol",    "iPrOH", 0.3, 0.0, 0.0, 0.0, 7.0, "Disinfectant"},

// -------- Sugars & Nutrients --------
{"Sugars & Nutrients", "nutrient", "Glucose",   "Glc",  0.0, 1.0, 0.0, 0.0, 7.0, "Primary carbon"},
{"Sugars & Nutrients", "nutrient", "Fructose",  "Fru",  0.0, 0.8, 0.0, 0.0, 7.0, "Carbon"},
{"Sugars & Nutrients", "nutrient", "Sucrose",   "Suc",  0.0, 0.7, 0.0, 0.0, 7.0, "Disaccharide"},
{"Sugars & Nutrients", "nutrient", "Glycerol",  "Gly",  0.0, 0.5, 0.0, 0.0, 7.0, "Osmolyte"},
{"Sugars & Nutrients", "nutrient", "Ribose",    "Rib",  0.0, 0.6, 0.0, 0.0, 7.0, "RNA sugar"},

// -------- Chelators & Inhibitors --------
{"Chelators & Inhibitors", "chelator", "EDTA",         "EDTA", 0.3, 0.0, 0.0, 0.0, 7.0, "Divalent chelator"},
{"Chelators & Inhibitors", "chelator", "EGTA",         "EGTA", 0.2, 0.0, 0.0, 0.0, 7.0, "Ca chelator"},
{"Chelators & Inhibitors", "chelator", "DTT",          "DTT",  0.2, 0.0, 0.0, 0.0, 7.0, "Reducing"},
{"Chelators & Inhibitors", "chelator", "BME",          "BME",  0.2, 0.0, 0.0, 0.0, 7.0, "Reducing"},
{"Chelators & Inhibitors", "inhibitor","Sodium azide", "NaN3", 0.8, 0.0, 0.0, 0.0, 7.0, "Cox inhibitor"},
{"Chelators & Inhibitors", "chelator", "Deferoxamine", "DFO",  0.3, 0.0, 0.0, 0.0, 7.0, "Fe chelator"},
{"Chelators & Inhibitors", "chelator", "Dimercaprol",  "BAL",  0.4, 0.0, 0.0, 0.0, 7.0, "Heavy-metal antidote"},

// -------- Radioactive --------
{"Radioactive", "radioactive", "Uranium-238",   "U238",  0.70, 0.0, 0.0, 0.30, 7.0, "Alpha"},
{"Radioactive", "radioactive", "Uranium-235",   "U235",  0.70, 0.0, 0.0, 0.40, 7.0, "Fissile"},
{"Radioactive", "radioactive", "Plutonium-239", "Pu239", 0.90, 0.0, 0.0, 0.60, 7.0, "Alpha+fission"},
{"Radioactive", "radioactive", "Radium-226",    "Ra226", 0.90, 0.0, 0.0, 0.80, 7.0, "Alpha"},
{"Radioactive", "radioactive", "Radon-222",     "Rn222", 0.80, 0.0, 0.0, 0.90, 7.0, "Noble gas"},
{"Radioactive", "radioactive", "Cesium-137",    "Cs137", 0.80, 0.0, 0.0, 0.70, 7.0, "Gamma"},
{"Radioactive", "radioactive", "Strontium-90",  "Sr90",  0.70, 0.0, 0.0, 0.50, 7.0, "Beta; bone"},
{"Radioactive", "radioactive", "Iodine-131",    "I131",  0.70, 0.0, 0.0, 0.60, 7.0, "Beta+gamma"},
{"Radioactive", "radioactive", "Cobalt-60",     "Co60",  0.60, 0.0, 0.0, 0.90, 7.0, "Gamma"},
{"Radioactive", "radioactive", "Americium-241", "Am241", 0.70, 0.0, 0.0, 0.40, 7.0, "Alpha"},
{"Radioactive", "radioactive", "Thorium-232",   "Th232", 0.60, 0.0, 0.0, 0.25, 7.0, "Alpha"},
{"Radioactive", "radioactive", "Phosphorus-32", "P32",   0.70, 0.0, 0.0, 0.60, 7.0, "Beta"},
{"Radioactive", "radioactive", "Polonium-210",  "Po210", 0.95, 0.0, 0.0, 0.85, 7.0, "Alpha"},
{"Radioactive", "radioactive", "Technetium-99m","Tc99m", 0.30, 0.0, 0.0, 0.55, 7.0, "SPECT tracer"},
{"Radioactive", "radioactive", "Gallium-67",    "Ga67",  0.40, 0.0, 0.0, 0.45, 7.0, "Imaging isotope"},
{"Radioactive", "radioactive", "Yttrium-90",    "Y90",   0.60, 0.0, 0.0, 0.75, 7.0, "Beta therapy"},
{"Radioactive", "radioactive", "Strontium-89",  "Sr89",  0.55, 0.0, 0.0, 0.55, 7.0, "Bone palliative"},
{"Radioactive", "radioactive", "Lutetium-177",  "Lu177", 0.50, 0.0, 0.0, 0.60, 7.0, "Therapy"},
}

type Chemex struct {
Concentrations map[string]float64
}

func NewChemex() *Chemex {
c := &Chemex{Concentrations: map[string]float64{}}
for _, s := range ChemSpeciesTable {
c.Concentrations[s.Formula] = 0
}
return c
}

func (c *Chemex) TotalToxicity() float64 {
total := 0.0
for _, s := range ChemSpeciesTable {
total += s.Toxicity * c.Concentrations[s.Formula]
}
return total
}

func (c *Chemex) TotalOxidant() float64 {
total := 0.0
for _, s := range ChemSpeciesTable {
total += s.Oxidant * c.Concentrations[s.Formula]
}
return total
}

func (c *Chemex) TotalRadiation() float64 {
total := 0.0
for _, s := range ChemSpeciesTable {
total += s.Radiation * c.Concentrations[s.Formula]
}
return total
}

func (c *Chemex) NutrientMultiplier() float64 {
m := 1.0
for _, s := range ChemSpeciesTable {
m += s.Nutrient * c.Concentrations[s.Formula] * 0.05
}
return m
}
