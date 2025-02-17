import { z } from "zod";

export const signupSchema = z
  .object({
    username: z
      .string()
      .min(3, "Username must be at least 3 characters")
      .max(24, "Username cannot be more than 24 characters")
      .regex(
        /^[a-z0-9_]+$/,
        "Username can only contain lowercase letters, numbers, and underscores"
      ),
    fullname: z.string().min(3, "Full name must be at least 3 characters"),
    email: z.string().email("Invalid email address"),
    password: z.string().min(6, "Password must be at least 6 characters"),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords must match",
    path: ["confirmPassword"],
  });

export type TSignupFormValues = z.infer<typeof signupSchema>;

export const loginFormSchema = z.object({
  identifier: z.string(),
  password: z.string(),
});
export type TLoginFormValues = z.infer<typeof loginFormSchema>;
