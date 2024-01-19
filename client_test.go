package xenapi

import (
	"flag"
	"fmt"
	"os"

	// "strconv"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestAuthentication(t *testing.T) {
	client, err := NewClient("http://10.71.56.85")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	sessionRef, err := client.Session.LoginWithPassword("root", "veIrti81XOaf", "1.0", "terraform")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	hostRefs, err := client.Host.GetAll(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	fmt.Println(hostRefs)

	sg, err := client.Host.GetSchedGran(sessionRef, hostRefs[0])
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	fmt.Println(sg)
	fmt.Println(reflect.TypeOf(sg))

	// records, err := client.Host.GetAllRecords(sessionRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// // fmt.Println(records)
	// fmt.Println(reflect.TypeOf(records))

	// editions, err := client.Host.GetEditions(sessionRef, hostRefs[0])
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(editions)
	// fmt.Println(reflect.TypeOf(editions))

	// hds, err := client.Host.GetDataSources(sessionRef, hostRefs[0])
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(hds)
	// fmt.Println(reflect.TypeOf(hds))

	// vmRefs, err := client.VM.GetAll(sessionRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(vmRefs)
	// fmt.Println(reflect.TypeOf(vmRefs))

	// for _, ref := range vmRefs {
	// 	vmName, _ := client.VM.GetNameLabel(sessionRef, ref)
	// 	fmt.Println(vmName)

	// 	vbds, err := client.VM.GetAllowedVBDDevices(sessionRef, ref)
	// 	if err != nil {
	// 		t.Log(err)
	// 		t.Fail()
	// 		return
	// 	}
	// 	fmt.Println(vbds)
	// 	fmt.Println(reflect.TypeOf(vbds))

	// 	// vds, err := client.VM.GetDataSources(sessionRef, ref)
	// 	// if err != nil {
	// 	// 	t.Log(err)
	// 	// 	t.Fail()
	// 	// 	return
	// 	// }
	// 	// fmt.Println(vds)
	// 	// fmt.Println(reflect.TypeOf(vds))
	// }

	// vmRecords, err := client.VM.GetAllRecords(sessionRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(vmRecords)
	// fmt.Println(reflect.TypeOf(vmRecords))

	//Create VM
	// templatesRefs, err := client.VM.GetByNameLabel(sessionRef, "CentOS 7")
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(templatesRefs)

	// var templates = []VMRef{}
	// for _, vm := range templatesRefs {
	// 	isATemplate, err := client.VM.GetIsATemplate(sessionRef, vm)
	// 	fmt.Println(isATemplate)
	// 	if err != nil {
	// 		break
	// 	}
	// 	if isATemplate {
	// 		templates = append(templates, vm)
	// 	}
	// }
	// if len(templates) != 1 {
	// 	fmt.Println(templates)
	// 	t.Fail()
	// 	return
	// }
	// vmRef, err := client.VM.Clone(sessionRef, templates[0], "test-CentOS7")
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(vmRef)

	// vmRecord, err := client.VM.GetRecord(sessionRef, vmRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(vmRecord)

	// platform, err := client.VM.GetPlatform(sessionRef, vmRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// fmt.Println(platform)

	// vm := &VMDescriptor{
	// 	VMRef:       vmRef,
	// 	UUID:        vmRecord.UUID,
	// 	Name:        vmRecord.NameLabel,
	// 	Description: vmRecord.NameDescription,
	// 	PowerState:  vmRecord.PowerState,
	// 	IsPV:        vmRecord.PVBootloader != "",
	// 	VCPUCount:   2,
	// 	StaticMemory: Range{
	// 		Min: 8589934592,
	// 		Max: 8589934592,
	// 	},
	// 	DynamicMemory: Range{
	// 		Min: 8589934592,
	// 		Max: 8589934592,
	// 	},
	// 	VIFCount:          len(vmRecord.VIFs),
	// 	VBDCount:          len(vmRecord.VBDs),
	// 	PCICount:          len(vmRecord.AttachedPCIs),
	// 	OtherConfig:       vmRecord.OtherConfig,
	// 	XenstoreData:      vmRecord.XenstoreData,
	// 	HVMBootParameters: vmRecord.HVMBootParams,
	// 	IsATemplate:       vmRecord.IsATemplate,
	// 	Platform:          platform,
	// }

	// otherConfig := vm.OtherConfig
	// otherConfig["base_template_name"] = "CentOS 7"
	// err = client.VM.SetOtherConfig(sessionRef, vmRef, otherConfig)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }

	// err = client.VM.SetMemoryLimits(sessionRef, vmRef, vm.StaticMemory.Min, vm.StaticMemory.Max, vm.DynamicMemory.Min, vm.DynamicMemory.Max)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }

	// err = client.VM.SetVCPUsMax(sessionRef, vmRef, vm.VCPUCount)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// err = client.VM.SetVCPUsAtStartup(sessionRef, vmRef, vm.VCPUCount)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }

	// var vifs []*VIFDescriptor
	// for _, vif := range vifs {
	// 	vif.VM = vm
	// 	if vif.DeviceOrder == 0 {
	// 		vif.DeviceOrder = vif.VM.VIFCount
	// 	}
	// 	vifObject := VIFRecord{
	// 		VM:               vif.VM.VMRef,
	// 		Network:          vif.Network.NetworkRef,
	// 		MTU:              vif.MTU,
	// 		MACAutogenerated: vif.IsAutogeneratedMAC,
	// 		MAC:              vif.MAC,
	// 		Device:           strconv.Itoa(vif.DeviceOrder),
	// 		OtherConfig:      vif.OtherConfig,
	// 		LockingMode:      VifLockingModeNetworkDefault,
	// 	}
	// 	vifRef, err := client.VIF.Create(sessionRef, vifObject)
	// 	fmt.Println(vifRef)
	// 	if err != nil {
	// 		t.Log(err)
	// 		t.Fail()
	// 		return
	// 	}

	// 	vif.VIFRef = vifRef
	// 	vifRecord, err := client.VIF.GetRecord(sessionRef, vif.VIFRef)
	// 	fmt.Println(vifRecord)
	// 	if err != nil {
	// 		t.Log(err)
	// 		t.Fail()
	// 		return
	// 	}
	// 	vif.UUID = vifRecord.UUID
	// 	vif.MTU = vifRecord.MTU
	// 	vif.DeviceOrder, _ = strconv.Atoi(vifRecord.Device) // Error ignored, should not occur
	// 	vif.IsAutogeneratedMAC = vifRecord.MACAutogenerated
	// 	vif.MAC = vifRecord.MAC
	// 	vif.OtherConfig = vifRecord.OtherConfig
	// 	network, err := client.Network.GetRecord(sessionRef, vifRecord.Network)
	// 	fmt.Println(network)
	// 	if err != nil {
	// 		t.Log(err)
	// 		t.Fail()
	// 		return
	// 	}
	// 	vif.Network = &NetworkDescriptor{
	// 		NetworkRef:  vifRecord.Network,
	// 		UUID:        network.UUID,
	// 		Name:        network.NameLabel,
	// 		Description: network.NameDescription,
	// 		MTU:         network.MTU,
	// 		Bridge:      network.Bridge,
	// 	}
	// 	if vif.VM.PowerState == VMPowerStateRunning {
	// 		err = client.VIF.Plug(sessionRef, vif.VIFRef)
	// 		if err != nil {
	// 			t.Log(err)
	// 			t.Fail()
	// 			return
	// 		}
	// 	}

	// }
	// err = client.VM.SetHVMBootParams(sessionRef, vm.VMRef, vm.HVMBootParameters)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }
	// err = client.VM.SetPlatform(sessionRef, vmRef, vm.Platform)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }

	// err = client.VM.Provision(sessionRef, vmRef)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Fail()
	// 	return
	// }

	err = client.Session.Logout(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}

type Range struct {
	Min int
	Max int
}
type VMDescriptor struct {
	UUID              string
	Name              string
	Description       string
	PowerState        VMPowerState
	IsPV              bool
	StaticMemory      Range
	DynamicMemory     Range
	VCPUCount         int
	VIFCount          int
	VBDCount          int
	PCICount          int
	OtherConfig       map[string]string
	XenstoreData      map[string]string
	HVMBootParameters map[string]string
	Platform          map[string]string
	IsATemplate       bool

	VMRef VMRef
}

type NetworkDescriptor struct {
	UUID        string
	Name        string
	Description string
	Bridge      string
	MTU         int

	NetworkRef NetworkRef
}

type VIFDescriptor struct {
	Network            *NetworkDescriptor
	VM                 *VMDescriptor
	UUID               string
	MTU                int
	MAC                string
	IsAutogeneratedMAC bool
	DeviceOrder        int
	OtherConfig        map[string]string

	VIFRef VIFRef
}
