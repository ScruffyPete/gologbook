import { Outlet } from 'react-router-dom'
import './App.css'
import { ThemeProvider } from './ui/theme-provider'

function App() {

  return (
    <ThemeProvider defaultTheme='dark' storageKey="vite-ui-theme">
      <Outlet />
    </ThemeProvider>
  )
}

export default App
