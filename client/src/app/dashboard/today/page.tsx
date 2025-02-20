"use client";
import Profile from "@/components/profile";
import { toast } from "@/hooks/use-toast";
import { classesStore } from "@/store/store";
import { IClassesAPIResponse } from "@/types/classes";
import { useEffect } from "react";
import { Checkbox } from "@/components/ui/checkbox";

export default () => {
  const { classes, setClasses } = classesStore();

  const fetchClasses = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/today`;
    try {
      const response = await fetch(url, {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) {
        toast({ title: "There was an error fetching today's classes" });
        return;
      }
      const data: IClassesAPIResponse = await response.json();
      console.log(data);
      toast({ title: data.message });
      setClasses(data.classes);
    } catch (error) {
      toast({ title: "There was an unexpected error fetching classes" });
    }
  };

  const markAttendance = async (subjectId: number) => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/mark`,
        {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ subjectId, status: true }),
        }
      );
      if (!response.ok) {
        toast({ title: "Error marking attendance" });
        return;
      }
      toast({ title: "Attendance marked successfully" });
    } catch (error) {
      toast({ title: "Unexpected error occurred" });
    }
  };

  useEffect(() => {
    if (classes == null) {
      fetchClasses();
    }
  }, []);

  return (
    <div className="flex flex-col items-center w-full min-h-screen p-5">
      <Profile />
      <h1 className="text-2xl font-bold mt-6 mb-4">Today's Classes</h1>
      <div className="w-full max-w-2xl space-y-4">
        {classes &&
          classes.map((c) => (
            <div
              key={c.ID}
              className="p-6 shadow-lg rounded-xl border flex items-center justify-between"
            >
              <div>
                <p className="text-lg font-semibold uppercase">{c.name}</p>
                <div className="mt-2 space-y-1">
                  <p>
                    <span className="font-medium">Attended classes:</span>{" "}
                    {c.attendedClasses}
                  </p>
                  <p>
                    <span className="font-medium">Total Taken:</span>{" "}
                    {c.totalTaken}
                  </p>
                  <p>
                    <span className="font-medium">Maximum classes:</span>{" "}
                    {c.maxClasses}
                  </p>
                  <p>
                    <span className="font-medium">Remaining classes:</span>{" "}
                    {c.maxClasses - c.totalTaken}
                  </p>
                </div>
              </div>
              <p className="flex justify-center items-center gap-2">
                <Checkbox
                  checked={c.status}
                  onCheckedChange={() => markAttendance(c.ID)}
                />{" "}
                {c.status ? "present" : ""}
              </p>
            </div>
          ))}
      </div>
    </div>
  );
};
