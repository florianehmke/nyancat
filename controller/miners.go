package controller

import "sync"

type MinerRegistry struct {
	sync.RWMutex
	miners map[string]bool
}

func NewMinerRegistry() *MinerRegistry {
	return &MinerRegistry{
		miners: make(map[string]bool),
	}
}

func (mr *MinerRegistry) Miners() []string {
	mr.RLock()
	keys := make([]string, len(mr.miners))

	i := 0
	for k := range mr.miners {
		keys[i] = k
		i++
	}

	mr.RUnlock()
	return keys
}

func (mr *MinerRegistry) RemoveMiner(miner string) {
	mr.Lock()
	delete(mr.miners, miner)
	mr.Unlock()
}

func (mr *MinerRegistry) AddMiner(miner string) {
	mr.Lock()
	mr.miners[miner] = true
	mr.Unlock()
}
