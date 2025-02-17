import { toast } from "@/hooks/use-toast";
import { signupSchema } from "@/schemas/user";
import { useRouter } from "next/router";
import { z } from "zod";

export const onSubmit = async (values: z.infer<typeof signupSchema>) => {
  const router = useRouter();
  const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/signup`;
  try {
    const response = await fetch(url, {
      method: "POST",
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
    console.log(error);
  }
};
