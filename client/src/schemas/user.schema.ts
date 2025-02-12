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

export const signUpFormSchema = z
  .object({
    email: z.string().email(),
    username: z
      .string()
      .min(4)
      .max(24)
      .regex(
        /^[a-z0-9_]+$/,
        "Username must be lowercase and can only contain letters, numbers, and underscores"
      ),
    password: z
      .string()
      .min(8)
      .max(24)
      .regex(
        /^(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/,
        "Password must contain at least one uppercase letter, one number, and one special character"
      ),
    confirmPassword: z.string().min(8).max(24),
    fullName: z.string().min(4).max(24),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

export const loginFormSchema = z.object({
  identifier: z.union([
    z.string().email(),
    z
      .string()
      .max(24)
      .regex(/^[a-zA-Z0-9_]+$/, "Invalid username format"),
  ]),
  password: z.string(),
});
