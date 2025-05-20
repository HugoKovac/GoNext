import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite';


const backendUrl = process.env.BACKEND_URL || "http://backend_dev:8080";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    port: 3000,
    proxy: {
      "/api": backendUrl,
    },
  },
  preview: {
    port: 3000,
    proxy: {
      "/api": backendUrl,
    },
    allowedHosts: "gonext.com",
  },
})
