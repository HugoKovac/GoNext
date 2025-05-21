import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const isDev = mode === "development";
  return {
    plugins: [react(), tailwindcss()],
    server: {
      port: 3000,
      proxy: isDev ? {
          "/api": "http://backend_dev:8080",
        }
      : undefined,
    },
  };
});
