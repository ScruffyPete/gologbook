import { useEffect, useState } from "react";
import { Entry } from "../types/Entry";
import EntryItem from "./Entry";
import { Project as ProjectType } from "../types/Project";

export default function Project({ project }: { project: ProjectType }) {
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