import { Outlet } from 'react-router-dom'
import './App.css'
import { AppSidebar } from './AppSidebar'
import { SidebarProvider, SidebarTrigger } from './ui/sidebar'
import { ThemeProvider } from './ui/theme-provider'

function App() {

  return (
    <ThemeProvider defaultTheme='dark' storageKey="vite-ui-theme">
      <SidebarProvider>
        <AppSidebar />
        <main>
          <SidebarTrigger />
          <Outlet />
        </main>
      </SidebarProvider>
    </ThemeProvider>
  )
}

export default App
