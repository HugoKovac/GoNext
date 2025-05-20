import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/public-home.tsx"),
  route("signin", "routes/login.tsx"),
  route("signup", "routes/signup.tsx"),
  route("protected", "utils/ProtectedRoute.tsx", [
    index("routes/protected-home.tsx"),
    // Add more protected routes here
  ]),
] satisfies RouteConfig;
