import { Note } from "@/app/types";

export const getNotes = (): Note[] => {
  if (typeof window === "undefined") return [];
  const notes = localStorage.getItem("notes");
  return notes ? JSON.parse(notes) : [];
};

export const saveNotes = (notes: Note[]) => {
  if (typeof window !== "undefined") {
    localStorage.setItem("notes", JSON.stringify(notes));
  }
};
