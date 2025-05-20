import { Outlet, Navigate, useLocation } from "react-router";


export function ProtectedRoute() {
  const location = useLocation();
  const user = null;

  if (!user) {
    return <Navigate to="/signin" state={{ from: location }} replace />;
  }

  return <Outlet />;
}