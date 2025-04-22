import { useState } from 'react'
import './App.css'
import ProjectList from './components/ProjectList'
import Project from './components/Project'
import { Project as ProjectType } from './types/Project'

function App() {
  const [selectedProject, setSelectedProject] = useState<ProjectType | null>(null)

  async function handleNewProject() {
    const title = prompt('Enter project title:')
    if (!title) return

    try {
      const res = await fetch('/api/projects', {
        method: 'POST',
        headers: { 'Content-Type': 'applicatio/json' },
        body: JSON.stringify({ title }),
      })

      if (!res.ok) throw new Error('Failed to cretae project')

      const newProject = await res.json()
      setSelectedProject(newProject)
    } catch (err) {
      console.error(err)
      alert('Could not create project')
    }
  }

  return (
    <main style={{ padding: '1rem', maxWidth: '600px', margin: '0 auto' }}>
      {!selectedProject ? (
        <>
          <button onClick={handleNewProject} style={{ marginBottom: '1rem' }}>
            + New Project
          </button>
          <ProjectList onSelect={setSelectedProject} />
        </>
      ) : (
        <Project project={selectedProject} />
      )}
    </main>
  )
}

export default App
