import { useEffect, useState } from "react";
import {
  fetchItems,
  addItemToCart,
  checkoutCart,
  fetchCart,
  fetchOrders,
} from "../api";

function Items({ token, onLogout }) {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const loadItems = async () => {
      try {
        const data = await fetchItems();
        setItems(data || []);
      } catch (err) {
        console.error(err);
        window.alert("Could not load items");
      }
    };
    loadItems();
  }, []);

  const handleAdd = async (itemId) => {
    if (!token) {
      window.alert("Please login again");
      return;
    }
    setLoading(true);
    try {
      await addItemToCart(itemId, token);
      window.alert("Item added to cart");
    } catch (err) {
      window.alert("Could not add item");
    } finally {
      setLoading(false);
    }
  };

  const handleCheckout = async () => {
    if (!token) {
      window.alert("Please login again");
      return;
    }
    try {
      await checkoutCart(token);
      window.alert("Order successful");
    } catch (err) {
      window.alert("Checkout failed");
    }
  };

  const handleCart = async () => {
    if (!token) {
      window.alert("Please login again");
      return;
    }
    try {
      const data = await fetchCart(token);
      const names = (data.cartItems || []).map(
        (ci) => `Cart ${ci.CartID}: item ${ci.ItemID} x${ci.Quantity}`
      );
      window.alert(names.length ? names.join("\n") : "Cart is empty");
    } catch (err) {
      window.alert("Could not load cart");
    }
  };

  const handleOrders = async () => {
    if (!token) {
      window.alert("Please login again");
      return;
    }
    try {
      const data = await fetchOrders(token);
      const ids = (data.orders || []).map((o) => `Order #${o.ID}`);
      window.alert(ids.length ? ids.join("\n") : "No orders yet");
    } catch (err) {
      window.alert("Could not load orders");
    }
  };

  return (
    <div className="page">
      <div className="top-bar">
        <button onClick={handleCheckout}>Checkout</button>
        <button onClick={handleCart}>Cart</button>
        <button onClick={handleOrders}>Orders</button>
        <button className="ghost" onClick={onLogout}>
          Logout
        </button>
      </div>
      <h2>Items</h2>
      <div className="list">
        {items.map((item) => (
          <div
            key={item.ID}
            className="card"
            onClick={() => handleAdd(item.ID)}
          >
            <p className="title">{item.Name}</p>
            <p>${Number(item.Price).toFixed(2)}</p>
            <p className="hint">
              {loading ? "..." : "Click to add to cart (qty 1)"}
            </p>
          </div>
        ))}
        {!items.length && <p>No items yet.</p>}
      </div>
    </div>
  );
}

export default Items;

