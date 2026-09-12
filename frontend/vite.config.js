import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  base: './', // Use relative paths for static hosting & GitHub Pages
  plugins: [vue()],
  build: {
    outDir: '../docs', // Build static website directly into /docs for GitHub Pages
    emptyOutDir: true
  },
  server: {
    port: 5173
  }
})
