package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	RegisterUserRoutes(router, db)
	RegisterProductRoutes(router, db)
	RegisterOrderRoutes(router, db)
	RegisterPaymentRoutes(router, db)
	RegisterProductImageRoutes(router, db)
	RegisterUserAddressRoutes(router, db)
	RegisterCacheRoutes(router, db)
	RegisterCategoryRoutes(router, db)
	RegisterReviewRoutes(router, db)
	RegisterRoleRoutes(router, db)
	RegisterOrderItemRoutes(router, db)
	RegisterSessionRoutes(router, db)
	RegisterAuditLogRoutes(router, db)
	RegisterShoppingCartRoutes(router, db)
	RegisterCartItemRoutes(router, db)
	RegisterAuthRoutes(router, db)
	RegisterAuthMiddlewareRoutes(router, db)
	RegisterDataRoutes(router)
}
