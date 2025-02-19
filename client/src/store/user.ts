import { IUser } from "@/types/user";
import { create } from "zustand";

interface IUserStore {
  user: IUser | null;
  setUser: (state: IUser) => void;
}

export const userStore = create<IUserStore>((set) => ({
  user: null,
  setUser: () => set((state) => ({ user: state.user })),
}));
