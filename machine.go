//funcoes para monitorar a maquina

package main
import( 
"fmt"
"os"
"net"
)

type StMachineResource struct{
	Memory *StMemory  
	Disk    []*StDisk 	
	Network []*StNetworkInterface 
	CPU  	[]*StCPU 
}

type StMemory struct {  ///procs/meminfo
	MemTotal	uint64
	MemFree		uint64
	MemUsed		uint64 //MemTotal - MemFree
	MemCached	uint64
	MemBuffers  uint64
	SwapTotal 	uint64
	SwapFree 	uint64
	SwapUsed	uint64 //SwapTotal - SwapFree
	SwapCached  uint64
}

type StDisk struct { // cat /proc/diskstats https://www.kernel.org/doc/Documentation/iostats.txt
	Size uint64	//cat /sys/block/sda/size
	Free uint64
	Used uint64
	Name [10]string
}

type StNetworkInterface struct {
	Name [20]string
	Send uint64 			//cat /proc/net/dev
	Received uint64			//cat /proc/net/dev
	Ips []net.IP
}

type StCPU struct {
	Usage float64 //cat /proc/cpuinfo TODO: confirmar isso 
}

func (this *StMachineResource) discoveryMemory(){

}
func (this *StMachineResource) discoveryDisk(){
	
}
func (this *StMachineResource) discoveryNetwork(){
	
}
func (this *StMachineResource) discoveryCpu(){
	
}



func discoveryMachineResource(mr *StMachineResource){
	if mr == nil {
		mr = &StMachineResource{}
	}

	mr.discoveryMemory()
	
	mr.discoveryDisk()
	
	mr.discoveryNetwork()
	
	mr.discoveryCpu()

}



