"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { signupSchema } from "@/schemas/user";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

// Default form values (avoids re-renders)
const defaultValues: z.infer<typeof signupSchema> = {
  username: "",
  fullname: "",
  email: "",
  password: "",
  confirmPassword: "",
};

// Reusable Form Input Component
const FormInput = ({
  name,
  label,
  type = "text",
  control,
}: {
  name: keyof typeof defaultValues;
  label: string;
  type?: string;
  control: any;
}) => (
  <FormField
    control={control}
    name={name}
    render={({ field }) => (
      <FormItem>
        <FormLabel>{label}</FormLabel>
        <FormControl>
          <Input
            type={type}
            placeholder={`Enter your ${label.toLowerCase()}`}
            {...field}
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    )}
  />
);

export default function SignupForm() {
  const form = useForm<z.infer<typeof signupSchema>>({
    resolver: zodResolver(signupSchema),
    defaultValues,
  });
  const { control, handleSubmit } = form;

  const onSubmit = (values: z.infer<typeof signupSchema>) =>
    console.log("Signup Data:", values);

  return (
    <Card className="max-w-md w-full mx-auto p-6 border-red-500 border-4">
      <CardHeader>
        <CardTitle className="text-center">Sign Up</CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <FormInput name="username" label="Username" control={control} />
            <FormInput name="fullname" label="Full Name" control={control} />
            <FormInput
              name="email"
              label="Email"
              type="email"
              control={control}
            />
            <FormInput
              name="password"
              label="Password"
              type="password"
              control={control}
            />
            <FormInput
              name="confirmPassword"
              label="Confirm Password"
              type="password"
              control={control}
            />

            <Button type="submit" className="w-full">
              Sign Up
            </Button>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
