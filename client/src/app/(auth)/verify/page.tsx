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
import { userStore } from "@/store/user";
import { MailIcon } from "lucide-react";

export default function Verify() {
  const [value, setValue] = React.useState<string>("");
  const { user } = userStore();

  const onSubmit = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/verify?code=${value}`;
    try {
      if (value.length < 6) return;
      const response = await fetch(url, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },

        body: JSON.stringify({ code: value }),
      });
      if (!response.ok) toast({ title: "failed to verify" });

      const data: IAPIResponse = await response.json();
      toast({ title: data.message });
      if (data.redirect) toast({ title: "Email verified successfully" });
    } catch (error: any) {
      toast({ title: error.message });
    }
  };

  const onResend = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/auth/verificationCode`;
    const response = await fetch(url, {
      method: "GET",
      credentials: "include",
    });

    const data: IAPIResponse = await response.json();
    if (!response.ok) {
      toast({ title: data.message });
      return;
    }
    toast({ title: data.message });
  };

  return (
    <div className="max-w-md mx-auto p-6 space-y-6 rounded-lg shadow-md">
      <div className="flex flex-col justify-start items-center ">
        <MailIcon className="size-32" />
        <h1 className="text-xl">Email Confirmation</h1>
      </div>
      <div className="text-center">
        <h2 className="text-2xl font-semibold">Hello, {user?.fullname}</h2>
        <p className="text-sm leading-relaxed">
          We've sent a verification code to your email address:
          <br />
          <span className="font-medium">{user?.email}</span> Please check your
          inbox
        </p>
      </div>
      <div className="space-y-4 flex flex-col justify-center items-center">
        <p className="text-lg font-medium text-center">
          Enter your verification code
        </p>
        <InputOTP
          maxLength={6}
          value={value}
          onChange={(value) => setValue(value)}
        >
          <InputOTPGroup>
            {[...Array(6)].map((_, index) => (
              <InputOTPSlot key={index} index={index} autoFocus={index === 0} />
            ))}
          </InputOTPGroup>
        </InputOTP>
      </div>
      <div className="text-center space-y-4">
        <Button
          onClick={onSubmit}
          className="w-full sm:w-auto px-6 py-2 rounded-full font-medium transition-transform hover:scale-105"
        >
          Verify
        </Button>
        <p className="text-sm">
          Didn't receive the code?{" "}
          <button onClick={onResend} className="underline font-medium">
            Resend
          </button>
        </p>
      </div>
    </div>
  );
}
