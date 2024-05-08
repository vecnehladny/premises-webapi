package premises_wl

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Copy following section to separate file, uncomment, and implemented as needed
// CreateRoomEntry - Saves new room into room list
func (this *implRoomsListAPI) CreateRoomEntry(ctx *gin.Context) {
	updateBuildingFunc(ctx, func(c *gin.Context, building *Building) (*Building, interface{}, int) {
		var room RoomEntry

		if err := c.ShouldBindJSON(&room); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if room.Type == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Room type is required",
			}, http.StatusBadRequest
		}

		if room.Id == "" || room.Id == "@new" {
			room.Id = uuid.NewString()
		}

		conflictIdx := slices.IndexFunc(building.Rooms, func(existingRoom RoomEntry) bool {
			return room.Id == existingRoom.Id
		})

		if conflictIdx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Room already exists",
			}, http.StatusConflict
		}

		building.Rooms = append(building.Rooms, room)

		// Since room is copied by value return reconciled value from the list
		roomIdx := slices.IndexFunc(building.Rooms, func(existingRoom RoomEntry) bool {
			return room.Id == existingRoom.Id
		})
		if roomIdx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save room",
			}, http.StatusInternalServerError
		}
		return building, building.Rooms[roomIdx], http.StatusOK
	})
}

// DeleteRoomEntry - Deletes specific room entry
func (this *implRoomsListAPI) DeleteRoomEntry(ctx *gin.Context) {
	updateBuildingFunc(ctx, func(c *gin.Context, building *Building) (*Building, interface{}, int) {
		roomId := ctx.Param("roomId")

		if roomId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		roomIdx := slices.IndexFunc(building.Rooms, func(existingRoom RoomEntry) bool {
			return roomId == existingRoom.Id
		})

		if roomIdx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		building.Rooms = append(building.Rooms[:roomIdx], building.Rooms[roomIdx+1:]...)
		return building, nil, http.StatusNoContent
	})
}

// GetRoomEntry - Provides details about room entry
func (this *implRoomsListAPI) GetRoomEntry(ctx *gin.Context) {
	updateBuildingFunc(ctx, func(c *gin.Context, building *Building) (*Building, interface{}, int) {
		roomId := ctx.Param("roomId")

		if roomId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		roomIdx := slices.IndexFunc(building.Rooms, func(existingRoom RoomEntry) bool {
			return roomId == existingRoom.Id
		})

		if roomIdx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		// return nil ambulance - no need to update it in db
		return nil, building.Rooms[roomIdx], http.StatusOK
	})
}

// GetRoomList - Provides the room list for a specific building
func (this *implRoomsListAPI) GetRoomList(ctx *gin.Context) {
	updateBuildingFunc(ctx, func(c *gin.Context, building *Building) (*Building, interface{}, int) {
		result := building.Rooms
		if result == nil {
			result = []RoomEntry{}
		}
		// return nil ambulance - no need to update it in db
		return nil, result, http.StatusOK
	})
}

// UpdateRoomEntry - Updates specific room entry
func (this *implRoomsListAPI) UpdateRoomEntry(ctx *gin.Context) {
	updateBuildingFunc(ctx, func(c *gin.Context, building *Building) (*Building, interface{}, int) {
		var room RoomEntry

		if err := c.ShouldBindJSON(&room); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		roomId := ctx.Param("roomId")

		if roomId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		roomIdx := slices.IndexFunc(building.Rooms, func(existingRoom RoomEntry) bool {
			return roomId == existingRoom.Id
		})

		if roomIdx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		if room.Type != "" {
			building.Rooms[roomIdx].Type = room.Type
		}

		if room.Id != "" {
			building.Rooms[roomIdx].Id = room.Id
		}

		if room.Status != "" {
			building.Rooms[roomIdx].Status = room.Status
		}

		if room.Capacity > 0 {
			building.Rooms[roomIdx].Capacity = room.Capacity
		}

		return building, building.Rooms[roomIdx], http.StatusOK
	})
}
