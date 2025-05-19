import { Outlet, Navigate } from "react-router";

export function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const user = null;

  if (!user) {
    return user ? <Outlet /> : <Navigate to="/login" replace />;
  }

  return children;
}