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
import { useState } from "react";
import { Loader2Icon } from "lucide-react";
import { userStore } from "@/store/store";
import { IAPIResponse } from "@/types/user";

const defaultValues: z.infer<typeof loginFormSchema> = {
  identifier: "",
  password: "",
};

export default function LoginForm() {
  const [isSubmittingForm, setIsSubmittingForm] = useState(false);
  const { setUser } = userStore();

  const router = useRouter();
  const form = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues,
  });

  const { control, handleSubmit } = form;

  const onSubmit = async (values: z.infer<typeof loginFormSchema>) => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/login`;
    try {
      setIsSubmittingForm(true);
      const response = await fetch(url, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });

      const data: IAPIResponse = await response.json();
      if (!response.ok) {
        toast({
          title: "login error",
          description: data.message,
        });
        return;
      }
      toast({
        title: "Successfully logged in",
      });

      if (data.user) setUser(data.user); // set user state

      if (data.user?.isVerified) router.push("/");
      else router.push("/verify");
    } catch (error) {
      toast({
        title: "error logging in, please try again later",
      });
    } finally {
      setIsSubmittingForm(false);
    }
  };

  return (
    <Card className="max-w-md w-full mx-auto p-6 ">
      <CardHeader>
        <CardTitle className="text-center text-3xl font-serif">
          Welcome Back
        </CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <FormInput
              name="identifier"
              label="Email or register no."
              control={control}
            />
            <FormInput
              name="password"
              label="Password"
              type="password"
              control={control}
            />
            <Button
              type="submit"
              disabled={isSubmittingForm}
              className="w-full flex justify-center items-center"
            >
              {isSubmittingForm ? (
                <p className="flex justify-center items-center  w-full h-full gap-0.5">
                  <Loader2Icon className="animate-spin" /> Signing in
                </p>
              ) : (
                "Sign in"
              )}
            </Button>
          </form>
        </Form>
      </CardContent>
      <CardFooter>
        <p className="text-xs">
          Don&apos;t have an account?{" "}
          <Link href="/signup" className="underline">
            Sign up
          </Link>
        </p>
      </CardFooter>
    </Card>
  );
}
