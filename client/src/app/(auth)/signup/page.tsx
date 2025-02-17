"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { signupSchema } from "@/schemas/user";
import { Form } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { FormInput } from "../_components/form-input"; // Import the reusable component
import { onSubmit } from "./onsubmit";

const defaultValues: z.infer<typeof signupSchema> = {
  username: "",
  fullname: "",
  email: "",
  password: "",
  confirmPassword: "",
};

export default function SignupForm() {
  const form = useForm<z.infer<typeof signupSchema>>({
    resolver: zodResolver(signupSchema),
    defaultValues,
  });

  const { control, handleSubmit } = form;

  return (
    <Card className="max-w-md w-full mx-auto p-6 ">
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
