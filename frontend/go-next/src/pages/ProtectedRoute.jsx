import { Navigate, Outlet } from "react-router";
import { useAuth } from "../context/AuthContext";

const ProtectedRoute = () => {
  const { user, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!user) {
    return <Navigate to="/signin" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;