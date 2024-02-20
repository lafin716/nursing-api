package router

func (c *container) RegisterPrescriptionRoute() {
	r := c.router.Group("/prescription")
	{
		r.Get("", c.AuthMiddleware(c.handler.prescription.GetList)...)
		r.Get("/:id", c.AuthMiddleware(c.handler.prescription.GetById)...)
		r.Post("", c.AuthMiddleware(c.handler.prescription.Register)...)
		r.Patch("", c.AuthMiddleware(c.handler.prescription.Update)...)
		r.Delete("/:id", c.AuthMiddleware(c.handler.prescription.Delete)...)
		r.Get("/items", c.AuthMiddleware(c.handler.prescription.GetItemList)...)
		r.Get("/items/:id", c.AuthMiddleware(c.handler.prescription.GetItemById)...)
		r.Post("/items", c.AuthMiddleware(c.handler.prescription.AddItem)...)
		r.Patch("/items", c.AuthMiddleware(c.handler.prescription.UpdateItem)...)
		r.Delete("/items/:id", c.AuthMiddleware(c.handler.prescription.DeleteItem)...)
	}
}
