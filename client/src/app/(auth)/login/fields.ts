export const fields: {
  name: "identifier" | "password";
  label: string;
  type: string;
  placeholder: string;
  description?: string;
}[] = [
  {
    name: "identifier",
    label: "Email/Username",
    type: "string",
    placeholder: "Enter your email or username",
  },
  {
    name: "password",
    label: "Password",
    type: "password",
    placeholder: "Enter a password",
  },
];
