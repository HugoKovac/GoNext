import { useAuth } from "../context/AuthContext";
import App from "../App";
import ProtectedHome from "./ProtectedHome";

export default function HomeSwitch() {
  const { user } = useAuth();
  return user ? <ProtectedHome /> : <App />;
}