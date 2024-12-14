// src/pages/OrdersPage.js
import React, { useEffect, useState } from "react";
import { fetchOrders } from "../api/orders";

const OrdersPage = () => {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    fetchOrders()
      .then((data) => setOrders(data))
      .catch((err) => console.error("Error fetching orders:", err));
  }, []);

  return (
    <div>
      <h1>Orders</h1>
      <ul>
        {orders.map((order) => (
          <li key={order.order_id}>
            Order #{order.order_id} - Status: {order.status}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default OrdersPage;
