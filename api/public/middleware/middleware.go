package middleware

// func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var (
// 			id    string
// 			objID primitive.ObjectID
// 			err   error
// 		)
// 		g := c.(*context.GlobalContext)

// 		if id, _ = context.CheckToken(g); err != nil {
// 			return g.JSON(http.StatusUnauthorized, customer.Response{
// 				ErrorMessage: "authorized error",
// 				StatusCode:   http.StatusUnauthorized,
// 			})
// 		}

// 		if objID, err = primitive.ObjectIDFromHex(id); err != nil || objID.IsZero() {
// 			return g.JSON(http.StatusUnauthorized, customer.Response{
// 				ErrorMessage: "authorized error",
// 				StatusCode:   http.StatusUnauthorized,
// 			})
// 		}
// 		return next(g)
// 	}
// }

// func CustomerAccess(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var (
// 			cust customer.Customer
// 			err  error
// 		)
// 		g := c.(*context.GlobalContext)

// 		if err = cust.Current(g); err != nil {
// 			return g.JSON(http.StatusForbidden, customer.Response{
// 				ErrorMessage: "invalid access",
// 				StatusCode:   http.StatusForbidden,
// 			})
// 		}

// 		if !cust.IsCompleted() {
// 			return g.JSON(http.StatusForbidden, customer.Response{
// 				ErrorMessage: "customer account not completed",
// 				StatusCode:   http.StatusForbidden,
// 			})
// 		}

// 		return next(g)
// 	}
// }
