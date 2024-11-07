import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()], // Plugin para React
  build: {
    target: 'esnext',  // Establece la versión de JavaScript para la compilación
    sourcemap: true    // Genera mapas de origen para depuración
  },
  server: {
    port: 3000,         // Define el puerto en el que corre el servidor de desarrollo
    open: true          // Abre el navegador automáticamente al iniciar el servidor
  }
});
