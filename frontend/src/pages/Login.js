function Login({ onLogin }) {
  const handleLogin = async (e) => {
    e.preventDefault();

    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();

    if (!username || !password) {
      alert("Please enter both username and password");
      return;
    }

    try {
      const res = await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
      });

      const data = await res.json();
      console.log("LOGIN RESPONSE:", data);

      if (!res.ok || !data.token) {
        alert("Invalid username/password");
        return;
      }

      localStorage.setItem("token", data.token);
      if (onLogin) {
        onLogin(data.token);
      } else {
        window.location.href = "/items";
      }
    } catch (err) {
      console.error("Login Error:", err);
      alert("Failed to connect to backend. Make sure the backend is running on http://localhost:8080");
    }
  };

  return (
    <div className="page">
      <h2>Login</h2>
      <form className="card" onSubmit={handleLogin}>
        <label>
          Username
          <input
            id="username"
            type="text"
            required
          />
        </label>
        <label>
          Password
          <input
            id="password"
            type="password"
            required
          />
        </label>
        <button type="submit">
          Login
        </button>
      </form>
      <p className="note">
        Use POST /users first to create a user if you need one.
      </p>
    </div>
  );
}

export default Login;

