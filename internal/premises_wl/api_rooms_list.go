/*
 * Room List API
 *
 * Hospital Room List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: pfx@google.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package premises_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomsListAPI interface {

	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateRoomEntry - Saves new room into room list
	CreateRoomEntry(ctx *gin.Context)

	// DeleteRoomEntry - Deletes specific room entry
	DeleteRoomEntry(ctx *gin.Context)

	// GetRoomEntry - Provides details about room entry
	GetRoomEntry(ctx *gin.Context)

	// GetRoomList - Provides the room list for a specific building
	GetRoomList(ctx *gin.Context)

	// UpdateRoomEntry - Updates specific room entry
	UpdateRoomEntry(ctx *gin.Context)

}

// partial implementation of RoomsListAPI - all functions must be implemented in add on files
type implRoomsListAPI struct {

}

func newRoomsListAPI() RoomsListAPI {
	return &implRoomsListAPI{}
}

func (this *implRoomsListAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle( http.MethodPost, "/rooms-list/:buildingId/rooms", this.CreateRoomEntry) 
	routerGroup.Handle( http.MethodDelete, "/rooms-list/:buildingId/rooms/:roomId", this.DeleteRoomEntry) 
	routerGroup.Handle( http.MethodGet, "/rooms-list/:buildingId/rooms/:roomId", this.GetRoomEntry) 
	routerGroup.Handle( http.MethodGet, "/rooms-list/:buildingId/rooms", this.GetRoomList) 
	routerGroup.Handle( http.MethodPut, "/rooms-list/:buildingId/rooms/:roomId", this.UpdateRoomEntry) 

}

// Copy following section to separate file, uncomment, and implemented as needed
// // CreateRoomEntry - Saves new room into room list
// func (this *implRoomsListAPI) CreateRoomEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteRoomEntry - Deletes specific room entry
// func (this *implRoomsListAPI) DeleteRoomEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetRoomEntry - Provides details about room entry
// func (this *implRoomsListAPI) GetRoomEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetRoomList - Provides the room list for a specific building
// func (this *implRoomsListAPI) GetRoomList(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateRoomEntry - Updates specific room entry
// func (this *implRoomsListAPI) UpdateRoomEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

