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
          <SidebarTrigger>

          </SidebarTrigger>
        </main>
      </SidebarProvider>
    </ThemeProvider>
  )

  // return (
  //   <main style={{ padding: '1rem', maxWidth: '600px', margin: '0 auto' }}>
  //     <Routes>
  //       <Route path='/' element={<Navbar />} />
  //       <Route path='projects/:projectId' element={<ProjectWrapper />} />
  //     </Routes>

  //   </main>
  // )
}

export default App
