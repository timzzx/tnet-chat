package object

import "sync"

type Rooms struct {
	Members map[int]*Member
	mu      sync.Mutex
}

var Room *Rooms

func init() {
	var once sync.Once
	once.Do(func() {
		Room = &Rooms{
			Members: make(map[int]*Member),
		}
	})

}

func (r *Rooms) Add(m *Member) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Members[m.UserId] = m
}

func (r *Rooms) Del(uid int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Members, uid)
}

func (r *Rooms) List() map[int]*Member {
	return r.Members
}
