"use client";

import * as React from "react";

import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp";
import { Button } from "@/components/ui/button";
import { toast } from "@/hooks/use-toast";
import { IAPIResponse } from "@/types/user";

export default function page() {
  const [value, setValue] = React.useState<string>("");
  const [verified, setVerified] = React.useState<boolean>(false);

  const handleSubmit = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/verify`;
    try {
      if (value.length < 6) return;
      const response = await fetch(url, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(value),
      });
      if (!response.ok) toast({ title: "failed to verify" });

      const data: IAPIResponse = await response.json();
      toast({ title: data.message });
      if (data.redirect) toast({ title: "Email verified successfully" });
    } catch (error: any) {
      toast({ title: error.message });
    }
  };
  return (
    <div className="space-y-2">
      <p className="text-xl">Enter your one time password</p>
      <InputOTP
        maxLength={6}
        value={value}
        onChange={(value) => setValue(value)}
      >
        <InputOTPGroup>
          <InputOTPSlot index={0} />
          <InputOTPSlot index={1} />
          <InputOTPSlot index={2} />
          <InputOTPSlot index={3} />
          <InputOTPSlot index={4} />
          <InputOTPSlot index={5} />
        </InputOTPGroup>
      </InputOTP>
      <div className="text-center text-sm">
        <Button className="" onClick={handleSubmit}>
          Submit
        </Button>
      </div>
    </div>
  );
}
