import { IUser } from "@/types/user";
import { persist } from "zustand/middleware";
import { create } from "zustand";

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
