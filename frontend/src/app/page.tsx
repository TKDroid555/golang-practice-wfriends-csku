"use client";

import { useEffect, useState } from "react";

interface Note {
  id: string;
  title: string;
  content: string;
  created_at: string;
}

export default function Home() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(true);

  const fetchNotes = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/notes");
      const json = await res.json();
      if (json.success && json.data) {
        setNotes(json.data);
      } else {
        setNotes([]);
      }
    } catch (error) {
      console.error("Failed to fetch notes:", error);
      setNotes([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchNotes();
  }, []);

  const handleAddNote = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title || !content) return;

    try {
      const res = await fetch("http://localhost:8080/api/notes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, content }),
      });
      const json = await res.json();
      if (json.success) {
        setTitle("");
        setContent("");
        fetchNotes();
      }
    } catch (error) {
      console.error("Failed to add note:", error);
    }
  };

  const handleDeleteNote = async (id: string) => {
    try {
      const res = await fetch(`http://localhost:8080/api/notes/${id}`, {
        method: "DELETE",
      });
      const json = await res.json();
      if (json.success) {
        fetchNotes();
      }
    } catch (error) {
      console.error("Failed to delete note:", error);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 p-8 font-[family-name:var(--font-geist-sans)]">
      <main className="max-w-2xl mx-auto bg-white p-6 rounded-lg shadow-md">
        <h1 className="text-2xl font-bold mb-6 text-center text-gray-800">Note App (Next.js + Golang)</h1>

        <form onSubmit={handleAddNote} className="mb-8 bg-gray-50 p-4 rounded-md border">
          <h2 className="text-lg font-semibold mb-3 text-gray-700">Add New Note</h2>
          <div className="mb-3">
            <input
              type="text"
              placeholder="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full p-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black"
              required
            />
          </div>
          <div className="mb-3">
            <textarea
              placeholder="Content"
              value={content}
              onChange={(e) => setContent(e.target.value)}
              className="w-full p-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black"
              rows={3}
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-600 text-white p-2 rounded-md hover:bg-blue-700 transition"
          >
            Add Note
          </button>
        </form>

        <div>
          <h2 className="text-xl font-semibold mb-4 text-gray-800">Your Notes</h2>
          {loading ? (
            <p className="text-gray-500 text-center">Loading notes...</p>
          ) : notes.length === 0 ? (
            <p className="text-gray-500 text-center">No notes found.</p>
          ) : (
            <ul className="space-y-4">
              {notes.map((note) => (
                <li key={note.id} className="p-4 border rounded-md bg-gray-50 relative group">
                  <h3 className="text-lg font-bold text-gray-800">{note.title}</h3>
                  <p className="text-gray-600 mt-1 whitespace-pre-wrap">{note.content}</p>
                  <div className="mt-2 text-xs text-gray-400">
                    {new Date(note.created_at).toLocaleString()}
                  </div>
                  <button
                    onClick={() => handleDeleteNote(note.id)}
                    className="absolute top-4 right-4 text-red-500 hover:text-red-700 opacity-0 group-hover:opacity-100 transition"
                    title="Delete Note"
                  >
                    Delete
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>
      </main>
    </div>
  );
}
