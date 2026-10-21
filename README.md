# RELIX v0.01
Biological simulator meant for testing synthetic biology. 
Will be updated soon. Feel free to fork and commit as you wish.

# DESCRIPTION
RELIX is a biological simulator meant for testing synbio.

It contains a truckload of features, meant to make synbio and cellular testing fun.
It can be used as a game, a educational piece of software, a base for a similar simulator, or more.
It is as much a game as a serious piece of software.

Built in Go (With Fyne, GofPDF), This is a amazing piece of software which pushes the limits of the language.
It contains multiple different features, ranging from high-level, easy to understand features, to low-level, lab features.
One of my favorite features out of the bunch is chemex. Containing chemical experimentation, From water to sulfur.

As it is a very young tool (published october 2026) it contains a somewhat small amount of things in chemex, which is subject to change in v0.02 and higher.

It is one of my favorite pieces of software i have made, and i am very proud of it.

# FEATURES

SIMULATION CORE
- Tick-based cell simulation
- Cell division (ATP-triggered)
- Stochastic cell death
- Starvation death
- Environmental glucose pool
- ATP metabolism (Michaelis-Menten)
- Generation tracking
- Parent lineage tracking
- Per-cell age tracking
- Deterministic simulation (seeded RNG)
- Stochastic simulation mode (Poisson sampling)

GENETICS
- Nucleotide-level genome (A/T/G/C)
- Gene model (promoter / RBS / CDS)
- Gene activation by transcription factors
- Gene repression by transcription factors
- Hill kinetics for promoter occupancy
- Basal transcription rate
- Activation coefficient (K)
- Repression coefficient (K)
- Hill coefficient (n)
- RBS strength calculation (G-fraction)
- Codon translation with stop-codon detection
- Point mutation
- Insertion mutation
- Deletion mutation
- Gene coordinate shifting on indel
- Gene knockout (disabled flag)
- CRISPRi graded knockdown
- Operon grouping (shared promoter activity)
- Per-gene knockdown percentage

GENE REGULATION
- Inducer molecules (IPTG, aTc, Arabinose)
- Inducer activation mode
- Inducer repression mode
- Per-gene inducer binding constant
- Quorum sensing via AHL
- AHL production (LuxI-like synthase)
- AHL dilution per tick
- Quorum activation mode
- Quorum repression mode

METABOLISM
- Transcription (DNA -> mRNA)
- Translation (mRNA -> protein)
- mRNA degradation
- Protein degradation
- Metabolite pools (per-cell map)
- Metabolic pathway steps (enzyme / substrate / product)
- Pathway Michaelis-Menten kinetics (Kcat, Km)
- Pathway substrate concentration tracking
- Pathway product accumulation
- Plasmid copy-number expression model

ANALYTICS
- Total population tracking
- Max generation tracking
- Global protein totals
- Global mRNA totals
- Time series history (alive count, ATP, glucose)
- Species history (per-tick species snapshot)
- Steady-state detection (windowed variance)
- Entropy / regime classification

FLUX AND KINETICS
- Flux Balance Analysis
- FBA objective evaluation (ATP)
- FBA bottleneck detection
- Gillespie SSA (stochastic transcription)
- Per-gene SSA trajectories
- Parameter estimation (grid search by SSE)
- Parameter estimation curve fitting
- Bifurcation scan
- Regime classification (extinct / stable / oscillating / growing / declining / unstable)
- Hardy-Weinberg equilibrium analysis
- Chi-square deviation
- Allele frequency estimation

SEQUENCE ANALYSIS
- Restriction enzyme table (20 enzymes)
- IUPAC-aware recognition site matching
- Circular DNA matching (origin-spanning)
- Linear DNA matching
- Cut site position calculation
- Fragment size computation
- Multi-enzyme digest
- ASCII gel rendering (log-scale fragment widths)
- Fragment color-coding by size bracket

VISUALIZATION
- Cell colony grid
- Cell color by generation
- Cell color by energy
- Selected-cell highlight
- Genome linear view (nucleotide row)
- Gene annotation overlay
- Codon color coding (A/T/G/C)
- Per-gene selection highlight
- Circular plasmid map
- Gene arcs
- Origin marker
- Restriction site markers
- Plasmid legend
- Lineage tree
- Phase portrait (2-species state space)
- Trajectory age gradient
- Dose-response Hill fit curve
- Bifurcation curve (regime-colored)
- Time series sparkline
- Species trajectory output (Series tab)

EDITING
- Raw DNA editor
- Nucleotide cycling (A -> T -> G -> C)
- Gene name selection
- Gene knockout toggle
- Gene knockdown toggle
- Pathway step addition
- Pathway step removal
- Parameter sliders (mutation, death, division ATP, speed)
- Max cell count
- Inducer sliders (IPTG, aTc, Arabinose)
- Chemical concentration sliders
- DNA apply button
- Reset button
- Play / pause
- Step
- Tick speed control
- Cell selection (list)
- Cell selection (grid)

CHEMISTRY
- 12 chemical species table
- Per-species toxicity coefficient
- Per-species nutrient coefficient
- Per-species oxidant coefficient
- Concentration sliders per species
- Total toxicity computation
- Nutrient multiplier computation
- Toxicity -> death probability
- Nutrient -> ATP production multiplier

PERSISTENCE
- Save project to JSON
- Load project from JSON
- Full state round-trip (cells, genomes, params, history, chemistry, inducers, AHL, stochastic flag)

EXPORT
- Export to .txt
- Export to .pdf
- Export to CSV (time-course wide format)
- Export to SBML L3V1
- Export to SBOL 2.3.0
- SBOL plasmid ComponentDefinition
- SBOL gene ComponentDefinitions
- SBOL SequenceAnnotation with Range
- SBOL per-CDS Sequence

UI
- Lab-themed dark UI
- Menu bar (File, Simulate)
- Top toolbar (Play/Step/Reset)
- Live status (tick, alive, max generation, env pool, AHL)
- Cell list panel
- Genome editor panel
- Inspector tab (details + DNA + factors + pathways)
- Lineage tab
- Phase tab
- Dose tab
- Chemex tab
- FBA tab
- Series tab
- Bifurcation tab
- Genetics tab
- Digest tab
- Plasmid tab
- Per-tab controls
- Settings dialog (max instructions, min string length, max strings, hex dump size, theme)

