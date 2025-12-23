package Hotel

import "fmt"

type Hotel struct {
	Rooms map[string]Room
}

func (h *Hotel) AddRoom(room Room) {
	if h.Rooms[room.RoomNumber].RoomNumber != "" {
		fmt.Println("Room already exists")
		return
	}

	h.Rooms[room.RoomNumber] = room
	fmt.Println("Room added")
}

func (h *Hotel) CheckIn(roomNumber string) {
	room, ok := h.Rooms[roomNumber]
	if !ok {
		fmt.Println("Room not found")
		return
	}

	if room.IsOccupied {
		fmt.Println("Room is already occupied")
		return
	}

	room.IsOccupied = true
	h.Rooms[roomNumber] = room
	fmt.Println("Checked in")
}

func (h *Hotel) CheckOut(roomNumber string) {
	room, ok := h.Rooms[roomNumber]
	if !ok {
		fmt.Println("Room not found")
		return
	}

	if !room.IsOccupied {
		fmt.Println("Room is already free")
		return
	}

	room.IsOccupied = false
	h.Rooms[roomNumber] = room
	fmt.Println("Checked out")
}

func (h *Hotel) ListVacantRooms() {
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Println("Room:", room.RoomNumber, "Type:", room.Type, "Price:", room.PricePerNight)
		}
	}
}
