"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { signupSchema } from "@/schemas/user";
import { Form } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { FormInput } from "../_components/form-input"; // Import the reusable component
import { toast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";

const defaultValues: z.infer<typeof signupSchema> = {
  username: "",
  fullname: "",
  email: "",
  password: "",
  confirmPassword: "",
};

export default function SignupForm() {
  const router = useRouter();
  const form = useForm<z.infer<typeof signupSchema>>({
    resolver: zodResolver(signupSchema),
    defaultValues,
  });

  const { control, handleSubmit } = form;

  const onSubmit = async (values: z.infer<typeof signupSchema>) => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/signup`;
    try {
      const response = await fetch(url, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });

      const data = await response.json();
      if (!response.ok) {
        toast({
          title: "Error",
          description: data.message,
        });
        return;
      }
      toast({
        title: "Successfully registered",
      });
      router.push("/");
    } catch (error) {
      toast({
        title: "error registering the user",
      });
    }
  };

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
