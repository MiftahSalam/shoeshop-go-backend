```
query products {
  getProducts(input: {keyword: "Blue"}) {
    id
    name
  }
}

query product {
  getProduct(id: "9b5cd22c-5220-4f42-8fcf-71c5169923b6") {
    id
    name
    rating
    reviews {
      rating
      comment
      user {
        name
      }
    }
  }
}

query userLogin {
  login(input: {email: "admin@example.com", password: "123456"}) {
    id
    name
    isAdmin
    token
  }
}

query userProfile {
  getUserProfile {
    name
  }
}

mutation createReview {
  createProductReview(
    input: {productId: "9b5cd22c-5220-4f42-8fcf-71c5169923b6", rating: 5, comment: "good"}
  )
}

mutation pay {
  payOrder(
    id: "53ebdc30-a679-4116-8d7d-46fbd90be57c"
    payment: {id: "123", status: "Paid", email: "admin@example.com", updateTime: "2023-01-31T22:06:43.641733+07:00"}
  ) {
    paymentStatus {
      id
      status
      updateTime
      email
    }
  }
}

query getUserOrder {
  getUserOrders {
    user {
      name
    }
    items {
      product {
        id
        name
        price
        countInStock
      }
    }
    shippingAddress {
      Address
      City
    }
    paymentMethod
    paymentStatus {
      status
      email
      updateTime
    }
    taxPrice
    shippingPrice
    isPaid
    paidAt
    isDelivered
    deliveredAt
    createdAt
  }
}

query getOrder {
  getOrder(id: "53ebdc30-a679-4116-8d7d-46fbd90be57c") {
    user {
      name
    }
    items {
      product {
        id
        name
        price
        countInStock
      }
    }
    shippingAddress {
      Address
      City
    }
    paymentMethod
    paymentStatus {
      status
      email
      updateTime
    }
    taxPrice
    shippingPrice
    isPaid
    paidAt
    isDelivered
    deliveredAt
    createdAt
  }
}

mutation createOrder {
  createOrder(
    input: {items: [{productId: "9b5cd22c-5220-4f42-8fcf-71c5169923b7", Quantity: 10, Price: 1}], shippingAddress: {Address: "mel", City: "KBB", PostalCode: "123456", Country: "Indonesia"}, paymentMethod: "paypal", taxPrice: 1.23, shippingPrice: 43.22, totalPrice: 45.00}
  ) {
    id
    user {
      id
    }
    items {
      name
      quantity
      price
    }
    shippingAddress {
      Address
      City
      PostalCode
      Country
    }
    paymentMethod
    paymentStatus {
      status
      email
      updateTime
    }
  }
}
```
