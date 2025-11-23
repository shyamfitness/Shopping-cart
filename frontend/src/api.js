export const BASE_URL = "http://localhost:8080";

async function request(path, method = "GET", body, token) {
  const headers = { "Content-Type": "application/json" };
  if (token) {
    headers["Authorization"] = token;
  }

  const url = `${BASE_URL}${path}`;
  const options = {
    method,
    headers,
  };

  if (body) {
    options.body = JSON.stringify(body);
  }

  const res = await fetch(url, options);

  if (!res.ok) {
    let errorMessage = "request failed";
    try {
      const errorData = await res.json();
      errorMessage = errorData.error || errorMessage;
    } catch {
      const text = await res.text();
      errorMessage = text || errorMessage;
    }
    throw new Error(errorMessage);
  }

  if (res.status === 204) {
    return null;
  }

  return res.json();
}

export function loginUser(username, password) {
  return request("/users/login", "POST", { username, password });
}

export function fetchItems() {
  return request("/items", "GET");
}

export function addItemToCart(itemId, token) {
  return request("/carts", "POST", { item_id: itemId, quantity: 1 }, token);
}

export function checkoutCart(token) {
  return request("/orders", "POST", null, token);
}

export function fetchCart(token) {
  return request("/carts", "GET", null, token);
}

export function fetchOrders(token) {
  return request("/orders", "GET", null, token);
}

