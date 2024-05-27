package engine

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

// LuaStatePool 是 Lua 状态机的资源池
type LuaStatePool struct {
	mu       sync.Mutex
	pool     []*lua.LState
	maxCount int
}

var luaEngine *LuaStatePool

func init() {
	luaEngine = newLuaStatePool(10)
}

// NewLuaStatePool 创建一个新的 Lua 状态机资源池
func newLuaStatePool(maxCount int) *LuaStatePool {
	return &LuaStatePool{
		pool:     make([]*lua.LState, 0),
		maxCount: maxCount,
	}
}

// Get 获取一个 Lua 状态机
func (p *LuaStatePool) Get() *lua.LState {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.pool) == 0 {
		// 如果池中没有可用的状态机，则创建新的状态机
		return lua.NewState()
	}

	// 从池中取出一个状态机
	state := p.pool[len(p.pool)-1]
	p.pool = p.pool[:len(p.pool)-1]
	return state
}

// Put 归还一个 Lua 状态机到资源池
func (p *LuaStatePool) Put(state *lua.LState) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.pool) < p.maxCount {
		// 将状态机归还到池中
		p.pool = append(p.pool, state)
	} else {
		// 如果池中状态机数量已达到最大限制，则关闭状态机
		state.Close()
	}
}
