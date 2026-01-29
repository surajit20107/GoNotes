"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Plus, LogOut } from "lucide-react";
import { apiFetch, Note } from "./utils";

export default function HomePage() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  const fetchNotes = async () => {
    try {
      const data = await apiFetch("/notes");
      setNotes(data.data || []);
    } catch (error) {
      console.error(error);
      router.push("/login");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchNotes();
  }, []);

  const handleAddNote = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiFetch("/notes", {
        method: "POST",
        body: JSON.stringify({ title, content }),
      });
      setTitle("");
      setContent("");
      fetchNotes();
    } catch (error) {
      console.error(error);
    }
  };

  const handleLogout = async () => {
    try {
      await apiFetch("/auth/logout", { method: "POST" });
      router.push("/login");
    } catch (error) {
      console.error(error);
    }
  };

  if (loading) return <div className="min-h-screen flex items-center justify-center text-black">Loading...</div>;

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 p-4 text-black">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-4xl font-bold text-gray-900 mb-2">Go Notes</h1>
            <p className="text-gray-600">Manage your ideas</p>
          </div>
          <button onClick={handleLogout} className="flex items-center gap-2 text-red-600 hover:underline">
            <LogOut className="h-5 w-5" /> Logout
          </button>
        </div>

        <div className="bg-white rounded-2xl shadow-xl p-6 mb-8 border border-gray-100">
          <h2 className="text-2xl font-semibold text-gray-900 mb-4">Add New Note</h2>
          <form onSubmit={handleAddNote} className="space-y-4">
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Note Title"
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 outline-none bg-gray-50"
              required
            />
            <textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
              placeholder="Note Content"
              rows={4}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 outline-none bg-gray-50"
              required
            />
            <button
              type="submit"
              className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-3 px-6 rounded-lg font-semibold hover:opacity-90 flex items-center gap-2"
            >
              <Plus className="h-5 w-5" /> Add Note
            </button>
          </form>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {notes.map((note) => (
            <div
              key={note.id}
              onClick={() => router.push(`/notes/${note.id}`)}
              className="bg-white rounded-xl shadow-lg p-6 border border-gray-100 cursor-pointer hover:shadow-xl transition-all"
            >
              <h3 className="text-xl font-semibold text-gray-900 mb-2 truncate">{note.title}</h3>
              <p className="text-gray-600 line-clamp-3">{note.content}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
