package premises_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Copy following section to separate file, uncomment, and implemented as needed
// CreateRoomEntry - Saves new room into room list
func (this *implRoomsListAPI) CreateRoomEntry(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// DeleteRoomEntry - Deletes specific room entry
func (this *implRoomsListAPI) DeleteRoomEntry(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetRoomEntry - Provides details about room entry
func (this *implRoomsListAPI) GetRoomEntry(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetRoomList - Provides the room list for a specific building
func (this *implRoomsListAPI) GetRoomList(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateRoomEntry - Updates specific room entry
func (this *implRoomsListAPI) UpdateRoomEntry(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}
