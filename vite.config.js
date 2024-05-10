import { defineConfig } from "vite";
import gobest from "./resources/assets/script/tools/vite-plugin-backend";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react(), ...gobest({
    input: ["resources/assets/script/main.ts"],
    refresh: false,
  })],
});
