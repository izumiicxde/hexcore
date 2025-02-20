import { IUser } from "@/types/user";
import { persist } from "zustand/middleware";
import { create } from "zustand";
import { IClasses, Summary } from "@/types/classes";

interface IUserStore {
  user: IUser | null;
  setUser: (state: IUser) => void;
}

export const userStore = create<IUserStore>()(
  persist(
    (set) => ({
      user: null,
      setUser: (user) => set(() => ({ user })),
    }),
    { name: "user-storage" }
  )
);

interface IClassesStore {
  classes: IClasses[] | null;
  setClasses: (classes: any) => void;
}
export const classesStore = create<IClassesStore>()(
  persist(
    (set) => ({
      classes: null,
      setClasses: (classes: IClasses[]) => set(() => ({ classes })),
    }),
    { name: "classes-storage" }
  )
);

interface ISubjectSummaryStore {
  summary: Summary | null;
  setSummary: (arg0: Summary) => void;
}
export const subjectSummaryStore = create<ISubjectSummaryStore>()(
  persist(
    (set) => ({
      summary: null,
      setSummary: (summary: Summary) => set(() => ({ summary })),
    }),
    { name: "subject-summary" }
  )
);
