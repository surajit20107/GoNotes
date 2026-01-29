"use client";

import { useState, useEffect } from "react";
import { useRouter, useParams } from "next/navigation";
import { Edit, Trash, Save, X } from "lucide-react";
import { Note } from "@/app/types";
import { getNotes, saveNotes } from "@/app/utils";

export default function NotePage() {
  const { noteid } = useParams();
  const router = useRouter();
  const [note, setNote] = useState<Note | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [editTitle, setEditTitle] = useState("");
  const [editContent, setEditContent] = useState("");

  useEffect(() => {
    const notes = getNotes();
    const foundNote = notes.find((n) => n.id === noteid);
    if (foundNote) {
      setNote(foundNote);
      setEditTitle(foundNote.title);
      setEditContent(foundNote.content);
    } else {
      router.push("/");
    }
  }, [noteid, router]);

  const handleSave = () => {
    if (!note) return;
    const notes = getNotes();
    const updatedNotes = notes.map((n) =>
      n.id === note.id
        ? { ...n, title: editTitle.trim(), content: editContent.trim() }
        : n,
    );
    saveNotes(updatedNotes);
    setNote({ ...note, title: editTitle.trim(), content: editContent.trim() });
    setIsEditing(false);
  };

  const handleDelete = () => {
    if (!note || !confirm("Are you sure you want to delete this note?")) return;
    const notes = getNotes().filter((n) => n.id !== note.id);
    saveNotes(notes);
    router.push("/");
  };

  if (!note) return <div>Loading...</div>;

  useEffect(() => {
    const token = localStorage.getItem("token")
    if (!token) {
      router.push("/login")
    }
  }, [])

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 p-4">
      <div className="max-w-2xl mx-auto bg-white rounded-2xl shadow-xl p-8 border border-gray-100">
        {isEditing ? (
          <div className="space-y-4">
            <input
              type="text"
              value={editTitle}
              onChange={(e) => setEditTitle(e.target.value)}
              className="w-full text-black outline-none p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500"
            />
            <textarea
              value={editContent}
              onChange={(e) => setEditContent(e.target.value)}
              rows={10}
              className="w-full text-black outline-none p-3 border border-gray-300 rounded-lg focus:ring-1 focus:ring-blue-500"
            />
            <div className="flex gap-4">
              <button
                onClick={handleSave}
                className="bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 flex items-center gap-2"
              >
                <Save className="h-5 w-5" />
                Save
              </button>
              <button
                onClick={() => setIsEditing(false)}
                className="bg-gray-600 text-white py-2 px-4 rounded-lg hover:bg-gray-700 flex items-center gap-2"
              >
                <X className="h-5 w-5" />
                Cancel
              </button>
            </div>
          </div>
        ) : (
          <>
            <h1 className="text-3xl font-bold text-gray-900 mb-4">
              {note.title}
            </h1>
            <p className="text-gray-700 mb-6 whitespace-pre-wrap">
              {note.content}
            </p>
            <div className="flex gap-4">
              <button
                onClick={() => setIsEditing(true)}
                className="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 flex items-center gap-2"
              >
                <Edit className="h-5 w-5" />
                Edit
              </button>
              <button
                onClick={handleDelete}
                className="bg-red-600 text-white py-2 px-4 rounded-lg hover:bg-red-700 flex items-center gap-2"
              >
                <Trash className="h-5 w-5" />
                Delete
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
