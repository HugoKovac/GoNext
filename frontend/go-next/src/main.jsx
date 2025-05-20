import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import { AuthProvider } from "./context/AuthContext";
import ProtectedRoute from "./pages/ProtectedRoute";
import Login from "./pages/Login";
import SignUp from "./pages/Signup";
import ProtectedHome from "./pages/ProtectedHome";
import "./index.css";
import App from "./App.jsx";
import Navbar from "./components/Navbar";

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <BrowserRouter>
      <AuthProvider>
        <Navbar />
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/signin" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/protected" element={<ProtectedRoute />}>
            <Route path="/protected/home" element={<ProtectedHome />} />
          </Route>
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  </StrictMode>
);
