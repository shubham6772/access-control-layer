package routes

import "access-control-layer/internal/handlers"

type Router struct {
	AuthHandler     *handlers.AuthHandler
	ActuatorHandler *handlers.ActuatorHandler
}
