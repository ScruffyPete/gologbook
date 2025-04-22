import { Entry } from "../types/Entry";

export default function EntryItem({ entry }: { entry: Entry }) {
    return (
        <div className="entry">
            <div className="meta">
                <span className="timestamp">{entry.createdAt}</span>
            </div>
            <div className="body">{entry.body}</div>
        </div>
    )
}