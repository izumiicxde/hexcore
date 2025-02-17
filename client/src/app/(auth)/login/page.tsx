"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Form } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { FormInput } from "../_components/form-input"; // Import the reusable component
import { toast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";
import { loginFormSchema } from "@/schemas/user";
import Link from "next/link";

const defaultValues: z.infer<typeof loginFormSchema> = {
  identifier: "",
  password: "",
};

export default function SignupForm() {
  const router = useRouter();
  const form = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues,
  });

  const { control, handleSubmit } = form;

  const onSubmit = async (values: z.infer<typeof loginFormSchema>) => {
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
        <CardTitle className="text-center text-3xl font-serif">
          Welcome Back,
        </CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <FormInput
              name="identifier"
              label="Email or Username"
              control={control}
            />
            <FormInput
              name="password"
              label="Password"
              type="password"
              control={control}
            />
            <Button type="submit" className="w-full">
              Sign Up
            </Button>
          </form>
        </Form>
      </CardContent>
      <CardFooter>
        <p>
          Don&apos;t have an account?{" "}
          <Link href="/login" className="underline">
            Sign up
          </Link>
        </p>
      </CardFooter>
    </Card>
  );
}
