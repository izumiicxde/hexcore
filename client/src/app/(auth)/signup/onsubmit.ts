import { toast } from "@/hooks/use-toast";
import { signupSchema } from "@/schemas/user";
import { z } from "zod";

export const onSubmit = async (values: z.infer<typeof signupSchema>) => {
  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/signup`, {
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

    console.log(data);
  } catch (error) {
    toast({
      title: "error registering the user",
    });
    console.log(error);
  }
};
