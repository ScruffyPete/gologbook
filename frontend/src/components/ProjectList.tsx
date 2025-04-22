import { useEffect, useState } from "react"
import { Project } from "../types/Project"

interface ProjectListProps {
    onSelect: (project: Project) => void
}

export default function ProjectList({ onSelect }: ProjectListProps) {
    const [projects, setProjects] = useState<Project[]>([])

    useEffect(() => {
        fetch('projects/')
            .then(res => res.json())
            .then(setProjects)
            .catch(err => console.error('Failed to load projects:', err))
    }, [])

    return (
        <div>
            <h2>Select a project:</h2>
            <ul>
                {projects.map((project) => (
                    <li key={project.id}>
                        <button onClick={() => onSelect(project)}>
                            {project.title}
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    )
}