import { useEffect, useState } from "react"
import { Project } from "../types/Project"
import { Link, useNavigate } from "react-router-dom"


export default function ProjectList() {
    const [projects, setProjects] = useState<Project[]>([])
    const navigate = useNavigate()

    useEffect(() => {
        fetch('/api/projects')
            .then(res => res.json())
            .then(setProjects)
            .catch(err => console.error('Failed to load projects:', err))
    }, [])

    async function handleNewProject() {
        const title = prompt('Enter project title:')
        if (!title) return

        try {
            const res = await fetch('/api/projects', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title }),
            })

            if (!res.ok) throw new Error('Failed to cretae project')

            const newProject = await res.json()
            navigate(`/projects/${newProject.id}`)
        } catch (err) {
            console.error(err)
            alert('Could not create project')
        }
    }


    return (
        <>
            <button onClick={handleNewProject} style={{ marginBottom: '1rem' }}>
                + New Project
            </button>
            <div>
                <h2>Select a project:</h2>
                <ul>
                    {projects.map((project) => (
                        <li key={project.id}>
                            <Link to={`/projects/${project.id}`}>{project.title}</Link>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    )
}