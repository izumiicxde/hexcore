import { z } from "zod";

export interface IUser {
  email: string;
  password: string;
  confirmPassword: string;
  username: string;
  fullname: string;
  role: string;
  id: number;
}
export const signUpFormSchema = z.object({
  email: z.string().email(),
  username: z.string().min(4).max(24),
  password: z.string().min(8).max(24),
  confirmPassword: z.string().min(8).max(24),
  fullName: z.string().min(4).max(24),
});
