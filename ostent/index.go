package ostent

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os/user"
	"sort"
	"strings"
	"sync"

	"github.com/ostrost/ostent/assets"
	"github.com/ostrost/ostent/client"
	"github.com/ostrost/ostent/cpu"
	"github.com/ostrost/ostent/format"
	"github.com/ostrost/ostent/getifaddrs"
	"github.com/ostrost/ostent/templates"
	"github.com/ostrost/ostent/types"
	metrics "github.com/rcrowley/go-metrics"
	sigar "github.com/rzab/gosigar"
)

func interfaceMeta(ifdata getifaddrs.IfData) types.InterfaceMeta {
	return interfaceMetaFromString(ifdata.Name)
}

func interfaceMetaFromString(name string) types.InterfaceMeta {
	return types.InterfaceMeta{
		NameKey:  name,
		NameHTML: tooltipable(12, name),
	}
}

type diskInfo struct {
	DevName     string
	Total       uint64
	Used        uint64
	Avail       uint64
	UsePercent  float64
	Inodes      uint64
	Iused       uint64
	Ifree       uint64
	IusePercent float64
	DirName     string
}

var TooltipableTemplate *templates.BinTemplate

func tooltipable(limit int, full string) template.HTML {
	html := "ERROR"
	if len(full) > limit {
		short := full[:limit]
		if TooltipableTemplate == nil {
			log.Printf("tooltipableTemplate hasn't been set")
		} else if buf, err := TooltipableTemplate.CloneExecute(struct {
			Full, Short string
		}{
			Full:  full,
			Short: short,
		}); err == nil {
			html = buf.String()
		}
	} else {
		html = template.HTMLEscapeString(full)
	}
	return template.HTML(html)
}

func diskMeta(disk MetricDF) types.DiskMeta {
	devname := disk.DevName.Snapshot().Value()
	dirname := disk.DirName.Snapshot().Value()
	return types.DiskMeta{
		DiskNameHTML: tooltipable(12, devname),
		DirNameHTML:  tooltipable(6, dirname),
		DirNameKey:   dirname,
		DevName:      devname,
	}
}

func username(uids map[uint]string, uid uint) string {
	if s, ok := uids[uid]; ok {
		return s
	}
	s := fmt.Sprintf("%d", uid)
	if usr, err := user.LookupId(s); err == nil {
		s = usr.Username
	}
	uids[uid] = s
	return s
}

func orderProc(procs []types.ProcInfo, cl *client.Client, send *client.SendClient) []types.ProcData {
	if len(procs) > 1 {
		sort.Sort(procOrder{ // not sort.Stable
			procs:   procs,
			seq:     cl.PSSEQ,
			reverse: client.PSBIMAP.SEQ2REVERSE[cl.PSSEQ],
		})
	}

	limitPS := cl.PSlimit
	notdec := limitPS <= 1
	notexp := limitPS >= len(procs)

	if limitPS >= len(procs) { // notexp
		limitPS = len(procs) // NB modified limitPS
	} else {
		procs = procs[:limitPS]
	}

	client.SetBool(&cl.PSnotDecreasable, &send.PSnotDecreasable, notdec)
	client.SetBool(&cl.PSnotExpandable, &send.PSnotExpandable, notexp)
	client.SetString(&cl.PSplusText, &send.PSplusText, fmt.Sprintf("%d+", limitPS))

	uids := map[uint]string{}
	var list []types.ProcData
	for _, proc := range procs {
		list = append(list, types.ProcData{
			PID:      proc.PID,
			Priority: proc.Priority,
			Nice:     proc.Nice,
			Time:     format.FormatTime(proc.Time),
			NameHTML: tooltipable(42, proc.Name),
			UserHTML: tooltipable(12, username(uids, proc.UID)),
			Size:     format.HumanB(proc.Size),
			Resident: format.HumanB(proc.Resident),
		})
	}
	return list
}

type last struct {
	lastinfo
	mutex sync.Mutex
}

type lastinfo struct {
	Generic  generic
	ProcList []types.ProcInfo
}

type IndexData struct {
	Generic generic
	CPU     cpu.CPUInfo
	MEM     types.MEM

	PStable PStable
	PSlinks *PSlinks `json:",omitempty"`

	DFlinks  *DFlinks       `json:",omitempty"`
	DFbytes  types.DFbytes  `json:",omitempty"`
	DFinodes types.DFinodes `json:",omitempty"`

	IFbytes   types.Interfaces
	IFerrors  types.Interfaces
	IFpackets types.Interfaces

	VagrantMachines *vagrantMachines
	VagrantError    string
	VagrantErrord   bool

	DISTRIB        string
	VERSION        string
	PeriodDuration types.Duration // default refresh value for placeholder

	Client client.Client

	IFTABS client.IFtabs
	DFTABS client.DFtabs
}

type IndexUpdate struct {
	Generic  *generic        `json:",omitempty"`
	CPU      *cpu.CPUInfo    `json:",omitempty"`
	MEM      *types.MEM      `json:",omitempty"`
	DFlinks  *DFlinks        `json:",omitempty"`
	DFbytes  *types.DFbytes  `json:",omitempty"`
	DFinodes *types.DFinodes `json:",omitempty"`
	PSlinks  *PSlinks        `json:",omitempty"`
	PStable  *PStable        `json:",omitempty"`

	IFbytes   *types.Interfaces `json:",omitempty"`
	IFerrors  *types.Interfaces `json:",omitempty"`
	IFpackets *types.Interfaces `json:",omitempty"`

	VagrantMachines *vagrantMachines `json:",omitempty"`
	VagrantError    string
	VagrantErrord   bool

	Client *client.SendClient `json:",omitempty"`
}

var lastInfo last

func (la *last) collect() {
	gch := make(chan generic, 1)
	pch := make(chan []types.ProcInfo, 1)
	ifch := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(4) // four so far
	go getRAM(&Reg1s, &wg)
	go getSwap(&Reg1s, &wg)

	go getGeneric(&Reg1s, gch)
	go read_disks(&Reg1s, &wg)
	go read_procs(pch)
	go getInterfaces(&Reg1s, ifch)
	go CollectCPU(&Reg1s, &wg)

	la.mutex.Lock()
	defer la.mutex.Unlock()

	// NB .mutex unchanged
	la.lastinfo = lastinfo{
		Generic:  <-gch,
		ProcList: <-pch,
	}

	la.Generic.IP = <-ifch
	wg.Wait()
}

// ListMetricInterface is a list of types.MetricInterface type. Used for sorting.
type ListMetricInterface []MetricInterface  // satisfying sort.Interface
func (x ListMetricInterface) Len() int      { return len(x) }
func (x ListMetricInterface) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x ListMetricInterface) Less(i, j int) bool {
	a := rx_lo.Match([]byte(x[i].Name))
	b := rx_lo.Match([]byte(x[j].Name))
	if !(a && b) {
		if a {
			return false
		} else if b {
			return true
		}
	}
	return x[i].Name < x[j].Name
}

// MetricInterface is a set of interface metrics.
type MetricInterface struct {
	metrics.Healthcheck // derive from one of (go-)metric types, otherwise it won't be registered
	Name                string
	BytesIn             types.GaugeDiff
	BytesOut            types.GaugeDiff
	ErrorsIn            types.GaugeDiff
	ErrorsOut           types.GaugeDiff
	PacketsIn           types.GaugeDiff
	PacketsOut          types.GaugeDiff
}

// Update reads ifdata and updates the corresponding fields in MetricInterface.
func (mi *MetricInterface) Update(ifdata getifaddrs.IfData) {
	mi.BytesIn.UpdateAbsolute(int64(ifdata.InBytes))
	mi.BytesOut.UpdateAbsolute(int64(ifdata.OutBytes))
	mi.ErrorsIn.UpdateAbsolute(int64(ifdata.InErrors))
	mi.ErrorsOut.UpdateAbsolute(int64(ifdata.OutErrors))
	mi.PacketsIn.UpdateAbsolute(int64(ifdata.InPackets))
	mi.PacketsOut.UpdateAbsolute(int64(ifdata.OutPackets))
}

func (mi MetricInterface) FormatInterface(ip InterfaceParts) types.Interface {
	ing, outg, isbytes := ip(mi)
	deltain, in := ing.Values()
	deltaout, out := outg.Values()
	form := format.HumanUnitless
	deltaForm := format.HumanUnitless // format.Ps
	if isbytes {
		form = format.HumanB
		deltaForm = func(c uint64) string { // , p uint64
			// return format.Bps(8, c, p) // format.Bps64(8, {in,out}, 0)
			return format.HumanBits(c * 8) // passing the bits
		}
	}
	return types.Interface{
		InterfaceMeta: interfaceMetaFromString(mi.Name),
		In:            form(uint64(in)),            // format.HumanB(uint64(in)),  // with units
		Out:           form(uint64(out)),           // format.HumanB(uint64(out)), // with units
		DeltaIn:       deltaForm(uint64(deltain)),  // format.Bps64(8, in, 0),     // with units
		DeltaOut:      deltaForm(uint64(deltaout)), // format.Bps64(8, out, 0),    // with units
	}
}

type InterfaceParts func(MetricInterface) (types.GaugeDiff, types.GaugeDiff, bool)

func (_ IndexRegistry) InterfaceBytes(mi MetricInterface) (types.GaugeDiff, types.GaugeDiff, bool) {
	return mi.BytesIn, mi.BytesOut, true
}
func (_ IndexRegistry) InterfaceErrors(mi MetricInterface) (types.GaugeDiff, types.GaugeDiff, bool) {
	return mi.ErrorsIn, mi.ErrorsOut, false
}
func (_ IndexRegistry) InterfacePackets(mi MetricInterface) (types.GaugeDiff, types.GaugeDiff, bool) {
	return mi.PacketsIn, mi.PacketsOut, false
}

func (ir IndexRegistry) Interfaces(cli *client.Client, send *client.SendClient, ip InterfaceParts) []types.Interface {
	private := ir.ListPrivateInterface()

	client.SetBool(&cli.ExpandableIF, &send.ExpandableIF, len(private) > cli.Toprows)
	client.SetString(&cli.ExpandtextIF, &send.ExpandtextIF, fmt.Sprintf("Expanded (%d)", len(private)))

	if len(private) == 0 {
		return []types.Interface{}
	}
	if len(private) == 1 {
		return []types.Interface{private[0].FormatInterface(ip)}
	}
	sort.Sort(ListMetricInterface(private))
	var public []types.Interface
	for i, mi := range private {
		if !*cli.ExpandIF && i >= cli.Toprows {
			break
		}
		public = append(public, mi.FormatInterface(ip))
	}
	return public
}

// ListPrivateInterface returns list of MetricInterface's by traversing the PrivateInterfaceRegistry.
func (ir *IndexRegistry) ListPrivateInterface() (lmi []MetricInterface) {
	ir.PrivateInterfaceRegistry.Each(func(name string, i interface{}) {
		lmi = append(lmi, i.(MetricInterface))
	})
	return lmi
}

// GetOrRegisterPrivateInterface produces a registered in PrivateInterfaceRegistry MetricInterface.
func (ir *IndexRegistry) GetOrRegisterPrivateInterface(name string) *MetricInterface {
	ir.PrivateMutex.Lock()
	defer ir.PrivateMutex.Unlock()
	if metric := ir.PrivateInterfaceRegistry.Get(name); metric != nil {
		i := metric.(MetricInterface)
		return &i
	}
	i := MetricInterface{
		Name:       name,
		BytesIn:    types.NewGaugeDiff("interface-"+name+".if_octets.rx", ir.Registry),
		BytesOut:   types.NewGaugeDiff("interface-"+name+".if_octets.tx", ir.Registry),
		ErrorsIn:   types.NewGaugeDiff("interface-"+name+".if_errors.rx", ir.Registry),
		ErrorsOut:  types.NewGaugeDiff("interface-"+name+".if_errors.tx", ir.Registry),
		PacketsIn:  types.NewGaugeDiff("interface-"+name+".if_packets.rx", ir.Registry),
		PacketsOut: types.NewGaugeDiff("interface-"+name+".if_packets.tx", ir.Registry),
	}
	ir.PrivateInterfaceRegistry.Register(name, i) // error is ignored
	// errs when the type is not derived from (go-)metrics types
	return &i
}

func (ir *IndexRegistry) GetOrRegisterPrivateDF(fs sigar.FileSystem) *MetricDF {
	ir.PrivateMutex.Lock()
	defer ir.PrivateMutex.Unlock()
	if fs.DirName == "/" {
		fs.DevName = "root"
	} else {
		fs.DevName = strings.Replace(strings.TrimPrefix(fs.DevName, "/dev/"), "/", "-", -1)
	}
	if metric := ir.PrivateDFRegistry.Get(fs.DevName); metric != nil {
		i := metric.(MetricDF)
		return &i
	}
	label := func(tail string) string {
		return fmt.Sprintf("df-%s.df_complex-%s", fs.DevName, tail)
	}
	r, unusedr := ir.Registry, metrics.NewRegistry()
	i := MetricDF{
		DevName:     &StandardMetricString{}, // unregistered
		DirName:     &StandardMetricString{}, // unregistered
		Free:        metrics.NewRegisteredGaugeFloat64(label("free"), r),
		Reserved:    metrics.NewRegisteredGaugeFloat64(label("reserved"), r),
		Total:       metrics.NewRegisteredGauge(label("total"), unusedr),
		Used:        metrics.NewRegisteredGaugeFloat64(label("used"), r),
		Avail:       metrics.NewRegisteredGauge(label("avail"), unusedr),
		UsePercent:  metrics.NewRegisteredGaugeFloat64(label("usepercent"), unusedr),
		Inodes:      metrics.NewRegisteredGauge(label("inodes"), unusedr),
		Iused:       metrics.NewRegisteredGauge(label("iused"), unusedr),
		Ifree:       metrics.NewRegisteredGauge(label("ifree"), unusedr),
		IusePercent: metrics.NewRegisteredGaugeFloat64(label("iusepercent"), unusedr),
	}
	ir.PrivateDFRegistry.Register(fs.DevName, i) // error is ignored
	// errs when the type is not derived from (go-)metrics types
	return &i
}

// ListMetricCPU is a list of types.MetricCPU type. Used for sorting.
type ListMetricCPU []types.MetricCPU  // satisfying sort.Interface
func (x ListMetricCPU) Len() int      { return len(x) }
func (x ListMetricCPU) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x ListMetricCPU) Less(i, j int) bool {
	var (
		juser = x[j].User.Percent.Snapshot().Value()
		jnice = x[j].Nice.Percent.Snapshot().Value()
		jsys  = x[j].Sys.Percent.Snapshot().Value()
		iuser = x[i].User.Percent.Snapshot().Value()
		inice = x[i].Nice.Percent.Snapshot().Value()
		isys  = x[i].Sys.Percent.Snapshot().Value()
	)
	return (juser + jnice + jsys) < (iuser + inice + isys)
}

func (ir IndexRegistry) DFbytes(seq types.SEQ, cli *client.Client, send *client.SendClient) []types.DiskBytes {
	private := ir.ListPrivateDisk()

	client.SetBool(&cli.ExpandableDF, &send.ExpandableDF, len(private) > cli.Toprows)
	client.SetString(&cli.ExpandtextDF, &send.ExpandtextDF, fmt.Sprintf("Expanded (%d)", len(private)))

	if len(private) == 0 {
		return []types.DiskBytes{}
	}
	if len(private) == 1 {
		return []types.DiskBytes{private[0].FormatDFbytes()}
	}
	sort.Stable(diskOrder{
		disks:   private,
		seq:     seq,
		reverse: client.DFBIMAP.SEQ2REVERSE[seq],
	})

	var public []types.DiskBytes
	for i, disk := range private {
		if !*cli.ExpandDF && i > cli.Toprows-1 {
			break
		}
		public = append(public, disk.FormatDFbytes())
	}
	return public
}

func (md MetricDF) FormatDFbytes() types.DiskBytes {
	var (
		diskTotal = md.Total.Snapshot().Value()
		diskUsed  = md.Used.Snapshot().Value()
		diskAvail = md.Avail.Snapshot().Value()
	)
	total, approxtotal, _ := format.HumanBandback(uint64(diskTotal))
	used, approxused, _ := format.HumanBandback(uint64(diskUsed))
	return types.DiskBytes{
		DiskMeta:        diskMeta(md),
		Total:           total,
		Used:            used,
		Avail:           format.HumanB(uint64(diskAvail)),
		UsePercent:      format.FormatPercent(approxused, approxtotal),
		UsePercentClass: format.LabelClassColorPercent(format.Percent(approxused, approxtotal)),
	}
}

func (ir IndexRegistry) DFinodes(seq types.SEQ, cli *client.Client, send *client.SendClient) []types.DiskInodes {
	private := ir.ListPrivateDisk()

	client.SetBool(&cli.ExpandableDF, &send.ExpandableDF, len(private) > cli.Toprows)
	client.SetString(&cli.ExpandtextDF, &send.ExpandtextDF, fmt.Sprintf("Expanded (%d)", len(private)))

	if len(private) == 0 {
		return []types.DiskInodes{}
	}
	if len(private) == 1 {
		return []types.DiskInodes{private[0].FormatDFinodes()}
	}
	sort.Stable(diskOrder{
		disks:   private,
		seq:     seq,
		reverse: client.DFBIMAP.SEQ2REVERSE[seq],
	})

	var public []types.DiskInodes
	for i, disk := range private {
		if !*cli.ExpandDF && i > cli.Toprows-1 {
			break
		}
		public = append(public, disk.FormatDFinodes())
	}
	return public
}

func (md MetricDF) FormatDFinodes() types.DiskInodes {
	var (
		diskInodes = md.Inodes.Snapshot().Value()
		diskIused  = md.Iused.Snapshot().Value()
		diskIfree  = md.Ifree.Snapshot().Value()
	)
	itotal, approxitotal, _ := format.HumanBandback(uint64(diskInodes))
	iused, approxiused, _ := format.HumanBandback(uint64(diskIused))
	return types.DiskInodes{
		DiskMeta:         diskMeta(md),
		Inodes:           itotal,
		Iused:            iused,
		Ifree:            format.HumanB(uint64(diskIfree)),
		IusePercent:      format.FormatPercent(approxiused, approxitotal),
		IusePercentClass: format.LabelClassColorPercent(format.Percent(approxiused, approxitotal)),
	}
}

func (ir IndexRegistry) CPU(cli *client.Client, send *client.SendClient) []cpu.CoreInfo {
	private := ir.ListPrivateCPU()

	client.SetBool(&cli.ExpandableCPU, &send.ExpandableCPU, len(private) > cli.Toprows) // one row reserved for "all N"
	client.SetString(&cli.ExpandtextCPU, &send.ExpandtextCPU, fmt.Sprintf("Expanded (%d)", len(private)))

	if len(private) == 0 { // no cpu?
		return []cpu.CoreInfo{}
	}
	if len(private) == 1 {
		return []cpu.CoreInfo{FormatCPU(private[0])}
	}
	sort.Sort(ListMetricCPU(private))
	var public []cpu.CoreInfo
	if !*cli.ExpandCPU {
		public = append(public, FormatCPU(ir.PrivateCPUAll))
	}
	for i, mc := range private {
		if !*cli.ExpandCPU && i > cli.Toprows-2 {
			// "collapsed" view, head of the list
			break
		}
		public = append(public, FormatCPU(mc))
	}
	return public
}

func FormatCPU(mc types.MetricCPU) cpu.CoreInfo {
	user := uint(mc.User.Percent.Snapshot().Value()) // rounding
	// .Nice is unused
	sys := uint(mc.Sys.Percent.Snapshot().Value())   // rounding
	idle := uint(mc.Idle.Percent.Snapshot().Value()) // rounding
	N := mc.N
	if prefix := "cpu-"; strings.HasPrefix(N, prefix) { // true for all but "all"
		N = "#" + N[len(prefix):] // fmt.Sprintf("#%d", n)
	}
	return cpu.CoreInfo{
		N:         N,
		User:      user,
		Sys:       sys,
		Idle:      idle,
		UserClass: format.TextClassColorPercent(user),
		SysClass:  format.TextClassColorPercent(sys),
		IdleClass: format.TextClassColorPercent(100 - idle),
	}
}

// ListPrivateCPU returns list of types.MetricCPU's by traversing the PrivateCPURegistry.
func (ir *IndexRegistry) ListPrivateCPU() (lmc []types.MetricCPU) {
	ir.PrivateCPURegistry.Each(func(name string, i interface{}) {
		lmc = append(lmc, i.(types.MetricCPU))
	})
	return lmc
}

// ListPrivateDisk returns list of types.MetricDF's by traversing the PrivateDFRegistry.
func (ir *IndexRegistry) ListPrivateDisk() (lmd []MetricDF) {
	ir.PrivateDFRegistry.Each(func(name string, i interface{}) {
		lmd = append(lmd, i.(MetricDF))
	})
	return lmd
}

// GetOrRegisterPrivateCPU produces a registered in PrivateCPURegistry MetricCPU.
func (ir *IndexRegistry) GetOrRegisterPrivateCPU(coreno int) *types.MetricCPU {
	ir.PrivateMutex.Lock()
	defer ir.PrivateMutex.Unlock()
	name := fmt.Sprintf("cpu-%d", coreno)
	if metric := ir.PrivateCPURegistry.Get(name); metric != nil {
		i := metric.(types.MetricCPU)
		return &i
	}
	i := types.NewMetricCPU(ir.Registry, name)
	ir.PrivateCPURegistry.Register(name, i) // error is ignored
	// errs when the type is not derived from (go-)metrics types
	return &i
}

func (ir IndexRegistry) MEM(client client.Client) *types.MEM {
	gr := ir.RAM
	mem := new(types.MEM)
	mem.List = []types.Memory{
		_getmem("RAM", sigar.Swap{
			Total: uint64(gr.Total.Snapshot().Value()),
			Free:  uint64(gr.Free.Snapshot().Value()),
			Used:  gr.UsedValue(), // == .Total - .Free
		}),
	}
	if !*client.HideSWAP {
		gs := ir.Swap
		mem.List = append(mem.List,
			_getmem("swap", sigar.Swap{
				Total: gs.TotalValue(),
				Free:  uint64(gs.Free.Snapshot().Value()),
				Used:  uint64(gs.Used.Snapshot().Value()),
			}))
	}
	return mem
}

func (ir IndexRegistry) LA() string {
	gl := ir.Load
	return gl.Short.Sparkline() + " " + fmt.Sprintf("%.2f %.2f %.2f",
		gl.Short.Snapshot().Value(),
		gl.Mid.Snapshot().Value(),
		gl.Long.Snapshot().Value())
}

func (ir *IndexRegistry) UpdateDF(fs sigar.FileSystem, usage sigar.FileSystemUsage) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	ir.GetOrRegisterPrivateDF(fs).Update(fs, usage)
}

func (ir *IndexRegistry) UpdateRAM(got sigar.Mem, extra1, extra2 uint64) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	ir.RAM.Update(got, extra1, extra2)
}

// UpdateSwap reads got and updates the ir.Swap. TODO Bad description.
func (ir *IndexRegistry) UpdateSwap(got sigar.Swap) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	ir.Swap.Update(got)
}

func (ir *IndexRegistry) UpdateLoadAverage(la sigar.LoadAverage) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	ir.Load.Short.Update(la.One)
	ir.Load.Mid.Update(la.Five)
	ir.Load.Long.Update(la.Fifteen)
}

func (ir *IndexRegistry) UpdateCPU(cpus []sigar.Cpu) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	all := sigar.Cpu{}
	for coreno, core := range cpus {
		ir.GetOrRegisterPrivateCPU(coreno).Update(core)
		types.CPUAdd(&all, core)
	}
	if ir.PrivateCPUAll.N == "all" {
		ir.PrivateCPUAll.N = fmt.Sprintf("all %d", len(cpus))
	}
	ir.PrivateCPUAll.Update(all)
}

func (ir *IndexRegistry) UpdateIFdata(ifdata getifaddrs.IfData) {
	ir.Mutex.Lock()
	defer ir.Mutex.Unlock()
	ir.GetOrRegisterPrivateInterface(ifdata.Name).Update(ifdata)
}

type IndexRegistry struct {
	Registry                 metrics.Registry
	PrivateCPUAll            types.MetricCPU
	PrivateCPURegistry       metrics.Registry // set of MetricCPUs is handled as a metric in this registry
	PrivateInterfaceRegistry metrics.Registry // set of MetricInterfaces is handled as a metric in this registry
	PrivateDFRegistry        metrics.Registry // set of MetricDFs is handled as a metric in this registry
	PrivateMutex             sync.Mutex

	RAM  types.MetricRAM
	Swap types.MetricSwap
	Load types.MetricLoad

	Mutex sync.Mutex
}

var Reg1s IndexRegistry

func init() {
	Reg1s = IndexRegistry{
		Registry:                 metrics.NewRegistry(),
		PrivateCPURegistry:       metrics.NewRegistry(),
		PrivateInterfaceRegistry: metrics.NewRegistry(),
		PrivateDFRegistry:        metrics.NewRegistry(),
	}
	// Reg1s.PrivateCPUAll = *Reg1s.RegisterCPU(metrics.NewRegistry(), "all")
	Reg1s.PrivateCPUAll = types.NewMetricCPU( /* pcreg := */ metrics.NewRegistry(), "all")
	// pcreg.Register("all", Reg1s.PrivateCPUAll)

	Reg1s.RAM = types.NewMetricRAM(Reg1s.Registry)
	Reg1s.Swap = types.NewMetricSwap(Reg1s.Registry)
	Reg1s.Load = types.NewMetricLoad(Reg1s.Registry)

	// addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2003")
	// go metrics.Graphite(reg, 1*time.Second, "ostent", addr)
}

func getUpdates(req *http.Request, cl *client.Client, send client.SendClient, forcerefresh bool) IndexUpdate {

	cl.RecalcRows() // before anything

	var ps_copy []types.ProcInfo
	iu := IndexUpdate{}
	func() {
		lastInfo.mutex.Lock()
		defer lastInfo.mutex.Unlock()

		ps_copy = make([]types.ProcInfo, len(lastInfo.ProcList))
		copy(ps_copy, lastInfo.ProcList)

		if true { // cl.RefreshGeneric.Refresh(forcerefresh)
			g := lastInfo.Generic
			g.LA = Reg1s.LA()
			iu.Generic = &g // &lastInfo.Generic
		}
	}()

	if req != nil {
		req.ParseForm() // do ParseForm even if req.Form == nil, otherwise *links won't be set for index requests without parameters
		base := url.Values{}
		iu.PSlinks = (*PSlinks)(types.NewLinkAttrs(req, base, "ps", client.PSBIMAP, &cl.PSSEQ))
		iu.DFlinks = (*DFlinks)(types.NewLinkAttrs(req, base, "df", client.DFBIMAP, &cl.DFSEQ))
	}

	if !*cl.HideMEM && cl.RefreshMEM.Refresh(forcerefresh) {
		iu.MEM = Reg1s.MEM(*cl)
	}
	if !*cl.HideCPU && cl.RefreshCPU.Refresh(forcerefresh) {
		iu.CPU = &cpu.CPUInfo{List: Reg1s.CPU(cl, &send)}
	}

	if !*cl.HideDF && cl.RefreshDF.Refresh(forcerefresh) {
		if *cl.TabDF == client.DFBYTES_TABID {
			iu.DFbytes = &types.DFbytes{List: Reg1s.DFbytes(cl.DFSEQ, cl, &send)}
		} else if *cl.TabDF == client.DFINODES_TABID {
			iu.DFinodes = &types.DFinodes{List: Reg1s.DFinodes(cl.DFSEQ, cl, &send)}
		}
	}

	if !*cl.HideIF && cl.RefreshIF.Refresh(forcerefresh) {
		switch *cl.TabIF {
		case client.IFBYTES_TABID:
			iu.IFbytes = &types.Interfaces{List: Reg1s.Interfaces(cl, &send, Reg1s.InterfaceBytes)}
		case client.IFERRORS_TABID:
			iu.IFerrors = &types.Interfaces{List: Reg1s.Interfaces(cl, &send, Reg1s.InterfaceErrors)}
		case client.IFPACKETS_TABID:
			iu.IFpackets = &types.Interfaces{List: Reg1s.Interfaces(cl, &send, Reg1s.InterfacePackets)}
		}
	}

	if !*cl.HidePS && cl.RefreshPS.Refresh(forcerefresh) {
		iu.PStable = &PStable{List: orderProc(ps_copy, cl, &send)}
	}

	if !*cl.HideVG && cl.RefreshVG.Refresh(forcerefresh) {
		machines, err := vagrantmachines()
		if err != nil {
			iu.VagrantError = err.Error()
			iu.VagrantErrord = true
		} else {
			iu.VagrantMachines = machines
			iu.VagrantErrord = false
		}
	}

	if send != (client.SendClient{}) {
		iu.Client = &send
	}
	return iu
}

func indexData(minrefresh types.Duration, req *http.Request) IndexData {
	if Connections.Len() == 0 {
		// collect when there're no active connections, so Loop does not collect
		lastInfo.collect()
	}

	cl := client.DefaultClient(minrefresh)
	updates := getUpdates(req, &cl, client.SendClient{}, true)

	data := IndexData{
		Client:  cl,
		Generic: *updates.Generic,
		CPU:     *updates.CPU,
		MEM:     *updates.MEM,

		DFlinks: updates.DFlinks,
		PSlinks: updates.PSlinks,

		PStable: *updates.PStable,

		DISTRIB: DISTRIB, // value set in init()
		VERSION: VERSION, // value from server.go

		PeriodDuration: minrefresh, // default refresh value for placeholder
	}

	if updates.DFbytes != nil {
		data.DFbytes = *updates.DFbytes
	} else if updates.DFinodes != nil {
		data.DFinodes = *updates.DFinodes
	}

	if updates.IFbytes != nil {
		data.IFbytes = *updates.IFbytes
	} else if updates.IFerrors != nil {
		data.IFerrors = *updates.IFerrors
	} else if updates.IFpackets != nil {
		data.IFpackets = *updates.IFpackets
	}
	data.VagrantMachines = updates.VagrantMachines
	data.VagrantError = updates.VagrantError
	data.VagrantErrord = updates.VagrantErrord

	data.DFTABS = client.DFTABS // const
	data.IFTABS = client.IFTABS // const

	return data
}

func statusLine(status int) string {
	return fmt.Sprintf("%d %s", status, http.StatusText(status))
}

func init() {
	DISTRIB = getDistrib()
}

var DISTRIB string

func IndexFunc(template *templates.BinTemplate, scripts []string, minrefresh types.Duration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		index(template, scripts, minrefresh, w, r)
	}
}

func index(template *templates.BinTemplate, scripts []string, minrefresh types.Duration, w http.ResponseWriter, r *http.Request) {
	response := template.Response(w, struct {
		Data      IndexData
		SCRIPTS   []string
		CLASSNAME string
	}{
		Data:    indexData(minrefresh, r),
		SCRIPTS: assets.FQscripts(scripts, r),
	})
	response.SetHeader("Content-Type", "text/html")
	response.SetContentLength()
	response.Send()
}

type MetricString interface {
	Snapshot() MetricString
	Value() string
	Update(string)
}

type StandardMetricString struct {
	string
	Mutex sync.Mutex
}

type MetricStringSnapshot StandardMetricString

func (mss MetricStringSnapshot) Snapshot() MetricString { return mss }
func (mss MetricStringSnapshot) Value() string          { return mss.string }
func (MetricStringSnapshot) Update(string)              { panic("Update called on a MetricStringSnapshot") }

func (sms StandardMetricString) Snapshot() MetricString { return MetricStringSnapshot(sms) }
func (sms StandardMetricString) Value() string {
	sms.Mutex.Lock()
	defer sms.Mutex.Unlock()
	return sms.string
}

func (sms *StandardMetricString) Update(new string) {
	sms.Mutex.Lock()
	defer sms.Mutex.Unlock()
	sms.string = new
}

type MetricDF struct {
	metrics.Healthcheck // derive from one of (go-)metric types, otherwise it won't be registered
	DevName             MetricString
	Free                metrics.GaugeFloat64
	Reserved            metrics.GaugeFloat64
	Total               metrics.Gauge
	Used                metrics.GaugeFloat64
	Avail               metrics.Gauge
	UsePercent          metrics.GaugeFloat64
	Inodes              metrics.Gauge
	Iused               metrics.Gauge
	Ifree               metrics.Gauge
	IusePercent         metrics.GaugeFloat64
	DirName             MetricString
}

// Update reads usage and fs and updates the corresponding fields in MetricDF.
func (md *MetricDF) Update(fs sigar.FileSystem, usage sigar.FileSystemUsage) {
	md.DevName.Update(fs.DevName)
	md.DirName.Update(fs.DirName)
	md.Free.Update(float64(usage.Free << 10))
	md.Reserved.Update(float64((usage.Free - usage.Avail) << 10))
	md.Total.Update(int64(usage.Total << 10))
	md.Used.Update(float64(usage.Used << 10))
	md.Avail.Update(int64(usage.Avail << 10))
	md.UsePercent.Update(usage.UsePercent())
	md.Inodes.Update(int64(usage.Files))
	md.Iused.Update(int64(usage.Files - usage.FreeFiles))
	md.Ifree.Update(int64(usage.FreeFiles))
	if iusePercent := 0.0; usage.Files != 0 {
		iusePercent = float64(100) * float64(usage.Files-usage.FreeFiles) / float64(usage.Files)
		md.IusePercent.Update(iusePercent)
	}
}