package router

func (c *container) RegisterMedicineRoute() {
	r := c.router.Group("/medicine")
	{
		r.Get("/search/:pillName", c.AuthMiddleware(c.handler.medicine.Search)...)
	}
}
