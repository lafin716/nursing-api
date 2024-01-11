package router

func (c *container) RegisterMedicineRoute() {
	userRouter := c.router.Group("/medicine")
	{
		userRouter.Get("/search/:pillName", c.AuthMiddleware(c.handler.medicine.Search)...)
	}
}
