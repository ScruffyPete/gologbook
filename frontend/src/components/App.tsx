import './App.css'
import ProjectWrapper from './Project'
import ProjectList from './ProjectList'
import { Route, Routes } from 'react-router-dom'

function App() {

  return (
    <main style={{ padding: '1rem', maxWidth: '600px', margin: '0 auto' }}>
      <Routes>
        <Route path='/' element={<ProjectList />} />
        <Route path='projects/:projectId' element={<ProjectWrapper />} />
      </Routes>

    </main>
  )
}

export default App
