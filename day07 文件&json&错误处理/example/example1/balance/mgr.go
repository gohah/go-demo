package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance(insts)
	return
}
