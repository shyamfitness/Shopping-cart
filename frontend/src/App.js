import { useEffect, useState } from "react";
import "./App.css";
import Login from "./pages/Login";
import Items from "./pages/Items";

function App() {
  const [token, setToken] = useState("");

  useEffect(() => {
    const saved = localStorage.getItem("token");
    if (saved) {
      setToken(saved);
    }
  }, []);

  const handleLogin = (newToken) => {
    localStorage.setItem("token", newToken);
    setToken(newToken);
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setToken("");
  };

  return (
    <div className="App">
      <header>
        <h1>Shopping Cart</h1>
      </header>
      {token ? (
        <Items token={token} onLogout={handleLogout} />
      ) : (
        <Login onLogin={handleLogin} />
      )}
    </div>
  );
}

export default App;
