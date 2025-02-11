import { IUser } from "@/schemas/user.schema";
import { create } from "zustand";

interface IUserStore {
  user: IUser | null;
  setUser: (user: IUser) => void;
}

export const userStore = create<IUserStore>((set) => ({
  user: null,
  setUser: (user) => set({ user }),
}));
