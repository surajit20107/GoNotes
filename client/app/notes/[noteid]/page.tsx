"use client";

import { useState, useEffect } from "react";
import { useRouter, useParams } from "next/navigation";
import { Edit, Trash, Save, X, ArrowLeft } from "lucide-react";
import { apiFetch, Note } from "../../utils";

export default function NotePage() {
  const { noteid } = useParams();
  const router = useRouter();
  const [note, setNote] = useState<Note | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [editTitle, setEditTitle] = useState("");
  const [editContent, setEditContent] = useState("");
  const [loading, setLoading] = useState(true);

  const fetchNote = async () => {
    try {
      const data = await apiFetch(`/notes/${noteid}`);
      setNote(data.data);
      setEditTitle(data.data.title);
      setEditContent(data.data.content);
    } catch (error) {
      console.error(error);
      router.push("/");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (noteid) fetchNote();
  }, [noteid]);

  const handleSave = async () => {
    try {
      await apiFetch(`/notes/${noteid}`, {
        method: "PUT",
        body: JSON.stringify({ title: editTitle, content: editContent }),
      });
      setIsEditing(false);
      fetchNote();
    } catch (error) {
      console.error(error);
    }
  };

  const handleDelete = async () => {
    if (!confirm("Are you sure?")) return;
    try {
      await apiFetch(`/notes/${noteid}`, { method: "DELETE" });
      router.push("/");
    } catch (error) {
      console.error(error);
    }
  };

  if (loading) return <div className="min-h-screen flex items-center justify-center text-black">Loading...</div>;
  if (!note) return null;

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 p-4 text-black">
      <div className="max-w-2xl mx-auto bg-white rounded-2xl shadow-xl p-8 border border-gray-100">
        <button onClick={() => router.push("/")} className="mb-6 flex items-center gap-2 text-gray-600 hover:text-black">
          <ArrowLeft className="h-5 w-5" /> Back to Notes
        </button>

        {isEditing ? (
          <div className="space-y-4">
            <input
              type="text"
              value={editTitle}
              onChange={(e) => setEditTitle(e.target.value)}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 outline-none"
            />
            <textarea
              value={editContent}
              onChange={(e) => setEditContent(e.target.value)}
              rows={10}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500 outline-none"
            />
            <div className="flex gap-4">
              <button onClick={handleSave} className="bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 flex items-center gap-2">
                <Save className="h-5 w-5" /> Save
              </button>
              <button onClick={() => setIsEditing(false)} className="bg-gray-600 text-white py-2 px-4 rounded-lg hover:bg-gray-700 flex items-center gap-2">
                <X className="h-5 w-5" /> Cancel
              </button>
            </div>
          </div>
        ) : (
          <>
            <h1 className="text-3xl font-bold text-gray-900 mb-4">{note.title}</h1>
            <p className="text-gray-700 mb-6 whitespace-pre-wrap">{note.content}</p>
            <div className="flex gap-4">
              <button onClick={() => setIsEditing(true)} className="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 flex items-center gap-2">
                <Edit className="h-5 w-5" /> Edit
              </button>
              <button onClick={handleDelete} className="bg-red-600 text-white py-2 px-4 rounded-lg hover:bg-red-700 flex items-center gap-2">
                <Trash className="h-5 w-5" /> Delete
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
