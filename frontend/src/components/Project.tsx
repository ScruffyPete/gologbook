import { useEffect, useState } from "react";
import { Entry } from "../types/Entry";
import EntryItem from "./Entry";
import { Project as ProjectType } from "../types/Project";
import { useNavigate, useParams } from "react-router-dom";


export default function ProjectWrapper() {
    const { projectId } = useParams()
    const [project, setProject] = useState<ProjectType | null>(null)
    const navigate = useNavigate()

    useEffect(() => {
        fetch(`/api/projects/${projectId}`)
            .then(res => res.json())
            .then(setProject)
            .catch(err => console.error('Failed to load project:', err))
    }, [projectId])

    if (!projectId) return <div>Missing project ID</div>
    if (!project) return <div>Loading...</div>

    return <>
        <button onClick={() => navigate(-1)}>← Back</button>
        <Project project={project}></Project>
    </>
}

export function Project({ project }: { project: ProjectType }) {
    const [entries, setEntries] = useState<Entry[]>([])

    useEffect(() => {
        fetch(`/api/projects/${project.id}/entries`)
            .then(res => res.json())
            .then(setEntries)
            .catch(err => console.error('Failed to fetch entries:', err))
    }, [project])

    return (
        <div className="entry-log">
            <header style={{ marginBottom: '1rem' }}>
                <h2>{project.title}</h2>
            </header>
            {entries.map(entry => (
                <EntryItem key={entry.id} entry={entry} />
            ))}
        </div>
    )
}
