"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { signupSchema } from "@/schemas/user";
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
import Link from "next/link";
import { useState } from "react";
import { Loader2Icon } from "lucide-react";
import { userStore } from "@/store/user";
import { IAPIResponse } from "@/types/user";

const defaultValues: z.infer<typeof signupSchema> = {
  register_no: "",
  fullname: "",
  email: "",
  password: "",
  confirmPassword: "",
};

export default function SignupForm() {
  const [isSubmitingForm, setIsSubmitingForm] = useState<boolean>(false);
  const { user, setUser } = userStore();

  const router = useRouter();
  const form = useForm<z.infer<typeof signupSchema>>({
    resolver: zodResolver(signupSchema),
    defaultValues,
  });

  const { control, handleSubmit } = form;

  const onSubmit = async (values: z.infer<typeof signupSchema>) => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/signup`;
    try {
      setIsSubmitingForm(true);
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
          title: "Error",
          description: data.message,
        });
        return;
      }

      toast({
        title: "Successfully registered",
      });

      if (data.user) setUser(data.user); // set the user state
      router.push("/");
    } catch (error) {
      toast({
        title: "error registering the user",
      });
    } finally {
      setIsSubmitingForm(false);
    }
  };

  return (
    <Card className="max-w-md w-full mx-auto p-6 ">
      <CardHeader>
        <CardTitle className="text-center text-3xl font-serif">
          Welcome to Zen0
        </CardTitle>
      </CardHeader>
      <CardContent className="">
        <Form {...form}>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 ">
            <FormInput
              name="register_no"
              label="Register Number"
              control={control}
            />
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

            <Button type="submit" disabled={isSubmitingForm} className="w-full">
              {isSubmitingForm ? (
                <p className="flex justify-center items-center gap-0.5">
                  <Loader2Icon className="animate-spin " />
                  Signing up
                </p>
              ) : (
                "Sign up"
              )}
            </Button>
          </form>
        </Form>
      </CardContent>
      <CardFooter>
        <p className="text-xs">
          Already have an account?{" "}
          <Link href="/login" className="underline">
            Sign in
          </Link>
        </p>
      </CardFooter>
    </Card>
  );
}
