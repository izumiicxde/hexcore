"use client";
import { signUpFormSchema } from "@/schemas/user.schema";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { fields } from "./fields";
import { useRouter } from "next/navigation";
import { toast } from "@/hooks/use-toast";

const page = () => {
  const router = useRouter();

  const form = useForm<z.infer<typeof signUpFormSchema>>({
    resolver: zodResolver(signUpFormSchema),
    defaultValues: fields.reduce(
      (acc, field) => ({ ...acc, [field.name]: "" }),
      {}
    ),
  });

  const onSubmit = async (values: z.infer<typeof signUpFormSchema>) => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/users/register`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(values),
          credentials: "include",
        }
      );
      const data = await response.json();
      if (!response.ok) {
        toast({
          title: "Error signing up",
          description: data.message,
        });
      } else {
        toast({
          title: "Successfully signed up",
        });
        router.push("/");
      }
    } catch (err) {
      toast({
        title: "Error signing up",
        description: "Something went wrong",
      });
    }
  };

  return (
    <div className="flex flex-col justify-center items-center w-full h-full">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="flex flex-col gap-5 w-full max-w-md"
        >
          {fields.map(({ name, label, type, placeholder, description }) => (
            <FormField
              key={name}
              control={form.control}
              name={name}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{label}</FormLabel>
                  <FormControl>
                    <Input type={type} placeholder={placeholder} {...field} />
                  </FormControl>
                  {description && (
                    <FormDescription>{description}</FormDescription>
                  )}
                  <FormMessage />
                </FormItem>
              )}
            />
          ))}
          <Button type="submit">Sign Up</Button>
        </form>
      </Form>
    </div>
  );
};

export default page;
