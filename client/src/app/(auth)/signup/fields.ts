export const fields: {
  name: "email" | "username" | "fullName" | "password" | "confirmPassword";
  label: string;
  type: string;
  placeholder: string;
  description?: string;
}[] = [
  {
    name: "email",
    label: "Email",
    type: "email",
    placeholder: "Enter your email",
  },
  {
    name: "username",
    label: "Username",
    type: "text",
    placeholder: "Choose a username",
    description: "This is your public display name.",
  },
  {
    name: "fullName",
    label: "Full Name",
    type: "text",
    placeholder: "Enter your full name",
  },
  {
    name: "password",
    label: "Password",
    type: "password",
    placeholder: "Enter a password",
  },
  {
    name: "confirmPassword",
    label: "Confirm Password",
    type: "password",
    placeholder: "Confirm your password",
  },
];
