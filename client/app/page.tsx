"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Plus, Eye } from "lucide-react";
import { Note } from "@/app/types";
import { getNotes, saveNotes } from "@/app/utils";
import axios from "axios";

export default function HomePage() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const router = useRouter();

  useEffect(() => {
    setNotes(getNotes());
  }, []);

  const handleAddNote = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API}/notes`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title,
          content,
        }),
        credentials: "include",
      })
      const data = await res.json()
      console.log(data)
    } catch (error) {
      console.error(error)
    }
  };

  const fetchNotes = async ()=> {
    try {
      const res = await axios.get(`https://gonotes-7d8s.onrender.com/api/v1/notes`, {
      withCredentials: true,
    })
    console.log(res.data)
    } catch (error) {
      console.error(error)
    }
  }

  useEffect(()=> {
    fetchNotes();
  }, [])

  const displayedNotes = notes.slice(0, 6); // Show first 6

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 p-4">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold text-gray-900 mb-2">Go Notes</h1>
          <p className="text-gray-600">Create and manage your notes</p>
        </div>

        {/* Add Note Form */}
        <div className="bg-white rounded-2xl shadow-xl p-6 mb-8 border border-gray-100">
          <h2 className="text-2xl font-semibold text-gray-900 mb-4">
            Add New Note
          </h2>
          <form onSubmit={handleAddNote} className="space-y-4">
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Note Title"
              className="w-full text-black outline-none p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 hover:bg-white"
              required
            />
            <textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
              placeholder="Note Content"
              rows={4}
              className="w-full text-black outline-none p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 hover:bg-white"
              required
            />
            <button
              type="submit"
              className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-3 px-6 rounded-lg font-semibold hover:from-blue-700 hover:to-purple-700 flex items-center gap-2"
            >
              <Plus className="h-5 w-5" />
              Add Note
            </button>
          </form>
        </div>

        {/* Notes Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
          {displayedNotes.map((note) => (
            <div
              key={note.id}
              onClick={() => router.push(`/notes/${note.id}`)}
              className="bg-white rounded-xl shadow-lg p-6 border border-gray-100 cursor-pointer hover:shadow-xl transition-shadow duration-200"
            >
              <h3 className="text-xl font-semibold text-gray-900 mb-2 truncate">
                {note.title}
              </h3>
              <p className="text-gray-600 line-clamp-3">{note.content}</p>
            </div>
          ))}
        </div>

        {/* View All Button */}
        {notes.length > 6 && (
          <div className="text-center">
            <button
              onClick={() => router.push("/notes")}
              className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-3 px-6 rounded-lg font-semibold hover:from-blue-700 hover:to-purple-700 flex items-center gap-2"
            >
              <Eye className="h-5 w-5" />
              View All Notes
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
